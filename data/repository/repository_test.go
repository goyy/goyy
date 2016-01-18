// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package repository_test

import (
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

func TestRepositoryDelete(t *testing.T) {
	log.SetPriority(log.Perror)
	var dml string
	if repo.Dialect().Type() == dialect.MYSQL {
		dml = "delete from users where version = ?"
	} else {
		dml = "delete from users where version = :1"
	}
	repo.Exec(dml, 0)
}

func TestRepositoryInsert(t *testing.T) {
	repo.Insert(buildUser("01"))
	repo.Insert(buildUser("02"))
	repo.Insert(buildUser("03"))
	repo.Insert(buildUser("04"))
	repo.Insert(buildUser("05"))
	repo.Insert(buildUser("06"))
	repo.Insert(buildUser("07"))
	repo.Insert(buildUser("08"))
	repo.Insert(buildUser("09"))
	repo.Insert(buildUser("10"))
	repo.Insert(buildUser("11"))
	repo.Insert(buildUser("12"))
	repo.Insert(buildUser("13"))
	repo.Insert(buildUser("14"))
	repo.Insert(buildUser("15"))
	repo.Insert(buildUser("16"))
	repo.Insert(buildUser("17"))
	repo.Insert(buildUser("18"))
	repo.Insert(buildUser("19"))
	repo.Insert(buildUser("20"))
	repo.Insert(buildUser("21"))
	repo.Insert(buildUser("22"))
	repo.Insert(buildUser("23"))
	repo.Insert(buildUser("24"))
	repo.Insert(buildUser("25"))
}

func TestSelectPageByNamed(t *testing.T) {
	s := []struct {
		dql      string
		args     map[string]interface{}
		expected int
	}{
		{
			"select * from users",
			map[string]interface{}{"name": "1%", "version": 0, "memo": "memo"},
			26,
		},
		{
			"select * from users where name like #{name}",
			map[string]interface{}{"name": "1%", "version": 0, "memo": "memo"},
			10,
		},
		{
			"select * from users where version = #{version}",
			map[string]interface{}{"name": "1%", "version": 0, "memo": "memo"},
			25,
		},
		{
			"select * from users where name like #{name}{{if gt .version 0}} and version = #{version}{{end}}",
			map[string]interface{}{"name": "1%", "version": 0, "memo": "memo"},
			10,
		},
	}
	for _, v := range s {
		us := NewUserEntities(100)
		p := domain.NewPageable(1, 20)
		page, err := repo.SelectPageByNamed(us, p, v.dql, v.args)
		if err != nil {
			t.Error(err.Error())
			return
		}
		if page.TotalElements() != v.expected {
			t.Errorf(`SelectPageByNamed:"%v", want:"%v"`, page.TotalElements(), v.expected)
		}
	}
}
