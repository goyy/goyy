// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controller

import (
	"strconv"

	"gopkg.in/goyy/goyy.v0/comm/xtype"
	"gopkg.in/goyy/goyy.v0/data/domain"
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/result"
	"gopkg.in/goyy/goyy.v0/data/service"
	"gopkg.in/goyy/goyy.v0/util/errors"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/util/times"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

type baseController struct {
	sqlOfIndex    string
	isSqlOfIndex  bool
	sqlOfExport   string
	isSqlOfExport bool
}

func (me *baseController) SetSqlOfIndex(sql string) {
	if strings.IsNotBlank(sql) {
		me.isSqlOfIndex = true
		me.sqlOfIndex = sql
	}
}

func (me *baseController) SetSqlOfExport(sql string) {
	if strings.IsNotBlank(sql) {
		me.isSqlOfExport = true
		me.sqlOfExport = sql
	}
}

func (me *baseController) Index(c xhttp.Context, mgr service.Service, pre func(c xhttp.Context) error, post func(c xhttp.Context, r *result.Page) error) (out *result.Page, err error) {
	if pre != nil {
		if err = pre(c); err != nil {
			return
		}
	}
	out, err = me.Page(c, mgr)
	if err != nil {
		return
	}
	if post != nil {
		if err = post(c, out); err != nil {
			return
		}
	}
	return
}

func (me *baseController) Show(c xhttp.Context, mgr service.Service, pre func(c xhttp.Context) error, post func(c xhttp.Context, r *result.Entity) error) (out *result.Entity, err error) {
	if pre != nil {
		if err = pre(c); err != nil {
			return
		}
	}
	sift, _ := domain.NewSift(siftId, c.Param(siftIdTR))
	data := mgr.NewEntity()
	err = mgr.SelectOneBySift(data, sift)
	if err != nil {
		return
	}
	err = c.Bind(data)
	if err != nil {
		return
	}
	out = &result.Entity{Success: true, Data: data}
	if post != nil {
		if err = post(c, out); err != nil {
			return
		}
	}
	return
}

func (me *baseController) Add(c xhttp.Context, mgr service.Service, pre func(c xhttp.Context) error, post func(c xhttp.Context, r *result.Entity) error) (out *result.Entity, err error) {
	if pre != nil {
		if err = pre(c); err != nil {
			return
		}
	}
	data := mgr.NewEntity()
	err = c.Bind(data)
	if err != nil {
		return
	}
	out = &result.Entity{Success: true, Data: data}
	if post != nil {
		if err = post(c, out); err != nil {
			return
		}
	}
	return
}

func (me *baseController) Edit(c xhttp.Context, mgr service.Service, pre func(c xhttp.Context) error, post func(c xhttp.Context, r *result.Entity) error) (out *result.Entity, err error) {
	if pre != nil {
		if err = pre(c); err != nil {
			return
		}
	}
	sift, _ := domain.NewSift(siftId, c.Param(siftIdTR))
	data := mgr.NewEntity()
	err = mgr.SelectOneBySift(data, sift)
	if err != nil {
		return
	}
	err = c.Bind(data)
	if err != nil {
		return
	}
	out = &result.Entity{Success: true, Data: data}
	if post != nil {
		if err = post(c, out); err != nil {
			return
		}
	}
	return
}

func (me *baseController) Save(c xhttp.Context, mgr service.Service, pre func(c xhttp.Context) error, post func(c xhttp.Context, r *result.Entity) error) (out *result.Entity, err error) {
	if pre != nil {
		if err = pre(c); err != nil {
			return
		}
	}
	data := mgr.NewEntity()
	err = c.Bind(data)
	if err != nil {
		return
	}
	err = mgr.Save(c, data)
	if err != nil {
		return
	}
	out = &result.Entity{Success: true, Data: data}
	if post != nil {
		if err = post(c, out); err != nil {
			return
		}
	}
	return
}

func (me *baseController) Disable(c xhttp.Context, mgr service.Service, pre func(c xhttp.Context) error, post func(c xhttp.Context, r *result.Entity) error) (out *result.Entity, err error) {
	if pre != nil {
		if err = pre(c); err != nil {
			return
		}
	}
	// Disable data
	sift, _ := domain.NewSift(siftId, c.Param(siftIdTR))
	data := mgr.NewEntity()
	err = mgr.SelectOneBySift(data, sift)
	if err != nil {
		return
	}
	_, err = mgr.Disable(c, data)
	if err != nil {
		return
	}
	out = &result.Entity{Success: true, Data: data}
	if post != nil {
		if err = post(c, out); err != nil {
			return
		}
	}
	return
}

func (me *baseController) Box(c xhttp.Context, mgr service.Service) (out []xtype.Box, err error) {
	datas := mgr.NewEntities()
	sifts, err := domain.NewSifts(c.Params())
	if err != nil {
		return
	}
	err = mgr.SelectListBySift(datas, sifts...)
	if err != nil {
		return
	}
	out = make([]xtype.Box, 0)
	var boxId, boxName, boxTime string
	for i := 0; i < datas.Len(); i++ {
		if "sys_dict" == datas.Index(i).Table().Name() {
			boxId = datas.Index(i).Get("mkey").(string)
			boxName = datas.Index(i).Get("mval").(string)
		} else {
			boxId = datas.Index(i).Get("id").(string)
			boxName = datas.Index(i).Get("name").(string)
		}
		boxTime = times.FormatUnixYYMDHMS(datas.Index(i).Get("created").(int64))
		out = append(out, xtype.Box{
			Id:   boxId,
			Name: boxName,
			Time: boxTime,
		})
	}
	return
}

func (me *baseController) Page(c xhttp.Context, mgr service.Service) (out *result.Page, err error) {
	// If the query condition does not deletion the field,
	// the default query is not logically deleted.
	if c.Param(siftDeletion) == "" {
		c.Params().Set(siftDeletion, strconv.Itoa(entity.DeletionEnable))
	}
	v := mgr.NewEntities()
	p, err := domain.NewPageableHTTP(c.ResponseWriter(), c.Request())
	if err != nil {
		return nil, err
	}
	sifts, err := domain.NewSifts(c.Params())
	if err != nil {
		return nil, err
	}
	if me.isSqlOfIndex {
		params := domain.SiftsToMap(sifts...)
		data, err := mgr.SelectPageByNamed(v, p, me.sqlOfIndex, params)
		if err != nil {
			return nil, err
		}
		return &result.Page{Success: true, Data: data}, nil
	} else {
		data, err := mgr.SelectPageBySift(v, p, sifts...)
		if err != nil {
			return nil, err
		}
		return &result.Page{Success: true, Data: data}, nil
	}
}

func (me *baseController) Export(c xhttp.Context, mgr service.Service, pre func(c xhttp.Context) error, post func(c xhttp.Context, r entity.Interfaces) error) (out entity.Interfaces, err error) {
	if pre != nil {
		if err = pre(c); err != nil {
			return
		}
	}
	// If the query condition does not deletion the field,
	// the default query is not logically deleted.
	if c.Param(siftDeletion) == "" {
		c.Params().Set(siftDeletion, strconv.Itoa(entity.DeletionEnable))
	}
	out = mgr.NewEntities()
	sifts, err := domain.NewSifts(c.Params())
	if err != nil {
		return nil, err
	}
	if me.isSqlOfExport {
		params := domain.SiftsToMap(sifts...)
		count, err := mgr.SelectCountByNamed(me.sqlOfExport, params)
		if err != nil {
			return nil, err
		}
		if count > expLimit {
			return nil, errors.New(i18N.Message("exp.limit"))
		}
		err = mgr.SelectListByNamed(out, me.sqlOfExport, params)
		if err != nil {
			return nil, err
		}
	} else {
		count, err := mgr.SelectCountBySift(sifts...)
		if err != nil {
			return nil, err
		}
		if count > expLimit {
			return nil, errors.New(i18N.Message("exp.limit"))
		}
		err = mgr.SelectListBySift(out, sifts...)
		if err != nil {
			return nil, err
		}
	}
	if post != nil {
		if err = post(c, out); err != nil {
			return
		}
	}
	return
}
