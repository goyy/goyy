// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xtype

type IntHeap []int

func (me IntHeap) Len() int {
	return len(me)
}

func (me *IntHeap) Push(x interface{}) {
	*me = append(*me, x.(int))
}

func (me *IntHeap) Pop() interface{} {
	old := *me
	n := len(old)
	x := old[n-1]
	*me = old[0 : n-1]
	return x
}
