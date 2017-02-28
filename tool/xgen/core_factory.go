// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"path/filepath"

	enti "gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/util/envs"
	"gopkg.in/goyy/goyy.v0/util/files"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

// factory is a file generation factory.
type factory struct {
	Project           string
	PackageName       string
	EntiPath          string
	AdmPath           string
	APIPath           string
	TstPath           string
	HasGenProj        bool
	HasGenEntity      bool
	HasGenService     bool
	HasGenController  bool
	HasGenDto         bool
	HasGenAPI         bool
	HasGenSQL         bool
	HasGenLog         bool
	HasGenUtil        bool
	HasGenConst       bool
	HasGenHTML        bool
	HasGenJs          bool
	IsTimeField       bool
	IsValidationField bool
	IsExtend          bool
	NewProjName       string // Name of new project
	NewProjPath       string // Path of new project
	NewProjPkg        string // Pkg Path of new project
	NewProjTitle      string // Title of new project
	NewProjHost       string // Host of new project
	Entities          []*entity
	SysColumns        []string // goyy>data>entity:SysColumns
	SysFields         []string // goyy>data>entity:SysFields
	TreeColumns       []string // goyy>data>entity:TreeColumns
	TreeFields        []string // goyy>data>entity:TreeFields
}

// Init initializes an File from a path.
func (me *factory) Init(path string) error {
	// Set up the inheritance list of column names
	me.SysColumns = enti.SysColumns[:]
	me.SysFields = enti.SysFields[:]
	me.TreeColumns = enti.TreeColumns[:]
	me.TreeFields = enti.TreeFields[:]

	// set the path
	if strings.HasSuffix(path, ".go") {
		me.EntiPath = strings.TrimSuffix(path, ".go")
	} else {
		return fmt.Errorf("File '%s' is not a Go file.", path)
	}

	f, err := parser.ParseFile(
		token.NewFileSet(),
		path,
		nil,
		parser.ParseComments,
	)
	if err != nil {
		return fmt.Errorf("Unable to parse '%s': %s", path, err)
	}

	// get package name
	if f.Name != nil {
		me.PackageName = f.Name.Name
	} else {
		return fmt.Errorf("Missing package name in '%s'", path)
	}

	// build list of entities
	var isEntity bool
	for _, decl := range f.Decls {

		// get the type declaration
		tdecl, ok := decl.(*ast.GenDecl)
		if !ok || tdecl.Doc == nil {
			continue
		}

		// find the @entity decorator
		isEntity = false
		project := ""
		module := ""
		extend := ""
		relationship := ""
		for _, comment := range tdecl.Doc.List {
			if strings.Contains(comment.Text, "@entity") {
				isEntity = true
				// get entity.Project and entity.Extend
				c := strings.Between(comment.Text, "@entity(", ")")
				if strings.IsNotBlank(c) {
					val := convertUTF8(c)
					if strings.IsNotBlank(val) {
						project = tagItemValue(val, "project")
						module = tagItemValue(val, "module")
						relationship = tagItemValue(val, "relationship")
						if relationship != "slave" {
							relationship = "master"
						}
					}
				}
				break
			}
		}
		if !isEntity {
			continue
		}

		e := &entity{Project: project, Module: module, Relationship: relationship}
		if strings.IsBlank(me.Project) {
			me.Project = project
		}

		// get the name of the entity
		for _, spec := range tdecl.Specs {
			if ts, ok := spec.(*ast.TypeSpec); ok {
				if ts.Name == nil {
					continue
				}
				e.Name = ts.Name.Name
				break
			}
		}
		if e.Name == "" {
			return fmt.Errorf("Unable to extract name from a entity struct.")
		}

		sdecl := tdecl.Specs[0].(*ast.TypeSpec).Type.(*ast.StructType)
		fields := sdecl.Fields.List
		for _, f := range fields {
			typ := me.printerType(f.Type)
			if typ == "pk" || typ == "sys" || typ == "tree" {
				extend = typ
				e.Extend = extend
				if "pk" == extend {
					e.AllColumnMaxLen = 2
					e.AllFieldMaxLen = 2
					e.AllTypeMaxLen = 6
				} else if "sys" == extend {
					e.AllColumnMaxLen = 9
					e.AllFieldMaxLen = 9
					e.AllTypeMaxLen = 6
				} else if "tree" == extend {
					e.AllColumnMaxLen = 12
					e.AllFieldMaxLen = 11
					e.AllTypeMaxLen = 6
				}
				switch extend {
				case "pk", "sys", "tree":
					col := &field{Name: "id", Type: "string", Column: "id", IsPrimary: true}
					e.PrimaryKeys = append(e.PrimaryKeys, col)
				}
				break
			}
		}
		// parse the xgen tag and build columns
		for _, f := range fields {
			typ := me.printerType(f.Type)

			if typ == "err" || typ == "pk" || typ == "sys" || typ == "tree" {
				continue
			}

			var items string

			if f.Tag == nil || strings.IsBlank(f.Tag.Value) {
				if typ == "table" {
					items = fmt.Sprintf("table=%s", strings.UnCamel(e.Name, "_"))
				} else {
					items = fmt.Sprintf("column=%s", strings.UnCamel(f.Names[0].Name, "_"))
				}
			} else {
				items = tagItemValue(f.Tag.Value, "db")
			}

			if strings.IsNotBlank(items) {
				if typ == "table" {
					// parse attributes
					attributes := strings.Split(items, "&")
					for _, attr := range attributes {
						pair := strings.Split(attr, "=")
						if len(pair) != 2 {
							return fmt.Errorf("Malformed tag: '%s'", attr)
						}
						switch strings.ToLower(pair[0]) {
						case "table":
							e.Table = pair[1]
						case "comment":
							e.Comment = pair[1]
						}
					}
					continue
				}

				col := &field{}
				if err := col.Init(f.Names[0].Name, typ, items); err != nil {
					return fmt.Errorf(
						"Unable to parse tag '%s' from entity '%s' in '%s': %v",
						items,
						e.Name,
						path,
						err,
					)
				}
				// validation init
				if f.Tag != nil && strings.IsNotBlank(f.Tag.Value) {
					items = tagItemValue(f.Tag.Value, "validation")
					if strings.IsNotBlank(items) {
						if v, ok := newValidations(items); ok {
							col.Validations = v
						} else {
							return fmt.Errorf(
								"Unable to parse tag '%s' from entity '%s' in '%s': %v",
								items,
								e.Name,
								path,
								err,
							)
						}
					}
				}
				// excel init
				if f.Tag != nil && strings.IsNotBlank(f.Tag.Value) {
					items = tagItemValue(f.Tag.Value, "excel")
					if strings.IsNotBlank(items) {
						if v, ok := newExcelField(col, items); ok {
							col.Excel = v
							col.IsExcel = true
						} else {
							return fmt.Errorf(
								"Unable to parse tag '%s' from entity '%s' in '%s': %v",
								items,
								e.Name,
								path,
								err,
							)
						}
					}
				}
				// json init
				if f.Tag != nil && strings.IsNotBlank(f.Tag.Value) {
					items = tagItemValue(f.Tag.Value, "json")
					if strings.IsBlank(items) {
						items = col.Name
					}
					if v, ok := newJSONField(col, items); ok {
						col.JSON = v
						col.IsJSON = true
					} else {
						return fmt.Errorf(
							"Unable to parse tag '%s' from entity '%s' in '%s': %v",
							items,
							e.Name,
							path,
							err,
						)
					}
				}
				// xml init
				if f.Tag != nil && strings.IsNotBlank(f.Tag.Value) {
					items = tagItemValue(f.Tag.Value, "xml")
					if strings.IsBlank(items) {
						items = col.Name
					}
					if v, ok := newXMLField(col, items); ok {
						col.XML = v
						col.IsXML = true
					} else {
						return fmt.Errorf(
							"Unable to parse tag '%s' from entity '%s' in '%s': %v",
							items,
							e.Name,
							path,
							err,
						)
					}
				}
				e.Fields = append(e.Fields, col)
				if col.IsPrimary {
					e.PrimaryKeys = append(e.PrimaryKeys, col)
				}
				if len(col.Name) > e.FieldMaxLen {
					e.FieldMaxLen = len(col.Name)
				}
				if len(col.Name) > e.AllFieldMaxLen {
					e.AllFieldMaxLen = len(col.Name)
				}
				if len(col.Column) > e.ColumnMaxLen {
					e.ColumnMaxLen = len(col.Column)
				}
				if len(col.Column) > e.AllColumnMaxLen {
					e.AllColumnMaxLen = len(col.Column)
				}
				if len(col.Type) > e.TypeMaxLen {
					e.TypeMaxLen = len(col.Type)
				}
				if len(col.Type) > e.AllTypeMaxLen {
					e.AllTypeMaxLen = len(col.Type)
				}
			}
		}
		if len(e.PrimaryKeys) > 0 {
			me.Entities = append(me.Entities, e)
		}
	}

	me.isTimeField()
	me.isValidationField()
	me.isExtend()

	return nil
}

func (me factory) Write() error {
	if me.HasGenProj {
		if err := me.writeNewProj(); err != nil {
			return err
		}
	}
	if me.HasGenEntity {
		if err := me.writeEntityXgen(); err != nil {
			return err
		}
		if err := me.writeEntitiesXgen(); err != nil {
			return err
		}
	}
	if me.HasGenService {
		if err := me.writeServiceXgen(); err != nil {
			return err
		}
		if err := me.writeServiceMain(); err != nil {
			return err
		}
		if strings.IsNotBlank(me.TstPath) {
			if err := me.writeServiceTest(); err != nil {
				return err
			}
		}
	}
	if me.HasGenController {
		if strings.IsNotBlank(me.APIPath) {
			if err := me.writeControllerXgen(); err != nil {
				return err
			}
			if err := me.writeControllerMain(); err != nil {
				return err
			}
			if strings.IsNotBlank(me.TstPath) {
				if err := me.writeControllerTest(); err != nil {
					return err
				}
			}
			if err := me.writeControllerReg(); err != nil {
				return err
			}
		}
	}
	if me.HasGenDto {
		if err := me.writeDtoXgen(); err != nil {
			return err
		}
	}
	if me.HasGenAPI {
		if strings.IsNotBlank(me.APIPath) {
			if err := me.writeAPIMain(); err != nil {
				return err
			}
		}
	}
	if me.HasGenSQL {
		if err := me.writeSQLMain(); err != nil {
			return err
		}
	}
	if me.HasGenLog {
		if err := me.writeLogJSONXgen(); err != nil {
			return err
		}
		if strings.IsNotBlank(me.APIPath) {
			if err := me.writeLogAPIXgen(); err != nil {
				return err
			}
		}
	}
	if me.HasGenUtil {
		if err := me.writeUtilMain(); err != nil {
			return err
		}
	}
	if me.HasGenConst {
		if err := me.writeConstMain(); err != nil {
			return err
		}
	}
	if me.HasGenHTML {
		if err := me.writeHTMLMain(); err != nil {
			return err
		}
	}
	if me.HasGenJs {
		if err := me.writeJsMain(); err != nil {
			return err
		}
	}
	return nil
}

func (me factory) writeBy(typ, content string) error {
	var dir, dstfile string
	if me.HasGenEntity {
		admdir := "../../../" + strings.AfterLast(me.AdmPath, "/")
		// get the destination file
		_, file := filepath.Split(me.EntiPath)
		f, name := me.genFileName(typ, file)
		switch typ {
		case xgenDto:
			dir = admdir + "/internal/" + me.Project + "/" + me.PackageName
		case mainAPI, xgenLogAPI:
			dir = "../../api/" + me.PackageName
		case mainHTML:
			dir = "../../templates/adm/" + name
		case mainJs:
			dir = "../../static/adm/js"
		case xgenCtlReg:
			dir = "../../"
		}
		dstfile = files.Join(dir, f)
	}
	if me.HasGenProj {
		switch typ {
		case newProj + ".README":
			dir = "."
			dstfile = "README.md"
		case newProj + ".tst":
			dir = me.NewProjPath + "/" + me.NewProjName + "-tst/"
			dstfile = "tst.go"
		case newProj + ".tst.settings":
			dir = me.NewProjPath + "/" + me.NewProjName + "-tst/conf/env/"
			dstfile = "settings.xml"
		case newProj + ".tst.session":
			dir = me.NewProjPath + "/" + me.NewProjName + "-tst/conf/env/"
			dstfile = "session.xml"
		case newProj + ".tst.db":
			dir = me.NewProjPath + "/" + me.NewProjName + "-tst/conf/env/"
			dstfile = "db.xml"
		case newProj + ".schema.bin.db.sh":
			dir = me.NewProjPath + "/" + me.NewProjName + "-schema/bin/"
			dstfile = "exp-db.sh"
		case newProj + ".schema.bin.db.bat":
			dir = me.NewProjPath + "/" + me.NewProjName + "-schema/bin/"
			dstfile = "exp-db.bat"
		case newProj + ".schema.bin.sql.sh":
			dir = me.NewProjPath + "/" + me.NewProjName + "-schema/bin/"
			dstfile = "exp-sql.sh"
		case newProj + ".schema.bin.sql.bat":
			dir = me.NewProjPath + "/" + me.NewProjName + "-schema/bin/"
			dstfile = "exp-sql.bat"
		case newProj + ".schema.bin.menu.sh":
			dir = me.NewProjPath + "/" + me.NewProjName + "-schema/bin/"
			dstfile = "exp-menu.sh"
		case newProj + ".schema.bin.menu.bat":
			dir = me.NewProjPath + "/" + me.NewProjName + "-schema/bin/"
			dstfile = "exp-menu.bat"
		case newProj + ".schema.bin.entity.sh":
			dir = me.NewProjPath + "/" + me.NewProjName + "-schema/bin/"
			dstfile = "exp-entity.sh"
		case newProj + ".schema.bin.entity.bat":
			dir = me.NewProjPath + "/" + me.NewProjName + "-schema/bin/"
			dstfile = "exp-entity.bat"
		case newProj + ".schema.conf.db":
			dir = me.NewProjPath + "/" + me.NewProjName + "-schema/conf/env/"
			dstfile = "db.xml"
		case newProj + ".schema.conf.project":
			dir = me.NewProjPath + "/" + me.NewProjName + "-schema/conf/schema/"
			dstfile = "projects.xml"
		case newProj + ".schema.sql.dml.merge.sh":
			dir = me.NewProjPath + "/" + me.NewProjName + "-schema/sql/dml/"
			dstfile = "merge-file.sh"
		case newProj + ".schema.sql.dml.merge.bat":
			dir = me.NewProjPath + "/" + me.NewProjName + "-schema/sql/dml/"
			dstfile = "merge-file.bat"
		case newProj + ".web":
			dir = me.NewProjPath + "/" + me.NewProjName + "-web/"
			dstfile = "web.go"
		case newProj + ".web.bin.restart":
			dir = me.NewProjPath + "/" + me.NewProjName + "-web/bin/"
			dstfile = "restart.sh"
		case newProj + ".web.bin.startup":
			dir = me.NewProjPath + "/" + me.NewProjName + "-web/bin/"
			dstfile = "startup.sh"
		case newProj + ".web.bin.shutdown":
			dir = me.NewProjPath + "/" + me.NewProjName + "-web/bin/"
			dstfile = "shutdown.sh"
		case newProj + ".web.conf.api":
			dir = me.NewProjPath + "/" + me.NewProjName + "-web/conf/env/"
			dstfile = "api.xml"
		case newProj + ".web.conf.db":
			dir = me.NewProjPath + "/" + me.NewProjName + "-web/conf/env/"
			dstfile = "db.xml"
		case newProj + ".web.conf.export":
			dir = me.NewProjPath + "/" + me.NewProjName + "-web/conf/env/"
			dstfile = "export.xml"
		case newProj + ".web.conf.log":
			dir = me.NewProjPath + "/" + me.NewProjName + "-web/conf/env/"
			dstfile = "log.xml"
		case newProj + ".web.conf.secure":
			dir = me.NewProjPath + "/" + me.NewProjName + "-web/conf/env/"
			dstfile = "secure.xml"
		case newProj + ".web.conf.sensitive":
			dir = me.NewProjPath + "/" + me.NewProjName + "-web/conf/env/"
			dstfile = "sensitive.xml"
		case newProj + ".web.conf.session":
			dir = me.NewProjPath + "/" + me.NewProjName + "-web/conf/env/"
			dstfile = "session.xml"
		case newProj + ".web.conf.settings":
			dir = me.NewProjPath + "/" + me.NewProjName + "-web/conf/env/"
			dstfile = "settings.xml"
		case newProj + ".web.conf.static":
			dir = me.NewProjPath + "/" + me.NewProjName + "-web/conf/env/"
			dstfile = "static.xml"
		case newProj + ".web.conf.template":
			dir = me.NewProjPath + "/" + me.NewProjName + "-web/conf/env/"
			dstfile = "template.xml"
		case newProj + ".web.conf.upload":
			dir = me.NewProjPath + "/" + me.NewProjName + "-web/conf/env/"
			dstfile = "upload.xml"
		case newProj + ".web.internal.doc":
			dir = me.NewProjPath + "/" + me.NewProjName + "-web/internal/"
			dstfile = "doc.go"
		case newProj + ".web.static.css.comm":
			dir = me.NewProjPath + "/" + me.NewProjName + "-web/static/css/comm/"
			dstfile = "comm.css"
		case newProj + ".web.static.css.core":
			dir = me.NewProjPath + "/" + me.NewProjName + "-web/static/css/core/"
			dstfile = "global.css"
		case newProj + ".web.static.img.comm":
			dir = me.NewProjPath + "/" + me.NewProjName + "-web/static/images/comm/"
			dstfile = "README.md"
		case newProj + ".web.static.img.core":
			dir = me.NewProjPath + "/" + me.NewProjName + "-web/static/images/core/"
			dstfile = "README.md"
		case newProj + ".web.static.js.comm":
			dir = me.NewProjPath + "/" + me.NewProjName + "-web/static/js/comm/"
			dstfile = "README.md"
		case newProj + ".web.static.js.core":
			dir = me.NewProjPath + "/" + me.NewProjName + "-web/static/js/core/"
			dstfile = "README.md"
		case newProj + ".web.templates.comm":
			dir = me.NewProjPath + "/" + me.NewProjName + "-web/templates/comm/"
			dstfile = "README.md"
		case newProj + ".web.templates.core":
			dir = me.NewProjPath + "/" + me.NewProjName + "-web/templates/core/"
			dstfile = "README.md"
		case newProj + ".web.templates.home":
			dir = me.NewProjPath + "/" + me.NewProjName + "-web/templates/"
			dstfile = "home.html"
		case newProj + ".web.templates.login":
			dir = me.NewProjPath + "/" + me.NewProjName + "-web/templates/"
			dstfile = "login.html"
		case newProj + ".web.templates.title":
			dir = me.NewProjPath + "/" + me.NewProjName + "-web/templates/"
			dstfile = "title.html"
		case newProj + ".web.templates.ver":
			dir = me.NewProjPath + "/" + me.NewProjName + "-web/templates/"
			dstfile = "version.html"
		case newProj + ".adm":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/"
			dstfile = "adm.go"
		case newProj + ".adm.bin.restart":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/bin/"
			dstfile = "restart.sh"
		case newProj + ".adm.bin.startup":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/bin/"
			dstfile = "startup.sh"
		case newProj + ".adm.bin.shutdown":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/bin/"
			dstfile = "shutdown.sh"
		case newProj + ".adm.conf.api":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/conf/env/"
			dstfile = "api.xml"
		case newProj + ".adm.conf.db":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/conf/env/"
			dstfile = "db.xml"
		case newProj + ".adm.conf.export":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/conf/env/"
			dstfile = "export.xml"
		case newProj + ".adm.conf.log":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/conf/env/"
			dstfile = "log.xml"
		case newProj + ".adm.conf.secure":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/conf/env/"
			dstfile = "secure.xml"
		case newProj + ".adm.conf.sensitive":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/conf/env/"
			dstfile = "sensitive.xml"
		case newProj + ".adm.conf.session":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/conf/env/"
			dstfile = "session.xml"
		case newProj + ".adm.conf.settings":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/conf/env/"
			dstfile = "settings.xml"
		case newProj + ".adm.conf.static":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/conf/env/"
			dstfile = "static.xml"
		case newProj + ".adm.conf.template":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/conf/env/"
			dstfile = "template.xml"
		case newProj + ".adm.conf.upload":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/conf/env/"
			dstfile = "upload.xml"
		case newProj + ".adm.internal.doc":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/internal/"
			dstfile = "doc.go"
		case newProj + ".adm.static.css":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/static/css/"
			dstfile = "README.md"
		case newProj + ".adm.static.img":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/static/images/"
			dstfile = "README.md"
		case newProj + ".adm.static.js.sys.area":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/static/sys/js/"
			dstfile = "area.js"
		case newProj + ".adm.static.js.sys.blacklist":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/static/sys/js/"
			dstfile = "blacklist.js"
		case newProj + ".adm.static.js.sys.cache":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/static/sys/js/"
			dstfile = "cache.js"
		case newProj + ".adm.static.js.sys.conf":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/static/sys/js/"
			dstfile = "conf.js"
		case newProj + ".adm.static.js.sys.dict":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/static/sys/js/"
			dstfile = "dict.js"
		case newProj + ".adm.static.js.sys.menu":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/static/sys/js/"
			dstfile = "menu.js"
		case newProj + ".adm.static.js.sys.org":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/static/sys/js/"
			dstfile = "org.js"
		case newProj + ".adm.static.js.sys.post":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/static/sys/js/"
			dstfile = "post.js"
		case newProj + ".adm.static.js.sys.role":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/static/sys/js/"
			dstfile = "role.js"
		case newProj + ".adm.static.js.sys.user":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/static/sys/js/"
			dstfile = "user.js"
		case newProj + ".adm.static.js.home":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/static/js/"
			dstfile = "home.js"
		case newProj + ".adm.static.js.login":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/static/js/"
			dstfile = "login.js"
		case newProj + ".adm.static.lib":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/static/dev"
			zipfile := me.uizip()
			if strings.IsNotBlank(zipfile) {
				if err := files.Unzip(zipfile, dir); err != nil {
					return err
				}
			}
			return nil
		case newProj + ".adm.templates.home":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/templates/"
			dstfile = "home.html"
		case newProj + ".adm.templates.login":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/templates/"
			dstfile = "login.html"
		case newProj + ".adm.templates.title":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/templates/"
			dstfile = "title.html"
		case newProj + ".adm.templates.ver":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/templates/"
			dstfile = "version.html"
		case newProj + ".adm.templates.core.comm.action":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/templates/core/comm/"
			dstfile = "action.html"
		case newProj + ".adm.templates.core.comm.alert":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/templates/core/comm/"
			dstfile = "alert.html"
		case newProj + ".adm.templates.core.comm.breadcrumb":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/templates/core/comm/"
			dstfile = "breadcrumb.html"
		case newProj + ".adm.templates.core.comm.ckeditor":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/templates/core/comm/"
			dstfile = "ckeditor.html"
		case newProj + ".adm.templates.core.comm.dialog":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/templates/core/comm/"
			dstfile = "dialog.html"
		case newProj + ".adm.templates.core.comm.disable":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/templates/core/comm/"
			dstfile = "disable.html"
		case newProj + ".adm.templates.core.comm.formtree":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/templates/core/comm/"
			dstfile = "formtree.html"
		case newProj + ".adm.templates.core.comm.navtabs":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/templates/core/comm/"
			dstfile = "navtabs.html"
		case newProj + ".adm.templates.core.comm.navtabs.list":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/templates/core/comm/"
			dstfile = "navtabs.list.html"
		case newProj + ".adm.templates.core.comm.page":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/templates/core/comm/"
			dstfile = "page.html"
		case newProj + ".adm.templates.core.include.footer":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/templates/core/include/"
			dstfile = "footer.html"
		case newProj + ".adm.templates.core.include.head":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/templates/core/include/"
			dstfile = "head.html"
		case newProj + ".adm.templates.err.401":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/templates/err/"
			dstfile = "401.html"
		case newProj + ".adm.templates.err.403":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/templates/err/"
			dstfile = "403.html"
		case newProj + ".adm.templates.err.404":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/templates/err/"
			dstfile = "404.html"
		case newProj + ".adm.templates.err.500":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/templates/err/"
			dstfile = "500.html"
		case newProj + ".adm.templates.sys.area":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/templates/sys/area/"
			dstfile = "area.html"
		case newProj + ".adm.templates.sys.blacklist":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/templates/sys/blacklist/"
			dstfile = "blacklist.html"
		case newProj + ".adm.templates.sys.cache":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/templates/sys/cache/"
			dstfile = "cache.html"
		case newProj + ".adm.templates.sys.conf":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/templates/sys/conf/"
			dstfile = "conf.html"
		case newProj + ".adm.templates.sys.dict":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/templates/sys/dict/"
			dstfile = "dict.html"
		case newProj + ".adm.templates.sys.menu":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/templates/sys/menu/"
			dstfile = "menu.html"
		case newProj + ".adm.templates.sys.org":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/templates/sys/org/"
			dstfile = "org.html"
		case newProj + ".adm.templates.sys.post":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/templates/sys/post/"
			dstfile = "post.html"
		case newProj + ".adm.templates.sys.role":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/templates/sys/role/"
			dstfile = "role.html"
		case newProj + ".adm.templates.sys.user":
			dir = me.NewProjPath + "/" + me.NewProjName + "-adm/templates/sys/user/"
			dstfile = "user.html"
		}
		dstfile = filepath.Join(dir, dstfile)
	}
	if files.IsExist(dstfile) {
		if strings.HasPrefix(typ, typXgen) {
			files.Remove(dstfile)
		} else {
			return nil
		}
	} else {
		files.MkdirAll(dir, 0755)
	}

	buf := bytes.Buffer{}
	tmpl := newTmpl(content)
	tmpl.Execute(&buf, me)
	err := ioutil.WriteFile(dstfile, buf.Bytes(), 0755)
	if err != nil {
		return errors.New("typ=" + typ + " err=" + err.Error())
	}
	return nil
}

func (me factory) writeNewProj() error {
	if err := me.writeBy(newProj+".README", tmplNewProjReadme); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".tst", tmplNewProjTst); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".tst.settings", tmplNewProjTstSettings); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".tst.session", tmplNewProjTstSession); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".tst.db", tmplNewProjTstDB); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".schema.bin.db.sh", tmplNewProjSchemaBinDbSh); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".schema.bin.db.bat", tmplNewProjSchemaBinDbBat); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".schema.bin.sql.sh", tmplNewProjSchemaBinSqlSh); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".schema.bin.sql.bat", tmplNewProjSchemaBinSqlBat); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".schema.bin.menu.sh", tmplNewProjSchemaBinMenuSh); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".schema.bin.menu.bat", tmplNewProjSchemaBinMenuBat); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".schema.bin.entity.sh", tmplNewProjSchemaBinEntitySh); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".schema.bin.entity.bat", tmplNewProjSchemaBinEntityBat); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".schema.conf.db", tmplNewProjTstDB); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".schema.conf.project", tmplNewProjSchemaConfProject); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".schema.sql.dml.merge.sh", tmplNewProjSchemaSQLMergeSh); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".schema.sql.dml.merge.bat", tmplNewProjSchemaSQLMergeBat); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".web", tmplNewProjWeb); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".web.bin.restart", tmplNewProjWebBinRestart); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".web.bin.startup", tmplNewProjWebBinStartup); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".web.bin.shutdown", tmplNewProjWebBinShutdown); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".web.conf.api", tmplNewProjWebConfAPI); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".web.conf.db", tmplNewProjTstDB); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".web.conf.export", tmplNewProjWebConfExport); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".web.conf.log", tmplNewProjWebConfLog); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".web.conf.secure", tmplNewProjWebConfSecure); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".web.conf.sensitive", tmplNewProjWebConfSensitive); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".web.conf.session", tmplNewProjWebConfSession); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".web.conf.settings", tmplNewProjWebConfSettings); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".web.conf.static", tmplNewProjWebConfStatic); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".web.conf.template", tmplNewProjWebConfTemplate); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".web.conf.upload", tmplNewProjWebConfUpload); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".web.internal.doc", tmplNewProjWebInternalDoc); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".web.static.css.comm", tmplNewProjWebStaticCssComm); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".web.static.css.core", tmplNewProjWebStaticCssCore); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".web.static.img.comm", tmplNewProjWebStaticImgComm); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".web.static.img.core", tmplNewProjWebStaticImgCore); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".web.static.js.comm", tmplNewProjWebStaticJSComm); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".web.static.js.core", tmplNewProjWebStaticJSCore); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".web.templates.comm", tmplNewProjWebTemplatesComm); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".web.templates.core", tmplNewProjWebTemplatesCore); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".web.templates.home", tmplNewProjWebTemplatesHome); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".web.templates.login", tmplNewProjWebTemplatesLogin); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".web.templates.title", tmplNewProjWebTemplatesTitle); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".web.templates.ver", tmplNewProjWebTemplatesVer); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm", tmplNewProjAdm); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.bin.restart", tmplNewProjAdmBinRestart); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.bin.startup", tmplNewProjAdmBinStartup); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.bin.shutdown", tmplNewProjAdmBinShutdown); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.conf.api", tmplNewProjAdmConfAPI); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.conf.db", tmplNewProjTstDB); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.conf.export", tmplNewProjAdmConfExport); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.conf.log", tmplNewProjAdmConfLog); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.conf.secure", tmplNewProjAdmConfSecure); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.conf.sensitive", tmplNewProjAdmConfSensitive); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.conf.session", tmplNewProjAdmConfSession); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.conf.settings", tmplNewProjAdmConfSettings); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.conf.static", tmplNewProjAdmConfStatic); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.conf.template", tmplNewProjAdmConfTemplate); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.conf.upload", tmplNewProjAdmConfUpload); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.internal.doc", tmplNewProjAdmInternalDoc); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.static.css", tmplNewProjAdmStaticCss); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.static.img", tmplNewProjAdmStaticImg); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.static.js.sys.area", tmplNewProjAdmStaticJSSysArea); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.static.js.sys.blacklist", tmplNewProjAdmStaticJSSysBlacklist); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.static.js.sys.cache", tmplNewProjAdmStaticJSSysCache); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.static.js.sys.conf", tmplNewProjAdmStaticJSSysConf); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.static.js.sys.dict", tmplNewProjAdmStaticJSSysDict); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.static.js.sys.menu", tmplNewProjAdmStaticJSSysMenu); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.static.js.sys.org", tmplNewProjAdmStaticJSSysOrg); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.static.js.sys.post", tmplNewProjAdmStaticJSSysPost); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.static.js.sys.role", tmplNewProjAdmStaticJSSysRole); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.static.js.sys.user", tmplNewProjAdmStaticJSSysUser); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.static.js.home", tmplNewProjAdmStaticJSHome); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.static.js.login", tmplNewProjAdmStaticJSLogin); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.static.lib", ""); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.templates.home", tmplNewProjAdmTemplatesHome); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.templates.login", tmplNewProjAdmTemplatesLogin); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.templates.title", tmplNewProjAdmTemplatesTitle); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.templates.ver", tmplNewProjAdmTemplatesVer); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.templates.core.comm.action", tmplNewProjAdmTemplatesCoreCommAction); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.templates.core.comm.alert", tmplNewProjAdmTemplatesCoreCommAlert); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.templates.core.comm.breadcrumb", tmplNewProjAdmTemplatesCoreCommBreadcrumb); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.templates.core.comm.ckeditor", tmplNewProjAdmTemplatesCoreCommCkeditor); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.templates.core.comm.dialog", tmplNewProjAdmTemplatesCoreCommDialog); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.templates.core.comm.disable", tmplNewProjAdmTemplatesCoreCommDisable); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.templates.core.comm.formtree", tmplNewProjAdmTemplatesCoreCommFormtree); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.templates.core.comm.navtabs", tmplNewProjAdmTemplatesCoreCommNavtabs); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.templates.core.comm.navtabs.list", tmplNewProjAdmTemplatesCoreCommNavtabsList); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.templates.core.comm.page", tmplNewProjAdmTemplatesCoreCommPage); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.templates.core.include.footer", tmplNewProjAdmTemplatesCoreIncludeFooter); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.templates.core.include.head", tmplNewProjAdmTemplatesCoreIncludeHead); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.templates.err.401", tmplNewProjAdmTemplatesErr401); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.templates.err.403", tmplNewProjAdmTemplatesErr403); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.templates.err.404", tmplNewProjAdmTemplatesErr404); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.templates.err.500", tmplNewProjAdmTemplatesErr500); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.templates.sys.area", tmplNewProjAdmTemplatesSysArea); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.templates.sys.blacklist", tmplNewProjAdmTemplatesSysBlacklist); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.templates.sys.cache", tmplNewProjAdmTemplatesSysCache); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.templates.sys.conf", tmplNewProjAdmTemplatesSysConf); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.templates.sys.dict", tmplNewProjAdmTemplatesSysDict); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.templates.sys.menu", tmplNewProjAdmTemplatesSysMenu); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.templates.sys.org", tmplNewProjAdmTemplatesSysOrg); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.templates.sys.post", tmplNewProjAdmTemplatesSysPost); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.templates.sys.role", tmplNewProjAdmTemplatesSysRole); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".adm.templates.sys.user", tmplNewProjAdmTemplatesSysUser); err != nil {
		return err
	}
	return nil
}

func (me factory) writeEntityXgen() error {
	return me.writeBy(xgenEntity, tmplEntity)
}

func (me factory) writeEntitiesXgen() error {
	return me.writeBy(xgenEntities, tmplEntities)
}

func (me factory) writeServiceXgen() error {
	return me.writeBy(xgenService, tmplServiceXgen)
}

func (me factory) writeServiceMain() error {
	return me.writeBy(mainService, tmplServiceMain)
}

func (me factory) writeServiceTest() error {
	return me.writeBy(testService, tmplServiceTest)
}

func (me factory) writeControllerXgen() error {
	return me.writeBy(xgenCtl, tmplControllerXgen)
}

func (me factory) writeControllerMain() error {
	return me.writeBy(mainCtl, tmplControllerMain)
}

func (me factory) writeControllerTest() error {
	return me.writeBy(testCtl, tmplControllerTest)
}

func (me factory) writeControllerReg() error {
	return me.writeBy(xgenCtlReg, tmplControllerReg)
}

func (me factory) writeSQLMain() error {
	return me.writeBy(mainSQL, tmplSQLMain)
}

func (me factory) writeLogJSONXgen() error {
	return me.writeBy(xgenLogJSON, tmplLogXgen)
}

func (me factory) writeLogAPIXgen() error {
	return me.writeBy(xgenLogAPI, tmplLogXgen)
}

func (me factory) writeUtilMain() error {
	return me.writeBy(mainUtil, tmplUtilMain)
}

func (me factory) writeConstMain() error {
	return me.writeBy(mainConst, tmplConstMain)
}

func (me factory) writeHTMLMain() error {
	return me.writeBy(mainHTML, tmplHTMLMain)
}

func (me factory) writeJsMain() error {
	return me.writeBy(mainJs, tmplJsMain)
}

func (me factory) writeDtoXgen() error {
	if strings.IsBlank(me.AdmPath) {
		return nil
	}
	return me.writeBy(xgenDto, tmplDtoXgen)
}

func (me factory) writeAPIMain() error {
	return me.writeBy(mainAPI, tmplAPIMain)
}

func (me factory) genFileName(typ, name string) (string, string) {
	switch typ {
	case xgenLogJSON, xgenLogAPI:
		return "log_xgen.go", name
	case mainAPI:
		return me.PackageName + ".go", name
	case mainUtil:
		return me.PackageName + "_util.go", name
	case mainConst:
		return me.PackageName + "_const.go", name
	case xgenCtlReg:
		return me.PackageName + "_register_xgen.go", name
	}
	if strings.HasPrefix(name, typMain) {
		name = strings.After(name, typMain)
	}
	if strings.HasSuffix(name, "_entity") {
		name = strings.Before(name, "_entity")
	}
	if name == "domain" || name == "entity" || name == "main.domain" || name == "main.entity" {
		name = ""
	} else {
		if strings.IsNotBlank(name) && typ != mainHTML && typ != mainJs {
			name = name + "_"
		}
	}
	switch typ {
	case mainCtl:
		typ = "controller"
	case mainService:
		typ = "manager"
	case mainSQL:
		typ = "sql"
	case xgenCtl:
		typ, name = me.resetTypAndName("controller_xgen", name)
	case xgenService:
		typ, name = me.resetTypAndName("manager_xgen", name)
	case xgenEntity:
		typ, name = me.resetTypAndName("entity_xgen", name)
	case xgenEntities:
		typ, name = me.resetTypAndName("entities_xgen", name)
	case xgenDto:
		typ, name = me.resetTypAndName("dto_xgen", name)
	}
	switch typ {
	case mainHTML:
		return fmt.Sprintf("%s.html", name), name
	case mainJs:
		return fmt.Sprintf("%s.js", name), name
	case testCtl:
		return fmt.Sprintf("%scontroller_test.go", name), name
	case testService:
		return fmt.Sprintf("%smanager_test.go", name), name
	default:
		return fmt.Sprintf("%s%s.go", name, typ), name
	}
}

func (me factory) resetTypAndName(typ, name string) (string, string) {
	if strings.HasSuffix(name, "_test_") {
		name = strings.Before(name, "test_")
		typ = typ + "_test"
	}
	return typ, name
}

func (me *factory) isTimeField() {
	for _, e := range me.Entities {
		for _, f := range e.Fields {
			if f.Type == "time.Time" {
				me.IsTimeField = true
			}
		}
	}
}

func (me *factory) isValidationField() {
	for _, e := range me.Entities {
		for _, f := range e.Fields {
			if len(f.Validations) > 0 {
				me.IsValidationField = true
			}
		}
	}
}

func (me *factory) isExtend() {
	for _, e := range me.Entities {
		if strings.IsNotBlank(e.Extend) {
			me.IsExtend = true
		}
	}
}

func (me factory) printerType(e ast.Expr) string {
	var b bytes.Buffer
	printer.Fprint(&b, token.NewFileSet(), e)
	switch b.String() {
	case "schema.Table":
		return "table"
	case "entity.String":
		return "string"
	case "entity.Bool":
		return "bool"
	case "entity.Float32":
		return "float32"
	case "entity.Float64":
		return "float64"
	case "entity.Int":
		return "int"
	case "entity.Int8":
		return "int8"
	case "entity.Int16":
		return "int16"
	case "entity.Int32":
		return "int32"
	case "entity.Int64":
		return "int64"
	case "entity.Time":
		return "time.Time"
	case "entity.Uint":
		return "uint"
	case "entity.Uint8":
		return "uint8"
	case "entity.Uint16":
		return "uint16"
	case "entity.Uint32":
		return "uint32"
	case "entity.Uint64":
		return "uint64"
	case "entity.Pk":
		return "pk"
	case "entity.Sys":
		return "sys"
	case "entity.Tree":
		return "tree"
	default:
		return "err"
	}
}

func (me *factory) uizip() string {
	path := "%GOPATH%/src/gopkg.in/goyy/goyy.v0/tool/xgen/assets/ui.zip"
	zippath := envs.ParseGOPATH(path)
	if files.IsExist(zippath) {
		return zippath
	}
	return ""
}
