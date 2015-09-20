// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package properties implements properties utility functions.
package properties

import (
	"bufio"
	"errors"
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/files"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Properties is the interface that wraps the operation properties file method.
type Properties interface {
	Property(key string) string
	Propertyf(key string, args ...interface{}) string
	SetProperty(key, value string) error
}

// New creates a new Properties by filename.
func New(filename string) (p Properties, err error) {
	if files.IsExist(filename) == false {
		err = errors.New(`properties.New("` + filename + `") file not exists`)
		return
	}
	p = &defaultProperties{filename}
	return
}

type defaultProperties struct {
	filename string
}

// Searches for the property with the specified key in this property list.
// If the property can't be found, err will be set.
func (me *defaultProperties) Property(key string) (value string) {
	f, ferr := os.Open(me.filename)
	defer f.Close()
	if ferr != nil {
		log.Println(ferr)
		return
	}
	r := bufio.NewReader(f)
	str, rerr := r.ReadString('\n')
	if rerr != nil {
		log.Println(rerr)
		return
	}
	isFind := false
	for rerr != io.EOF {
		if strings.HasPrefix(str, key+"=") {
			value = strings.Replace(strings.TrimSpace(str), key+"=", "", 1)
			isFind = true
			break
		}
		str, rerr = r.ReadString('\n')
	}
	if isFind == false {
		value = ""
		log.Println("the [" + key + "] property can't be found")
		return
	}
	return
}

// Searches for the property with the specified key in this property list.
// If the property can't be found, err will be set.
func (me *defaultProperties) Propertyf(key string, args ...interface{}) (value string) {
	value = me.Property(key)
	if value == "" {
		return
	}
	value = fmt.Sprintf(value, args...)
	return
}

// Sets for the property with the specified key in this property list.
// If the property can't be found, err will be set.
func (me *defaultProperties) SetProperty(key, value string) (err error) {
	b, rerr := ioutil.ReadFile(me.filename)
	if rerr != nil {
		err = errors.New("properties.SetProperty:ReadFile [" + me.filename + "] failure")
		log.Println(rerr)
		return
	}
	v := me.Property(key)
	b = []byte(strings.Replace(string(b), key+"="+v, key+"="+value, 1))
	werr := ioutil.WriteFile(me.filename, b, 0777)
	if werr != nil {
		err = errors.New("properties.SetProperty:WriteFile [" + me.filename + "] failure")
		log.Println(werr)
		return
	}
	return
}
