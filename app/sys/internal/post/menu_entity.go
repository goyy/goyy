package post

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
)

//go:generate xgen -entity=$GOFILE -scaffold -clidir=../../../bms -clipath=gopkg.in/goyy/goyy.v0/app/bms -apipath=gopkg.in/goyy/goyy.v0/app/sys

// POST MENU Entity.
// @entity(project:"sys" relationship:"slave")
type MenuEntity struct {
	entity.Sys
	table  schema.Table  `db:"table=sys_post_menu&comment=POST MENU"`
	postId entity.String `db:"column=post_id"`
	menuId entity.String `db:"column=menu_id"`
}
