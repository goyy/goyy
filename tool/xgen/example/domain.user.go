package entity

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
)

//go:generate xgen -entity=$GOFILE -scaffold

// User stores user account information.
// @entity(project:"sys")
type User struct {
	table    schema.Table  `orm:"table=users&comment=user"`
	id       entity.String `orm:"column=id&primary=true"`
	email    entity.String `orm:"column=email"`
	roles    entity.String `orm:"column=roles&transient=true"`
	creater  entity.String `orm:"column=creater&creater=true"`
	created  entity.Time   `orm:"column=created&created=true&default=-62135596800"`
	modifier entity.String `orm:"column=modifier&modifier=true"`
	modified entity.Time   `orm:"column=modified&modified=true&default=-62135596800"`
	version  entity.Int    `orm:"column=version&version=true"`
	deletion entity.Int    `orm:"column=deletion&deletion=true"`
}
