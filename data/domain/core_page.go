// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package domain

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
)

// Abstract interface for pagination information.
type page struct {
	pageable
	pagination
	totalElements int // The total number of records, Is set to "-1" means no query total
	content       entity.Interfaces
}

// Returns the previous Pageable or the first Pageable
// if the current one already is the first one.
// @return
func (me *page) Previous() Pageable {
	return NewPageable(me.PageNo()-1, me.PageSize())
}

// Returns the Pageable requesting the next Page.
// @return
func (me *page) Next() Pageable {
	if me.PageNo() == me.TotalPages() {
		return NewPageable(me.TotalPages(), me.PageSize())
	} else {
		return NewPageable(me.PageNo()+1, me.PageSize())
	}
}

// Returns the Pageable requesting the first page.
// @return
func (me *page) First() Pageable {
	return NewPageable(defaultPageNo, me.PageSize())
}

// Returns the Pageable requesting the last page.
// @return
func (me *page) Last() Pageable {
	return NewPageable(me.TotalPages(), me.PageSize())
}

// Returns whether there's a previous Pageable we can access from
// the current one. Will return false in case the current
// Pageable already refers to the first page.
// @return
func (me *page) HasPrevious() bool {
	return me.PageNo() > defaultPageNo
}

// Returns whether there's a next Pageable we can access from
// the current one. Will return false in case the current
// Pageable already refers to the last page.
// @return
func (me *page) HasNext() bool {
	return me.PageNo() < me.TotalPages()
}

// Returns whether the current Pageable is the first one.
//@return
func (me *page) IsFirst() bool {
	return me.PageNo() == defaultPageNo
}

// Returns whether the current Pageable is the last one.
// @return
func (me *page) IsLast() bool {
	return me.PageNo() == me.TotalPages()
}

// Returns the number of total pages.
// @return the number of total pages
func (me *page) TotalPages() int {
	var totalPages int
	if me.totalElements == 0 {
		totalPages = defaultPageNo
	} else {
		last := me.totalElements / me.PageSize()
		if me.totalElements%me.PageSize() == 0 {
			totalPages = last
		} else {
			totalPages = last + 1
		}
	}
	return totalPages
}

// Returns the total amount of elements.
// @return the total amount of elements
func (me *page) TotalElements() int {
	return me.totalElements
}

// Returns the total amount of elements.
// @return the total amount of elements
func (me *page) SetTotalElements(totalElements int) {
	me.totalElements = totalElements
}

// Returns the page content as entity.Interfaces.
// @return
func (me *page) Content() entity.Interfaces {
	return me.content
}

// Returns the page content as entity.Interfaces.
// @return
func (me *page) SetContent(content entity.Interfaces) {
	me.content = content
}

// Returns the page content as slice.
// @return
func (me *page) Slice() interface{} {
	if me.content == nil {
		return nil
	}
	return me.content.Slice()
}

// Set pageable information.
func (me *page) SetPageable(pageable Pageable) {
	me.SetPageNo(pageable.PageNo())
	me.SetPageSize(pageable.PageSize())
}

// Set pagination information.
func (me *page) SetPagination(pagination Pagination) {
	me.SetLength(pagination.Length())
	me.SetSlider(pagination.Slider())
	me.SetPageFn(pagination.PageFn())
}
