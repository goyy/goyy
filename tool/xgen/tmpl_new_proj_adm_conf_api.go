// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjAdmConfAPI = `<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE env PUBLIC "-//GOYY//DTD ENV 1.0//EN" "http://goyy.org/dtd/env-1.0.dtd">
<configuration>
	<environments default="development">
		<environment id="development">
			<api name="<%.NewProjName%>">
				<url>/apis</url>
			</api>
		</environment>
		<environment id="test">
			<api name="<%.NewProjName%>">
				<url>/apis</url>
			</api>
		</environment>
		<environment id="production">
			<api name="<%.NewProjName%>">
				<url>/apis</url>
			</api>
		</environment>
	</environments>
</configuration>
`
