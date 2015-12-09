// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package domain

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
)

// A page is a sublist of a list of objects. It allows gain information about
// the position of it in the containing entire list.
type Page interface {
	Pageable
	Pagination

	// Returns the previous Pageable or the first Pageable
	// if the current one already is the first one.
	// @return
	Previous() Pageable

	// Returns the Pageable requesting the next Page.
	// @return
	Next() Pageable

	// Returns the Pageable requesting the first page.
	// @return
	First() Pageable

	// Returns the Pageable requesting the last page.
	// @return
	Last() Pageable

	// Returns whether there's a previous Pageable we can access from
	// the current one. Will return false in case the current
	// Pageable already refers to the first page.
	// @return
	HasPrevious() bool

	// Returns whether there's a next Pageable we can access from
	// the current one. Will return false in case the current
	// Pageable already refers to the last page.
	// @return
	HasNext() bool

	// Returns whether the current Pageable is the first one.
	//@return
	IsFirst() bool

	// Returns whether the current Pageable is the last one.
	// @return
	IsLast() bool

	// Returns the number of total pages.
	// @return the number of total pages
	TotalPages() int

	// Returns the total amount of elements.
	// @return the total amount of elements
	TotalElements() int

	// Returns the page content as entity.Interfaces.
	// @return
	Content() entity.Interfaces

	// Returns the page content as slice.
	// @return
	Slice() interface{}

	// Set total amount of elements.
	SetTotalElements(totalElements int)

	// Set pageable information.
	SetPageable(pageable Pageable)

	// Set pagination information.
	SetPagination(pagination Pagination)
}

// NewPage returns the Page from pageable, content, totalElements.
func NewPage(pageable Pageable, content entity.Interfaces, totalElements int) Page {
	p := &page{}
	p.SetPageNo(pageable.PageNo())
	p.SetPageSize(pageable.PageSize())
	p.SetTotalElements(totalElements)
	p.SetContent(content)
	p.SetPagination(NewPaginationDefault())
	return p
}

// NewPageDefault returns the default Page.
func NewPageDefault(content entity.Interfaces) Page {
	p := &page{}
	p.SetPageNo(defaultPageNo)
	p.SetPageSize(defaultPageSize)
	p.SetContent(content)
	p.SetPagination(NewPaginationDefault())
	return p
}
