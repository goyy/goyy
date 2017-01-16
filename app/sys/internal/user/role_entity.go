package user

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
)

//go:generate xgen -entity=$GOFILE -scaffold -clipath=gopkg.in/goyy/goyy.v0/app/adm -apipath=gopkg.in/goyy/goyy.v0/app/sys -tstpath=gopkg.in/goyy/goyy.v0/app/tst

// USER ROLE Entity.
// @entity(module:"user_role" project:"sys" relationship:"slave")
type RoleEntity struct {
	entity.Sys
	table  schema.Table  `db:"table=sys_user_role&comment=USER ROLE"`
	userId entity.String `db:"column=user_id&comment=USER_ID"`
	roleId entity.String `db:"column=role_id&comment=ROLE_ID"`
}
