// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*
Package xhttp implements web utility functions.

Usage

	xhttp.GET("/", func(ctx xhttp.Context) {
		ctx.TEXT(xhttp.StatusOK, "index")
	})
	xhttp.Run()
*/
package xhttp
