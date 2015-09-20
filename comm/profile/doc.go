// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*
Package profile implements profile utility functions.

Usage

	profile.Actives()
	profile.SetActives(profile.PROD, profile.TEST)
	profile.Accepts(profile.PROD, profile.DEV)
	profile.Accepts("tmp", "!development")
*/
package profile
