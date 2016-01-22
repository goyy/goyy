// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xtype

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (me IntHeap) Len() int           { return len(me) }
func (me IntHeap) Less(i, j int) bool { return me[i] < me[j] }
func (me IntHeap) Swap(i, j int)      { me[i], me[j] = me[j], me[i] }

func (me *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*me = append(*me, x.(int))
}

func (me *IntHeap) Pop() interface{} {
	old := *me
	n := len(old)
	if n == 0 {
		return nil
	}
	x := old[n-1]
	*me = old[0 : n-1]
	return x
}
