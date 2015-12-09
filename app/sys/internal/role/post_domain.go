package role

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
)

//go:generate xgen -entity=$GOFILE -scaffold -clidir=../../../bms -clipath=gopkg.in/goyy/goyy.v0/app/bms -apipath=gopkg.in/goyy/goyy.v0/app/sys

// ROLE POST Entity.
// @entity(project:"sys" relationship:"slave")
type PostEntity struct {
	entity.Sys
	table  schema.Table  `db:"table=sys_role_post&comment=ROLE POST"`
	roleId entity.String `db:"column=role_id"`
	postId entity.String `db:"column=post_id"`
}
