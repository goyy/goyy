// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjWebConfUpload = `<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE env PUBLIC "-//GOYY//DTD ENV 1.0//EN" "http://goyy.org/dtd/env-1.0.dtd">
<configuration>
	<environments default="development">
		<environment id="development">
			<upload name="<%.NewProjName%>">
				<enable>true</enable>
				<dir>/app/assets/<%.NewProjName%>upl</dir>
				<url>/<%.NewProjName%>upl</url>
				<maxSize>5242880</maxSize>
			</upload>
		</environment>
		<environment id="test">
			<upload name="<%.NewProjName%>">
				<enable>false</enable>
				<dir>/app/assets/<%.NewProjName%>upl</dir>
				<url>http://assets.test.<%.NewProjHost%>/<%.NewProjName%>upl</url>
				<maxSize>5242880</maxSize>
			</upload>
		</environment>
		<environment id="production">
			<upload name="<%.NewProjName%>">
				<enable>false</enable>
				<dir>/app/assets/<%.NewProjName%>upl</dir>
				<url>http://assets.<%.NewProjHost%>/<%.NewProjName%>upl</url>
				<maxSize>5242880</maxSize>
			</upload>
		</environment>
	</environments>
</configuration>
`
