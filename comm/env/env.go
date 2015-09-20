// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package env

import (
	"encoding/xml"
	"errors"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"os"
)

// Get the database link parameters based on the environment profiles.
func Database(name string) (database xDatabase, err error) {
	xEnv, perr := parse("./conf/env/db.xml")
	if perr != nil {
		err = perr
		return
	}
	dbs := xEnv.Databases
	for _, db := range dbs {
		if db.Name == name {
			database = db
			return
		}
	}
	err = errors.New(i18N.Messagef("empty.environments.environment.database", name))
	return
}

// Get the mail server parameters based on the environment profiles.
func Mail(name string) (mail xMail, err error) {
	xEnv, perr := parse("./conf/env/mail.xml")
	if perr != nil {
		err = perr
		return
	}
	mails := xEnv.Mails
	for _, m := range mails {
		if m.Name == name {
			mail = m
			return
		}
	}
	err = errors.New(i18N.Messagef("empty.environments.environment.mail", name))
	return
}

func parse(file string) (xEnv *xEnvironment, err error) {
	f, ferr := os.Open(file)
	defer f.Close()
	if ferr != nil {
		err = ferr
		return
	}
	decoder := xml.NewDecoder(f)
	var xConf *xConfiguration
	if derr := decoder.Decode(&xConf); derr != nil {
		err = derr
		return
	}
	defaultName := xConf.Environments.Default
	if strings.IsBlank(defaultName) {
		err = errors.New(i18N.Message("empty.environments.default"))
		return
	}
	xenvs := xConf.Environments.Environment
	for _, xenv := range xenvs {
		if xenv.Id == defaultName {
			xEnv = &xenv
			return
		}
	}
	return
}
