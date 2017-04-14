package order

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
)

//go:generate xgen -entity=$GOFILE -scaffold -admpath=creditease.com/yxhn/hn-adm -apipath=creditease.com/yxhn/hn-example -tstpath=gopkg.in/goyy/goyy.v0/app/tst

// 订单信息实体结构.
// @entity(module:"order" project:"example")
type Entity struct {
	entity.Sys
	table      schema.Table   `db:"table=eg_order&comment=订单信息"`
	productId  entity.String  `db:"column=product_id&comment=产品"`
	discountId entity.String  `db:"column=discount_id&comment=折扣"`
	num        entity.Int     `db:"column=num&comment=购买数&default=1"`
	price      entity.Float64 `db:"column=price&comment=价格"`
}
