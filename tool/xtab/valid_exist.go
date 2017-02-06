// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
)

type valids struct{}

func (me *valids) IsExistXML(name, pkg string) {
	isExit := false
	if me.IsExistSettings() == false {
		isExit = true
	}
	if me.IsExistEnvironments(name) == false {
		isExit = true
	}
	if me.IsExistProjects(name, pkg) == false {
		isExit = true
	}
	if me.IsExistModules(name, pkg) == false {
		isExit = true
	}
	if me.IsExistButtons() == false {
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

func (me *valids) IsExistEnvironments(name string) bool {
	content := `<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE env PUBLIC "-//GOYY//DTD ENV 1.0//EN" "http://goyy.org/dtd/env-1.0.dtd">
<configuration>
	<environments default="development">
		<environment id="development">
			<database name="` + name + `">
				<driverName>mysql</driverName>
				<dataSourceName>root:root@tcp(127.0.0.1:3306)/` + name + `_development</dataSourceName>
				<maxIdleConns>10</maxIdleConns>
				<maxOpenConns>100</maxOpenConns>
			</database>
		</environment>
		<environment id="test">
			<database name="` + name + `">
				<driverName>mysql</driverName>
				<dataSourceName>root:root@tcp(127.0.0.1:3306)/` + name + `_test</dataSourceName>
				<maxIdleConns>10</maxIdleConns>
				<maxOpenConns>100</maxOpenConns>
			</database>
		</environment>
		<environment id="production">
			<database name="` + name + `">
				<driverName>mysql</driverName>
				<dataSourceName>root:root@tcp(127.0.0.1:3306)/` + name + `_production</dataSourceName>
				<maxIdleConns>10</maxIdleConns>
				<maxOpenConns>100</maxOpenConns>
			</database>
		</environment>
	</environments>
</configuration>
`
	return util.InitFile(xenvironments, content)
}

func (me *valids) IsExistProjects(name, pkg string) bool {
	content := `<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE xtab PUBLIC "-//GOYY//DTD XTAB 1.0//EN" "http://goyy.org/dtd/xtab-1.0.dtd">
<configuration>
	<projects>
		<project id="` + name + `" name="{{message "tmpl.proj.name"}}" database="` + name + `" generate="true" comment="{{message "tmpl.proj.memo"}}" admpath="` + pkg + `/` + name + `-adm" tstpath="` + pkg + `/` + name + `-tst"/>
	</projects>
</configuration>
`
	return util.InitFile(xprojects, content)
}

func (me *valids) IsExistModules(name, pkg string) bool {
	content := `<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE xtab PUBLIC "-//GOYY//DTD XTAB 1.0//EN" "http://goyy.org/dtd/xtab-1.0.dtd">
<configuration>
	<modules>
		<module id="sys" name="{{message "tmpl.mod.sys.name"}}" prefix="sys" project="` + name + `" generate="true" comment="{{message "tmpl.mod.sys.memo"}}" apipath="` + pkg + `/` + name + `-sys"/>
		<module id="app" name="{{message "tmpl.mod.app.name"}}" prefix="app" project="` + name + `" generate="true" comment="{{message "tmpl.mod.app.memo"}}" apipath="` + pkg + `/` + name + `-app"/>
	</modules>
</configuration>
`
	return util.InitFile(xmodules, content)
}

func (me *valids) IsExistButtons() bool {
	content := `<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE xtab PUBLIC "-//GOYY//DTD XTAB 1.0//EN" "http://goyy.org/dtd/xtab-1.0.dtd">
<configuration>
	<buttons>
		<button id="view"    name="{{message "tmpl.btn.view.name"}}" comment="{{comments "tmpl.btn.view.name" "tmpl.btn.memo"}}"/>
		<button id="add"     name="{{message "tmpl.btn.add.name"}}" comment="{{comments "tmpl.btn.add.name" "tmpl.btn.memo"}}"/>
		<button id="edit"    name="{{message "tmpl.btn.edit.name"}}" comment="{{comments "tmpl.btn.edit.name" "tmpl.btn.memo"}}"/>
		<button id="disable" name="{{message "tmpl.btn.disable.name"}}" comment="{{comments "tmpl.btn.disable.name" "tmpl.btn.memo"}}"/>
		<button id="export"  name="{{message "tmpl.btn.export.name"}}" comment="{{comments "tmpl.btn.export.name" "tmpl.btn.memo"}}"/>
	</buttons>
</configuration>
`
	return util.InitFile(xbuttons, content)
}

func (me *valids) IsExistDomains() bool {
	content := `<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE xtab PUBLIC "-//GOYY//DTD XTAB 1.0//EN" "http://goyy.org/dtd/xtab-1.0.dtd">
<configuration>
	<domains>
		<domain id="id"     name="ID"     types="string" length="50"/>
		<domain id="flag"   name="FLAG"   types="string" length="10"/>
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
		<domain id="bytes"  name="BYTES"  types="bytes"/>
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
		<column id="id"           domain="id"     name="{{message "tmpl.col.id.name"}}" comment="{{message "tmpl.col.id.memo"}}"/>
		<column id="code"         domain="code"   name="{{message "tmpl.col.code.name"}}" comment="{{message "tmpl.col.code.memo"}}"/>
		<column id="name"         domain="name"   name="{{message "tmpl.col.name.name"}}" comment="{{message "tmpl.col.name.memo"}}"/>
		<column id="nickname"     domain="code"   name="{{message "tmpl.col.nickname.name"}}" comment="{{message "tmpl.col.nickname.memo"}}"/>
		<column id="fullname"     domain="memo"   name="{{message "tmpl.col.fullname.name"}}" comment="{{message "tmpl.col.fullname.memo"}}"/>
		<column id="genre"        domain="code"   name="{{message "tmpl.col.genre.name"}}" comment="{{message "tmpl.col.genre.memo"}}"/>
		<column id="classify"     domain="code"   name="{{message "tmpl.col.classify.name"}}" comment="{{message "tmpl.col.classify.memo"}}"/>
		<column id="content"      domain="descr"  name="{{message "tmpl.col.content.name"}}" comment="{{message "tmpl.col.content.memo"}}"/>
		<column id="memo"         domain="memo"   name="{{message "tmpl.col.memo.name"}}" comment="{{message "tmpl.col.memo.memo"}}"/>
		<column id="descr"        domain="descr"  name="{{message "tmpl.col.descr.name"}}" comment="{{message "tmpl.col.descr.memo"}}"/>
		<column id="remark"       domain="remark" name="{{message "tmpl.col.remark.name"}}" comment="{{message "tmpl.col.remark.memo"}}"/>
		<column id="timed"        domain="time"   name="{{message "tmpl.col.timed.name"}}" comment="{{message "tmpl.col.timed.memo"}}"/>
		<column id="passwd"       domain="name"   name="{{message "tmpl.col.passwd.name"}}" comment="{{message "tmpl.col.passwd.memo"}}"/>
		<column id="email"        domain="name"   name="{{message "tmpl.col.email.name"}}" comment="{{message "tmpl.col.email.memo"}}"/>
		<column id="mobile"       domain="code"   name="{{message "tmpl.col.mobile.name"}}" comment="{{message "tmpl.col.mobile.memo"}}"/>
		<column id="tel"          domain="code"   name="{{message "tmpl.col.tel.name"}}" comment="{{message "tmpl.col.tel.memo"}}"/>
		<column id="contact"      domain="name"   name="{{message "tmpl.col.contact.name"}}" comment="{{message "tmpl.col.contact.memo"}}"/>
		<column id="sex"          domain="flag"   name="{{message "tmpl.col.sex.name"}}" comment="{{message "tmpl.col.sex.memo"}}"/>
		<column id="birthday"     domain="long"   name="{{message "tmpl.col.birthday.name"}}" comment="{{message "tmpl.col.birthday.memo"}}"/>
		<column id="receiver"     domain="code"   name="{{message "tmpl.col.receiver.name"}}" comment="{{message "tmpl.col.receiver.memo"}}"/>
		<column id="address"      domain="memo"   name="{{message "tmpl.col.address.name"}}" comment="{{message "tmpl.col.address.memo"}}"/>
		<column id="zipcode"      domain="code"   name="{{message "tmpl.col.zipcode.name"}}" comment="{{message "tmpl.col.zipcode.memo"}}"/>
		<column id="intro"        domain="remark" name="{{message "tmpl.col.intro.name"}}" comment="{{message "tmpl.col.intro.memo"}}"/>
		<column id="weight"       domain="int"    name="{{message "tmpl.col.weight.name"}}" comment="{{message "tmpl.col.weight.memo"}}"/>
		<column id="city"         domain="id"     name="{{message "tmpl.col.city.name"}}" comment="{{message "tmpl.col.city.memo"}}"/>
		<column id="ordinal"      domain="code"   name="{{message "tmpl.col.ordinal.name"}}" comment="{{message "tmpl.col.ordinal.memo"}}"/>
		<column id="img"          domain="name"   name="{{message "tmpl.col.img.name"}}" comment="{{message "tmpl.col.img.memo"}}"/>
		<column id="usable"       domain="flag"   name="{{message "tmpl.col.usable.name"}}" comment="{{message "tmpl.col.usable.memo"}}"/>
		<column id="status"       domain="flag"   name="{{message "tmpl.col.status.name"}}" comment="{{message "tmpl.col.status.memo"}}"/>
		<column id="mark"         domain="flag"   name="{{message "tmpl.col.mark.name"}}" comment="{{message "tmpl.col.mark.memo"}}"/>
		<column id="source"       domain="flag"   name="{{message "tmpl.col.source.name"}}" comment="{{message "tmpl.col.source.memo"}}"/>
		<column id="mkey"         domain="flag"   name="{{message "tmpl.col.mkey.name"}}" comment="{{message "tmpl.col.mkey.memo"}}"/>
		<column id="mval"         domain="name"   name="{{message "tmpl.col.mval.name"}}" comment="{{message "tmpl.col.mval.memo"}}"/>
		<column id="params"       domain="name"   name="{{message "tmpl.col.params.name"}}" comment="{{message "tmpl.col.params.memo"}}"/>
		<column id="filters"      domain="name"   name="{{message "tmpl.col.filters.name"}}" comment="{{message "tmpl.col.filters.memo"}}"/>
		<column id="begin_time"   domain="long"   name="{{message "tmpl.col.begin_time.name"}}" comment="{{message "tmpl.col.begin_time.memo"}}"/>
		<column id="end_time"     domain="long"   name="{{message "tmpl.col.end_time.name"}}" comment="{{message "tmpl.col.end_time.memo"}}"/>
		<column id="area_id"      domain="id"     name="{{message "tmpl.col.area_id.name"}}" comment="{{message "tmpl.col.area_id.memo"}}"/>
		<column id="org_id"       domain="id"     name="{{message "tmpl.col.org_id.name"}}" comment="{{message "tmpl.col.org_id.memo"}}"/>
		<column id="user_id"      domain="id"     name="{{message "tmpl.col.user_id.name"}}" comment="{{message "tmpl.col.user_id.memo"}}"/>
		<column id="role_id"      domain="id"     name="{{message "tmpl.col.role_id.name"}}" comment="{{message "tmpl.col.role_id.memo"}}"/>
		<column id="post_id"      domain="id"     name="{{message "tmpl.col.post_id.name"}}" comment="{{message "tmpl.col.post_id.memo"}}"/>
		<column id="menu_id"      domain="id"     name="{{message "tmpl.col.menu_id.name"}}" comment="{{message "tmpl.col.menu_id.memo"}}"/>
		<column id="parent_id"    domain="id"     name="{{message "tmpl.col.parent_id.name"}}" comment="{{message "tmpl.col.parent_id.memo"}}"/>
		<column id="parent_ids"   domain="memo"   name="{{message "tmpl.col.parent_ids.name"}}" comment="{{message "tmpl.col.parent_ids.memo"}}"/>
		<column id="parent_codes" domain="memo"   name="{{message "tmpl.col.parent_codes.name"}}" comment="{{message "tmpl.col.parent_codes.memo"}}"/>
		<column id="parent_names" domain="memo"   name="{{message "tmpl.col.parent_names.name"}}" comment="{{message "tmpl.col.parent_names.memo"}}"/>
		<column id="leaf"         domain="int"    name="{{message "tmpl.col.leaf.name"}}" comment="{{message "tmpl.col.leaf.memo"}}"/>
		<column id="grade"        domain="int"    name="{{message "tmpl.col.grade.name"}}" comment="{{message "tmpl.col.grade.memo"}}"/>
		<column id="creates"      domain="id"     name="{{message "tmpl.col.creates.name"}}" comment="{{message "tmpl.col.creates.memo"}}"/>
		<column id="creater"      domain="id"     name="{{message "tmpl.col.creater.name"}}" comment="{{message "tmpl.col.creater.memo"}}"/>
		<column id="created"      domain="long"   name="{{message "tmpl.col.created.name"}}" comment="{{message "tmpl.col.created.memo"}}"/>
		<column id="modifier"     domain="id"     name="{{message "tmpl.col.modifier.name"}}" comment="{{message "tmpl.col.modifier.memo"}}"/>
		<column id="modified"     domain="long"   name="{{message "tmpl.col.modified.name"}}" comment="{{message "tmpl.col.modified.memo"}}"/>
		<column id="version"      domain="int"    name="{{message "tmpl.col.version.name"}}" comment="{{message "tmpl.col.version.memo"}}" default="0"/>
		<column id="deletion"     domain="bool"   name="{{message "tmpl.col.deletion.name"}}" comment="{{message "tmpl.col.deletion.memo"}}" default="0"/>
		<column id="artifical"    domain="bool"   name="{{message "tmpl.col.artifical.name"}}" comment="{{message "tmpl.col.artifical.memo"}}" default="0"/>
		<column id="history"      domain="bool"   name="{{message "tmpl.col.history.name"}}" comment="{{message "tmpl.col.history.memo"}}" default="0"/>
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
		<table id="pk" name="{{message "tmpl.tab.pk.name"}}" comment="{{message "tmpl.tab.pk.memo"}}">
			<column extends="id"/>
		</table>
		<table id="sys" name="{{message "tmpl.tab.sys.name"}}" extends="pk" comment="{{message "tmpl.tab.sys.memo"}}">
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
		<table id="tree" name="{{message "tmpl.tab.tree.name"}}" extends="sys" comment="{{message "tmpl.tab.tree.memo"}}">
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
		if v.Project == "demo" && v.ID == "app" {
			if util.InitFile("./conf/schema/tables-demo-app.xml", app) == false {
				isExist = false
			}
		} else if v.Project == "demo" && v.ID == "sys" {
			if util.InitFile("./conf/schema/tables-demo-sys.xml", sys) == false {
				isExist = false
			}
		} else {
			if util.InitFile("./conf/schema/tables-"+v.Project+"-"+v.ID+".xml", content) == false {
				isExist = false
			}
		}
	}
	return isExist
}
