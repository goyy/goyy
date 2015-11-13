package entity

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
)

//go:generate xgen -entity=$GOFILE -scaffold

// User stores user account information.
// @entity(project:"sys")
type User struct {
	entity.Sys
	table schema.Table  `orm:"table=users&comment=user"`
	email entity.String `orm:"column=email"`
	roles entity.String `orm:"column=roles&transient=true"`
}
