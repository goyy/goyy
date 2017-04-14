// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xtype

// Stack FILO : First In Last Out
type Stack []interface{}

// Len returns the lenght of Stack.
func (me Stack) Len() int {
	return len(me)
}

// Push push value.
func (me *Stack) Push(v interface{}) {
	*me = append(*me, v)
}

// Pop pop value.
func (me *Stack) Pop() interface{} {
	old := *me
	n := len(old)
	if n == 0 {
		return nil
	}
	s := old[n-1]
	*me = old[0 : n-1]
	return s
}
