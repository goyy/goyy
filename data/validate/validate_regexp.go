// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package validate

import (
	"regexp"
)

// Validation rules
var rules = map[string]*regexp.Regexp{
	typFloat:       regexp.MustCompile(`^[0-9\.]+$`),
	typInteger:     regexp.MustCompile(`^[0-9]+$`),
	typAlpha:       regexp.MustCompile(`^[a-zA-Z]+$`),
	typAlrod:       regexp.MustCompile(`^[a-zA-Z\-_]+$`),
	typAlnum:       regexp.MustCompile(`^[a-zA-Z0-9]+$`),
	typAlnumrod:    regexp.MustCompile(`^[a-zA-Z0-9\-_]+$`),
	typAlnumhan:    regexp.MustCompile(`^[a-zA-Z0-9\p{Han}]+$`),
	typAlnumhanrod: regexp.MustCompile(`^[a-zA-Z0-9\p{Han}\-_]+$`),
	typAlhan:       regexp.MustCompile(`^[a-zA-Z\p{Han}]+$`),
	typAlhanrod:    regexp.MustCompile(`^[a-zA-Z\p{Han}\-_]+$`),
	typHan:         regexp.MustCompile(`^[\p{Han}]+$`),
	typHanrod:      regexp.MustCompile(`^[\p{Han}\-_]+$`),
	typEmail:       regexp.MustCompile(`^[a-zA-Z0-9\+\-\.]+@[a-zA-Z0-9\.\-]+$`),
	typURL:         regexp.MustCompile(`^[a-zA-Z0-9]+:\/\/.+`),
	typIP:          regexp.MustCompile(`^(([01]?[0-9]{1,2}|2[0-4][0-9]|25[0-5])(\.([01]?[0-9]{1,2}|2[0-4][0-9]|25[0-5])){3}|([0-9a-fA-F]{1,4}:)+:?([0-9a-fA-F]{1,4}:)*[0-9a-fA-F]{1,4})$`),
	typMobile:      regexp.MustCompile(`^((\+86)|(86))?(1[3,4,5,7,8]\d{9})$`),
	typTel:         regexp.MustCompile(`^(0\d{2,3}(\-)?)?\d{7,8}$`),
	typZipcode:     regexp.MustCompile(`^\d{6}$`),
}
