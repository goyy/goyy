// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package schema

type Column interface {
	Table() Table
	Name() string
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
