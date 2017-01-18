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
	if me.IsOkButtons() {
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
		if strings.TrimSpace(xp.ID) == "" {
			isExit = true
			log.Println(i18N.Message("project.id.empty"))
		}
		if strings.TrimSpace(xp.Name) == "" {
			isExit = true
			log.Println(i18N.Message("project.name.empty"))
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
		if strings.TrimSpace(xm.ID) == "" {
			isExit = true
			log.Println(i18N.Message("module.id.empty"))
		}
		if strings.TrimSpace(xm.Name) == "" {
			isExit = true
			log.Println(i18N.Message("module.name.empty"))
		}
		if strings.TrimSpace(xm.Project) == "" {
			isExit = true
			log.Println(i18N.Message("module.project.empty"))
		}
		hasExist := false
		xcfg := util.DecodeXML(xprojects)
		for _, xp := range xcfg.Projects.Project {
			if xp.ID == xm.Project {
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

func (me *valids) IsOkButtons() (isExit bool) {
	xconf := util.DecodeXML(xbuttons)
	if xconf.Buttons == nil {
		log.Fatal(i18N.Message("button.empty"))
	}
	if xconf.Buttons.Button == nil {
		log.Fatal(i18N.Message("button.button.empty"))
	}
	for _, xb := range xconf.Buttons.Button {
		if strings.TrimSpace(xb.ID) == "" {
			isExit = true
			log.Println(i18N.Message("button.id.empty"))
		}
		if strings.TrimSpace(xb.Name) == "" {
			isExit = true
			log.Println(i18N.Message("button.name.empty"))
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
		if strings.TrimSpace(xd.ID) == "" {
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
				log.Println(i18N.Messagef("domain.length.emptyf", xd.ID))
			}
		case "float":
			if xd.Length <= 0 {
				isExit = true
				log.Println(i18N.Messagef("domain.length.emptyf", xd.ID))
			}
			if xd.Precision < 0 {
				isExit = true
				log.Println(i18N.Messagef("domain.precision.emptyf", xd.ID))
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
		if strings.TrimSpace(xc.ID) == "" {
			isExit = true
			log.Println(i18N.Message("column.id.empty"))
		}
		if strings.TrimSpace(xc.Name) == "" {
			isExit = true
			log.Println(i18N.Message("column.name.empty"))
		}
		if strings.TrimSpace(xc.Comment) == "" {
			isExit = true
			log.Println(i18N.Message("column.comment.empty"))
		}
		if strings.TrimSpace(xc.Domain) == "" {
			isExit = true
			log.Println(i18N.Message("column.domain.empty"))
		}
		hasExist := false
		xcfg := util.DecodeXML(xdomains)
		for _, xd := range xcfg.Domains.Domain {
			if xd.ID == xc.Domain {
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
		if strings.TrimSpace(xt.ID) == "" {
			isExit = true
			log.Println(i18N.Message("table.id.empty"))
		}
		if strings.TrimSpace(xt.Extends) == "" {
			if strings.TrimSpace(xt.Name) == "" {
				isExit = true
				log.Println(i18N.Message("table.name.empty"))
			}
			if strings.TrimSpace(xt.Comment) == "" {
				isExit = true
				log.Println(i18N.Message("table.comment.empty"))
			}
			for _, xc := range xt.Columns {
				if strings.TrimSpace(xc.Extends) == "" {
					if strings.TrimSpace(xc.ID) == "" {
						isExit = true
						log.Println(i18N.Messagef("table.column.id.emptyf", xtables, xt.ID))
					}
					if strings.TrimSpace(xc.Name) == "" {
						isExit = true
						log.Println(i18N.Messagef("table.column.name.emptyf", xtables, xt.ID))
					}
					if strings.TrimSpace(xc.Comment) == "" {
						isExit = true
						log.Println(i18N.Messagef("table.column.comment.emptyf", xtables, xt.ID))
					}
					if strings.TrimSpace(xc.Domain) == "" {
						isExit = true
						log.Println(i18N.Messagef("table.column.domain.emptyf", xtables, xt.ID))
					}
				}
				me.IsExistColumn(xtables, xt.ID, xc.Extends, &isExit)
				me.IsExistDomain(xtables, xt.ID, xc.Domain, &isExit)
			}
			me.IsOkChildTables(xconf, xt, &isExit)
		}
		me.IsExistButton(xtables, xt.ID, xt.Buttons, &isExit)
	}
	return
}

func (me *valids) IsOkChildTables(xconf *xConfiguration, parent *xTable, isExit *bool) {
	for _, xt := range xconf.Tables.Table {
		if strings.TrimSpace(xt.Extends) == parent.ID {
			for _, xc := range xt.Columns {
				if strings.TrimSpace(xc.Extends) == "" {
					if strings.TrimSpace(xc.ID) == "" {
						*isExit = true
						log.Println(i18N.Messagef("table.column.id.emptyf", xtables, xt.ID))
					}
					if strings.TrimSpace(xc.Name) == "" {
						*isExit = true
						log.Println(i18N.Messagef("table.column.name.emptyf", xtables, xt.ID))
					}
					if strings.TrimSpace(xc.Comment) == "" {
						*isExit = true
						log.Println(i18N.Messagef("table.column.comment.emptyf", xtables, xt.ID))
					}
					if strings.TrimSpace(xc.Domain) == "" {
						*isExit = true
						log.Println(i18N.Messagef("table.column.domain.emptyf", xtables, xt.ID))
					}
				}
				me.IsExistColumn(xtables, xt.ID, xc.Extends, isExit)
				me.IsExistDomain(xtables, xt.ID, xc.Domain, isExit)
			}
			me.IsOkChildTables(xconf, xt, isExit)
		}
		me.IsExistButton(xtables, xt.ID, xt.Buttons, isExit)
	}
	return
}

func (me *valids) IsOkProjectTables() (isExit bool) {
	xconf := util.DecodeXML(xmodules)
	for _, m := range xconf.Modules.Module {
		filename := "./conf/schema/tables-" + m.Project + "-" + m.ID + ".xml"
		xcfg := util.DecodeXML(filename)
		if xcfg.Tables == nil {
			log.Fatal(i18N.Messagef("table.emptyf", filename))
		}
		if xcfg.Tables.Table == nil {
			log.Fatal(i18N.Messagef("table.table.emptyf", filename))
		}
		for _, xt := range xcfg.Tables.Table {
			if strings.TrimSpace(xt.ID) == "" {
				isExit = true
				log.Println(i18N.Messagef("table.id.emptyf", filename))
			}
			me.IsExistTable(filename, xt.ID, xt.Extends, &isExit)
			for _, xc := range xt.Columns {
				if strings.TrimSpace(xc.Extends) == "" {
					if strings.TrimSpace(xc.ID) == "" {
						isExit = true
						log.Println(i18N.Messagef("table.column.id.emptyf", filename, xt.ID))
					}
					if strings.TrimSpace(xc.Name) == "" {
						isExit = true
						log.Println(i18N.Messagef("table.column.name.emptyf", filename, xt.ID))
					}
					if strings.TrimSpace(xc.Comment) == "" {
						isExit = true
						log.Println(i18N.Messagef("table.column.comment.emptyf", filename, xt.ID))
					}
					if strings.TrimSpace(xc.Domain) == "" {
						isExit = true
						log.Println(i18N.Messagef("table.column.domain.emptyf", filename, xt.ID))
					}
				}
				me.IsExistColumn(filename, xt.ID, xc.Extends, &isExit)
				me.IsExistDomain(filename, xt.ID, xc.Domain, &isExit)
			}
			me.IsExistButton(xtables, xt.ID, xt.Buttons, &isExit)
		}
	}
	return
}

func (me *valids) IsExistTable(filename, tableID, extends string, isExit *bool) {
	if strings.TrimSpace(extends) != "" {
		hasExist := false
		xconf := util.DecodeXML(xtables)
		for _, xc := range xconf.Tables.Table {
			if xc.ID == extends {
				hasExist = true
			}
		}
		if hasExist == false {
			*isExit = true
			log.Println(i18N.Messagef("table.extends.errorf", filename, tableID, extends))
		}
	}
}

func (me *valids) IsExistColumn(filename, tableID, columnID string, isExit *bool) {
	if strings.TrimSpace(columnID) != "" {
		hasExist := false
		xconf := util.DecodeXML(xcolumns)
		for _, xc := range xconf.Columns.Column {
			if xc.ID == columnID {
				hasExist = true
			}
		}
		if hasExist == false {
			*isExit = true
			log.Println(i18N.Messagef("table.column.extends.errorf", filename, tableID, columnID))
		}
	}
}

func (me *valids) IsExistDomain(filename, tableID, domainID string, isExit *bool) {
	if strings.TrimSpace(domainID) != "" {
		hasExist := false
		xconf := util.DecodeXML(xdomains)
		for _, xd := range xconf.Domains.Domain {
			if xd.ID == domainID {
				hasExist = true
			}
		}
		if hasExist == false {
			*isExit = true
			log.Println(i18N.Messagef("table.column.domain.errorf", filename, tableID, domainID))
		}
	}
}

func (me *valids) IsExistButton(filename, tableID, buttonIDs string, isExit *bool) {
	if strings.TrimSpace(buttonIDs) != "" {
		xconf := util.DecodeXML(xbuttons)
		buttons := strings.Split(buttonIDs, ",")
		for _, button := range buttons {
			if strings.TrimSpace(button) == "" {
				continue
			}
			hasExist := false
			for _, xb := range xconf.Buttons.Button {
				if xb.ID == button {
					hasExist = true
				}
			}
			if hasExist == false {
				*isExit = true
				log.Println(i18N.Messagef("table.buttons.errorf", filename, tableID, button))
			}
		}
	}
}
