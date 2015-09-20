// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"gopkg.in/goyy/goyy.v0/comm/env"
	"log"
	"strings"
)

var conf *configuration = &configuration{}
var util *utils = &utils{}

func init() {
	valid := &valids{}
	valid.IsExistXML()
	valid.IsOkXML()

	data := &inits{}
	data.Init()
}

type inits struct{}

func (me *inits) Init() {
	me.Settings()
	me.Projects()
	me.Modules()
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
		db, err := env.Database(xp.Database)
		if err != nil {
			log.Fatal(err)
		}
		d := &database{xp.Database, db.DriverName, db.DataSourceName}
		p := &project{id: xp.Id, name: xp.Name, database: d, generate: xp.Generate, comment: xp.Comment}
		conf.projects = append(conf.projects, p)
	}
}

func (me *inits) Modules() {
	xconf := util.DecodeXML(xmodules)
	for _, xm := range xconf.Modules.Module {
		var p *project
		for _, cp := range conf.projects {
			if cp.Id() == xm.Project {
				p = cp
			}
		}
		m := &module{id: xm.Id, name: xm.Name, prefix: xm.Prefix, project: p, generate: xm.Generate, comment: xm.Comment}
		conf.modules = append(conf.modules, m)
	}
}

func (me *inits) Domains() {
	xconf := util.DecodeXML(xdomains)
	for _, xd := range xconf.Domains.Domain {
		d := &domain{id: xd.Id, name: xd.Name, types: xd.Types, length: xd.Length, precision: xd.Precision, comment: xd.Comment, defaults: xd.Defaults, nullable: xd.Nullable}
		conf.domains = append(conf.domains, d)
	}
}

func (me *inits) Columns() {
	xconf := util.DecodeXML(xcolumns)
	for _, xc := range xconf.Columns.Column {
		var d *domain
		for _, cd := range conf.domains {
			if cd.Id() == xc.Domain {
				d = cd
			}
		}
		c := &column{id: xc.Id, name: xc.Name, domain: d, index: xc.Index, comment: xc.Comment, defaults: xc.Defaults, nullable: xc.Nullable}
		conf.columns = append(conf.columns, c)
	}
}

func (me *inits) Tables() {
	xconf := util.DecodeXML(xtables)
	m := &module{project: &project{}}
	for _, xt := range xconf.Tables.Table {
		if strings.TrimSpace(xt.Extends) == "" {
			t := &table{module: m, id: xt.Id, name: xt.Name, comment: xt.Comment}
			for _, xc := range xt.Columns {
				var ec *column
				if strings.TrimSpace(xc.Extends) != "" {
					for _, cc := range conf.columns {
						if cc.Id() == xc.Extends {
							ec = cc
						}
					}
				}
				var d *domain
				if strings.TrimSpace(xc.Domain) != "" {
					for _, cd := range conf.domains {
						if cd.Id() == xc.Domain {
							d = cd
						}
					}
				}
				c := &column{parent: ec, id: xc.Id, name: xc.Name, domain: d, index: xc.Index, comment: xc.Comment, defaults: xc.Defaults, nullable: xc.Nullable}
				t.columns = append(t.columns, c)
			}
			conf.parent = append(conf.parent, t)
			me.ChildTables(xconf, t, xtables)
		}
	}
}

func (me *inits) ChildTables(xconf *xConfiguration, parent *table, filename string) {
	for _, xt := range xconf.Tables.Table {
		if strings.TrimSpace(xt.Extends) == parent.Id() {
			t := &table{module: parent.module, id: xt.Id, name: xt.Name, parent: parent, comment: xt.Comment}
			for _, xc := range xt.Columns {
				var ec *column
				if strings.TrimSpace(xc.Extends) != "" {
					for _, cc := range conf.columns {
						if cc.Id() == xc.Extends {
							ec = cc
						}
					}
				}
				var d *domain
				if strings.TrimSpace(xc.Domain) != "" {
					for _, cd := range conf.domains {
						if cd.Id() == xc.Domain {
							d = cd
						}
					}
				}
				c := &column{parent: ec, id: xc.Id, name: xc.Name, domain: d, index: xc.Index, comment: xc.Comment, defaults: xc.Defaults, nullable: xc.Nullable}
				t.columns = append(t.columns, c)
			}
			conf.parent = append(conf.parent, t)
			me.ChildTables(xconf, t, filename)
		}
	}
}

func (me *inits) ProjectTables() {
	for _, m := range conf.modules {
		filename := "./conf/schema/tables-" + m.project.Id() + "-" + m.Id() + ".xml"
		xconf := util.DecodeXML("./" + filename)
		for _, xt := range xconf.Tables.Table {
			var p *table
			if strings.TrimSpace(xt.Extends) != "" {
				for _, t := range conf.parent {
					if t.Id() == xt.Extends {
						p = t
					}
				}
			}
			t := &table{module: m, parent: p, id: xt.Id, name: xt.Name, prefix: xt.Prefix, generate: xt.Generate, comment: xt.Comment}
			for _, xc := range xt.Columns {
				var ec *column
				if strings.TrimSpace(xc.Extends) != "" {
					for _, cc := range conf.columns {
						if cc.Id() == xc.Extends {
							ec = cc
						}
					}
				}
				var d *domain
				if strings.TrimSpace(xc.Domain) != "" {
					for _, cd := range conf.domains {
						if cd.Id() == xc.Domain {
							d = cd
						}
					}
				}
				c := &column{parent: ec, id: xc.Id, name: xc.Name, domain: d, index: xc.Index, comment: xc.Comment, defaults: xc.Defaults, nullable: xc.Nullable}
				t.columns = append(t.columns, c)
			}
			conf.tables = append(conf.tables, t)
		}
	}
}
