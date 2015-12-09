// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"os"
	"strings"
)

func (me *valids) IsOkXML() {
	isExit := false
	if me.IsOkSettings() {
		isExit = true
	}
	if me.IsOkProjects() {
		isExit = true
	}
	if me.IsOkModules() {
		isExit = true
	}
	if me.IsOkDomains() {
		isExit = true
	}
	if me.IsOkColumns() {
		isExit = true
	}
	if me.IsOkTables() {
		isExit = true
	}
	if me.IsOkProjectTables() {
		isExit = true
	}
	if isExit == true {
		os.Exit(0)
	}
}

func (me *valids) IsOkSettings() (isExit bool) {
	xconf := util.DecodeXML(xsettings)
	if xconf.Settings == nil {
		log.Fatal(i18N.Message("setting.empty"))
	}
	if xconf.Settings.Statement == nil {
		log.Fatal(i18N.Message("setting.statement.empty"))
	}
	if strings.TrimSpace(xconf.Settings.Statement.Seperator) == "" {
		isExit = true
		log.Println(i18N.Message("setting.statement.seperator.empty"))
	}
	if strings.TrimSpace(xconf.Settings.Statement.Case) == "" {
		isExit = true
		log.Println(i18N.Message("setting.statement.case.empty"))
	}
	return
}

func (me *valids) IsOkProjects() (isExit bool) {
	xconf := util.DecodeXML(xprojects)
	if xconf.Projects == nil {
		log.Fatal(i18N.Message("project.empty"))
	}
	if xconf.Projects.Project == nil {
		log.Fatal(i18N.Message("project.project.empty"))
	}
	for _, xp := range xconf.Projects.Project {
		if strings.TrimSpace(xp.Id) == "" {
			isExit = true
			log.Println(i18N.Message("project.id.empty"))
		}
		if strings.TrimSpace(xp.Database) == "" {
			isExit = true
			log.Println(i18N.Message("project.database.empty"))
		}
	}
	return
}

func (me *valids) IsOkModules() (isExit bool) {
	xconf := util.DecodeXML(xmodules)
	if xconf.Modules == nil {
		log.Fatal(i18N.Message("module.empty"))
	}
	if xconf.Modules.Module == nil {
		log.Fatal(i18N.Message("module.module.empty"))
	}
	for _, xm := range xconf.Modules.Module {
		if strings.TrimSpace(xm.Id) == "" {
			isExit = true
			log.Println(i18N.Message("module.id.empty"))
		}
		if strings.TrimSpace(xm.Project) == "" {
			isExit = true
			log.Println(i18N.Message("module.project.empty"))
		}
		hasExist := false
		xcfg := util.DecodeXML(xprojects)
		for _, xp := range xcfg.Projects.Project {
			if xp.Id == xm.Project {
				hasExist = true
			}
		}
		if hasExist == false {
			isExit = true
			log.Println(i18N.Messagef("module.project.errorf", xm.Project))
		}
	}
	return
}

func (me *valids) IsOkDomains() (isExit bool) {
	xconf := util.DecodeXML(xdomains)
	if xconf.Domains == nil {
		log.Fatal(i18N.Message("domain.empty"))
	}
	if xconf.Domains.Domain == nil {
		log.Fatal(i18N.Message("domain.domain.empty"))
	}
	for _, xd := range xconf.Domains.Domain {
		if strings.TrimSpace(xd.Id) == "" {
			isExit = true
			log.Println(i18N.Message("domain.id.empty"))
		}
		if strings.TrimSpace(xd.Types) == "" {
			isExit = true
			log.Println(i18N.Message("domain.types.empty"))
		}
		switch xd.Types {
		case "string":
			if xd.Length == 0 {
				isExit = true
				log.Println(i18N.Messagef("domain.length.emptyf", xd.Id))
			}
		case "float":
			if xd.Length <= 0 {
				isExit = true
				log.Println(i18N.Messagef("domain.length.emptyf", xd.Id))
			}
			if xd.Precision < 0 {
				isExit = true
				log.Println(i18N.Messagef("domain.precision.emptyf", xd.Id))
			}
		}
	}
	return
}

func (me *valids) IsOkColumns() (isExit bool) {
	xconf := util.DecodeXML(xcolumns)
	if xconf.Columns == nil {
		log.Fatal(i18N.Message("column.empty"))
	}
	if xconf.Columns.Column == nil {
		log.Fatal(i18N.Message("column.column.empty"))
	}
	for _, xc := range xconf.Columns.Column {
		if strings.TrimSpace(xc.Id) == "" {
			isExit = true
			log.Println(i18N.Message("column.id.empty"))
		}
		if strings.TrimSpace(xc.Domain) == "" {
			isExit = true
			log.Println(i18N.Message("column.domain.empty"))
		}
		hasExist := false
		xcfg := util.DecodeXML(xdomains)
		for _, xd := range xcfg.Domains.Domain {
			if xd.Id == xc.Domain {
				hasExist = true
			}
		}
		if hasExist == false {
			isExit = true
			log.Println(i18N.Messagef("column.domain.errorf", xc.Domain))
		}
	}
	return
}

func (me *valids) IsOkTables() (isExit bool) {
	xconf := util.DecodeXML(xtables)
	if xconf.Tables == nil {
		log.Fatal(i18N.Messagef("table.emptyf", xtables))
	}
	if xconf.Tables.Table == nil {
		log.Fatal(i18N.Messagef("table.table.emptyf", xtables))
	}
	for _, xt := range xconf.Tables.Table {
		if strings.TrimSpace(xt.Id) == "" {
			isExit = true
			log.Println(i18N.Message("table.id.empty"))
		}
		if strings.TrimSpace(xt.Extends) == "" {
			for _, xc := range xt.Columns {
				if strings.TrimSpace(xc.Id) == "" {
					if strings.TrimSpace(xc.Extends) == "" {
						isExit = true
						log.Println(i18N.Messagef("table.column.id.emptyf", xtables, xt.Id))
					}
				} else {
					if strings.TrimSpace(xc.Extends) == "" && strings.TrimSpace(xc.Domain) == "" {
						isExit = true
						log.Println(i18N.Messagef("table.column.domain.emptyf", xtables, xt.Id))
					}
				}
				me.IsExistColumn(xtables, xt.Id, xc.Extends, &isExit)
				me.IsExistDomain(xtables, xt.Id, xc.Domain, &isExit)
			}
			me.IsOkChildTables(xconf, xt, &isExit)
		}
	}
	return
}

func (me *valids) IsOkChildTables(xconf *xConfiguration, parent *xTable, isExit *bool) {
	for _, xt := range xconf.Tables.Table {
		if strings.TrimSpace(xt.Extends) == parent.Id {
			for _, xc := range xt.Columns {
				if strings.TrimSpace(xc.Id) == "" {
					if strings.TrimSpace(xc.Extends) == "" {
						*isExit = true
						log.Println(i18N.Messagef("table.column.id.emptyf", xtables, xt.Id))
					}
				} else {
					if strings.TrimSpace(xc.Extends) == "" && strings.TrimSpace(xc.Domain) == "" {
						*isExit = true
						log.Println(i18N.Messagef("table.column.domain.emptyf", xtables, xt.Id))
					}
				}
				me.IsExistColumn(xtables, xt.Id, xc.Extends, isExit)
				me.IsExistDomain(xtables, xt.Id, xc.Domain, isExit)
			}
			me.IsOkChildTables(xconf, xt, isExit)
		}
	}
	return
}

func (me *valids) IsOkProjectTables() (isExit bool) {
	xconf := util.DecodeXML(xmodules)
	for _, m := range xconf.Modules.Module {
		filename := "./conf/schema/tables-" + m.Project + "-" + m.Id + ".xml"
		xcfg := util.DecodeXML(filename)
		if xcfg.Tables == nil {
			log.Fatal(i18N.Messagef("table.emptyf", filename))
		}
		if xcfg.Tables.Table == nil {
			log.Fatal(i18N.Messagef("table.table.emptyf", filename))
		}
		for _, xt := range xcfg.Tables.Table {
			if strings.TrimSpace(xt.Id) == "" {
				isExit = true
				log.Println(i18N.Messagef("table.id.emptyf", filename))
			}
			me.IsExistTable(filename, xt.Id, xt.Extends, &isExit)
			for _, xc := range xt.Columns {
				if strings.TrimSpace(xc.Id) == "" {
					if strings.TrimSpace(xc.Extends) == "" {
						isExit = true
						log.Println(i18N.Messagef("table.column.id.emptyf", filename, xt.Id))
					}
				} else {
					if strings.TrimSpace(xc.Extends) == "" && strings.TrimSpace(xc.Domain) == "" {
						isExit = true
						log.Println(i18N.Messagef("table.column.domain.emptyf", filename, xt.Id))
					}
				}
				me.IsExistColumn(filename, xt.Id, xc.Extends, &isExit)
				me.IsExistDomain(filename, xt.Id, xc.Domain, &isExit)
			}
		}
	}
	return
}

func (me *valids) IsExistTable(filename, tableId, extends string, isExit *bool) {
	if strings.TrimSpace(extends) != "" {
		hasExist := false
		xconf := util.DecodeXML(xtables)
		for _, xc := range xconf.Tables.Table {
			if xc.Id == extends {
				hasExist = true
			}
		}
		if hasExist == false {
			*isExit = true
			log.Println(i18N.Messagef("table.extends.errorf", filename, tableId, extends))
		}
	}
}

func (me *valids) IsExistColumn(filename, tableId, columnId string, isExit *bool) {
	if strings.TrimSpace(columnId) != "" {
		hasExist := false
		xconf := util.DecodeXML(xcolumns)
		for _, xc := range xconf.Columns.Column {
			if xc.Id == columnId {
				hasExist = true
			}
		}
		if hasExist == false {
			*isExit = true
			log.Println(i18N.Messagef("table.column.extends.errorf", filename, tableId, columnId))
		}
	}
}

func (me *valids) IsExistDomain(filename, tableId, domainId string, isExit *bool) {
	if strings.TrimSpace(domainId) != "" {
		hasExist := false
		xconf := util.DecodeXML(xdomains)
		for _, xd := range xconf.Domains.Domain {
			if xd.Id == domainId {
				hasExist = true
			}
		}
		if hasExist == false {
			*isExit = true
			log.Println(i18N.Messagef("table.column.domain.errorf", filename, tableId, domainId))
		}
	}
}
