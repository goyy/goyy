package entity

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
)

//go:generate xgen -entity=$GOFILE

// Profile stores user attributes.
// @entity
type Profile struct {
	table     schema.Table
	id        entity.String
	userId    entity.String
	attribute entity.String
	val       entity.String
}
