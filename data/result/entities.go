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

type Entities struct {
	Success bool              `json:"success"`
	Token   string            `json:"token"`
	Code    string            `json:"code"`
	Message string            `json:"message"`
	Memo    string            `json:"memo"`
	Tag     string            `json:"tag"`
	Data    entity.Interfaces `json:"data"`
}

func (me *Entities) JSON() string {
	var b bytes.Buffer
	b.WriteString(`{"success":` + strconv.FormatBool(me.Success) + ",")
	b.WriteString(`"token":"` + jsons.Format(me.Token) + `",`)
	b.WriteString(`"code":"` + jsons.Format(me.Code) + `",`)
	b.WriteString(`"message":"` + jsons.Format(me.Message) + `",`)
	b.WriteString(`"memo":"` + jsons.Format(me.Memo) + `",`)
	b.WriteString(`"tag":"` + jsons.Format(me.Tag) + `",`)
	b.WriteString(`"data":[`)
	if me.Data != nil {
		for i := 0; i < me.Data.Len(); i++ {
			if i > 0 {
				b.WriteString(",")
			}
			b.WriteString(entity.FormatJSON(me.Data.Index(i)))
		}
	}
	b.WriteString("]}")
	return b.String()
}

func (me *Entities) ParseJSON(json string) error {
	if strings.IsBlank(json) {
		return nil
	}
	// JSON format validation
	if !strings.HasPrefix(json, `{"success":`) || !strings.HasSuffix(json, "]}") {
		return errors.New("JSON format is not legal : no success")
	}
	if !strings.Contains(json, `,"data":[`) {
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
	datajson := strings.Between(json, `,"data":[{`, "}]}")
	datas := strings.Split(datajson, "},{")
	for _, data := range datas {
		e := me.Data.New()
		if err := entity.ParseJSON(e, "{"+data+"}"); err != nil {
			return err
		}
		me.Data.Append(e)
	}
	return nil
}
