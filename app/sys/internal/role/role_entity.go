package role

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
)

//go:generate xgen -entity=$GOFILE -scaffold -clidir=../../../bms -clipath=gopkg.in/goyy/goyy.v0/app/bms -apipath=gopkg.in/goyy/goyy.v0/app/sys -tstpath=gopkg.in/goyy/goyy.v0/app/tst

// ROLE Entity.
// @entity(module:"role" project:"sys")
type Entity struct {
	entity.Sys
	table   schema.Table  `db:"table=sys_role&comment=ROLE"`
	name    entity.String `db:"column=name&comment=NAME"`
	code    entity.String `db:"column=code&comment=CODE"`
	genre   entity.String `db:"column=genre&comment=GENRE"`
	ordinal entity.String `db:"column=ordinal&comment=ORDINAL"`
}
