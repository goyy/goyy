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
	//BenchmarkConcat-4 	30000000	        39.9 ns/op
	//BenchmarkReplace-4	 5000000	       253 ns/op
	//BenchmarkSprintf-4	 5000000	       339 ns/op
	//ok  	gopkg.in/goyy/goyy.v0/test/bench/string	5.077s
}

func BenchmarkConcat(b *testing.B) {
	var b1 bytes.Buffer
	for i := 0; i < b.N; i++ {
		b1.WriteString(`"key":` + "value")
	}
}

func BenchmarkReplace(b *testing.B) {
	var b1 bytes.Buffer
	for i := 0; i < b.N; i++ {
		b1.WriteString(strings.Replace(`"key":%s`, `%s`, "value", -1))
	}
}

func BenchmarkSprintf(b *testing.B) {
	var b2 bytes.Buffer
	for i := 0; i < b.N; i++ {
		b2.WriteString(fmt.Sprintf(`"key":%s`, "value"))
	}
}
