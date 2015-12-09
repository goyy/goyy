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
	table schema.Table  `db:"table=users&comment=user"`
	email entity.String `db:"column=email"`
	roles entity.String `db:"column=roles&transient=true"`
}
