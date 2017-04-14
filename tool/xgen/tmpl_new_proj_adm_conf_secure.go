// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjAdmConfSecure = `<?xml version="1.0" encoding="UTF-8" ?>
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
					<intercept-url pattern="/**.(css|js|map)" access="anon"/>
					<intercept-url pattern="/**.(jpg|gif|png|bmp|ico)" access="anon"/>
					<intercept-url pattern="/login.html" access="anon"/>
					<intercept-url pattern="/login" access="anon"/>
					<intercept-url pattern="/signin" access="anon"/>
					<intercept-url pattern="/logout" access="anon"/>
					<intercept-url pattern="/captcha/build" access="anon"/>
					<intercept-url pattern="/err/**" access="anon"/>
					<intercept-url pattern="/**" access="authc"/>
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
					<intercept-url pattern="/**.(css|js|map)" access="anon"/>
					<intercept-url pattern="/**.(jpg|gif|png|bmp|ico)" access="anon"/>
					<intercept-url pattern="/login.html" access="anon"/>
					<intercept-url pattern="/login" access="anon"/>
					<intercept-url pattern="/signin" access="anon"/>
					<intercept-url pattern="/logout" access="anon"/>
					<intercept-url pattern="/captcha/build" access="anon"/>
					<intercept-url pattern="/err/**" access="anon"/>
					<intercept-url pattern="/**" access="authc"/>
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
					<intercept-url pattern="/**.(css|js|map)" access="anon"/>
					<intercept-url pattern="/**.(jpg|gif|png|bmp|ico)" access="anon"/>
					<intercept-url pattern="/login.html" access="anon"/>
					<intercept-url pattern="/login" access="anon"/>
					<intercept-url pattern="/signin" access="anon"/>
					<intercept-url pattern="/logout" access="anon"/>
					<intercept-url pattern="/captcha/build" access="anon"/>
					<intercept-url pattern="/err/**" access="anon"/>
					<intercept-url pattern="/**" access="authc"/>
				</filters>
			</secure>
		</environment>
	</environments>
</configuration>
`
