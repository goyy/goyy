package blacklist

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
)

//go:generate xgen -entity=$GOFILE -scaffold -admpath=gopkg.in/goyy/goyy.v0/app/adm -apipath=gopkg.in/goyy/goyy.v0/app/sys -tstpath=gopkg.in/goyy/goyy.v0/app/tst

// BLACKLIST Entity.
// @entity(module:"blacklist" project:"sys")
type Entity struct {
	entity.Sys
	table  schema.Table  `db:"table=sys_blacklist&comment=BLACKLIST"`
	name   entity.String `db:"column=name&comment=NAME"`
	genre  entity.String `db:"column=genre&comment=GENRE"`
	usable entity.String `db:"column=usable&comment=USABLE"`
}
