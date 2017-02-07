package user

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
)

//go:generate xgen -entity=$GOFILE -scaffold -admpath=gopkg.in/goyy/goyy.v0/app/adm -apipath=gopkg.in/goyy/goyy.v0/app/sys -tstpath=gopkg.in/goyy/goyy.v0/app/tst

// USER Entity.
// @entity(module:"user" project:"sys")
type Entity struct {
	entity.Sys
	table     schema.Table  `db:"table=sys_user&comment=USER"`
	name      entity.String `db:"column=name&comment=NAME"`
	code      entity.String `db:"column=code&comment=CODE"`
	passwd    entity.String `db:"column=passwd&comment=PASSWD"`
	genre     entity.String `db:"column=genre&comment=GENRE"`
	email     entity.String `db:"column=email&comment=EMAIL"`
	tel       entity.String `db:"column=tel&comment=TEL"`
	mobile    entity.String `db:"column=mobile&comment=MOBILE"`
	areaId    entity.String `db:"column=area_id&comment=AREA_ID"`
	orgId     entity.String `db:"column=org_id&comment=ORG_ID"`
	loginName entity.String `db:"column=login_name&comment=LOGIN NAME"`
	loginIp   entity.String `db:"column=login_ip&comment=LOGIN IP"`
	loginTime entity.Int64  `db:"column=login_time&comment=LOGIN TIME"`
	roleIds   entity.String `db:"comment=ROLE_IDS&transient=true"`
}
