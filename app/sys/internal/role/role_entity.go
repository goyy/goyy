package role

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
)

//go:generate xgen -entity=$GOFILE -scaffold -clidir=../../../bms -clipath=gopkg.in/goyy/goyy.v0/app/bms -apipath=gopkg.in/goyy/goyy.v0/app/sys

// ROLE Entity.
// @entity(project:"sys")
type Entity struct {
	entity.Sys
	table   schema.Table  `db:"table=sys_role&comment=ROLE"`
	name    entity.String `db:"column=name"`
	code    entity.String `db:"column=code"`
	genre   entity.String `db:"column=genre"`
	ordinal entity.String `db:"column=ordinal"`
}
