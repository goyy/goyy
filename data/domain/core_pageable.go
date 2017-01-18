// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package domain

// pageable abstract interface for pagination information.
type pageable struct {
	pageNo   int // Current page number
	pageSize int // Page size, Is set to "-1" means no paging (paging invalid)
}

// PageNo returns the page to be returned.
// @return the page to be returned.
func (me *pageable) PageNo() int {
	if me.pageNo < 2 {
		return defaultPageNo
	}
	return me.pageNo
}

// SetPageNo sets the page to be returned.
func (me *pageable) SetPageNo(pageNo int) {
	if pageNo < 2 {
		me.pageNo = defaultPageNo
	} else {
		me.pageNo = pageNo
	}
}

// PageSize returns the number of items to be returned.
// @return the number of items of that page
func (me *pageable) PageSize() int {
	if me.pageSize < 1 {
		return defaultPageSize
	} else if me.pageSize > defaultPageSizeMax {
		return defaultPageSizeMax
	}
	return me.pageSize
}

// SetPageSize sets the number of items to be returned.
func (me *pageable) SetPageSize(pageSize int) {
	if pageSize < 1 {
		me.pageSize = defaultPageSize
	} else if pageSize > defaultPageSizeMax {
		me.pageSize = defaultPageSizeMax
	} else {
		me.pageSize = pageSize
	}
}

// Offset returns the offset to be taken according to the underlying page and page size.
// @return the offset to be taken
func (me *pageable) Offset() int {
	return (me.PageNo() - 1) * me.PageSize()
}
