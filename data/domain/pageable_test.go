// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package domain_test

import (
	"gopkg.in/goyy/goyy.v0/data/domain"
	"testing"
)

func TestNewPageable(t *testing.T) {
	s := []struct {
		pageNo, pageSize, got, expected int
	}{
		{1, 10, 1, 10},
		{2, 10, 2, 10},
		{-3, -5, 1, 10},
		{5, 600, 5, 500},
	}
	for _, v := range s {
		out := domain.NewPageable(v.pageNo, v.pageSize)
		if out.PageNo() != v.got {
			t.Errorf(`domain.NewPageable().PageNo():"%v", want:"%v"`, out.PageNo(), v.got)
		}
		if out.PageSize() != v.expected {
			t.Errorf(`domain.NewPageable().PageSize():"%v", want:"%v"`, out.PageSize(), v.expected)
		}
	}
}
