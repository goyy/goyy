// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjAdmConfStatic = `<?xml version="1.0" encoding="UTF-8" ?>
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
				<url>/static</url>
				<mappings>
					<mapping path="/sys" dir="%GOPATH%/src/gopkg.in/goyy/goyy.v0/app/sys/static/adm"/>
					<mapping path="/example" dir="%GOPATH%/src/gopkg.in/goyy/goyy.v0/app/example/static/adm"/>
				</mappings>
			</static>
			<developer name="<%.NewProjName%>">
				<enable>true</enable>
				<dir>/app/assets/<%.NewProjName%>dev</dir>
				<url>/<%.NewProjName%>dev</url>
				<mappings>
					<mapping path="/ui" dir="%GOPATH%/src/gopkg.in/goyy/goyy.v0/app/ui/static"/>
				</mappings>
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
				<url>http://assets.test.<%.NewProjHost%></url>
			</asset>
			<static name="<%.NewProjName%>">
				<enable>false</enable>
				<dir>/app/assets/<%.NewProjName%>adm</dir>
				<url>http://assets.test.<%.NewProjHost%>/<%.NewProjName%>adm</url>
			</static>
			<developer name="<%.NewProjName%>">
				<enable>false</enable>
				<dir>/app/assets/<%.NewProjName%>dev</dir>
				<url>http://assets.test.<%.NewProjHost%>/<%.NewProjName%>dev</url>
			</developer>
			<operation name="<%.NewProjName%>">
				<enable>false</enable>
				<dir>/app/assets/<%.NewProjName%>opr</dir>
				<url>http://assets.test.<%.NewProjHost%>/<%.NewProjName%>opr</url>
			</operation>
		</environment>
		<environment id="production">
			<asset name="<%.NewProjName%>">
				<enable>false</enable>
				<dir>/app/assets</dir>
				<url>http://assets.<%.NewProjHost%></url>
			</asset>
			<static name="<%.NewProjName%>">
				<enable>false</enable>
				<dir>/app/assets/<%.NewProjName%>adm</dir>
				<url>http://assets.<%.NewProjHost%>/<%.NewProjName%>adm</url>
			</static>
			<developer name="<%.NewProjName%>">
				<enable>false</enable>
				<dir>/app/assets/<%.NewProjName%>dev</dir>
				<url>http://assets.<%.NewProjHost%>/<%.NewProjName%>dev</url>
			</developer>
			<operation name="<%.NewProjName%>">
				<enable>false</enable>
				<dir>/app/assets/<%.NewProjName%>opr</dir>
				<url>http://assets.<%.NewProjHost%>/<%.NewProjName%>opr</url>
			</operation>
		</environment>
	</environments>
</configuration>
`
