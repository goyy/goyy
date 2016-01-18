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

func TestSessionDelete(t *testing.T) {
	log.SetPriority(log.Perror)
	var dml string
	if session.DBType() == dialect.MYSQL {
		dml = "delete from users where version = ?"
	} else {
		dml = "delete from users where version = :1"
	}
	session.Exec(dml, 0)
}

func TestSessionInsert(t *testing.T) {
	session.Insert(buildUser("01"))
	session.Insert(buildUser("02"))
	session.Insert(buildUser("03"))
	session.Insert(buildUser("04"))
	session.Insert(buildUser("05"))
	session.Insert(buildUser("06"))
	session.Insert(buildUser("07"))
	session.Insert(buildUser("08"))
	session.Insert(buildUser("09"))
	session.Insert(buildUser("10"))
	session.Insert(buildUser("11"))
	session.Insert(buildUser("12"))
	session.Insert(buildUser("13"))
	session.Insert(buildUser("14"))
	session.Insert(buildUser("15"))
	session.Insert(buildUser("16"))
	session.Insert(buildUser("17"))
	session.Insert(buildUser("18"))
	session.Insert(buildUser("19"))
	session.Insert(buildUser("20"))
	session.Insert(buildUser("21"))
	session.Insert(buildUser("22"))
	session.Insert(buildUser("23"))
	session.Insert(buildUser("24"))
	session.Insert(buildUser("25"))
}

func TestSessionGet(t *testing.T) {
	user := NewUser()
	user.SetId("aa")
	expected := "aa"
	if _ = session.Get(user); user.Name() != expected {
		t.Errorf(`session.Get():"%v", want:"%v"`, user.Name(), expected)
	}
}

func TestSessionSelectOne(t *testing.T) {
	s, _ := domain.NewSift("sNameEQ", "11")
	user := NewUser()
	err := session.SelectOne(user, s)
	if err != nil {
		t.Error(err.Error())
		return
	}
	expected := "11"
	if out := user.Creater(); out != expected {
		t.Errorf(`session.SelectOne():"%v", want:"%v"`, out, expected)
	}
}

func TestSessionSelectList(t *testing.T) {
	s1, _ := domain.NewSift("sNameGT", "11")
	s2, _ := domain.NewSift("sVersionEQ", "0")
	s3, _ := domain.NewSift("sNameOA", "asc")
	users := NewUserEntities(20)
	err := session.SelectList(users, s1, s2, s3)
	if err != nil {
		t.Error(err.Error())
		return
	}
	got := 14
	if out := users.Len(); out != got {
		t.Errorf(`session.SelectList().Len():"%v", want:"%v"`, out, got)
	}
	expected := "12"
	if out := users.Index(0).(*User).Name(); out != expected {
		t.Errorf(`session.SelectList().Index(0):"%v", want:"%v"`, out, expected)
	}
}

func TestSessionSelectPage(t *testing.T) {
	sVersionEQ, _ := domain.NewSift("sVersionEQ", "0")
	sIdOA, _ := domain.NewSift("sIdOA", "asc")
	pageable := domain.NewPageable(2, 10)
	content := NewUserEntities(30)
	out, err := session.SelectPage(content, pageable, sVersionEQ, sIdOA)
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
	want := "11"
	name := out.Content().Index(0).(*User).Name()
	if name != want {
		t.Errorf(`page.Content():"%v", want:"%v"`, name, want)
	}
}

func TestSessionQueryRows(t *testing.T) {
	var dql string
	if session.DBType() == dialect.MYSQL {
		dql = "select * from users where name like ?"
	} else {
		dql = "select * from users where name like :1"
	}
	users := NewUserEntities(30)
	err := session.Query(dql, "2%").Rows(users)
	if err != nil {
		t.Error(err.Error())
		return
	}
	expected := 6
	if out := users.Len(); out != expected {
		t.Errorf(`query.Rows():"%v", want:"%v"`, out, expected)
	}
	for i := 0; i < users.Len(); i++ {
		want := strconv.Itoa(20 + i)
		if out := users.Value(i); out.Code() != want {
			t.Errorf(`get(%v).Code():"%v", want:"%v"`, i, out.Code(), want)
		}
	}
}

func TestSessionQueryRow(t *testing.T) {
	var dql string
	if session.DBType() == dialect.MYSQL {
		dql = "select * from users where name = ?"
	} else {
		dql = "select * from users where name = :1"
	}
	user := NewUser()
	err := session.Query(dql, "12").Row(user)
	if err != nil {
		t.Error(err.Error())
		return
	}
	expected := "12"
	if out := user.Creater(); out != expected {
		t.Errorf(`query.Row():"%v", want:"%v"`, out, expected)
	}
}

func TestSessionQueryInt(t *testing.T) {
	var dql string
	if session.DBType() == dialect.MYSQL {
		dql = "select count(*) from users where name like ?"
	} else {
		dql = "select count(*) from users where name like :1"
	}
	out, err := session.Query(dql, "1%").Int()
	if err != nil {
		t.Error(err.Error())
		return
	}
	expected := 10
	if out != expected {
		t.Errorf(`query.Int():"%v", want:"%v"`, out, expected)
	}
}

func TestSessionNamedQueryInt(t *testing.T) {
	s := []struct {
		dql      string
		args     map[string]interface{}
		expected int
	}{
		{
			"select count(*) from users",
			map[string]interface{}{"name": "1%", "version": 0, "memo": "memo"},
			26,
		},
		{
			"select count(*) from users where name like #{name}",
			map[string]interface{}{"name": "1%", "version": 0, "memo": "memo"},
			10,
		},
		{
			"select count(*) from users where version = #{version}",
			map[string]interface{}{"name": "1%", "version": 0, "memo": "memo"},
			25,
		},
		{
			"select count(*) from users where name like #{name}{{if gt .version 0}} and version = #{version}{{end}}",
			map[string]interface{}{"name": "1%", "version": 0, "memo": "memo"},
			10,
		},
	}
	for _, v := range s {
		query, err := session.NamedQuery(v.dql, v.args)
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

func TestSessionQueryStr(t *testing.T) {
	var dql string
	if session.DBType() == dialect.MYSQL {
		dql = "select code from users where name = ?"
	} else {
		dql = "select code from users where name = :1"
	}
	out, err := session.Query(dql, "03").Str()
	if err != nil {
		t.Error(err.Error())
		return
	}
	expected := "03"
	if out != expected {
		t.Errorf(`query.Str():"%v", want:"%v"`, out, expected)
	}
}

func TestSessionQueryTime(t *testing.T) {
	var dql string
	if session.DBType() == dialect.MYSQL {
		dql = "select created from users where name = ?"
	} else {
		dql = "select created from users where name = :1"
	}
	out, err := session.Query(dql, "03").Int()
	if err != nil {
		t.Error(err.Error())
		return
	}
	expected := times.FormatUnixYYMDHMS(created)
	if times.FormatUnixYYMDHMS(int64(out)) != expected {
		t.Errorf(`query.Time():"%v", want:"%v"`, times.FormatUnixYYMDHMS(int64(out)), expected)
	}
}

func TestSessionQueryIn(t *testing.T) {
	var dql string
	if session.DBType() == dialect.MYSQL {
		dql = "select count(*) from users where name in (?,?)"
	} else {
		dql = "select count(*) from users where name in (:1,:2)"
	}
	out, err := session.Query(dql, "01", "02").Int()
	if err != nil {
		t.Error(err.Error())
		return
	}
	expected := 2
	if out != expected {
		t.Errorf(`query:in:"%v", want:"%v"`, out, expected)
	}
}

func TestSessionQueryPage(t *testing.T) {
	var dql string
	if session.DBType() == dialect.MYSQL {
		dql = "select * from users where version = ? order by id"
	} else {
		dql = "select * from users where version = :1 order by id"
	}
	pageable := domain.NewPageable(2, 10)
	content := NewUserEntities(30)
	out, err := session.Query(dql, 0).Page(content, pageable)
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
	want := "11"
	name := out.Content().Index(0).(*User).Name()
	if name != want {
		t.Errorf(`page.Content():"%v", want:"%v"`, name, want)
	}
}

func TestSessionUpdate(t *testing.T) {
	var dql string
	if session.DBType() == dialect.MYSQL {
		dql = "select * from users where name = ?"
	} else {
		dql = "select * from users where name = :1"
	}
	user := NewUser()
	err := session.Query(dql, "22").Row(user)
	if err != nil {
		t.Error(err.Error())
		return
	}
	expected := "user2"
	user.SetCode(expected)
	session.Update(user)
	user = NewUser()
	err = session.Query(dql, "22").Row(user)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if out := user.Code(); out != expected {
		t.Errorf(`query.Update:"%v", want:"%v"`, out, expected)
	}
}

func TestSessionDisable(t *testing.T) {
	var dql string
	if session.DBType() == dialect.MYSQL {
		dql = "select * from users where name = ?"
	} else {
		dql = "select * from users where name = :1"
	}
	user := NewUser()
	err := session.Query(dql, "23").Row(user)
	if err != nil {
		t.Error(err.Error())
		return
	}
	session.Disable(user)
	expected := entity.DeletionDisable
	user = NewUser()
	err = session.Query(dql, "23").Row(user)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if out := user.Deletion(); out != expected {
		t.Errorf(`query.Disable():"%v", want:"%v"`, out, expected)
	}
}
