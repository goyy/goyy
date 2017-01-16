package post

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
)

//go:generate xgen -entity=$GOFILE -scaffold -clipath=gopkg.in/goyy/goyy.v0/app/adm -apipath=gopkg.in/goyy/goyy.v0/app/sys -tstpath=gopkg.in/goyy/goyy.v0/app/tst

// POST MENU Entity.
// @entity(module:"post_menu" project:"sys" relationship:"slave")
type MenuEntity struct {
	entity.Sys
	table  schema.Table  `db:"table=sys_post_menu&comment=POST MENU"`
	postId entity.String `db:"column=post_id&comment=POST_ID"`
	menuId entity.String `db:"column=menu_id&comment=MENU_ID"`
}
