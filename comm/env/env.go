// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package env

import (
	"encoding/xml"
	"errors"
	"os"

	"gopkg.in/goyy/goyy.v0/comm/profile"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

var conf = "./conf"

// Set the root directory of the configuration file.
// Default value is "./conf".
func SetConf(path string) {
	conf = path
}

// Get the settings link parameters based on the configuration file.
func Settings() (out xSettings, err error) {
	f, ferr := os.Open(conf + "/env/settings.xml")
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
	return xConf.Settings, nil
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

// Get the api based on the environment profiles.
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

// Get the export based on the environment profiles.
func Export(name string) (out xExport, err error) {
	xEnv, err := parse(conf + "/env/export.xml")
	if err != nil {
		return
	}
	outs := xEnv.Exports
	for _, v := range outs {
		if v.Name == name {
			out = v
			return
		}
	}
	err = errors.New(i18N.Messagef("empty.environments.environment.export", name))
	return
}

// Get the html based on the environment profiles.
func Html(name string) (out xTemplate, err error) {
	xEnv, err := parse(conf + "/env/template.xml")
	if err != nil {
		return
	}
	outs := xEnv.Htmls
	for _, v := range outs {
		if v.Name == name {
			out = v
			return
		}
	}
	err = errors.New(i18N.Messagef("empty.environments.environment.html", name))
	return
}

// Get the template based on the environment profiles.
func Template(name string) (out xTemplate, err error) {
	xEnv, err := parse(conf + "/env/template.xml")
	if err != nil {
		return
	}
	outs := xEnv.Templates
	for _, v := range outs {
		if v.Name == name {
			out = v
			return
		}
	}
	err = errors.New(i18N.Messagef("empty.environments.environment.template", name))
	return
}

// Get the illegal based on the environment profiles.
func Illegal(name string) (out xIllegal, err error) {
	xEnv, err := parse(conf + "/env/illegal.xml")
	if err != nil {
		return
	}
	outs := xEnv.Illegals
	for _, v := range outs {
		if v.Name == name {
			out = v
			return
		}
	}
	err = errors.New(i18N.Messagef("empty.environments.environment.illegal", name))
	return
}

// Get the secure based on the environment profiles.
func Secure(name string) (out xSecure, err error) {
	xEnv, err := parse(conf + "/env/secure.xml")
	if err != nil {
		return
	}
	outs := xEnv.Secures
	for _, v := range outs {
		if v.Name == name {
			out = v
			return
		}
	}
	err = errors.New(i18N.Messagef("empty.environments.environment.secure", name))
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
	xenvs := xConf.Environments.Environment
	for _, xenv := range xenvs {
		if xenv.Id == profile.Default() {
			xEnv = &xenv
			return
		}
	}
	defaultName := xConf.Environments.Default
	if strings.IsBlank(defaultName) {
		err = errors.New(i18N.Message("empty.environments.default"))
		return
	}
	for _, xenv := range xenvs {
		if xenv.Id == defaultName {
			xEnv = &xenv
			return
		}
	}
	return
}
