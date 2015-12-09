package org

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
)

//go:generate xgen -entity=$GOFILE -scaffold -clidir=../../../bms -clipath=gopkg.in/goyy/goyy.v0/app/bms -apipath=gopkg.in/goyy/goyy.v0/app/sys

// ORG Entity.
// @entity(project:"sys")
type Entity struct {
	entity.Tree
	table  schema.Table  `db:"table=sys_org&comment=ORG"`
	areaId entity.String `db:"column=area_id"`
}
