// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
)

type valids struct{}

func (me *valids) IsExistXML() {
	isExit := false
	if me.IsExistSettings() == false {
		isExit = true
	}
	if me.IsExistEnvironments() == false {
		isExit = true
	}
	if me.IsExistProjects() == false {
		isExit = true
	}
	if me.IsExistModules() == false {
		isExit = true
	}
	if me.IsExistDomains() == false {
		isExit = true
	}
	if me.IsExistColumns() == false {
		isExit = true
	}
	if me.IsExistTables() == false {
		isExit = true
	}
	if me.IsExistProjectTables() == false {
		isExit = true
	}
	if isExit == true {
		log.Fatal("Create and initialize the xmlfiles")
	}
}

func (me *valids) IsExistSettings() bool {
	content := `<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE xtab PUBLIC "-//GOYY//DTD XTAB 1.0//EN" "http://goyy.org/dtd/xtab-1.0.dtd">
<configuration>
	<settings>
		<statement seperator=";" case="upper" comment="true"/>
	</settings>
</configuration>
`
	return util.InitFile(xsettings, content)
}

func (me *valids) IsExistEnvironments() bool {
	content := `<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE env PUBLIC "-//GOYY//DTD ENV 1.0//EN" "http://goyy.org/dtd/env-1.0.dtd">
<configuration>
	<environments default="development">
		<environment id="development">
			<dataSource name="xtab">
				<driverName>mysql</driverName>
				<dataSourceName>root:root@/xtab_development?charset=utf8</dataSourceName>
			</dataSource>
		</environment>
		<environment id="test">
			<dataSource name="xtab">
				<driverName>mysql</driverName>
				<dataSourceName>xtab:xtab@/xtab_test?charset=utf8</dataSourceName>
			</dataSource>
		</environment>
		<environment id="production">
			<dataSource name="xtab">
				<driverName>mysql</driverName>
				<dataSourceName>xtab:xtab@tcp(localhost:3306)/xtab_production?charset=utf8</dataSourceName>
			</dataSource>
		</environment>
	</environments>
</configuration>
`
	return util.InitFile(xenvironments, content)
}

func (me *valids) IsExistProjects() bool {
	content := `<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE xtab PUBLIC "-//GOYY//DTD XTAB 1.0//EN" "http://goyy.org/dtd/xtab-1.0.dtd">
<configuration>
	<projects>
		<project id="demo" name="DEMO" database="xtab" generate="true" comment="Project demo"/>
	</projects>
</configuration>
`
	return util.InitFile(xprojects, content)
}

func (me *valids) IsExistModules() bool {
	content := `<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE xtab PUBLIC "-//GOYY//DTD XTAB 1.0//EN" "http://goyy.org/dtd/xtab-1.0.dtd">
<configuration>
	<modules>
		<module id="sys" name="SYS" prefix="sys" project="demo" generate="true" comment="System tables"/>
		<module id="app" name="APP" prefix="app" project="demo" generate="true" comment="Application business tables"/>
	</modules>
</configuration>
`
	return util.InitFile(xmodules, content)
}

func (me *valids) IsExistDomains() bool {
	content := `<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE xtab PUBLIC "-//GOYY//DTD XTAB 1.0//EN" "http://goyy.org/dtd/xtab-1.0.dtd">
<configuration>
	<domains>
		<domain id="id"     name="ID"     types="string" length="50"/>
		<domain id="flag"   name="flag"   types="string" length="10"/>
		<domain id="code"   name="CODE"   types="string" length="100"/>
		<domain id="name"   name="NAME"   types="string" length="255"/>
		<domain id="memo"   name="MEMO"   types="string" length="1000"/>
		<domain id="descr"  name="DESCR"  types="string" length="2000"/>
		<domain id="remark" name="REMARK" types="string" length="4000"/>
		<domain id="int"    name="INT"    types="int"/>
		<domain id="long"   name="LONG"   types="long"/>
		<domain id="float"  name="FLOAT"  types="float"  length="11" precision="3"/>
		<domain id="bool"   name="BOOL"   types="bool"/>
		<domain id="time"   name="TIME"   types="time"/>
		<domain id="text"   name="TEXT"   types="text"/>
		<domain id="bytes"  name="bytes"  types="bytes"/>
	</domains>
</configuration>
`
	return util.InitFile(xdomains, content)
}

func (me *valids) IsExistColumns() bool {
	content := `<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE xtab PUBLIC "-//GOYY//DTD XTAB 1.0//EN" "http://goyy.org/dtd/xtab-1.0.dtd">
<configuration>
	<columns>
		<column id="id"           name="ID"           domain="id"     comment="primary key"/>
		<column id="code"         name="CODE"         domain="code"   comment="code"/>
		<column id="name"         name="NAME"         domain="name"   comment="name"/>
		<column id="nickname"     name="NICKNAME"     domain="code"   comment="nickname"/>
		<column id="fullname"     name="FULLNAME"     domain="memo"   comment="fullname"/>
		<column id="genre"        name="GENRE"        domain="code"   comment="genre"/>
		<column id="content"      name="CONTENT"      domain="descr"  comment="content"/>
		<column id="memo"         name="MEMO"         domain="memo"   comment="memo"/>
		<column id="descr"        name="DESCR"        domain="descr"  comment="description"/>
		<column id="remark"       name="REMARK"       domain="remark" comment="remark"/>
		<column id="timed"        name="TIME"         domain="time"   comment="time"/>
		<column id="passwd"       name="PASSWD"       domain="name"   comment="password"/>
		<column id="email"        name="EMAIL"        domain="name"   comment="email"/>
		<column id="mobile"       name="MOBILE"       domain="code"   comment="mobile"/>
		<column id="tel"          name="TEL"          domain="code"   comment="telephone"/>
		<column id="contact"      name="CONTACT"      domain="name"   comment="contact"/>
		<column id="sex"          name="SEX"          domain="flag"   comment="sex"/>
		<column id="birthday"     name="BIRTHDAY"     domain="long"   comment="birthday"/>
		<column id="receiver"     name="RECEIVER"     domain="code"   comment="receiver"/>
		<column id="address"      name="ADDRESS"      domain="memo"   comment="address"/>
		<column id="zipcode"      name="ZIPCODE"      domain="code"   comment="zipcode"/>
		<column id="intro"        name="INTRODUCTION" domain="remark" comment="introduction"/>
		<column id="weight"       name="WEIGHT"       domain="int"    comment="weight"/>
		<column id="city"         name="CITY"         domain="id"     comment="city"/>
		<column id="ordinal"      name="ORDINAL"      domain="code"   comment="ordinal"/>
		<column id="img"          name="IMAGES"       domain="name"   comment="images"/>
		<column id="usable"       name="USABLE"       domain="flag"   comment="usable"/>
		<column id="status"       name="STATUS"       domain="flag"   comment="status"/>
		<column id="mark"         name="MARK"         domain="flag"   comment="mark"/>
		<column id="source"       name="SOURCE"       domain="flag"   comment="source"/>
		<column id="mkey"         name="MKEY"         domain="flag"   comment="map key"/>
		<column id="mval"         name="MVAL"         domain="name"   comment="map value"/>
		<column id="params"       name="PARAMS"       domain="name"   comment="params"/>
		<column id="filters"      name="FILTERS"      domain="name"   comment="filters"/>
		<column id="begin_time"   name="BEGIN_TIME"   domain="long"   comment="begin time"/>
		<column id="end_time"     name="END_TIME"     domain="long"   comment="end time"/>
		<column id="area_id"      name="AREA_ID"      domain="id"     comment="the identity of the area table"/>
		<column id="org_id"       name="ORG_ID"       domain="id"     comment="the identity of the org table"/>
		<column id="user_id"      name="USER_ID"      domain="id"     comment="the identity of the user table"/>
		<column id="role_id"      name="ROLE_ID"      domain="id"     comment="the identity of the role table"/>
		<column id="post_id"      name="POST_ID"      domain="id"     comment="the identity of the post table"/>
		<column id="menu_id"      name="MENU_ID"      domain="id"     comment="the identity of the menu table"/>
		<column id="parent_id"    name="PARENT_ID"    domain="id"     comment="the identity of the parent table"/>
		<column id="parent_ids"   name="PARENT_IDS"   domain="memo"   comment="all the identity of the parent table"/>
		<column id="parent_codes" name="PARENT_CODES" domain="memo"   comment="all the code of the parent table"/>
		<column id="parent_names" name="PARENT_NAMES" domain="memo"   comment="all the name of the parent table"/>
		<column id="leaf"         name="LEAF"         domain="int"    comment="whether is the leaf node of the tree"/>
		<column id="grade"        name="GRADE"        domain="int"    comment="The level of the tree node"/>
		<column id="creates"      name="CREATES"      domain="id"     comment="created org"/>
		<column id="creater"      name="CREATER"      domain="id"     comment="created user"/>
		<column id="created"      name="CREATED"      domain="long"   comment="created time"/>
		<column id="modifier"     name="MODIFIER"     domain="id"     comment="modified user"/>
		<column id="modified"     name="MODIFIED"     domain="long"   comment="modified time"/>
		<column id="version"      name="VERSION"      domain="int"    comment="optimistic locking" default="0"/>
		<column id="deletion"     name="DELETION"     domain="bool"   comment="logic delete flag"  default="0"/>
		<column id="artifical"    name="ARTIFICAL"    domain="bool"   comment="artificial data"    default="0"/>
		<column id="history"      name="HISTORY"      domain="bool"   comment="history data"       default="0"/>
	</columns>
</configuration>
`
	return util.InitFile(xcolumns, content)
}

func (me *valids) IsExistTables() bool {
	content := `<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE xtab PUBLIC "-//GOYY//DTD XTAB 1.0//EN" "http://goyy.org/dtd/xtab-1.0.dtd">
<configuration>
	<tables>
		<table id="pk" name="PK" comment="primary key table">
			<column extends="id"/>
		</table>
		<table id="sys" name="SYS" extends="pk" comment="sys table">
			<column extends="memo"/>
			<column extends="creates"/>
			<column extends="creater"/>
			<column extends="created"/>
			<column extends="modifier"/>
			<column extends="modified"/>
			<column extends="version"/>
			<column extends="deletion"/>
			<column extends="artifical"/>
			<column extends="history"/>
		</table>
		<table id="tree" name="TREE" extends="sys" comment="tree table">
			<column extends="code"/>
			<column extends="name"/>
			<column extends="fullname"/>
			<column extends="genre"/>
			<column extends="ordinal"/>
			<column extends="parent_id"/>
			<column extends="parent_ids"/>
			<column extends="parent_codes"/>
			<column extends="parent_names"/>
			<column extends="leaf"/>
			<column extends="grade"/>
		</table>
	</tables>
</configuration>
`
	return util.InitFile(xtables, content)
}

func (me *valids) IsExistProjectTables() bool {
	app := `<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE xtab PUBLIC "-//GOYY//DTD XTAB 1.0//EN" "http://goyy.org/dtd/xtab-1.0.dtd">
<configuration>
	<tables>
		<table id="product" name="PRODUCT" extends="base" comment="product table">
			<column id="title"       name="TITLE"       domain="memo"   comment="product title"/>
			<column id="description" name="DESCRIPTION" domain="remark" comment="product description"/>
			<column id="price"       name="PRICE"       domain="float"  comment="product price"/>
		</table>
		<table id="order" name="ORDER" extends="base" comment="order table">
			<column extends="name"/>
			<column extends="email"/>
			<column extends="address"/>
		</table>
		<table id="order_product" name="ORDER_PRODUCT" extends="base" comment="order product">
			<column id="order_id"   name="ORDER_ID"   domain="id"    comment="product title"/>
			<column id="product_id" name="PRODUCT_ID" domain="id"    comment="product description"/>
			<column id="quantity"   name="QUANTITY"   domain="int"   comment="product quantity"/>
			<column id="unit_price" name="UNIT_PRICE" domain="float" comment="product unit price"/>
		</table>
	</tables>
</configuration>
`
	sys := `<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE xtab PUBLIC "-//GOYY//DTD XTAB 1.0//EN" "http://goyy.org/dtd/xtab-1.0.dtd">
<configuration>
	<tables>
		<table id="user" name="USER" extends="base" comment="user table">
			<column extends="name" index="true"/>
			<column extends="passwd"/>
		</table>
		<table id="role" name="ROLE" extends="base" comment="role table">
			<column extends="name"/>
		</table>
		<table id="user_role" name="USER_ROLE" extends="pk" comment="user role associated table">
			<column id="user_id" name="USER_ID" domain="id" comment="user table identifies"/>
			<column id="role_id" name="ROLE_ID" domain="id" comment="role table identifies"/>
		</table>
	</tables>
</configuration>
`
	content := `<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE xtab PUBLIC "-//GOYY//DTD XTAB 1.0//EN" "http://goyy.org/dtd/xtab-1.0.dtd">
<configuration>
	<tables>
		<table id="table1" name="TABLE1" extends="base" comment="table1">
			<column extends="memo"/>
			<column id="column1" name="COLUMN1" domain="code" comment="column1"/>
		</table>
	</tables>
</configuration>
`
	xconf := util.DecodeXML(xmodules)
	isExist := true
	for _, v := range xconf.Modules.Module {
		if v.Project == "demo" && v.Id == "app" {
			if util.InitFile("./conf/schema/tables-demo-app.xml", app) == false {
				isExist = false
			}
		} else if v.Project == "demo" && v.Id == "sys" {
			if util.InitFile("./conf/schema/tables-demo-sys.xml", sys) == false {
				isExist = false
			}
		} else {
			if util.InitFile("./conf/schema/tables-"+v.Project+"-"+v.Id+".xml", content) == false {
				isExist = false
			}
		}
	}
	return isExist
}
