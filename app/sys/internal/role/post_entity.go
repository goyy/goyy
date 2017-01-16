package role

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
)

//go:generate xgen -entity=$GOFILE -scaffold -clipath=gopkg.in/goyy/goyy.v0/app/adm -apipath=gopkg.in/goyy/goyy.v0/app/sys -tstpath=gopkg.in/goyy/goyy.v0/app/tst

// ROLE POST Entity.
// @entity(module:"role_post" project:"sys" relationship:"slave")
type PostEntity struct {
	entity.Sys
	table  schema.Table  `db:"table=sys_role_post&comment=ROLE POST"`
	roleId entity.String `db:"column=role_id&comment=ROLE_ID"`
	postId entity.String `db:"column=post_id&comment=POST_ID"`
}
