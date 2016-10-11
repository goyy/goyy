package menu

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
)

//go:generate xgen -entity=$GOFILE -scaffold -clipath=gopkg.in/goyy/goyy.v0/app/bms -apipath=gopkg.in/goyy/goyy.v0/app/sys -tstpath=gopkg.in/goyy/goyy.v0/app/tst

// MENU Entity.
// @entity(module:"menu" project:"sys")
type Entity struct {
	entity.Tree
	table      schema.Table  `db:"table=sys_menu&comment=MENU"`
	href       entity.String `db:"column=href&comment=HREF"`
	target     entity.String `db:"column=target&comment=TARGET"`
	icon       entity.String `db:"column=icon&comment=ICON"`
	hidden     entity.String `db:"column=hidden&comment=HIDDEN"`
	permission entity.String `db:"column=permission&comment=PERMISSION"`
}
