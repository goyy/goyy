// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*
Package validate implements validate utility functions.

Usage

	err := validate.Required("ABc123")
	err = validate.Min("30", 20);
	err = validate.Max("30", 50);
	err = validate.Float("1.23");
	err = validate.Integer("123");
	err = validate.Alphanumeric("ABc123");
	err = validate.Alphabetic("ABc");
	err = validate.Email("admin@goyy.org");
	err = validate.URL("ftp://admin:123456@goyy.org/");
*/
package validate
