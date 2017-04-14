// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xsql

import (
	"gopkg.in/goyy/goyy.v0/data/domain"
	"gopkg.in/goyy/goyy.v0/data/entity"
)

// add space string:{before is add:yes, after is add:no}
type sifter struct {
	db    *db
	sifts []domain.Sift
}

// Retrieve a list of mapped entities from the dql and args.
func (me *sifter) Rows(out entity.Interfaces) error {
	dql, args, err := me.db.dql.SelectListBySift(out.New(), me.sifts...)
	if err != nil {
		logger.Debug(err.Error())
		return err
	}
	err = me.db.Query(dql, args...).Rows(out)
	if err != nil {
		logger.Debug(err.Error())
		return err
	}
	return nil
}

// Retrieve a single row mapped from the dql and args.
func (me *sifter) Row(out entity.Interface) error {
	dql, args, err := me.db.dql.SelectListBySift(out, me.sifts...)
	if err != nil {
		logger.Debug(err.Error())
		return err
	}
	err = me.db.Query(dql, args...).Row(out)
	if err != nil {
		logger.Debug(err.Error())
		return err
	}
	return nil
}

// Retrieve a int mapped from the dql and args.
func (me *sifter) Count(out entity.Interface) (int, error) {
	dql, args, err := me.db.dql.SelectCountBySift(out, me.sifts...)
	if err != nil {
		logger.Debug(err.Error())
		return 0, err
	}
	return me.db.Query(dql, args...).Int()
}

// Retrieve a single row mapped from the dql and args.
func (me *sifter) Page(content entity.Interfaces, pageable domain.Pageable) (domain.Page, error) {
	dql, args, err := me.db.dql.SelectListBySift(content.New(), me.sifts...)
	if err != nil {
		logger.Debug(err.Error())
		return nil, err
	}
	page, err := me.db.Query(dql, args...).Page(content, pageable)
	if err != nil {
		logger.Debug(err.Error())
		return nil, err
	}
	return page, nil
}
