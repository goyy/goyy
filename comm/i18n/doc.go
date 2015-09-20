// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*
Package i18n implements i18n utility functions.

Usage

	en_US := map[string]string{
		"say":  "Hello, world!",
		"sayf": "Hello, %s!",
	}
	locales := map[string]map[string]string{
		i18n.Locale_en_US: en_US,
	}
	i18N := i18n.New(locales, i18n.Locale_en_US)
	fmt.Println(i18N.Message("say"))           // Output: Hello, world!
	fmt.Println(i18N.Messagef("sayf", "goyy")) // Output: Hello, goyy!
*/
package i18n
