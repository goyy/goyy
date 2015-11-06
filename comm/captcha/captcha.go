// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package captcha

import (
	"io"
)

const (
	// Default number of digits in captcha solution.
	DefaultLen = 4
)

// WriteImage writes PNG-encoded image representation of the captcha with the
// given id. The image will have the given width and height.
func WriteImage(w io.Writer, id string, length, width, height int) ([]byte, error) {
	v := RandomDigits(length)
	_, err := NewImage(id, v, width, height).WriteTo(w)
	return v, err
}
