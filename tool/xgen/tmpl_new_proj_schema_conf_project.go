// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjSchemaConfProject = `<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE xtab PUBLIC "-//GOYY//DTD XTAB 1.0//EN" "http://goyy.org/dtd/xtab-1.0.dtd">
<configuration>
	<projects>
		<project id="<%.NewProjName%>" name="<%.NewProjTitle%>" database="<%.NewProjName%>" generate="true" comment="<%.NewProjTitle%>" admpath="<%.NewProjPkg%>/<%.NewProjName%>-adm" tstpath="<%.NewProjPkg%>/<%.NewProjName%>-tst"/>
	</projects>
</configuration>
`
