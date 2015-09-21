// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package jsons

import (
	"gopkg.in/goyy/goyy.v0/util/errors"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

type Parser interface {
	Value(key string) string
	Len(key string) int
}

func NewParser(json string) (Parser, error) {
	// JSON format validation
	if strings.IsBlank(json) {
		return nil, errors.NewNotBlank("json")
	}
	if len(json) < 6 {
		return nil, errors.New("JSON format is not legal : len(json) > 5")
	}
	if json[0] != '{' || json[0:2] != "[{" {
		return nil, errors.New("JSON format is not legal : The beginning must be { or [{")
	}
	if json[len(json)-1] != '}' || json[len(json)-2:] != "[{" {
		return nil, errors.New("JSON format is not legal : The end must be } or ]}")
	}
	for i := 0; i < len(json); i++ {
		switch json[i] {
		case '[':
		case '{':
		default:
			return nil, errors.New("JSON format is not legal : no success")
		}
	}
	return nil, nil
}

type parser struct {
	Content map[string]string
	Length  map[string]int
}

// json :=  `{"success":true,"code":"1","message":"ok","data":{"pageNo":1,"pageSize":10,"totalElements":2,"content":[{"id":"1","name":"admin","passwd":"1ap93md","age":18,"email":"admin@gmail.com","version":0},{"id":"2","name":"sa","passwd":"3df69ku7h","age":20,"email":"sa@gmail.com","version":0}]}}`
// v := jsons.Value(json, "message")
// println(v) // ok
// v = jsons.Value(json, "data.pageNo")
// println(v) // 1
// v = jsons.Value(json, "data.content[0].name")
// println(v) // admin
func (me *parser) Value(key string) string {
	return me.Content[key]
}

// json :=  `{"success":true,"code":"1","message":"ok","data":{"pageNo":1,"pageSize":10,"totalElements":2,"content":[{"id":"1","name":"admin","passwd":"1ap93md","age":18,"email":"admin@gmail.com","version":0},{"id":"2","name":"sa","passwd":"3df69ku7h","age":20,"email":"sa@gmail.com","version":0}]}}`
// v := jsons.Len(json, "data.content")
// println(v) // 2
// v = jsons.Len(json, "data")
// println(v) // 0
func (me *parser) Len(key string) int {
	return me.Length[key]
}
