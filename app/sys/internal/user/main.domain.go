package user

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
)

//go:generate xgen -entity=$GOFILE -scaffold -clidir=../../../bms -clipath=gopkg.in/goyy/goyy.v0/app/bms -apipath=gopkg.in/goyy/goyy.v0/app/sys

// USER Entity.
// @entity(project:"sys")
type Entity struct {
	entity.Sys
	table     schema.Table  `db:"table=sys_user&comment=USER"`
	name      entity.String `db:"column=name"`
	code      entity.String `db:"column=code"`
	passwd    entity.String `db:"column=passwd"`
	genre     entity.String `db:"column=genre"`
	email     entity.String `db:"column=email"`
	tel       entity.String `db:"column=tel"`
	mobile    entity.String `db:"column=mobile"`
	areaId    entity.String `db:"column=area_id"`
	orgId     entity.String `db:"column=org_id"`
	loginName entity.String `db:"column=login_name"`
	loginIp   entity.String `db:"column=login_ip"`
	loginTime entity.Int64  `db:"column=login_time"`
}
