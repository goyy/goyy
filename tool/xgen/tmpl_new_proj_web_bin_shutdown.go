// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjWebBinShutdown = `#!/bin/sh
kill -9 ` + "`" + `ps -ef | grep /<%.NewProjName%>web | grep -v 'grep' | awk '{print $2}'` + "`" + `
exit
`
