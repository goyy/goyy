// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
)

// PreError controller.PreError.
type PreError struct {
	Code    string
	Message string
}

func (me *PreError) Error() string {
	return fmt.Sprintf("code:%s, message:%s", me.Code, me.Message)
}
