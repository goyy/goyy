// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xtype

type IntStack []int

func (me IntStack) Len() int {
	return len(me)
}

func (me *IntStack) Push(x interface{}) {
	*me = append(*me, x.(int))
}

func (me *IntStack) Pop() interface{} {
	old := *me
	n := len(old)
	x := old[n-1]
	*me = old[0 : n-1]
	return x
}
