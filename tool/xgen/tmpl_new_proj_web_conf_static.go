// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjWebConfStatic = `<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE env PUBLIC "-//GOYY//DTD ENV 1.0//EN" "http://goyy.org/dtd/env-1.0.dtd">
<configuration>
	<environments default="development">
		<environment id="development">
			<asset name="<%.NewProjName%>">
				<enable>true</enable>
				<dir>/app/assets</dir>
				<url>/assets</url>
			</asset>
			<static name="<%.NewProjName%>">
				<enable>true</enable>
				<dir>static</dir>
				<url>/statics</url>
			</static>
			<developer name="<%.NewProjName%>">
				<enable>true</enable>
				<dir>/app/assets/<%.NewProjName%>dev</dir>
				<url>/<%.NewProjName%>dev</url>
			</developer>
			<operation name="<%.NewProjName%>">
				<enable>true</enable>
				<dir>/app/assets/<%.NewProjName%>opr</dir>
				<url>/<%.NewProjName%>opr</url>
			</operation>
		</environment>
		<environment id="test">
			<asset name="<%.NewProjName%>">
				<enable>false</enable>
				<dir>/app/assets</dir>
				<url>http://assets.test.<%.NewProjName%>.com</url>
			</asset>
			<static name="<%.NewProjName%>">
				<enable>false</enable>
				<dir>/app/assets/<%.NewProjName%>web</dir>
				<url>http://assets.test.<%.NewProjName%>.com/<%.NewProjName%>web</url>
			</static>
			<developer name="<%.NewProjName%>">
				<enable>false</enable>
				<dir>/app/assets/<%.NewProjName%>dev</dir>
				<url>http://assets.test.<%.NewProjName%>.com/<%.NewProjName%>dev</url>
			</developer>
			<operation name="<%.NewProjName%>">
				<enable>false</enable>
				<dir>/app/assets/<%.NewProjName%>opr</dir>
				<url>http://assets.test.<%.NewProjName%>.com/<%.NewProjName%>opr</url>
			</operation>
		</environment>
		<environment id="production">
			<asset name="<%.NewProjName%>">
				<enable>false</enable>
				<dir>/app/assets</dir>
				<url>http://assets.<%.NewProjName%>.com</url>
			</asset>
			<static name="<%.NewProjName%>">
				<enable>false</enable>
				<dir>/app/assets/<%.NewProjName%>web</dir>
				<url>http://assets.<%.NewProjName%>.com/<%.NewProjName%>web</url>
			</static>
			<developer name="<%.NewProjName%>">
				<enable>false</enable>
				<dir>/app/assets/<%.NewProjName%>dev</dir>
				<url>http://assets.<%.NewProjName%>.com/<%.NewProjName%>dev</url>
			</developer>
			<operation name="<%.NewProjName%>">
				<enable>false</enable>
				<dir>/app/assets/<%.NewProjName%>opr</dir>
				<url>http://assets.<%.NewProjName%>.com/<%.NewProjName%>opr</url>
			</operation>
		</environment>
	</environments>
</configuration>
`
