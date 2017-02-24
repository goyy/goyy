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
				<mappings url="/assets" dir="/app/assets"/>
			</asset>
			<static name="<%.NewProjName%>">
				<enable>true</enable>
				<mappings url="/static" dir="static"/>
			</static>
			<developer name="<%.NewProjName%>">
				<enable>true</enable>
				<mappings url="/<%.NewProjName%>dev" dir="/app/assets/<%.NewProjName%>dev"/>
			</developer>
			<operation name="<%.NewProjName%>">
				<enable>true</enable>
				<mappings url="/<%.NewProjName%>opr" dir="/app/assets/<%.NewProjName%>opr"/>
			</operation>
		</environment>
		<environment id="test">
			<asset name="<%.NewProjName%>">
				<enable>false</enable>
				<mappings url="http://assets.test.<%.NewProjHost%>" dir="/app/assets"/>
			</asset>
			<static name="<%.NewProjName%>">
				<enable>false</enable>
				<mappings url="http://assets.test.<%.NewProjHost%>/<%.NewProjName%>adm" dir="/app/assets/<%.NewProjName%>adm"/>
			</static>
			<developer name="<%.NewProjName%>">
				<enable>false</enable>
				<mappings url="http://assets.test.<%.NewProjHost%>/<%.NewProjName%>dev" dir="/app/assets/<%.NewProjName%>dev"/>
			</developer>
			<operation name="<%.NewProjName%>">
				<enable>false</enable>
				<mappings url="http://assets.test.<%.NewProjHost%>/<%.NewProjName%>opr" dir="/app/assets/<%.NewProjName%>opr"/>
			</operation>
		</environment>
		<environment id="production">
			<asset name="<%.NewProjName%>">
				<enable>false</enable>
				<mappings url="http://assets.<%.NewProjHost%>" dir="/app/assets"/>
			</asset>
			<static name="<%.NewProjName%>">
				<enable>false</enable>
				<mappings url="http://assets.<%.NewProjHost%>/<%.NewProjName%>adm" dir="/app/assets/<%.NewProjName%>adm"/>
			</static>
			<developer name="<%.NewProjName%>">
				<enable>false</enable>
				<mappings url="http://assets.<%.NewProjHost%>/<%.NewProjName%>dev" dir="/app/assets/<%.NewProjName%>dev"/>
			</developer>
			<operation name="<%.NewProjName%>">
				<enable>false</enable>
				<mappings url="http://assets.<%.NewProjHost%>/<%.NewProjName%>opr" dir="/app/assets/<%.NewProjName%>opr"/>
			</operation>
		</environment>
	</environments>
</configuration>
`
