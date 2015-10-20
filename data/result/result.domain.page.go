// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package result

import (
	"bytes"
	"gopkg.in/goyy/goyy.v0/data/domain"
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/util/errors"
	"gopkg.in/goyy/goyy.v0/util/jsons"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"strconv"
)

type Page struct {
	Success bool        `json:"success"`
	Id      string      `json:"id"`
	Token   string      `json:"token"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Memo    string      `json:"memo"`
	Tag     string      `json:"tag"`
	Data    domain.Page `json:"data"`
}

func (me *Page) JSON() string {
	var b bytes.Buffer
	b.WriteString(`{"success":` + strconv.FormatBool(me.Success) + ",")
	b.WriteString(`"id":"` + jsons.Format(me.Id) + `",`)
	b.WriteString(`"token":"` + jsons.Format(me.Token) + `",`)
	b.WriteString(`"code":"` + jsons.Format(me.Code) + `",`)
	b.WriteString(`"message":"` + jsons.Format(me.Message) + `",`)
	b.WriteString(`"memo":"` + jsons.Format(me.Memo) + `",`)
	b.WriteString(`"tag":"` + jsons.Format(me.Tag) + `",`)
	b.WriteString(`"data":{`)
	if me.Data != nil {
		b.WriteString(`"pageNo":` + strconv.Itoa(me.Data.PageNo()) + `,`)
		b.WriteString(`"pageSize":` + strconv.Itoa(me.Data.PageSize()) + `,`)
		b.WriteString(`"pageFn":"` + me.Data.PageFn() + `",`)
		b.WriteString(`"totalPages":` + strconv.Itoa(me.Data.TotalPages()) + `,`)
		b.WriteString(`"totalElements":` + strconv.Itoa(me.Data.TotalElements()) + `,`)
		b.WriteString(`"length":` + strconv.Itoa(me.Data.Length()) + `,`)
		b.WriteString(`"slider":` + strconv.Itoa(me.Data.Slider()) + `,`)
		b.WriteString(`"slice":[`)
		for i := 0; i < me.Data.Content().Len(); i++ {
			if i > 0 {
				b.WriteString(",")
			}
			b.WriteString(entity.FormatJSON(me.Data.Content().Index(i)))
		}
		b.WriteString("]")
	}
	b.WriteString("}}")
	return b.String()
}

func (me *Page) ParseJSON(json string) error {
	if strings.IsBlank(json) {
		return nil
	}
	// JSON format validation
	if !strings.HasPrefix(json, `{"success":`) || !strings.HasSuffix(json, "}]}}") {
		return errors.New("JSON format is not legal : no success")
	}
	if !strings.Contains(json, `,"data":{`) {
		return errors.New("JSON format is not legal : no data")
	}
	if !strings.Contains(json, `"pageNo":`) {
		return errors.New("JSON format is not legal : no pageNo")
	}
	if !strings.Contains(json, `"pageSize":`) {
		return errors.New("JSON format is not legal : no pageSize")
	}
	if !strings.Contains(json, `"totalElements":`) {
		return errors.New("JSON format is not legal : no totalElements")
	}
	if !strings.Contains(json, `,"slice":[{`) {
		return errors.New("JSON format is not legal : no slice")
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
	// page info
	pagejson := strings.After(json, `,"data":{`)
	pagejson = strings.BeforeLast(pagejson, "}}")
	strPageNo := strings.Between(pagejson, `"pageNo":`, `,"`)
	strPageSize := strings.Between(pagejson, `"pageSize":`, `,"`)
	strTotalElements := strings.Between(pagejson, `"totalElements":`, `,"`)
	if pageNo, err := strconv.Atoi(strPageNo); err != nil {
		if pageSize, err := strconv.Atoi(strPageSize); err != nil {
			pageable := domain.NewPageable(pageNo, pageSize)
			me.Data.SetPageable(pageable)
		}
	}
	if totalElements, err := strconv.Atoi(strTotalElements); err != nil {
		me.Data.SetTotalElements(totalElements)
	}
	// slice info
	datajson := strings.Between(json, `,"slice":[{`, "}]}}")
	datas := strings.Split(datajson, "},{")
	for _, data := range datas {
		e := me.Data.Content().New()
		if err := entity.ParseJSON(e, "{"+data+"}"); err != nil {
			return err
		}
		me.Data.Content().Append(e)
	}
	return nil
}
