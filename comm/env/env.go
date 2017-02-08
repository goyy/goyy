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

// ParseSettings get the settings link parameters based on the configuration file.
func ParseSettings() (out Settings, err error) {
	f, ferr := os.Open(conf + "/env/settings.xml")
	if ferr != nil {
		err = ferr
		return
	}
	defer f.Close()
	decoder := xml.NewDecoder(f)
	var xConf *Configuration
	if derr := decoder.Decode(&xConf); derr != nil {
		err = derr
		return
	}
	return xConf.Settings, nil
}

// ParseDatabase get the database link parameters based on the environment profiles.
func ParseDatabase(name string) (out Database, err error) {
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

// ParseMail get the mail server parameters based on the environment profiles.
func ParseMail(name string) (out Mail, err error) {
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

// ParseSession get the session based on the environment profiles.
func ParseSession(name string) (out Session, err error) {
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

// ParseApi get the api based on the environment profiles.
func ParseApi(name string) (out Api, err error) {
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

// ParseAsset get the asset based on the environment profiles.
func ParseAsset(name string) (out Static, err error) {
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

// ParseStatic get the static based on the environment profiles.
func ParseStatic(name string) (out Static, err error) {
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

// ParseDeveloper get the developer based on the environment profiles.
func ParseDeveloper(name string) (out Static, err error) {
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

// ParseOperation get the operation based on the environment profiles.
func ParseOperation(name string) (out Static, err error) {
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

// ParseUpload get the upload based on the environment profiles.
func ParseUpload(name string) (out Upload, err error) {
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

// ParseExport get the export based on the environment profiles.
func ParseExport(name string) (out Export, err error) {
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

// ParseHtml get the html based on the environment profiles.
func ParseHtml(name string) (out Template, err error) {
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

// ParseTemplate get the template based on the environment profiles.
func ParseTemplate(name string) (out Template, err error) {
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

// ParseSensitive get the sensitive word based on the environment profiles.
func ParseSensitive(name string) (out Sensitive, err error) {
	xEnv, err := parse(conf + "/env/sensitive.xml")
	if err != nil {
		return
	}
	outs := xEnv.Sensitives
	for _, v := range outs {
		if v.Name == name {
			out = v
			return
		}
	}
	err = errors.New(i18N.Messagef("empty.environments.environment.sensitive", name))
	return
}

// ParseLog get the log based on the environment profiles.
func ParseLog(name string) (out Log, err error) {
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

// ParseSecure get the secure based on the environment profiles.
func ParseSecure(name string) (out Secure, err error) {
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

func parse(file string) (xEnv *Environment, err error) {
	f, ferr := os.Open(file)
	if ferr != nil {
		err = ferr
		return
	}
	defer f.Close()
	decoder := xml.NewDecoder(f)
	var xConf *Configuration
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
