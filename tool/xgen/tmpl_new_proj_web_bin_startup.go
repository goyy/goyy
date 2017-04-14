// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjWebBinStartup = `#!/bin/sh
cd /app/webapps/<%.NewProjName%>web/
nohup ./<%.NewProjName%>web &
`
