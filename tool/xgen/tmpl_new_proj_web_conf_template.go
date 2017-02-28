// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjWebConfTemplate = `<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE env PUBLIC "-//GOYY//DTD ENV 1.0//EN" "http://goyy.org/dtd/env-1.0.dtd">
<configuration>
	<environments default="development">
		<environment id="development">
			<html name="<%.NewProjName%>">
				<enable>true</enable>
				<reloaded>true</reloaded>
				<mappings>
					<mapping path="/example" dir="../<%.NewProjName%>-example/templates/web"/>
				</mappings>
			</html>
			<template name="<%.NewProjName%>">
				<enable>true</enable>
				<reloaded>true</reloaded>
			</template>
		</environment>
		<environment id="test">
			<html name="<%.NewProjName%>">
				<enable>true</enable>
				<reloaded>false</reloaded>
			</html>
			<template name="<%.NewProjName%>">
				<enable>true</enable>
				<reloaded>false</reloaded>
			</template>
		</environment>
		<environment id="production">
			<html name="<%.NewProjName%>">
				<enable>true</enable>
				<reloaded>false</reloaded>
			</html>
			<template name="<%.NewProjName%>">
				<enable>true</enable>
				<reloaded>false</reloaded>
			</template>
		</environment>
	</environments>
</configuration>
`
