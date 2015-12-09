package menu

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
)

//go:generate xgen -entity=$GOFILE -scaffold -clidir=../../../bms -clipath=gopkg.in/goyy/goyy.v0/app/bms -apipath=gopkg.in/goyy/goyy.v0/app/sys

// MENU Entity.
// @entity(project:"sys")
type Entity struct {
	entity.Tree
	table      schema.Table  `db:"table=sys_menu&comment=MENU"`
	href       entity.String `db:"column=href"`
	target     entity.String `db:"column=target"`
	icon       entity.String `db:"column=icon"`
	hidden     entity.Bool   `db:"column=hidden"`
	permission entity.String `db:"column=permission"`
}
