// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"gopkg.in/goyy/goyy.v0/util/files"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"io/ioutil"
	"path/filepath"
)

// factory is a file generation factory.
type factory struct {
	Project           string
	PackageName       string
	Epath             string
	Htmpath           string
	Clidir            string
	Clipath           string
	Apipath           string
	HasGenService     bool
	HasGenController  bool
	HasGenDto         bool
	HasGenApi         bool
	HasGenSql         bool
	HasGenLog         bool
	HasGenUtil        bool
	HasGenConst       bool
	IsTimeField       bool
	IsValidationField bool
	IsExtend          bool
	Entities          []entity
}

// Init initializes an File from a path.
func (me *factory) Init(path string) error {
	// set the path
	if strings.HasSuffix(path, ".go") {
		me.Epath = strings.TrimSuffix(path, ".go")
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
		fmt.Errorf("Unable to parse '%s': %s", path, err)
	}

	// get package name
	if f.Name != nil {
		me.PackageName = f.Name.Name
	} else {
		fmt.Errorf("Missing package name in '%s'", path)
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

		e := entity{Project: project, Relationship: relationship}
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
					col := field{Name: "id", Type: "string", Column: "id", IsPrimary: true}
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
				items = tagItemValue(f.Tag.Value, "orm")
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

				col := field{}
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
					if err := col.InitValidation(items); err != nil {
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

func (me factory) Write() error {
	if err := me.writeEntityXgen(); err != nil {
		return err
	}
	if err := me.writeEntitiesXgen(); err != nil {
		return err
	}
	if me.HasGenService {
		if err := me.writeServiceXgen(); err != nil {
			return err
		}
		if err := me.writeServiceMain(); err != nil {
			return err
		}
	}
	if me.HasGenController {
		if strings.IsBlank(me.Clidir) {
			if strings.IsBlank(me.Apipath) {
				if err := me.writeControllerHtmlXgen(); err != nil {
					return err
				}
				if err := me.writeControllerHtmlMain(); err != nil {
					return err
				}
				if err := me.writeControllerHtmlReg(); err != nil {
					return err
				}
			} else {
				if err := me.writeControllerJsonXgen(); err != nil {
					return err
				}
				if err := me.writeControllerJsonMain(); err != nil {
					return err
				}
				if err := me.writeControllerJsonReg(); err != nil {
					return err
				}
			}
		} else {
			if err := me.writeControllerJsonXgen(); err != nil {
				return err
			}
			if err := me.writeControllerJsonMain(); err != nil {
				return err
			}
			if err := me.writeControllerJsonReg(); err != nil {
				return err
			}
			if err := me.writeControllerClientXgen(); err != nil {
				return err
			}
			if err := me.writeControllerClientMain(); err != nil {
				return err
			}
			if err := me.writeControllerClientReg(); err != nil {
				return err
			}
		}
	}
	if me.HasGenDto {
		if err := me.writeDtoXgen(); err != nil {
			return err
		}
	}
	if me.HasGenApi {
		if strings.IsNotBlank(me.Apipath) {
			if err := me.writeApiMain(); err != nil {
				return err
			}
		}
	}
	if me.HasGenSql {
		if err := me.writeSqlMain(); err != nil {
			return err
		}
	}
	if me.HasGenLog {
		if err := me.writeLogJsonXgen(); err != nil {
			return err
		}
		if strings.IsNotBlank(me.Clidir) {
			if err := me.writeLogClientXgen(); err != nil {
				return err
			}
		}
		if strings.IsNotBlank(me.Apipath) {
			if err := me.writeLogApiXgen(); err != nil {
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
	return nil
}

func (me factory) writeBy(typ, content string) error {
	// get the destination file
	dir, file := filepath.Split(me.Epath)
	if typ == "xgen.dto" || typ == "xgen.controller.client" || typ == "main.controller.client" || typ == "xgen.log.client" {
		dir = me.Clidir + "/internal/" + me.Project + "/" + me.PackageName
	}
	if typ == "main.api" || typ == "xgen.log.api" {
		dir = "../../api/" + me.PackageName
	}
	if typ == "reg.controller.client" {
		dir = me.Clidir + "/internal/" + me.Project
	}
	if typ == "reg.proj.controller.client" {
		dir = me.Clidir + "/internal"
	}
	if typ == "reg.controller.json" {
		dir = "../../"
	}
	if typ == "reg.controller.html" {
		dir = ".."
	}
	dstfile := filepath.Join(dir, me.genFileName(typ, file))
	if files.IsExist(dstfile) {
		if strings.HasPrefix(typ, typMain) {
			return nil
		} else {
			files.Remove(dstfile)
		}
	} else {
		files.MkdirAll(dir, 0644)
	}

	buf := bytes.Buffer{}
	tmpl := newTmpl(content)
	tmpl.Execute(&buf, me)
	return ioutil.WriteFile(dstfile, buf.Bytes(), 0644)
}

func (me factory) writeEntityXgen() error {
	return me.writeBy("xgen.entity", tmplEntity)
}

func (me factory) writeEntitiesXgen() error {
	return me.writeBy("xgen.entities", tmplEntities)
}

func (me factory) writeServiceXgen() error {
	return me.writeBy("xgen.service", tmplServiceXgen)
}

func (me factory) writeServiceMain() error {
	return me.writeBy("main.service", tmplServiceMain)
}

func (me factory) writeControllerHtmlXgen() error {
	return me.writeBy("xgen.controller.html", tmplControllerHtmlXgen)
}

func (me factory) writeControllerHtmlMain() error {
	return me.writeBy("main.controller.html", tmplControllerHtmlMain)
}

func (me factory) writeControllerHtmlReg() error {
	if strings.IsBlank(me.Htmpath) {
		return nil
	}
	return me.writeBy("reg.controller.html", tmplControllerHtmlReg)
}

func (me factory) writeControllerJsonXgen() error {
	return me.writeBy("xgen.controller.json", tmplControllerJsonXgen)
}

func (me factory) writeControllerJsonMain() error {
	return me.writeBy("main.controller.json", tmplControllerJsonMain)
}

func (me factory) writeControllerJsonReg() error {
	return me.writeBy("reg.controller.json", tmplControllerJsonReg)
}

func (me factory) writeControllerClientXgen() error {
	return me.writeBy("xgen.controller.client", tmplControllerClientXgen)
}

func (me factory) writeControllerClientMain() error {
	return me.writeBy("main.controller.client", tmplControllerClientMain)
}

func (me factory) writeControllerClientReg() error {
	err := me.writeBy("reg.controller.client", tmplControllerClientReg)
	if err != nil {
		return err
	}
	return me.writeBy("reg.proj.controller.client", tmplControllerClientRegProj)
}

func (me factory) writeSqlMain() error {
	return me.writeBy("main.sql", tmplSqlMain)
}

func (me factory) writeLogJsonXgen() error {
	return me.writeBy("xgen.log.json", tmplLogXgen)
}

func (me factory) writeLogApiXgen() error {
	return me.writeBy("xgen.log.api", tmplLogXgen)
}

func (me factory) writeLogClientXgen() error {
	return me.writeBy("xgen.log.client", tmplLogXgen)
}

func (me factory) writeUtilMain() error {
	return me.writeBy("main.util", tmplUtilMain)
}

func (me factory) writeConstMain() error {
	return me.writeBy("main.const", tmplConstMain)
}

func (me factory) writeDtoXgen() error {
	if strings.IsBlank(me.Clidir) {
		return nil
	}
	return me.writeBy("xgen.dto", tmplDtoXgen)
}

func (me factory) writeApiMain() error {
	return me.writeBy("main.api", tmplApiMain)
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

func (me factory) genFileName(typ, name string) string {
	if typ == "xgen.log.json" || typ == "xgen.log.client" || typ == "xgen.log.api" {
		return "xgen.log.go"
	}
	if typ == "main.util" {
		return "main.util.go"
	}
	if typ == "main.const" {
		return "main.const.go"
	}
	if typ == "main.api" {
		return "main." + me.PackageName + ".go"
	}
	if typ == "reg.controller.html" || typ == "reg.controller.client" || typ == "reg.controller.json" {
		return "xgen.register." + me.PackageName + ".go"
	}
	if typ == "reg.proj.controller.client" {
		return "xgen.register." + me.Project + ".go"
	}
	if strings.HasPrefix(name, typMain) {
		name = strings.After(name, typMain)
	}
	if strings.HasPrefix(name, "domain.") {
		name = strings.After(name, "domain.")
	}
	if strings.HasPrefix(name, "entity.") {
		name = strings.After(name, "entity.")
	}
	if name == "domain" || name == "entity" || name == "main.domain" || name == "main.entity" {
		name = ""
	} else {
		if strings.IsNotBlank(name) {
			name = "." + name
		}
	}
	if strings.HasPrefix(typ, "main.controller") {
		typ = "main.controller"
	}
	if strings.HasPrefix(typ, "xgen.controller") {
		typ = "xgen.controller"
	}
	return fmt.Sprintf("%s%s.go", typ, name)
}
