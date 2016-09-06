// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xsql_test

import (
	"strconv"
	"testing"

	"gopkg.in/goyy/goyy.v0/comm/log"
	"gopkg.in/goyy/goyy.v0/data/dialect"
	"gopkg.in/goyy/goyy.v0/data/domain"
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/util/times"
)

var created = times.NowUnix()

func buildUser(i string) entity.Interface {
	user := NewUser()
	user.SetCode(i)
	user.SetName(i)
	user.SetPassword(i)
	user.SetMemo(i)
	user.SetGenre(i)
	user.SetStatus(i)
	user.SetRoles(i)
	user.SetPosts(i)
	user.SetOrg(i)
	user.SetArea(i)
	user.SetCreater(i)
	user.SetCreated(created)
	user.SetModifier(i)
	user.SetModified(times.NowUnix())
	user.SetVersion(0)
	user.SetDeletion(0)
	return user
}

func TestDBDelete(t *testing.T) {
	log.SetPriority(log.Perror)
	var dml string
	if db.Dialect().Type() == dialect.MYSQL {
		dml = "delete from users where version = ?"
	} else {
		dml = "delete from users where version = :1"
	}
	db.Exec(dml, 0)
}

func TestDBInsert(t *testing.T) {
	db.Insert(buildUser("a01"))
	db.Insert(buildUser("a02"))
	db.Insert(buildUser("a03"))
	db.Insert(buildUser("a04"))
	db.Insert(buildUser("a05"))
	db.Insert(buildUser("a06"))
	db.Insert(buildUser("a07"))
	db.Insert(buildUser("a08"))
	db.Insert(buildUser("a09"))
	db.Insert(buildUser("a10"))
	db.Insert(buildUser("a11"))
	db.Insert(buildUser("a12"))
	db.Insert(buildUser("a13"))
	db.Insert(buildUser("a14"))
	db.Insert(buildUser("a15"))
	db.Insert(buildUser("a16"))
	db.Insert(buildUser("a17"))
	db.Insert(buildUser("a18"))
	db.Insert(buildUser("a19"))
	db.Insert(buildUser("a20"))
	db.Insert(buildUser("a21"))
	db.Insert(buildUser("a22"))
	db.Insert(buildUser("a23"))
	db.Insert(buildUser("a24"))
	db.Insert(buildUser("a25"))
}

func TestDBGet(t *testing.T) {
	user := NewUser()
	user.SetId("aa")
	expected := "aa"
	if _ = db.Get(user); user.Name() != expected {
		t.Errorf(`db.Get():"%v", want:"%v"`, user.Name(), expected)
	}
}

func TestDBSifterRow(t *testing.T) {
	s, _ := domain.NewSift("sNameEQ", "a11")
	user := NewUser()
	err := db.Sifter(s).Row(user)
	if err != nil {
		t.Error(err.Error())
		return
	}
	expected := "a11"
	if out := user.Creater(); out != expected {
		t.Errorf(`db.Sifter().Row():"%v", want:"%v"`, out, expected)
	}
}

func TestDBSifterRows(t *testing.T) {
	s1, _ := domain.NewSift("sNameGT", "a11")
	s2, _ := domain.NewSift("sVersionEQ", "0")
	s3, _ := domain.NewSift("sNameOA", "asc")
	users := NewUserEntities(20)
	err := db.Sifter(s1, s2, s3).Rows(users)
	if err != nil {
		t.Error(err.Error())
		return
	}
	got := 14
	if out := users.Len(); out != got {
		t.Errorf(`db.Sifter().Rows().Len():"%v", want:"%v"`, out, got)
	}
	expected := "a12"
	if out := users.Index(0).(*User).Name(); out != expected {
		t.Errorf(`db.Sifter().Rows().Index(0):"%v", want:"%v"`, out, expected)
	}
}

func TestDBSifterPage(t *testing.T) {
	sVersionEQ, _ := domain.NewSift("sVersionEQ", "0")
	sIdOA, _ := domain.NewSift("sIdOA", "asc")
	pageable := domain.NewPageable(2, 10)
	content := NewUserEntities(30)
	out, err := db.Sifter(sVersionEQ, sIdOA).Page(content, pageable)
	if err != nil {
		t.Error(err.Error())
		return
	}
	expected := 25
	if out.TotalElements() != expected {
		t.Errorf(`page.TotalElements():"%v", want:"%v"`, out.TotalElements(), expected)
	}
	expected = 3
	if out.TotalPages() != expected {
		t.Errorf(`page.TotalPages():"%v", want:"%v"`, out.TotalPages(), expected)
	}
	expected = 2
	if out.PageNo() != expected {
		t.Errorf(`page.PageNo():"%v", want:"%v"`, out.PageNo(), expected)
	}
	expected = 10
	if out.PageSize() != expected {
		t.Errorf(`page.PageSize():"%v", want:"%v"`, out.PageSize(), expected)
	}
	want := "a11"
	name := out.Content().Index(0).(*User).Name()
	if name != want {
		t.Errorf(`page.Content():"%v", want:"%v"`, name, want)
	}
}

func TestDBSifterCount(t *testing.T) {
	s, _ := domain.NewSift("sNameEQ", "a11")
	user := NewUser()
	out, err := db.Sifter(s).Count(user)
	if err != nil {
		t.Error(err.Error())
		return
	}
	expected := 1
	if out != expected {
		t.Errorf(`db.Sifter().Count():"%v", want:"%v"`, out, expected)
	}
}

func TestDBQueryRows(t *testing.T) {
	var dql string
	if db.Dialect().Type() == dialect.MYSQL {
		dql = "select * from users where name like ?"
	} else {
		dql = "select * from users where name like :1"
	}
	users := NewUserEntities(30)
	err := db.Query(dql, "a2%").Rows(users)
	if err != nil {
		t.Error(err.Error())
		return
	}
	expected := 6
	if out := users.Len(); out != expected {
		t.Errorf(`query.Rows():"%v", want:"%v"`, out, expected)
	}
	for i := 0; i < users.Len(); i++ {
		want := "a" + strconv.Itoa(20+i)
		if out := users.Value(i); out.Code() != want {
			t.Errorf(`get(%v).Code():"%v", want:"%v"`, i, out.Code(), want)
		}
	}
}

func TestDBQueryRow(t *testing.T) {
	var dql string
	if db.Dialect().Type() == dialect.MYSQL {
		dql = "select * from users where name = ?"
	} else {
		dql = "select * from users where name = :1"
	}
	user := NewUser()
	err := db.Query(dql, "a12").Row(user)
	if err != nil {
		t.Error(err.Error())
		return
	}
	expected := "a12"
	if out := user.Creater(); out != expected {
		t.Errorf(`query.Row():"%v", want:"%v"`, out, expected)
	}
}

func TestDBQueryPage(t *testing.T) {
	var dql string
	if db.Dialect().Type() == dialect.MYSQL {
		dql = "select * from users where version = ? order by id"
	} else {
		dql = "select * from users where version = :1 order by id"
	}
	pageable := domain.NewPageable(2, 10)
	content := NewUserEntities(30)
	out, err := db.Query(dql, 0).Page(content, pageable)
	if err != nil {
		t.Error(err.Error())
		return
	}
	expected := 25
	if out.TotalElements() != expected {
		t.Errorf(`page.TotalElements():"%v", want:"%v"`, out.TotalElements(), expected)
	}
	expected = 3
	if out.TotalPages() != expected {
		t.Errorf(`page.TotalPages():"%v", want:"%v"`, out.TotalPages(), expected)
	}
	expected = 2
	if out.PageNo() != expected {
		t.Errorf(`page.PageNo():"%v", want:"%v"`, out.PageNo(), expected)
	}
	expected = 10
	if out.PageSize() != expected {
		t.Errorf(`page.PageSize():"%v", want:"%v"`, out.PageSize(), expected)
	}
	want := "a11"
	name := out.Content().Index(0).(*User).Name()
	if name != want {
		t.Errorf(`page.Content():"%v", want:"%v"`, name, want)
	}
}

func TestDBQueryInt(t *testing.T) {
	var dql string
	if db.Dialect().Type() == dialect.MYSQL {
		dql = "select count(*) from users where name like ?"
	} else {
		dql = "select count(*) from users where name like :1"
	}
	out, err := db.Query(dql, "a1%").Int()
	if err != nil {
		t.Error(err.Error())
		return
	}
	expected := 10
	if out != expected {
		t.Errorf(`query.Int():"%v", want:"%v"`, out, expected)
	}
}

func TestDBQueryStr(t *testing.T) {
	var dql string
	if db.Dialect().Type() == dialect.MYSQL {
		dql = "select code from users where name = ?"
	} else {
		dql = "select code from users where name = :1"
	}
	out, err := db.Query(dql, "a03").Str()
	if err != nil {
		t.Error(err.Error())
		return
	}
	expected := "a03"
	if out != expected {
		t.Errorf(`query.Str():"%v", want:"%v"`, out, expected)
	}
}

func TestDBQueryTime(t *testing.T) {
	var dql string
	if db.Dialect().Type() == dialect.MYSQL {
		dql = "select created from users where name = ?"
	} else {
		dql = "select created from users where name = :1"
	}
	out, err := db.Query(dql, "a03").Int()
	if err != nil {
		t.Error(err.Error())
		return
	}
	expected := times.FormatUnixYYMDHMS(created)
	if times.FormatUnixYYMDHMS(int64(out)) != expected {
		t.Errorf(`query.Time():"%v", want:"%v"`, times.FormatUnixYYMDHMS(int64(out)), expected)
	}
}

func TestDBQueryIn(t *testing.T) {
	var dql string
	if db.Dialect().Type() == dialect.MYSQL {
		dql = "select count(*) from users where name in (?,?)"
	} else {
		dql = "select count(*) from users where name in (:1,:2)"
	}
	out, err := db.Query(dql, "a01", "a02").Int()
	if err != nil {
		t.Error(err.Error())
		return
	}
	expected := 2
	if out != expected {
		t.Errorf(`query:in:"%v", want:"%v"`, out, expected)
	}
}

func TestDBNamedQueryInt(t *testing.T) {
	s := []struct {
		dql      string
		args     map[string]interface{}
		expected int
	}{
		{
			"select count(*) from users",
			map[string]interface{}{"name": "a1%", "version": 0, "memo": "memo"},
			26,
		},
		{
			"select count(*) from users where name like #{name}",
			map[string]interface{}{"name": "a1%", "version": 0, "memo": "memo"},
			10,
		},
		{
			"select count(*) from users where version = #{version}",
			map[string]interface{}{"name": "a1%", "version": 0, "memo": "memo"},
			25,
		},
		{
			"select count(*) from users where name like #{name}{{if gt .version 0}} and version = #{version}{{end}}",
			map[string]interface{}{"name": "a1%", "version": 0, "memo": "memo"},
			10,
		},
	}
	for _, v := range s {
		query, err := db.NamedQuery(v.dql, v.args)
		if err != nil {
			t.Error(err.Error())
			return
		}
		out, err := query.Int()
		if err != nil {
			t.Error(err.Error())
			return
		}
		if out != v.expected {
			t.Errorf(`namedQuery.Int():"%v", want:"%v"`, out, v.expected)
		}
	}
}

func TestDBUpdate(t *testing.T) {
	var dql string
	if db.Dialect().Type() == dialect.MYSQL {
		dql = "select * from users where name = ?"
	} else {
		dql = "select * from users where name = :1"
	}
	user := NewUser()
	err := db.Query(dql, "a22").Row(user)
	if err != nil {
		t.Error(err.Error())
		return
	}
	expected := "user2"
	user.SetCode(expected)
	db.Update(user)
	user = NewUser()
	err = db.Query(dql, "a22").Row(user)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if out := user.Code(); out != expected {
		t.Errorf(`query.Update:"%v", want:"%v"`, out, expected)
	}
}

func TestDBDisable(t *testing.T) {
	var dql string
	if db.Dialect().Type() == dialect.MYSQL {
		dql = "select * from users where name = ?"
	} else {
		dql = "select * from users where name = :1"
	}
	user := NewUser()
	err := db.Query(dql, "a23").Row(user)
	if err != nil {
		t.Error(err.Error())
		return
	}
	db.Disable(user)
	expected := entity.DeletionDisable
	user = NewUser()
	err = db.Query(dql, "a23").Row(user)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if out := user.Deletion(); out != expected {
		t.Errorf(`query.Disable():"%v", want:"%v"`, out, expected)
	}
}

func TestDBSifterCountIn(t *testing.T) {
	s, _ := domain.NewSift("sNameIN", "a11,a12,a13,a14")
	user := NewUser()
	out, err := db.Sifter(s).Count(user)
	if err != nil {
		t.Error(err.Error())
		return
	}
	expected := 4
	if out != expected {
		t.Errorf(`db.Sifter().Count():"%v", want:"%v"`, out, expected)
	}
}

func TestDBSifterCountNotIn(t *testing.T) {
	s, _ := domain.NewSift("sNameNI", "a11,a12,a13,a14")
	user := NewUser()
	out, err := db.Sifter(s).Count(user)
	if err != nil {
		t.Error(err.Error())
		return
	}
	expected := 22
	if out != expected {
		t.Errorf(`db.Sifter().Count():"%v", want:"%v"`, out, expected)
	}
}

func TestDBSifterCountBetween(t *testing.T) {
	s, _ := domain.NewSift("sNameBE", "a11,a12")
	user := NewUser()
	out, err := db.Sifter(s).Count(user)
	if err != nil {
		t.Error(err.Error())
		return
	}
	expected := 2
	if out != expected {
		t.Errorf(`db.Sifter().Count():"%v", want:"%v"`, out, expected)
	}
}

func TestDBSifterCountNotBetween(t *testing.T) {
	s, _ := domain.NewSift("sNameNB", "a11,a12")
	user := NewUser()
	out, err := db.Sifter(s).Count(user)
	if err != nil {
		t.Error(err.Error())
		return
	}
	expected := 24
	if out != expected {
		t.Errorf(`db.Sifter().Count():"%v", want:"%v"`, out, expected)
	}
}
