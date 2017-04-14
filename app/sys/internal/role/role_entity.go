package role

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
)

//go:generate xgen -entity=$GOFILE -scaffold -admpath=gopkg.in/goyy/goyy.v0/app/adm -apipath=gopkg.in/goyy/goyy.v0/app/sys -tstpath=gopkg.in/goyy/goyy.v0/app/tst

// ROLE Entity.
// @entity(module:"role" project:"sys")
type Entity struct {
	entity.Sys
	table    schema.Table  `db:"table=sys_role&comment=ROLE"`
	name     entity.String `db:"column=name&comment=NAME"`
	code     entity.String `db:"column=code&comment=CODE"`
	genre    entity.String `db:"column=genre&comment=GENRE"`
	classify entity.String `db:"column=classify&comment=CLASSIFY"`
	ordinal  entity.String `db:"column=ordinal&comment=ORDINAL"`
	postIds  entity.String `db:"comment=POST_IDS&transient=true"`
}
