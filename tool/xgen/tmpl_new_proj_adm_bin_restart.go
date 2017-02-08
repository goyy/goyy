// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjAdmBinRestart = `#!/bin/sh
kill -9 ` + "`" + `ps -ef | grep /<%.NewProjName%>adm | grep -v 'grep' | awk '{print $2}'` + "`" + `
cd /app/webapps/<%.NewProjName%>adm/
nohup ./<%.NewProjName%>adm &
`
