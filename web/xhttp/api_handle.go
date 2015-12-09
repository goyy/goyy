// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

// Handle is a function that can be registered to a route to handle HTTP
// requests. Like http.HandleFunc.
type Handle func(c Context)
