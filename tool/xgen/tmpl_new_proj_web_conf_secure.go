// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjWebConfSecure = `<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE env PUBLIC "-//GOYY//DTD ENV 1.0//EN" "http://goyy.org/dtd/env-1.0.dtd">
<configuration>
	<environments default="development">
		<environment id="development">
			<secure name="<%.NewProjName%>">
				<enable>true</enable>
				<login-url>/login.html</login-url>
				<forbid-url>/err/403.html</forbid-url>
				<success-url>/home.html</success-url>
				<filters>
					<intercept-url pattern="/apis/sys/**/(save|saved)" access="forbid"/>
					<intercept-url pattern="/apis/**/(disable|disabled)" access="forbid"/>
					<intercept-url pattern="/apis/sys/menu/**" access="forbid"/>
					<intercept-url pattern="/apis/sys/post/**" access="forbid"/>
					<intercept-url pattern="/apis/sys/user/role/**" access="forbid"/>
					<intercept-url pattern="/apis/sys/user/(index|show|add|edit|repwd)" access="forbid"/>
					<intercept-url pattern="/**" access="anon"/>
				</filters>
			</secure>
		</environment>
		<environment id="test">
			<secure name="<%.NewProjName%>">
				<enable>true</enable>
				<login-url>/login.html</login-url>
				<forbid-url>/err/403.html</forbid-url>
				<success-url>/home.html</success-url>
				<filters>
					<intercept-url pattern="/apis/sys/**/(save|saved)" access="forbid"/>
					<intercept-url pattern="/apis/**/(disable|disabled)" access="forbid"/>
					<intercept-url pattern="/apis/sys/menu/**" access="forbid"/>
					<intercept-url pattern="/apis/sys/post/**" access="forbid"/>
					<intercept-url pattern="/apis/sys/user/role/**" access="forbid"/>
					<intercept-url pattern="/apis/sys/user/(index|show|add|edit|repwd)" access="forbid"/>
					<intercept-url pattern="/**" access="anon"/>
				</filters>
			</secure>
		</environment>
		<environment id="production">
			<secure name="<%.NewProjName%>">
				<enable>true</enable>
				<login-url>/login.html</login-url>
				<forbid-url>/err/403.html</forbid-url>
				<success-url>/home.html</success-url>
				<filters>
					<intercept-url pattern="/apis/sys/**/(save|saved)" access="forbid"/>
					<intercept-url pattern="/apis/**/(disable|disabled)" access="forbid"/>
					<intercept-url pattern="/apis/sys/menu/**" access="forbid"/>
					<intercept-url pattern="/apis/sys/post/**" access="forbid"/>
					<intercept-url pattern="/apis/sys/user/role/**" access="forbid"/>
					<intercept-url pattern="/apis/sys/user/(index|show|add|edit|repwd)" access="forbid"/>
					<intercept-url pattern="/**" access="anon"/>
				</filters>
			</secure>
		</environment>
	</environments>
</configuration>
`
