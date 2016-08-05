// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package env

import (
	"encoding/xml"
	"errors"
	"os"

	"gopkg.in/goyy/goyy.v0/util/strings"
)

var conf = "./conf"

// Set the root directory of the configuration file.
// Default value is "./conf".
func SetConf(path string) {
	conf = path
}

// Get the database link parameters based on the environment profiles.
func Database(name string) (out xDatabase, err error) {
	xEnv, err := parse(conf + "/env/db.xml")
	if err != nil {
		return
	}
	outs := xEnv.Databases
	for _, v := range outs {
		if v.Name == name {
			out = v
			return
		}
	}
	err = errors.New(i18N.Messagef("empty.environments.environment.database", name))
	return
}

// Get the mail server parameters based on the environment profiles.
func Mail(name string) (out xMail, err error) {
	xEnv, err := parse(conf + "/env/mail.xml")
	if err != nil {
		return
	}
	outs := xEnv.Mails
	for _, v := range outs {
		if v.Name == name {
			out = v
			return
		}
	}
	err = errors.New(i18N.Messagef("empty.environments.environment.mail", name))
	return
}

// Get the session based on the environment profiles.
func Session(name string) (out xSession, err error) {
	xEnv, err := parse(conf + "/env/session.xml")
	if err != nil {
		return
	}
	outs := xEnv.Sessions
	for _, v := range outs {
		if v.Name == name {
			out = v
			return
		}
	}
	err = errors.New(i18N.Messagef("empty.environments.environment.session", name))
	return
}

// Get the apis based on the environment profiles.
func Api(name string) (out xApi, err error) {
	xEnv, err := parse(conf + "/env/api.xml")
	if err != nil {
		return
	}
	outs := xEnv.Apis
	for _, v := range outs {
		if v.Name == name {
			out = v
			return
		}
	}
	err = errors.New(i18N.Messagef("empty.environments.environment.api", name))
	return
}

// Get the asset based on the environment profiles.
func Asset(name string) (out xStatic, err error) {
	xEnv, err := parse(conf + "/env/static.xml")
	if err != nil {
		return
	}
	outs := xEnv.Assets
	for _, v := range outs {
		if v.Name == name {
			out = v
			return
		}
	}
	err = errors.New(i18N.Messagef("empty.environments.environment.asset", name))
	return
}

// Get the static based on the environment profiles.
func Static(name string) (out xStatic, err error) {
	xEnv, err := parse(conf + "/env/static.xml")
	if err != nil {
		return
	}
	outs := xEnv.Statics
	for _, v := range outs {
		if v.Name == name {
			out = v
			return
		}
	}
	err = errors.New(i18N.Messagef("empty.environments.environment.static", name))
	return
}

// Get the developer based on the environment profiles.
func Developer(name string) (out xStatic, err error) {
	xEnv, err := parse(conf + "/env/static.xml")
	if err != nil {
		return
	}
	outs := xEnv.Developers
	for _, v := range outs {
		if v.Name == name {
			out = v
			return
		}
	}
	err = errors.New(i18N.Messagef("empty.environments.environment.developer", name))
	return
}

// Get the operation based on the environment profiles.
func Operation(name string) (out xStatic, err error) {
	xEnv, err := parse(conf + "/env/static.xml")
	if err != nil {
		return
	}
	outs := xEnv.Operations
	for _, v := range outs {
		if v.Name == name {
			out = v
			return
		}
	}
	err = errors.New(i18N.Messagef("empty.environments.environment.operation", name))
	return
}

// Get the upload based on the environment profiles.
func Upload(name string) (out xUpload, err error) {
	xEnv, err := parse(conf + "/env/upload.xml")
	if err != nil {
		return
	}
	outs := xEnv.Uploads
	for _, v := range outs {
		if v.Name == name {
			out = v
			return
		}
	}
	err = errors.New(i18N.Messagef("empty.environments.environment.upload", name))
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
