// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package service

import (
	"errors"
	"gopkg.in/goyy/goyy.v0/comm/xtype"
	"gopkg.in/goyy/goyy.v0/data/domain"
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/util/times"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
	"strconv"
)

type TreeManager struct {
	Manager
}

func (me *TreeManager) Save(c xhttp.Context, e entity.Interface) error {
	if me.Pre != nil {
		me.Pre()
	}
	tx, err := me.Repository().Begin()
	if err != nil {
		return err
	}
	if strings.IsBlank(e.Get(e.Table().Primary().Name()).(string)) {
		if c != nil && c.Session().IsLogin() {
			if p, err := c.Session().Principal(); err == nil {
				e.SetString(creater, p.Id)
				e.SetString(modifier, p.Id)
			}
		}
		e.SetString(created, times.NowUnixStr())
		e.SetString(modified, times.NowUnixStr())
		_, err = me.Repository().Insert(e)
	} else {
		if c != nil && c.Session().IsLogin() {
			if p, err := c.Session().Principal(); err == nil {
				e.SetString(modifier, p.Id)
			}
		}
		e.SetString(modified, times.NowUnixStr())
		_, err = me.Repository().Update(e)
	}
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (me *TreeManager) Disable(c xhttp.Context, e entity.Interface) (int64, error) {
	if me.Pre != nil {
		me.Pre()
	}
	if strings.IsBlank(e.Get(e.Table().Primary().Name()).(string)) {
		return 0, errors.New("Gets the primary key value failed")
	}
	tx, err := me.Repository().Begin()
	if err != nil {
		return 0, err
	}
	if c != nil && c.Session().IsLogin() {
		if p, err := c.Session().Principal(); err == nil {
			e.SetString(modifier, p.Id)
			e.SetString(modified, times.NowUnixStr())
		}
	}
	r, err := me.Repository().Disable(e)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return r, nil
}

// Modify system field information : ParentIds、ParentCodes、ParentNames、fullname、leaf、grade
func (me *TreeManager) modifyParents(c xhttp.Context, e entity.Interface) {
	parentId := e.Get(defaultParentId).(string)
	if strings.IsBlank(parentId) {
		return
	}
	parent := me.NewEntity()
	if err := parent.SetString(defaultId, parentId); err != nil {
		logger.Debug(err.Error())
		return
	}
	if err := me.Get(parent); err != nil {
		logger.Debug(err.Error())
		return
	}
	// Whether the parent node is a leaf node.
	if leafYes == parent.Get(defaultLeaf).(string) {
		parent.SetString(defaultLeaf, leafNo)
		me.Save(c, parent)
	}
	parentIds := parent.Get(defaultParentIds).(string)
	parentCodes := parent.Get(defaultParentCodes).(string)
	parentNames := parent.Get(defaultParentNames).(string)
	fullname := parent.Get(defaultFullname).(string)
	if strings.IsNotBlank(fullname) {
		fullname = fullname + " - "
	}
	e.SetString(defaultParentIds, parentIds+parentId+",")
	parentCode := e.Get(defaultCode).(string)
	if strings.IsNotBlank(parentCode) {
		e.SetString(defaultParentCodes, parentCodes+parentCode+",")
	}
	parentName := e.Get(defaultName).(string)
	e.SetString(defaultParentNames, parentNames+parentName+",")
	e.SetString(defaultFullname, fullname+parentName)
	grade := e.Get(defaultGrade).(string)
	if i, err := strconv.Atoi(grade); err == nil {
		e.SetString(defaultGrade, strconv.Itoa(i+1))
	}
	// Set whether it is a leaf node.
	e.SetString(defaultLeaf, me.getLeaf(e.Get(defaultId).(string)))
}

func (me *TreeManager) modifyParents4Childs(c xhttp.Context, id string) {
	if strings.IsBlank(id) {
		return
	}
	// Query the child nodes of the current node.
	sDeletionEQ, _ := domain.NewSift(defaultDeletionSearch, defaultDeletionEnable)
	sParentIdEQ, _ := domain.NewSift(defaultParentIdSearch, id)
	childs := me.NewEntities()
	if err := me.SelectListBySift(childs, sDeletionEQ, sParentIdEQ); err != nil {
		logger.Debug(err.Error())
		return
	}
	for i := 0; i < childs.Len(); i++ {
		me.modifyParents(c, childs.Index(i))
		cid := childs.Index(i).Get(defaultId).(string)
		me.modifyParents4Childs(c, cid)
	}
}

// Modify system field information : ParentIds、ParentCodes、ParentNames、fullname、leaf、grade
func (me *TreeManager) modifyOldParents(c xhttp.Context, e entity.Interface) {
	id := e.Get(defaultId).(string)
	if strings.IsBlank(id) {
		return
	}
	current := me.NewEntity()
	if err := current.SetString(defaultId, id); err != nil {
		logger.Debug(err.Error())
		return
	}
	if err := me.Get(current); err != nil {
		logger.Debug(err.Error())
		return
	}
	newParentId := e.Get(defaultParentId).(string)
	oldParentId := current.Get(defaultParentId).(string)
	if strings.IsNotBlank(newParentId) && strings.IsNotBlank(oldParentId) && newParentId != oldParentId { // Modified parentId
		oldParent := me.NewEntity()
		if err := oldParent.SetString(defaultId, oldParentId); err != nil {
			logger.Debug(err.Error())
			return
		}
		if err := me.Get(oldParent); err != nil {
			logger.Debug(err.Error())
			return
		}
		me.modifyParents(c, oldParent)
	}
	return
}

// Whether the parent node is a leaf node.
func (me *TreeManager) setParentLeaf(c xhttp.Context, e entity.Interface) {
	parentId := e.Get(defaultParentId).(string)
	if strings.IsBlank(parentId) {
		return
	}
	parent := me.NewEntity()
	if err := parent.SetString(defaultId, parentId); err != nil {
		logger.Debug(err.Error())
		return
	}
	if err := me.Get(parent); err != nil {
		logger.Debug(err.Error())
		return
	}
	leaf := me.getLeaf(parentId)
	if leaf != parent.Get(defaultLeaf).(string) {
		parent.SetString(defaultLeaf, leaf)
		me.Save(c, parent)
	}
}

// Get the leaf value based on id.
func (me *TreeManager) getLeaf(id string) string {
	if strings.IsBlank(id) {
		return leafYes
	}
	sDeletionEQ, _ := domain.NewSift(defaultDeletionSearch, defaultDeletionEnable)
	sParentIdEQ, _ := domain.NewSift(defaultParentIdSearch, id)
	if count, err := me.SelectCountBySift(sDeletionEQ, sParentIdEQ); err == nil {
		if count > 0 {
			return leafNo
		}
	}
	return leafYes
}

// Recursive method for obtaining all the parent nodes of the current node.
func (me *TreeManager) getParents(currentId, rootId string) (parents []xtype.Tree, err error) {
	if strings.IsBlank(currentId) || strings.IsBlank(rootId) {
		return
	}
	current := me.NewEntity()
	if err = current.SetString(defaultId, currentId); err != nil {
		logger.Debug(err.Error())
		return
	}
	if err = me.Get(current); err != nil {
		logger.Debug(err.Error())
		return
	}
	parentId := current.Get(defaultParentId).(string)
	if strings.IsNotBlank(parentId) {
		parent := me.NewEntity()
		if err = parent.SetString(defaultId, parentId); err != nil {
			logger.Debug(err.Error())
			return
		}
		if err = me.Get(parent); err != nil {
			logger.Debug(err.Error())
			return
		}
		if ps, perr := me.getParents(parentId, rootId); perr != nil {
			err = perr
			logger.Debug(err.Error())
			return
		} else {
			for _, v := range ps {
				parents = append(parents, v)
			}
		}
	}
	checked := false
	if currentId == rootId {
		checked = true
	}
	tree := xtype.Tree{
		Id:      currentId,
		Name:    current.Get(defaultName).(string),
		Checked: checked,
	}
	parents = append(parents, tree)
	return
}
