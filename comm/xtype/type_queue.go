// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xtype

// Queue FIFO : First In First Out
type Queue []interface{}

// Len returns the lenght of Queue.
func (me Queue) Len() int {
	return len(me)
}

// Push push value.
func (me *Queue) Push(v interface{}) {
	*me = append(*me, v)
}

// Pop pop value.
func (me *Queue) Pop() interface{} {
	old := *me
	if len(old) == 0 {
		return nil
	}
	s := old[0]
	*me = old[1:]
	return s
}
