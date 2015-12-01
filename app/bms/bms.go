package main

import _ "gopkg.in/goyy/goyy.v0/app/bms/internal/conf"

import (
	_ "gopkg.in/goyy/goyy.v0/app/bms/internal"
	_ "gopkg.in/goyy/goyy.v0/app/sys"
	"gopkg.in/goyy/goyy.v0/comm/log"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func main() {
	xhttp.Conf.Addr = ":9097"
	xhttp.Use(xhttp.Recovery())
	err := xhttp.Run()
	if err != nil {
		log.Error(err.Error())
	}
}
