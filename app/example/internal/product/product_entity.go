package product

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
)

//go:generate xgen -entity=$GOFILE -scaffold -admpath=creditease.com/yxhn/hn-adm -apipath=creditease.com/yxhn/hn-example -tstpath=gopkg.in/goyy/goyy.v0/app/tst

// 产品信息实体结构.
// @entity(module:"product" project:"example")
type Entity struct {
	entity.Sys
	table schema.Table   `db:"table=eg_product&comment=产品信息"`
	name  entity.String  `db:"column=name&comment=名称" validation:"required=true"`
	num   entity.Int     `db:"column=num&comment=库存数" validation:"required=true&min=10"`
	price entity.Float64 `db:"column=price&comment=价格" validation:"required=true"`
}
