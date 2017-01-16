package conf

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
)

//go:generate xgen -entity=$GOFILE -scaffold -clipath=gopkg.in/goyy/goyy.v0/app/adm -apipath=gopkg.in/goyy/goyy.v0/app/sys -tstpath=gopkg.in/goyy/goyy.v0/app/tst

// CONF Entity.
// @entity(module:"conf" project:"sys")
type Entity struct {
	entity.Sys
	table   schema.Table  `db:"table=sys_conf&comment=CONF"`
	name    entity.String `db:"column=name&comment=NAME"`
	code    entity.String `db:"column=code&comment=CODE"`
	content entity.String `db:"column=content&comment=CONTENT"`
	genre   entity.String `db:"column=genre&comment=GENRE"`
	usable  entity.String `db:"column=usable&comment=USABLE"`
	ordinal entity.String `db:"column=ordinal&comment=ORDINAL"`
}
