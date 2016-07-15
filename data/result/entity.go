// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package result

import (
	"bytes"
	"strconv"

	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/util/errors"
	"gopkg.in/goyy/goyy.v0/util/jsons"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

type Entity struct {
	Success bool             `json:"success"`
	Token   string           `json:"token"`
	Code    string           `json:"code"`
	Message string           `json:"message"`
	Memo    string           `json:"memo"`
	Tag     string           `json:"tag"`
	Data    entity.Interface `json:"data"`
}

func (me *Entity) JSON() string {
	var b bytes.Buffer
	b.WriteString(`{"success":` + strconv.FormatBool(me.Success) + ",")
	b.WriteString(`"token":"` + jsons.Format(me.Token) + `",`)
	b.WriteString(`"code":"` + jsons.Format(me.Code) + `",`)
	b.WriteString(`"message":"` + jsons.Format(me.Message) + `",`)
	b.WriteString(`"memo":"` + jsons.Format(me.Memo) + `",`)
	b.WriteString(`"tag":"` + jsons.Format(me.Tag) + `",`)
	b.WriteString(`"data":` + entity.FormatJSON(me.Data) + "}")
	return b.String()
}

func (me *Entity) ParseJSON(json string) error {
	if strings.IsBlank(json) {
		return nil
	}
	// JSON format validation
	if !strings.HasPrefix(json, `{"success":`) || !strings.HasSuffix(json, "}}") {
		return errors.New("JSON format is not legal : no success")
	}
	if !strings.Contains(json, `,"data":`) {
		return errors.New("JSON format is not legal : no data")
	}
	// result info
	if "true" == strings.Between(json, `"success":`, `,"`) {
		me.Success = true
	}
	content := jreplace(json)
	me.Code = jparse(strings.Between(content, `"code":"`, `",`))
	me.Message = jparse(strings.Between(content, `"message":"`, `",`))
	me.Memo = jparse(strings.Between(content, `"memo":"`, `",`))
	me.Tag = jparse(strings.Between(content, `"tag":"`, `",`))
	datajson := strings.After(json[:len(json)-1], `,"data":`)
	err := entity.ParseJSON(me.Data, datajson)
	return err
}
