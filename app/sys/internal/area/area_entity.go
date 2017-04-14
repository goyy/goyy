package area

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
)

//go:generate xgen -entity=$GOFILE -scaffold -admpath=gopkg.in/goyy/goyy.v0/app/adm -apipath=gopkg.in/goyy/goyy.v0/app/sys -tstpath=gopkg.in/goyy/goyy.v0/app/tst

// AREA Entity.
// @entity(module:"area" project:"sys")
type Entity struct {
	entity.Tree
	table schema.Table `db:"table=sys_area&comment=AREA"`
}
