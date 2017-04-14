// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package captcha

import (
	"fmt"
	"os"
	"testing"
)

type byteCounter struct {
	n int64
}

func (bc *byteCounter) Write(b []byte) (int, error) {
	bc.n += int64(len(b))
	return len(b), nil
}

func TestNewImage(t *testing.T) {
	f, err := os.Create("/captcha.png")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	_, err = NewImage(key, RandomDigits(DefaultLen), StdWidth, StdHeight).WriteTo(f)
	if err != nil {
		fmt.Println(err)
	}
}

func BenchmarkNewImage(b *testing.B) {
	b.StopTimer()
	d := RandomDigits(DefaultLen)
	id := randomId()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		NewImage(id, d, StdWidth, StdHeight)
	}
}

func BenchmarkImageWriteTo(b *testing.B) {
	b.StopTimer()
	d := RandomDigits(DefaultLen)
	id := randomId()
	b.StartTimer()
	counter := &byteCounter{}
	for i := 0; i < b.N; i++ {
		img := NewImage(id, d, StdWidth, StdHeight)
		img.WriteTo(counter)
		b.SetBytes(counter.n)
		counter.n = 0
	}
}
