// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xtype

// Principal the subject's principal
type Principal struct {
	Id          string
	Name        string
	LoginName   string
	LoginTime   string
	Permissions string
	Roles       struct {
		Func string
		Data string
	}
}
