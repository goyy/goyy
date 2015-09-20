// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package profile_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/comm/profile"
)

func ExampleActives() {
	profile.SetActives(profile.PROD, profile.TEST)
	fmt.Println(profile.Actives())

	// Output:[production test]
}
