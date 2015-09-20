# data-xtab
To design the database tables by using XML files

# Setup
* `go get gopkg.in/goyy/goyy.v0/data/xtab`
* `go install gopkg.in/goyy/goyy.v0/data/xtab`

# Associations
environment `1 -> 1` project `1 -> *` module `1 -> *` table `1 -> *` column `1 -> 1` domain

# Convention
* Settings Info:`./xsettings.xml`
* Environment Info:`./xenvironments.xml`
* Project Info:`./xprojects.xml`
* Module Info:`./xmodules.xml`
* Base Table Info:`./xtables.xml`
* Derived Table Info:`./xtables-${project}-${module}.xml`
* Base Column Info:`./xcolumns.xml`
* Domain Info:`./xdomains.xml`

# Output Format
* SQL
* HTML
* Database

# Types
<table>
	<tr><th>xtab</th><th>mysql</th><th>postgres</th><th>oracle</th><th>sqlserver</th></tr>
	<tr><td>string</td><td>varchar</td><td>varchar</td><td>nvarchar2</td><td>nvarchar</td></tr>
	<tr><td>int</td><td>int</td><td>integer</td><td>integer</td><td>int</td></tr>
	<tr><td>long</td><td>bigint</td><td>bigint</td><td>number(19,0)</td><td>numeric(19,0)</td></tr>
	<tr><td>float</td><td>decimal</td><td>numeric</td><td>number</td><td>numeric</td></tr>
	<tr><td>bool</td><td>tinyint</td><td>boolean</td><td>number(1,0)</td><td>tinyint</td></tr>
	<tr><td>time</td><td>datetime</td><td>timestamp</td><td>date</td><td>datetime</td></tr>
	<tr><td>text</td><td>text</td><td>text</td><td>clob</td><td>text</td></tr>
	<tr><td>bytes</td><td>longblob</td><td>bytea</td><td>blob</td><td>image</td></tr>
</table>
