// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package domain

import (
	"gopkg.in/goyy/goyy.v0/util/strings"
)

// pagination abstract interface for pagination information.
type pagination struct {
	length int    // Page display length
	slider int    // Front page shows the length
	pageFn string // Click the name of the page set js function call, the default is page
}

// Length returns the display length of page to be returned.
// @return the display length of that page
func (me *pagination) Length() int {
	if me.length < 1 {
		return defaultPageLength
	}
	return me.length
}

// SetLength sets the display length of page to be returned.
func (me *pagination) SetLength(length int) {
	if length < 1 {
		me.length = defaultPageLength
	} else {
		me.length = length
	}
}

// Slider returns the front page shows the length to be returned.
// @return the front page shows the length of that page
func (me *pagination) Slider() int {
	if me.slider < 1 {
		return defaultPageSlider
	}
	return me.slider
}

// SetSlider sets the front page shows the length to be returned.
func (me *pagination) SetSlider(slider int) {
	if slider < 1 {
		me.slider = defaultPageSlider
	} else {
		me.slider = slider
	}
}

// PageFn click the name of the page to get js function calls.
func (me *pagination) PageFn() string {
	if strings.IsBlank(me.pageFn) {
		return defaultPageFn
	}
	return me.pageFn
}

// SetPageFn click the name of the page set js function call, the default is page,
// use a page has multiple tabs when objects
func (me *pagination) SetPageFn(pageFn string) {
	if strings.IsBlank(pageFn) {
		me.pageFn = defaultPageFn
	} else {
		me.pageFn = pageFn
	}
}
