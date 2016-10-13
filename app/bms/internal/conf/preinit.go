package conf

import (
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/goyy/goyy.v0/app/sys/api/dict"
	"gopkg.in/goyy/goyy.v0/comm/profile"
	"gopkg.in/goyy/goyy.v0/data/dialect"
	"gopkg.in/goyy/goyy.v0/data/schema"
	"gopkg.in/goyy/goyy.v0/data/service"
	"gopkg.in/goyy/goyy.v0/web/conf"
)

var _ = settings()

func settings() string {
	conf.Init("goyy", profile.DEV, profile.BMS)
	service.DB = service.NewDB(&dialect.MySQL{}, "goyy")
	schema.ParseDict = dict.Mval
	return ""
}
