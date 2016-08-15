// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strings_test

import (
	"fmt"
	"testing"

	"gopkg.in/goyy/goyy.v0/util/strings"
)

func TestFieldsSpace(t *testing.T) {
	s := []struct{ in, expected string }{
		{`a  b    c d`, `[a b c d]`},
		{`a`, `[a]`},
		{`   a  `, `[a]`},
		{`  `, `[]`},
		{`a 'b   c'`, `[a 'b   c']`},
		{`a 'b"  c'`, `[a 'b"  c']`},
		{`a  b    c d 'e   f'`, `[a b c d 'e   f']`},
		{`a  b    c d 'e "  f'`, `[a b c d 'e "  f']`},
		{`a  b    c d "e '  f"`, `[a b c d "e '  f"]`},
		{`id="div" class="nav"`, `[id="div" class="nav"]`},
		{`id="div" class='nav'`, `[id="div" class='nav']`},
		{"id=`div` class='nav'", "[id=`div` class='nav']"},
		{`id="div"`, `[id="div"]`},
		{`   class="nav"   data-attr='id="div"  class="nav"' `, `[class="nav" data-attr='id="div"  class="nav"']`},
	}
	for _, v := range s {
		if out := strings.FieldsSpace(v.in); fmt.Sprint(out) != v.expected {
			t.Errorf(`strings.FieldsSpace("%v") = "%v"; want "%v"`, v.in, out, v.expected)
		}
	}
}
