// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of me source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controller

import (
	"gopkg.in/goyy/goyy.v0/comm/xtype"
	"gopkg.in/goyy/goyy.v0/data/domain"
	"gopkg.in/goyy/goyy.v0/data/result"
	"gopkg.in/goyy/goyy.v0/data/service"
	"gopkg.in/goyy/goyy.v0/util/errors"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

type baseTreeController struct {
	baseController
}

func (me *baseTreeController) Index(c xhttp.Context, mgr service.Service, pre func(c xhttp.Context) error, post func(c xhttp.Context, r *result.Page) error) (out *result.Page, err error) {
	if ts, err := me.Breadcrumb(c, mgr); err == nil {
		c.SetAttribute(defaultParents, ts)
	} else {
		logger.Error(err.Error())
	}
	return me.baseController.Index(c, mgr, pre, post)
}

func (me *baseTreeController) Save(c xhttp.Context, mgr service.Service, pre func(c xhttp.Context) error, post func(c xhttp.Context, r *result.Page) error) (out *result.Page, err error) {
	if ts, err := me.Breadcrumb(c, mgr); err == nil {
		c.SetAttribute(defaultParents, ts)
	} else {
		logger.Error(err.Error())
	}
	sId := c.Param(siftIdTR)
	eId := c.Param(entityId)
	eParentId := c.Param(entityParentId)
	if strings.IsBlank(eParentId) && strings.IsBlank(eId) && strings.IsNotBlank(sId) {
		eParentId = sId
	}
	if strings.IsBlank(eParentId) {
		eParentId = defaultTreeRoot
	}
	p := mgr.NewEntity()
	p.SetString(entityId, eParentId)
	if err = mgr.Get(p); err != nil {
		return
	} else {
		c.Params().Set(entityParentId, eParentId)
		if eParentId == defaultTreeRoot {
			c.Params().Set(entityFullname, c.Param(entityName))
			c.Params().Set(entityParentIds, eParentId)
			c.Params().Set(entityParentNames, p.Get(colName).(string))
		} else {
			c.Params().Set(entityFullname, p.Get(colFullname).(string)+" - "+c.Param(entityName))
			c.Params().Set(entityParentIds, p.Get(colParentIds).(string)+","+eParentId)
			c.Params().Set(entityParentNames, p.Get(colParentNames).(string)+","+p.Get(colName).(string))
		}
	}
	return me.baseController.Save(c, mgr, pre, post)
}

func (me *baseTreeController) Disable(c xhttp.Context, mgr service.Service, pre func(c xhttp.Context) error, post func(c xhttp.Context, r *result.Page) error) (out *result.Page, err error) {
	if ts, err := me.Breadcrumb(c, mgr); err == nil {
		c.SetAttribute(defaultParents, ts)
	} else {
		logger.Error(err.Error())
	}
	return me.baseController.Disable(c, mgr, pre, post)
}

func (me *baseTreeController) Tree(c xhttp.Context, mgr service.Service) (out []xtype.Tree, err error) {
	datas := mgr.NewEntities()
	sifts, err := domain.NewSifts(c.Params())
	if err != nil {
		return
	}
	// If the query condition does not deletion the field,
	// the default query is not logically deleted.
	hasDeletion := false
	for _, sift := range sifts {
		if sift.Key() == "Deletion" {
			hasDeletion = true
		}
	}
	if !hasDeletion {
		siftDeletion, _ := domain.NewSift(siftDeletion, "0")
		sifts = append(sifts, siftDeletion)
	}
	err = mgr.SelectListBySift(datas, sifts...)
	if err != nil {
		return
	}
	out = make([]xtype.Tree, 0)
	var treeId, treePid, treeName string
	var isOpen bool
	for i := 0; i < datas.Len(); i++ {
		treeId = datas.Index(i).Get(colId).(string)
		treePid = datas.Index(i).Get(colParentId).(string)
		treeName = datas.Index(i).Get(colName).(string)
		if treeId == defaultTreeRoot {
			isOpen = true
		}
		out = append(out, xtype.Tree{
			Id:       treeId,
			ParentId: treePid,
			Name:     treeName,
			Open:     isOpen,
		})
	}
	return
}

func (me *baseTreeController) Breadcrumb(c xhttp.Context, mgr service.Service) ([]xtype.Box, error) {
	parentId := c.Param(siftParentId)
	e := mgr.NewEntity()
	if strings.IsNotBlank(parentId) {
		pk := e.Table().Primary().Name()
		if id, ok := e.Column(pk); ok {
			e.SetString(id.Name(), parentId)
			err := mgr.Get(e)
			if err != nil {
				return nil, err
			}
		}
	}
	parentIds := e.Get(colParentIds).(string)
	if strings.IsBlank(parentIds) {
		ts := make([]xtype.Box, 1)
		id := e.Get(colId).(string)
		name := e.Get(colName).(string)
		t := xtype.Box{
			Id:     id,
			Name:   name,
			Active: true,
		}
		ts[0] = t
		return ts, nil
	} else {
		parentNames := e.Get(colParentNames).(string)
		pIds := strings.Split(parentIds, ",")
		pNames := strings.Split(parentNames, ",")
		l := len(pIds)
		if l == len(pNames) {
			ts := make([]xtype.Box, l+1)
			for i, pId := range pIds {
				t := xtype.Box{
					Id:     pId,
					Name:   pNames[i],
					Active: false,
				}
				ts[i] = t
			}
			id := e.Get(colId).(string)
			name := e.Get(colName).(string)
			t := xtype.Box{
				Id:     id,
				Name:   name,
				Active: true,
			}
			ts[l] = t
			return ts, nil
		}
	}
	return nil, errors.Newf("'%s' does not match the parent node information", parentId)
}
