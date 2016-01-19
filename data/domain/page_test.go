// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package domain_test

import (
	"testing"

	"gopkg.in/goyy/goyy.v0/data/domain"
)

func TestNewPage(t *testing.T) {
	s := []struct {
		pageNo, pageSize, totalElements, totalPages int
		hasPrevious, hasNext, isFirst, isLast       bool
	}{
		{1, 10, 5, 1, false, false, true, true},
		{1, 10, 10, 1, false, false, true, true},
		{1, 10, 11, 2, false, true, true, false},
		{2, 10, 11, 2, true, false, false, true},
		{2, 10, 55, 6, true, true, false, false},
		{2, 10, 60, 6, true, true, false, false},
		{6, 10, 60, 6, true, false, false, true},
	}
	for _, v := range s {
		pageable := domain.NewPageable(v.pageNo, v.pageSize)
		out := domain.NewPage(pageable, nil, v.totalElements)
		if out.TotalPages() != v.totalPages {
			t.Errorf(`domain.NewPage().TotalPages():"%v", want:"%v"`, out.TotalPages(), v.totalPages)
		}
		if out.HasPrevious() != v.hasPrevious {
			t.Errorf(`domain.NewPage().HasPrevious():"%v", want:"%v"`, out.HasPrevious(), v.hasPrevious)
		}
		if out.HasNext() != v.hasNext {
			t.Errorf(`domain.NewPage().HasNext():"%v", want:"%v"`, out.HasNext(), v.hasNext)
		}
		if out.IsFirst() != v.isFirst {
			t.Errorf(`domain.NewPage().IsFirst():"%v", want:"%v"`, out.IsFirst(), v.isFirst)
		}
		if out.IsLast() != v.isLast {
			t.Errorf(`domain.NewPage().IsLast():"%v", want:"%v"`, out.IsLast(), v.isLast)
		}
	}
}
