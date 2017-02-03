// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"log"

	"gopkg.in/goyy/goyy.v0/comm/env"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

var conf = &configuration{}
var util = &utils{}

func initgen(name string) {
	if strings.IsBlank(name) {
		name = "goyy"
	}
	valid := &valids{}
	valid.IsExistXML(name)
	valid.IsOkXML()

	data := &inits{}
	data.Init()
}

type inits struct{}

func (me *inits) Init() {
	me.Settings()
	me.Projects()
	me.Modules()
	me.Buttons()
	me.Domains()
	me.Columns()
	me.Tables()
	me.ProjectTables()
}

func (me *inits) Settings() {
	xconf := util.DecodeXML(xsettings)
	s := &settings{Statement: &statement{}}
	s.Statement.Seperator = xconf.Settings.Statement.Seperator
	s.Statement.Case = xconf.Settings.Statement.Case
	s.Statement.Comment = xconf.Settings.Statement.Comment
	conf.Settings = s
}

func (me *inits) Projects() {
	xconf := util.DecodeXML(xprojects)
	for _, xp := range xconf.Projects.Project {
		db, err := env.ParseDatabase(xp.Database)
		if err != nil {
			log.Fatal(err)
		}
		d := &database{xp.Database, db.DriverName, db.DataSourceName}
		p := &project{
			id:       xp.ID,
			name:     xp.Name,
			database: d,
			generate: xp.Generate,
			menu:     xp.Menu,
			comment:  xp.Comment,
			clipath:  xp.Clipath,
			apipath:  xp.Apipath,
			tstpath:  xp.Tstpath,
		}
		conf.projects = append(conf.projects, p)
	}
}

func (me *inits) Modules() {
	xconf := util.DecodeXML(xmodules)
	for _, xm := range xconf.Modules.Module {
		var p *project
		for _, cp := range conf.projects {
			if cp.ID() == xm.Project {
				p = cp
			}
		}
		m := &module{
			id:       xm.ID,
			name:     xm.Name,
			prefix:   xm.Prefix,
			project:  p,
			generate: xm.Generate,
			menu:     xm.Menu,
			comment:  xm.Comment,
			clipath:  xm.Clipath,
			apipath:  xm.Apipath,
			tstpath:  xm.Tstpath,
		}
		conf.modules = append(conf.modules, m)
	}
}

func (me *inits) Buttons() {
	xconf := util.DecodeXML(xbuttons)
	for _, xd := range xconf.Buttons.Button {
		d := &button{
			id:      xd.ID,
			name:    xd.Name,
			comment: xd.Comment,
		}
		conf.buttons = append(conf.buttons, d)
	}
}

func (me *inits) Domains() {
	xconf := util.DecodeXML(xdomains)
	for _, xd := range xconf.Domains.Domain {
		d := &domain{
			id:        xd.ID,
			name:      xd.Name,
			types:     xd.Types,
			length:    xd.Length,
			precision: xd.Precision,
			comment:   xd.Comment,
			defaults:  xd.Defaults,
			nullable:  xd.Nullable,
			etype:     util.Etype(xd.Types),
		}
		conf.domains = append(conf.domains, d)
	}
}

func (me *inits) Columns() {
	xconf := util.DecodeXML(xcolumns)
	for _, xc := range xconf.Columns.Column {
		var d *domain
		for _, cd := range conf.domains {
			if cd.ID() == xc.Domain {
				d = cd
			}
		}
		c := &column{
			id:       xc.ID,
			name:     xc.Name,
			domain:   d,
			index:    xc.Index,
			comment:  xc.Comment,
			defaults: xc.Defaults,
			nullable: xc.Nullable,
		}
		c.field = strings.ToLowerFirst(strings.Camel(c.ID()))
		conf.columns = append(conf.columns, c)
	}
}

func (me *inits) Tables() {
	xconf := util.DecodeXML(xtables)
	m := &module{project: &project{}}
	for _, xt := range xconf.Tables.Table {
		if strings.TrimSpace(xt.Extends) == "" {
			t := &table{
				module:  m,
				id:      xt.ID,
				name:    xt.Name,
				comment: xt.Comment,
				master:  xt.Master,
				slave:   xt.Slave,
			}
			for _, xc := range xt.Columns {
				var ec *column
				if strings.TrimSpace(xc.Extends) != "" {
					for _, cc := range conf.columns {
						if cc.ID() == xc.Extends {
							ec = cc
						}
					}
				}
				var d *domain
				if strings.TrimSpace(xc.Domain) != "" {
					for _, cd := range conf.domains {
						if cd.ID() == xc.Domain {
							d = cd
						}
					}
				}
				c := &column{
					parent:   ec,
					id:       xc.ID,
					name:     xc.Name,
					domain:   d,
					index:    xc.Index,
					comment:  xc.Comment,
					defaults: xc.Defaults,
					nullable: xc.Nullable,
				}
				c.field = strings.ToLowerFirst(strings.Camel(c.ID()))
				t.columns = append(t.columns, c)
			}
			conf.parent = append(conf.parent, t)
			me.ChildTables(xconf, t, xtables)
		}
	}
}

func (me *inits) ChildTables(xconf *xConfiguration, parent *table, filename string) {
	for _, xt := range xconf.Tables.Table {
		if strings.TrimSpace(xt.Extends) == parent.ID() {
			t := &table{
				module:  parent.module,
				id:      xt.ID,
				name:    xt.Name,
				parent:  parent,
				comment: xt.Comment,
				master:  xt.Master,
				slave:   xt.Slave,
			}
			for _, xc := range xt.Columns {
				var ec *column
				if strings.TrimSpace(xc.Extends) != "" {
					for _, cc := range conf.columns {
						if cc.ID() == xc.Extends {
							ec = cc
						}
					}
				}
				var d *domain
				if strings.TrimSpace(xc.Domain) != "" {
					for _, cd := range conf.domains {
						if cd.ID() == xc.Domain {
							d = cd
						}
					}
				}
				c := &column{
					parent:   ec,
					id:       xc.ID,
					name:     xc.Name,
					domain:   d,
					index:    xc.Index,
					comment:  xc.Comment,
					defaults: xc.Defaults,
					nullable: xc.Nullable,
				}
				c.field = strings.ToLowerFirst(strings.Camel(c.ID()))
				t.columns = append(t.columns, c)
			}
			conf.parent = append(conf.parent, t)
			me.ChildTables(xconf, t, filename)
		}
	}
}

func (me *inits) ProjectTables() {
	for _, m := range conf.modules {
		filename := "./conf/schema/tables-" + m.project.ID() + "-" + m.ID() + ".xml"
		xconf := util.DecodeXML("./" + filename)
		for _, xt := range xconf.Tables.Table {
			var p *table
			if strings.TrimSpace(xt.Extends) != "" {
				for _, t := range conf.parent {
					if t.ID() == xt.Extends {
						p = t
					}
				}
			}
			t := &table{
				module:      m,
				parent:      p,
				id:          xt.ID,
				name:        xt.Name,
				prefix:      xt.Prefix,
				generate:    xt.Generate,
				menu:        xt.Menu,
				comment:     xt.Comment,
				master:      xt.Master,
				slave:       xt.Slave,
				permissions: xt.Permissions,
				href:        xt.Href,
				buttons:     xt.Buttons,
			}
			if strings.IsBlank(t.buttons) {
				t.buttons = defaultButtons
			}
			switch t.Super() {
			case "pk":
				t.allColumnMaxLen = 5
				t.allFieldMaxLen = 5
				t.allTypeMaxLen = 13
			case "sys":
				t.allColumnMaxLen = 9
				t.allFieldMaxLen = 9
				t.allTypeMaxLen = 13
			case "tree":
				t.allColumnMaxLen = 12
				t.allFieldMaxLen = 11
				t.allTypeMaxLen = 13
			default:
				t.allColumnMaxLen = 5
				t.allFieldMaxLen = 5
				t.allTypeMaxLen = 12
			}
			t.columnMaxLen = 5
			t.fieldMaxLen = 5
			t.typeMaxLen = 12
			for _, xc := range xt.Columns {
				var ec *column
				if strings.TrimSpace(xc.Extends) != "" {
					for _, cc := range conf.columns {
						if cc.ID() == xc.Extends {
							ec = cc
						}
					}
				}
				var d *domain
				if strings.TrimSpace(xc.Domain) != "" {
					for _, cd := range conf.domains {
						if cd.ID() == xc.Domain {
							d = cd
						}
					}
				}
				c := &column{
					parent:   ec,
					id:       xc.ID,
					name:     xc.Name,
					domain:   d,
					index:    xc.Index,
					comment:  xc.Comment,
					defaults: xc.Defaults,
					nullable: xc.Nullable,
				}
				c.field = strings.ToLowerFirst(strings.Camel(c.ID()))
				t.columns = append(t.columns, c)
			}
			for _, c := range t.Columns() {
				if util.IsSuperCol(c.ID(), t.Super()) {
					continue
				}
				if t.fieldMaxLen < len(c.field) {
					t.fieldMaxLen = len(c.field)
				}
				if t.allFieldMaxLen < len(c.field) {
					t.allFieldMaxLen = len(c.field)
				}
				if t.columnMaxLen < len(c.ID()) {
					t.columnMaxLen = len(c.ID())
				}
				if t.allColumnMaxLen < len(c.ID()) {
					t.allColumnMaxLen = len(c.ID())
				}
				if t.typeMaxLen < len(c.Etype()) {
					t.typeMaxLen = len(c.Etype())
				}
				if t.allTypeMaxLen < len(c.Etype()) {
					t.allTypeMaxLen = len(c.Etype())
				}
			}
			conf.tables = append(conf.tables, t)
		}
	}
}
