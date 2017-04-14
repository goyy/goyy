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
	// BenchmarkMapSet-4       10000000               177 ns/op
	// BenchmarkSliceSet-4     50000000               49.1 ns/op
	// BenchmarkMapGet-4       20000000               104 ns/op
	// BenchmarkSliceGet-4       500000            521592 ns/op
	// PASS
	// ok      gopkg.in/goyy/goyy.v0/test/benchmark/map        282.871s
}

func BenchmarkMapSet(b *testing.B) {
	dmap = make(map[int]int, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dmap[i] = i
	}
}

func BenchmarkSliceSet(b *testing.B) {
	dslice.Data = make([]*Map, b.N)
	b.ResetTimer()
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
