// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package domain

// Abstract interface for pagination information.
type Pagination interface {
	Length() int
	Slider() int
	Function() string
}

// NewPagination returns the Pagination from length, slider, function.
func NewPagination(length, slider int, function string) Pagination {
	p := &pagination{}
	p.SetLength(length)
	p.SetSlider(slider)
	p.SetFunction(function)
	return p
}

// NewPaginationLength returns the Pagination from length.
// slider defaults to 1
// function defaults to page
func NewPaginationLength(length int) Pagination {
	return NewPagination(length, defaultPageSlider, defaultPageFunction)
}

// NewPaginationDefault returns the default Pagination.
// length defaults to 8
// slider defaults to 1
// function defaults to page
func NewPaginationDefault() Pagination {
	return NewPaginationLength(defaultPageLength)
}
