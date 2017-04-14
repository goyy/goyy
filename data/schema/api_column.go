// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package schema

// Column schema.Column.
type Column interface {
	Table() Table
	Name() string
	Comment() string
	Dict() string
	SetDict(value string)
	IsPrimary() bool
	IsVersion() bool
	IsDeletion() bool
	IsCreater() bool
	IsCreated() bool
	IsModifier() bool
	IsModified() bool
	IsTransient() bool
	String() string
}

// ParseDict get dictionary values by genre and mkey.
var ParseDict func(genre, mkey string) string
