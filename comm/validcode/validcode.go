// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package validcode

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(int64(time.Second))
}

// Standard characters allowed in uniuri string.
var StdChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
