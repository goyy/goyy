// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*
Package schema implements db schema utility functions.

Notice

Table and column names in lowercase only support

Usage

	USER          = schema.TABLE("user")
	USER_ID       = USER.KEY("id")
	USER_VERSION  = USER.VERSION("version")
	USER_DELETION = USER.DELETION("deletion")
	USER_NAME     = USER.COLUMN("name")
*/
package schema
