// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package db_test

import (
	"../db"
	_ "github.com/ziutek/mymysql/godrv"
)

type User struct {
	db.Table
	Id       db.Id
	Code     db.String
	Name     db.String
	Password db.String
	Memo     db.String
	Genre    db.String
	Status   db.String
	Roles    db.String
	Posts    db.String
	Org      db.String
	Area     db.String
	Creater  db.Creater
	Created  db.Created
	Modifier db.Modifier
	Modified db.Modified
	Version  db.Version
	Disabled db.Disabled
}

func (me User) New() (entity *User) {
	entity = &User{Table: "user"}
	entity.Id.Column = "id"
	entity.Code.Column = "code"
	entity.Name.Column = "name"
	entity.Password.Column = "password"
	entity.Memo.Column = "memo"
	entity.Genre.Column = "genre"
	entity.Status.Column = "status"
	entity.Roles.Column = "roles"
	entity.Posts.Column = "posts"
	entity.Org.Column = "org"
	entity.Area.Column = "area"
	entity.Creater.Column = "creater"
	entity.Created.Column = "created"
	entity.Modifier.Column = "modifier"
	entity.Modified.Column = "modified"
	entity.Version.Column = "version"
	entity.Disabled.Column = "disabled"
	return
}

var factory db.Factory
var session db.Session
var query db.Query

func init() {
	factory, err := db.New("db")
	if err != nil {
		panic(err)
	}
	session, err = factory.Session()
	if err != nil {
		panic(err)
	}
	query = session.Query()
}
