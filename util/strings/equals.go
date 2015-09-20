// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strings

import (
	"strings"
)

// EqualFold reports whether s and t, interpreted as UTF-8 strings,
// are equal under Unicode case-folding.
func EqualFold(s, t string) bool { return strings.EqualFold(s, t) }
