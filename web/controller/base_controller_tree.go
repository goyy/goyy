// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controller

import (
	"strconv"

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

func (me *baseTreeController) Save(c xhttp.Context, mgr service.Service, pre func(c xhttp.Context) error, post func(c xhttp.Context, r *result.Entity) error) (out *result.Entity, err error) {
	if ts, err := me.Breadcrumb(c, mgr); err == nil {
		c.SetAttribute(defaultParents, ts)
	} else {
		logger.Error(err.Error())
	}
	if err = me.setTreeInfo(c, mgr); err != nil {
		return
	}
	return me.baseController.Save(c, mgr, pre, post)
}

func (me *baseTreeController) Disable(c xhttp.Context, mgr service.Service, pre func(c xhttp.Context) error, post func(c xhttp.Context, r *result.Entity) error) (out *result.Entity, err error) {
	if ts, err := me.Breadcrumb(c, mgr); err == nil {
		c.SetAttribute(defaultParents, ts)
	} else {
		logger.Errorln(err)
	}
	if err = me.setLeafOfDisable(c, mgr); err != nil {
		logger.Errorln(err)
		return
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
		} else {
			isOpen = false
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

func (me *baseTreeController) setTreeInfo(c xhttp.Context, mgr service.Service) error {
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
	if err := mgr.Get(p); err != nil {
		logger.Errorln(err)
		return err
	} else {
		c.Params().Set(entityParentId, eParentId)
		if eParentId == defaultTreeRoot {
			if strings.IsBlank(eId) {
				c.Params().Set(entityLeaf, "1")
			}
			c.Params().Set(entityGrade, "2")
			c.Params().Set(entityFullname, c.Param(entityName))
			c.Params().Set(entityParentIds, eParentId)
			c.Params().Set(entityParentNames, p.GetString(entityName))
			c.Params().Set(entityParentCodes, p.GetString(entityCode))
		} else {
			if strings.IsBlank(eId) {
				c.Params().Set(entityLeaf, "1")
			}
			grade := p.GetString(entityGrade)
			if strings.IsNotBlank(grade) {
				if v, err := strconv.Atoi(grade); err == nil {
					c.Params().Set(entityGrade, strconv.Itoa(v+1))
				} else {
					logger.Errorln(err)
					return err
				}
			}
			c.Params().Set(entityFullname, p.Get(colFullname).(string)+" - "+c.Param(entityName))
			c.Params().Set(entityParentIds, p.Get(colParentIds).(string)+","+eParentId)
			c.Params().Set(entityParentNames, p.Get(colParentNames).(string)+","+p.Get(colName).(string))
			c.Params().Set(entityParentCodes, p.Get(colParentCodes).(string)+","+p.Get(colCode).(string))
		}
		leaf := p.GetString(entityLeaf)
		if leaf != "0" {
			p.SetString(entityLeaf, "0")
			mgr.Save(c, p)
		}
	}
	return nil
}

func (me *baseTreeController) setLeafOfDisable(c xhttp.Context, mgr service.Service) error {
	id := c.Param(siftIdTR)
	sId, _ := domain.NewSift(siftId, id)
	data := mgr.NewEntity()
	if err := mgr.SelectOneBySift(data, sId); err != nil {
		return err
	}
	if data.GetString(entityLeaf) == "0" {
		return i18N.Error("msg.disable.leaf")
	}
	parentId := data.GetString(entityParentId)
	if strings.IsNotBlank(parentId) {
		sParentId, _ := domain.NewSift(siftParentId, parentId)
		sIdNE, _ := domain.NewSift(siftIdNE, id)
		sDeletion, _ := domain.NewSift(siftDeletion, "0")
		if count, err := mgr.SelectCountBySift(sParentId, sIdNE, sDeletion); err == nil {
			if count == 0 {
				sIdParent, _ := domain.NewSift(siftId, parentId)
				pEntity := mgr.NewEntity()
				if err := mgr.SelectOneBySift(pEntity, sIdParent); err == nil {
					pEntity.SetString(entityLeaf, "1")
					mgr.Save(c, pEntity)
				} else {
					return err
				}
			}
		}
	}
	return nil
}
