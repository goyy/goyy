// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplControllerClientRegProj = `// generated by xgen -- DO NOT EDIT
package internal

import (
	_ "{{.Clipath}}/internal/{{.Project}}"
)
`