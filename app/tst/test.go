package test

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/goyy/goyy.v0/comm/env"
	"gopkg.in/goyy/goyy.v0/comm/log"
	"gopkg.in/goyy/goyy.v0/comm/profile"
	"gopkg.in/goyy/goyy.v0/data/cache"
	"gopkg.in/goyy/goyy.v0/data/dialect"
	"gopkg.in/goyy/goyy.v0/data/service"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func init() {
	// profile
	profile.SetActives(profile.TEST)

	// DataBase
	env.SetConf("/gopkg.in/goyy/goyy.v0/app/tst/conf")
	service.DB = service.NewDB(&dialect.MySQL{}, "goyy")

	// cache
	if v, err := env.Session("goyy"); err == nil {
		xhttp.Conf.Session.Addr = v.Addr
		xhttp.Conf.Session.Password = v.Password
		cache.Init(cache.Conf{
			Address:     v.Addr,
			Password:    v.Password,
			MaxIdle:     80,
			MaxActive:   12000,
			IdleTimeout: 240 * time.Second,
		})
	} else {
		log.Println(err.Error())
	}
}
