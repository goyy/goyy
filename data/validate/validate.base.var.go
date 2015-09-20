// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package validate

import (
	"regexp"
)

var (
	// Validation rules
	Rules = map[string]*regexp.Regexp{
		typFloat:        regexp.MustCompile(`^[0-9\.]+$`),
		typInteger:      regexp.MustCompile(`^[0-9]+$`),
		typAlphanumeric: regexp.MustCompile(`^[a-zA-Z0-9]+$`),
		typAlphabetic:   regexp.MustCompile(`^[a-zA-Z]+$`),
		typEmail:        regexp.MustCompile(`^[a-zA-Z0-9\+\-\.]+@[a-zA-Z0-9\.\-]+$`),
		typURL:          regexp.MustCompile(`^[a-zA-Z0-9]+:\/\/.+`),
		typIP:           regexp.MustCompile(`^(([01]?[0-9]{1,2}|2[0-4][0-9]|25[0-5])(\.([01]?[0-9]{1,2}|2[0-4][0-9]|25[0-5])){3}|([0-9a-fA-F]{1,4}:)+:?([0-9a-fA-F]{1,4}:)*[0-9a-fA-F]{1,4})$`),
		typMobile:       regexp.MustCompile(`^((\\+86)|(86))?(1(([35][0-9])|(47)|[8][012356789]))\\d{8}$`),
		typTel:          regexp.MustCompile(`^(0\\d{2,3}(\\-)?)?\\d{7,8}$`),
		typZipcode:      regexp.MustCompile(`^[1-9]\\d{5}$`),
	}

	// Validation errors.
	Messages = map[string]string{
		typRequired:     "Expecting a non empty value.",
		typMin:          "Expecting a value greater than or equal to %d.",
		typMax:          "Expecting a value less than or equal to %d.",
		typRange:        "Expecting a value between %d and %d.",
		typMinlen:       "Expecting at least %d characters.",
		typMaxlen:       "Expecting no more than %d characters.",
		typRangelen:     "Expecting a value between %d and %d characters long.",
		typFloat:        "Expecting a floating point number.",
		typInteger:      "Expecting an integer.",
		typAlphanumeric: "Expecting an alphanumeric string.",
		typAlphabetic:   "Expecting an alphabetic string.",
		typEmail:        "Expecting an e-mail.",
		typURL:          "Expecting an URL.",
		typIP:           "Expecting a ip address",
		typMobile:       "Expecting a mobile number",
		typTel:          "Expecting a telephone number",
		typPhone:        "Expecting a telephone or mobile phone number",
		typZipcode:      "Expecting an zipcode",
	}
)
