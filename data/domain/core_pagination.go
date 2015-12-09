// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package domain

import (
	"gopkg.in/goyy/goyy.v0/util/strings"
)

// Abstract interface for pagination information.
type pagination struct {
	length int    // Page display length
	slider int    // Front page shows the length
	pageFn string // Click the name of the page set js function call, the default is page
}

// Returns the display length of page to be returned.
// @return the display length of that page
func (me *pagination) Length() int {
	if me.length < 1 {
		return defaultPageLength
	} else {
		return me.length
	}
}

func (me *pagination) SetLength(length int) {
	if length < 1 {
		me.length = defaultPageLength
	} else {
		me.length = length
	}
}

// Returns the front page shows the length to be returned.
// @return the front page shows the length of that page
func (me *pagination) Slider() int {
	if me.slider < 1 {
		return defaultPageSlider
	} else {
		return me.slider
	}
}

func (me *pagination) SetSlider(slider int) {
	if slider < 1 {
		me.slider = defaultPageSlider
	} else {
		me.slider = slider
	}
}

// Click the name of the page to get js function calls
func (me *pagination) PageFn() string {
	if strings.IsBlank(me.pageFn) {
		return defaultPageFn
	} else {
		return me.pageFn
	}
}

// Click the name of the page set js function call, the default is page,
// use a page has multiple tabs when objects
func (me *pagination) SetPageFn(pageFn string) {
	if strings.IsBlank(pageFn) {
		me.pageFn = defaultPageFn
	} else {
		me.pageFn = pageFn
	}
}
