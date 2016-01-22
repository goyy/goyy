// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplEntitySh = `#!/bin/sh

echo [INFO] run go generate.

#export I18N_LOCALE=en_US
go generate
`
