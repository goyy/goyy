// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package map_test

import (
	"testing"
)

type Map struct {
	Key   int
	Value int
}

type Maps struct {
	Data []*Map
}

func (me *Maps) Get(key int) int {
	for _, v := range me.Data {
		if key == v.Key {
			return v.Value
		}
	}
	return 0
}

var dmap map[int]int
var dslice Maps

func TestResult(t *testing.T) {
	//PASS
	//BenchmarkMapSet-4  	10000000	       180 ns/op
	//BenchmarkSliceSet-4	50000000	        76.5 ns/op
	//BenchmarkMapGet-4  	20000000	       110 ns/op
	//BenchmarkSliceGet-4	  300000	    216905 ns/op
	//ok  	gopkg.in/goyy/goyy.v0/test/bench/map	91.263s
}

func BenchmarkMapSet(b *testing.B) {
	dmap = make(map[int]int, b.N)
	for i := 0; i < b.N; i++ {
		dmap[i] = i
	}
}

func BenchmarkSliceSet(b *testing.B) {
	dslice.Data = make([]*Map, b.N)
	for i := 0; i < b.N; i++ {
		v := &Map{i, i}
		dslice.Data[i] = v
	}
}

func BenchmarkMapGet(b *testing.B) {
	var count int
	for i := 0; i < b.N; i++ {
		count = dmap[i]
	}
	count++
}

func BenchmarkSliceGet(b *testing.B) {
	var count int
	for i := 0; i < b.N; i++ {
		count = dslice.Get(i)
	}
	count++
}
