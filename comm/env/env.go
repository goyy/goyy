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

// SetConf set the root directory of the configuration file.
// Default value is "./conf".
func SetConf(path string) {
	conf = path
}

// Settings get the settings link parameters based on the configuration file.
func Settings() (out XMLSettings, err error) {
	f, ferr := os.Open(conf + "/env/settings.xml")
	if ferr != nil {
		err = ferr
		return
	}
	defer f.Close()
	decoder := xml.NewDecoder(f)
	var xConf *XMLConfiguration
	if derr := decoder.Decode(&xConf); derr != nil {
		err = derr
		return
	}
	return xConf.Settings, nil
}

// Database get the database link parameters based on the environment profiles.
func Database(name string) (out XMLDatabase, err error) {
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

// Mail get the mail server parameters based on the environment profiles.
func Mail(name string) (out XMLMail, err error) {
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

// Session get the session based on the environment profiles.
func Session(name string) (out XMLSession, err error) {
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

// API get the api based on the environment profiles.
func API(name string) (out XMLAPI, err error) {
	xEnv, err := parse(conf + "/env/api.xml")
	if err != nil {
		return
	}
	outs := xEnv.APIs
	for _, v := range outs {
		if v.Name == name {
			out = v
			return
		}
	}
	err = errors.New(i18N.Messagef("empty.environments.environment.api", name))
	return
}

// Asset get the asset based on the environment profiles.
func Asset(name string) (out XMLStatic, err error) {
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

// Static get the static based on the environment profiles.
func Static(name string) (out XMLStatic, err error) {
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

// Developer get the developer based on the environment profiles.
func Developer(name string) (out XMLStatic, err error) {
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

// Operation get the operation based on the environment profiles.
func Operation(name string) (out XMLStatic, err error) {
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

// Upload get the upload based on the environment profiles.
func Upload(name string) (out XMLUpload, err error) {
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

// Export get the export based on the environment profiles.
func Export(name string) (out XMLExport, err error) {
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

// HTML get the html based on the environment profiles.
func HTML(name string) (out XMLTemplate, err error) {
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

// Template get the template based on the environment profiles.
func Template(name string) (out XMLTemplate, err error) {
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

// Illegal get the illegal based on the environment profiles.
func Illegal(name string) (out XMLIllegal, err error) {
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

// Log get the log based on the environment profiles.
func Log(name string) (out XMLLog, err error) {
	xEnv, err := parse(conf + "/env/log.xml")
	if err != nil {
		return
	}
	outs := xEnv.Logs
	for _, v := range outs {
		if v.Name == name {
			out = v
			return
		}
	}
	err = errors.New(i18N.Messagef("empty.environments.environment.log", name))
	return
}

// Secure get the secure based on the environment profiles.
func Secure(name string) (out XMLSecure, err error) {
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

func parse(file string) (xEnv *XMLEnvironment, err error) {
	f, ferr := os.Open(file)
	if ferr != nil {
		err = ferr
		return
	}
	defer f.Close()
	decoder := xml.NewDecoder(f)
	var xConf *XMLConfiguration
	if derr := decoder.Decode(&xConf); derr != nil {
		err = derr
		return
	}
	xenvs := xConf.Environments.Environment
	for _, xenv := range xenvs {
		if xenv.ID == profile.Default() {
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
		if xenv.ID == defaultName {
			xEnv = &xenv
			return
		}
	}
	return
}
