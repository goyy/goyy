package post

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
)

//go:generate xgen -entity=$GOFILE -scaffold -clipath=gopkg.in/goyy/goyy.v0/app/bms -apipath=gopkg.in/goyy/goyy.v0/app/sys -tstpath=gopkg.in/goyy/goyy.v0/app/tst

// POST Entity.
// @entity(module:"post" project:"sys")
type Entity struct {
	entity.Tree
	table   schema.Table  `db:"table=sys_post&comment=POST"`
	menuIds entity.String `db:"comment=MENU_IDS&transient=true"`
}
