package org

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
)

//go:generate xgen -entity=$GOFILE -scaffold -admpath=gopkg.in/goyy/goyy.v0/app/adm -apipath=gopkg.in/goyy/goyy.v0/app/sys -tstpath=gopkg.in/goyy/goyy.v0/app/tst

// ORG Entity.
// @entity(module:"org" project:"sys")
type Entity struct {
	entity.Tree
	table  schema.Table  `db:"table=sys_org&comment=ORG"`
	areaId entity.String `db:"column=area_id&comment=AREA_ID"`
}
