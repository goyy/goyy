package dict

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
)

//go:generate xgen -entity=$GOFILE -scaffold -admpath=gopkg.in/goyy/goyy.v0/app/adm -apipath=gopkg.in/goyy/goyy.v0/app/sys -tstpath=gopkg.in/goyy/goyy.v0/app/tst

// DICT Entity.
// @entity(module:"dict" project:"sys")
type Entity struct {
	entity.Sys
	table   schema.Table  `db:"table=sys_dict&comment=DICT"`
	genre   entity.String `db:"column=genre&comment=GENRE"`
	descr   entity.String `db:"column=descr&comment=DESCR"`
	mkey    entity.String `db:"column=mkey&comment=MKEY"`
	mval    entity.String `db:"column=mval&comment=MVAL"`
	filters entity.String `db:"column=filters&comment=FILTERS"`
	ordinal entity.String `db:"column=ordinal&comment=ORDINAL"`
}
