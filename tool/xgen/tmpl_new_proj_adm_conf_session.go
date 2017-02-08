// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjAdmConfSession = `<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE env PUBLIC "-//GOYY//DTD ENV 1.0//EN" "http://goyy.org/dtd/env-1.0.dtd">
<configuration>
	<environments default="development">
		<environment id="development">
			<session name="<%.NewProjName%>">
				<addr>127.0.0.1:6379</addr>
			</session>
		</environment>
		<environment id="test">
			<session name="<%.NewProjName%>">
				<addr>127.0.0.1:6379</addr>
			</session>
		</environment>
		<environment id="production">
			<session name="<%.NewProjName%>">
				<addr>127.0.0.1:6379</addr>
				<password><%.NewProjName%></password>
			</session>
		</environment>
	</environments>
</configuration>
`
