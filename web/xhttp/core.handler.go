// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

type Handler func(Context)
type Handlers []Handler

func (me Handlers) Last() Handler {
	length := len(me)
	if length > 0 {
		return me[length-1]
	}
	return nil
}
