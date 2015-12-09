package user

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
)

//go:generate xgen -entity=$GOFILE -scaffold -clidir=../../../bms -clipath=gopkg.in/goyy/goyy.v0/app/bms -apipath=gopkg.in/goyy/goyy.v0/app/sys

// USER ROLE Entity.
// @entity(project:"sys" relationship:"slave")
type RoleEntity struct {
	entity.Sys
	table  schema.Table  `db:"table=sys_user_role&comment=USER ROLE"`
	userId entity.String `db:"column=user_id"`
	roleId entity.String `db:"column=role_id"`
}
