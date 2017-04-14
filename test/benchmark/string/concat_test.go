// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package string_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestResult(t *testing.T) {
	// BenchmarkConcat-4         1000000            132598 ns/op
	// BenchmarkSprintf-4         500000             75239 ns/op
	// BenchmarkJoin-4         100000000              25.7 ns/op
	// BenchmarkBuffer-4       100000000              18.6 ns/op
	// BenchmarkCopy-4         300000000              5.90 ns/op
	// PASS
	// ok      gopkg.in/goyy/goyy.v0/test/benchmark/string 180.231s
}

func BenchmarkConcat(b *testing.B) {
	var str string
	for n := 0; n < b.N; n++ {
		str += "x"
	}
	b.StopTimer()

	if s := strings.Repeat("x", b.N); str != s {
		b.Errorf("unexpected result; got=%s, want=%s", str, s)
	}
}

func BenchmarkSprintf(b *testing.B) {
	var str string
	for n := 0; n < b.N; n++ {
		str = fmt.Sprintf("%s%s", str, "x")
	}
	b.StopTimer()

	if s := strings.Repeat("x", b.N); str != s {
		b.Errorf("unexpected result; got=%s, want=%s", str, s)
	}
}

func BenchmarkJoin(b *testing.B) {
	bs := make([]string, b.N)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		bs[n] = "x"
	}
	str := strings.Join(bs, "")
	b.StopTimer()

	if s := strings.Repeat("x", b.N); str != s {
		b.Errorf("unexpected result; got=%s, want=%s", str, s)
	}
}

func BenchmarkBuffer(b *testing.B) {
	var buffer bytes.Buffer
	for n := 0; n < b.N; n++ {
		buffer.WriteString("x")
	}
	b.StopTimer()

	if s := strings.Repeat("x", b.N); buffer.String() != s {
		b.Errorf("unexpected result; got=%s, want=%s", buffer.String(), s)
	}
}

func BenchmarkCopy(b *testing.B) {
	bs := make([]byte, b.N)
	bl := 0

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		bl += copy(bs[bl:], "x")
	}
	b.StopTimer()

	if s := strings.Repeat("x", b.N); string(bs) != s {
		b.Errorf("unexpected result; got=%s, want=%s", string(bs), s)
	}
}
