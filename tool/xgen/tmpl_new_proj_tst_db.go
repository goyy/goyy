// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjTstDB = `<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE env PUBLIC "-//GOYY//DTD ENV 1.0//EN" "http://goyy.org/dtd/env-1.0.dtd">
<configuration>
	<environments default="development">
		<environment id="development">
			<database name="<%.NewProjName%>">
				<driverName>mysql</driverName>
				<dataSourceName>root:root@tcp(127.0.0.1:3306)/<%.NewProjName%>_development</dataSourceName>
				<maxIdleConns>10</maxIdleConns>
				<maxOpenConns>100</maxOpenConns>
			</database>
		</environment>
		<environment id="test">
			<database name="<%.NewProjName%>">
				<driverName>mysql</driverName>
				<dataSourceName>root:root@tcp(127.0.0.1:3306)/<%.NewProjName%>_test</dataSourceName>
				<maxIdleConns>10</maxIdleConns>
				<maxOpenConns>100</maxOpenConns>
			</database>
		</environment>
		<environment id="production">
			<database name="<%.NewProjName%>">
				<driverName>mysql</driverName>
				<dataSourceName>root:root@tcp(127.0.0.1:3306)/<%.NewProjName%>_production</dataSourceName>
				<maxIdleConns>10</maxIdleConns>
				<maxOpenConns>100</maxOpenConns>
			</database>
		</environment>
	</environments>
</configuration>
`
