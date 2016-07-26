package conf

import (
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/goyy/goyy.v0/comm/log"
	"gopkg.in/goyy/goyy.v0/comm/profile"
	"gopkg.in/goyy/goyy.v0/data/dialect"
	"gopkg.in/goyy/goyy.v0/data/repository"
	"gopkg.in/goyy/goyy.v0/data/service"
	"gopkg.in/goyy/goyy.v0/data/xsql"
)

var _ = settings()

func settings() string {
	profile.SetActives(profile.DEV)
	if profile.Accepts(profile.DEV) {
		log.DefaultOutput = log.Ostd
	} else {
		log.DefaultOutput = log.Odailyfile
	}
	if profile.Accepts(profile.PROD) {
		xsql.SetPriority(log.Perror)
	} else {
		xsql.SetPriority(log.Pdebug)
	}
	service.Repository = repository.New(&dialect.MySQL{}, "goyy")
	return ""
}
