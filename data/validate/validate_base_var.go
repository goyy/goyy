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
		typMobile:      regexp.MustCompile(`^((\+86)|(86))?(1[3,5,7,8]\d{9})$`),
		typTel:         regexp.MustCompile(`^(0\d{2,3}(\-)?)?\d{7,8}$`),
		typZipcode:     regexp.MustCompile(`^\d{6}$`),
	}

	// Validation errors.
	Messages = map[string]string{
		typRequired:    "Expecting a non empty value.",
		typMin:         "Expecting a value greater than or equal to %d.",
		typMax:         "Expecting a value less than or equal to %d.",
		typRange:       "Expecting a value between %d and %d.",
		typMinf:        "Expecting a value greater than or equal to %g.",
		typMaxf:        "Expecting a value less than or equal to %g.",
		typRangef:      "Expecting a value between %g and %g.",
		typMinlen:      "Expecting at least %d characters.",
		typMaxlen:      "Expecting no more than %d characters.",
		typRangelen:    "Expecting a value between %d and %d characters long.",
		typFloat:       "Expecting a floating point number.",
		typInteger:     "Expecting an integer.",
		typAlpha:       "Expecting an alphabetic string.",
		typAlrod:       "Expecting an alphabetic or rod string.",
		typAlnum:       "Expecting an alphanumeric string.",
		typAlnumrod:    "Expecting an alphanumeric or rod string.",
		typAlnumhan:    "Expecting an alphanumeric or chinese string.",
		typAlnumhanrod: "Expecting an alphanumeric or chinese or rod string.",
		typAlhan:       "Expecting an alphabetic or chinese string.",
		typAlhanrod:    "Expecting an alphabetic or chinese or rod string.",
		typHan:         "Expecting an chinese string.",
		typHanrod:      "Expecting an chinese or rod string.",
		typMatch:       "Expecting a value to match s%.",
		typNomatch:     "Expecting a value that does not match s%.",
		typEmail:       "Expecting an e-mail.",
		typURL:         "Expecting an URL.",
		typIP:          "Expecting a ip address",
		typMobile:      "Expecting a mobile number",
		typTel:         "Expecting a telephone number",
		typPhone:       "Expecting a telephone or mobile phone number",
		typZipcode:     "Expecting an zipcode",
	}
)
