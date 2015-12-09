// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package domain

// Abstract interface for pagination information.
type pageable struct {
	pageNo   int // Current page number
	pageSize int // Page size, Is set to "-1" means no paging (paging invalid)
}

// Returns the page to be returned.
// @return the page to be returned.
func (me *pageable) PageNo() int {
	if me.pageNo < 2 {
		return defaultPageNo
	} else {
		return me.pageNo
	}
}

func (me *pageable) SetPageNo(pageNo int) {
	if pageNo < 2 {
		me.pageNo = defaultPageNo
	} else {
		me.pageNo = pageNo
	}
}

// Returns the number of items to be returned.
// @return the number of items of that page
func (me *pageable) PageSize() int {
	if me.pageSize < 1 {
		return defaultPageSize
	} else if me.pageSize > defaultPageSizeMax {
		return defaultPageSizeMax
	}
	return me.pageSize
}

func (me *pageable) SetPageSize(pageSize int) {
	if pageSize < 1 {
		me.pageSize = defaultPageSize
	} else if pageSize > defaultPageSizeMax {
		me.pageSize = defaultPageSizeMax
	} else {
		me.pageSize = pageSize
	}
}

// Returns the offset to be taken according to the underlying page and page size.
// @return the offset to be taken
func (me *pageable) Offset() int {
	return (me.PageNo() - 1) * me.PageSize()
}
