// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

type valids struct{}

func (me *valids) IsExistXML(name, pkg string) {
	isPrint := false
	if me.IsExistSettings() == false {
		isPrint = true
	}
	if me.IsExistEnvironments(name) == false {
		isPrint = true
	}
	if me.IsExistProjects(name, pkg) == false {
		isPrint = true
	}
	if me.IsExistModules(name, pkg) == false {
		isPrint = true
	}
	if me.IsExistButtons() == false {
		isPrint = true
	}
	if me.IsExistDomains() == false {
		isPrint = true
	}
	if me.IsExistColumns() == false {
		isPrint = true
	}
	if me.IsExistTables() == false {
		isPrint = true
	}
	if me.IsExistProjectTables() == false {
		isPrint = true
	}
	if isPrint == true {
		logger.Println("Create and initialize the xmlfiles")
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
		<module id="example" name="{{message "tmpl.mod.eg.name"}}" prefix="eg" project="` + name + `" generate="true" comment="{{message "tmpl.mod.eg.memo"}}" apipath="` + pkg + `/` + name + `-example"/>
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
		<domain id="id"      name="ID"      types="string" length="50"/>
		<domain id="flag"    name="FLAG"    types="string" length="10"/>
		<domain id="code"    name="CODE"    types="string" length="100"/>
		<domain id="name"    name="NAME"    types="string" length="255"/>
		<domain id="memo"    name="MEMO"    types="string" length="1000"/>
		<domain id="descr"   name="DESCR"   types="string" length="2000"/>
		<domain id="remark"  name="REMARK"  types="string" length="4000"/>
		<domain id="int"     name="INT"     types="int"/>
		<domain id="long"    name="LONG"    types="long"/>
		<domain id="float"   name="FLOAT"   types="float"  length="11" precision="3"/>
		<domain id="decimal" name="DECIMAL" types="float"  length="11" precision="6"/>
		<domain id="bool"    name="BOOL"    types="bool"/>
		<domain id="time"    name="TIME"    types="time"/>
		<domain id="text"    name="TEXT"    types="text"/>
		<domain id="bytes"   name="BYTES"   types="bytes"/>
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
		<column id="kind"         domain="code"   name="{{message "tmpl.col.kind.name"}}" comment="{{message "tmpl.col.kind.memo"}}"/>
		<column id="genre"        domain="code"   name="{{message "tmpl.col.genre.name"}}" comment="{{message "tmpl.col.genre.memo"}}"/>
		<column id="classify"     domain="code"   name="{{message "tmpl.col.classify.name"}}" comment="{{message "tmpl.col.classify.memo"}}"/>
		<column id="content"      domain="descr"  name="{{message "tmpl.col.content.name"}}" comment="{{message "tmpl.col.content.memo"}}"/>
		<column id="title"        domain="name"   name="{{message "tmpl.col.title.name"}}" comment="{{message "tmpl.col.title.memo"}}"/>
		<column id="memo"         domain="memo"   name="{{message "tmpl.col.memo.name"}}" comment="{{message "tmpl.col.memo.memo"}}"/>
		<column id="descr"        domain="descr"  name="{{message "tmpl.col.descr.name"}}" comment="{{message "tmpl.col.descr.memo"}}"/>
		<column id="remark"       domain="remark" name="{{message "tmpl.col.remark.name"}}" comment="{{message "tmpl.col.remark.memo"}}"/>
		<column id="timed"        domain="time"   name="{{message "tmpl.col.timed.name"}}" comment="{{message "tmpl.col.timed.memo"}}"/>
		<column id="passwd"       domain="name"   name="{{message "tmpl.col.passwd.name"}}" comment="{{message "tmpl.col.passwd.memo"}}"/>
		<column id="email"        domain="name"   name="{{message "tmpl.col.email.name"}}" comment="{{message "tmpl.col.email.memo"}}"/>
		<column id="mobile"       domain="code"   name="{{message "tmpl.col.mobile.name"}}" comment="{{message "tmpl.col.mobile.memo"}}"/>
		<column id="tel"          domain="code"   name="{{message "tmpl.col.tel.name"}}" comment="{{message "tmpl.col.tel.memo"}}"/>
		<column id="contact"      domain="name"   name="{{message "tmpl.col.contact.name"}}" comment="{{message "tmpl.col.contact.memo"}}"/>
		<column id="age"          domain="int"    name="{{message "tmpl.col.age.name"}}" comment="{{message "tmpl.col.age.memo"}}"/>
		<column id="sex"          domain="flag"   name="{{message "tmpl.col.sex.name"}}" comment="{{message "tmpl.col.sex.memo"}}" dict="{{message "tmpl.col.sex.dict"}}"/>
		<column id="birthday"     domain="long"   name="{{message "tmpl.col.birthday.name"}}" comment="{{message "tmpl.col.birthday.memo"}}"/>
		<column id="inc"          domain="name"   name="{{message "tmpl.col.inc.name"}}" comment="{{message "tmpl.col.inc.memo"}}"/>
		<column id="position"     domain="name"   name="{{message "tmpl.col.position.name"}}" comment="{{message "tmpl.col.position.memo"}}"/>
		<column id="receiver"     domain="code"   name="{{message "tmpl.col.receiver.name"}}" comment="{{message "tmpl.col.receiver.memo"}}"/>
		<column id="address"      domain="memo"   name="{{message "tmpl.col.address.name"}}" comment="{{message "tmpl.col.address.memo"}}"/>
		<column id="zipcode"      domain="code"   name="{{message "tmpl.col.zipcode.name"}}" comment="{{message "tmpl.col.zipcode.memo"}}"/>
		<column id="intro"        domain="remark" name="{{message "tmpl.col.intro.name"}}" comment="{{message "tmpl.col.intro.memo"}}"/>
		<column id="weight"       domain="int"    name="{{message "tmpl.col.weight.name"}}" comment="{{message "tmpl.col.weight.memo"}}"/>
		<column id="price"        domain="float"  name="{{message "tmpl.col.price.name"}}" comment="{{message "tmpl.col.price.memo"}}"/>
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
		<column id="leaf"         domain="bool"   name="{{message "tmpl.col.leaf.name"}}" comment="{{message "tmpl.col.leaf.memo"}}"/>
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
	eg := `<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE xtab PUBLIC "-//GOYY//DTD XTAB 1.0//EN" "http://goyy.org/dtd/xtab-1.0.dtd">
<configuration>
	<tables>
		<table id="product" name="{{message "tmpl.tab.eg.product.name"}}" extends="sys" generate="true" comment="{{message "tmpl.tab.eg.product.memo"}}">
			<column extends="name"/>
			<column extends="genre"/>
			<column id="num" name="{{message "tmpl.tab.eg.product.num.name"}}" domain="int" comment="{{message "tmpl.tab.eg.product.num.memo"}}"/>
			<column extends="price"/>
		</table>
		<table id="discount" name="{{message "tmpl.tab.eg.discount.name"}}" extends="sys" generate="false" comment="{{message "tmpl.tab.eg.discount.memo"}}">
			<column id="extent" name="{{message "tmpl.tab.eg.discount.extent.name"}}" domain="float" comment="{{message "tmpl.tab.eg.discount.extent.memo"}}" default="1"/>
		</table>
		<table id="order" name="{{message "tmpl.tab.eg.order.name"}}" extends="sys" generate="true" comment="{{message "tmpl.tab.eg.order.memo"}}">
			<column id="product_id" name="{{message "tmpl.tab.eg.order.product_id.name"}}" domain="id" comment="{{message "tmpl.tab.eg.order.product_id.memo"}}"/>
			<column id="discount_id" name="{{message "tmpl.tab.eg.order.discount_id.name"}}" domain="id" comment="{{message "tmpl.tab.eg.order.discount_id.memo"}}"/>
			<column id="num" name="{{message "tmpl.tab.eg.order.num.name"}}" domain="int" comment="{{message "tmpl.tab.eg.order.num.memo"}}" default="1"/>
			<column extends="price"/>
		</table>
	</tables>
</configuration>
`
	sys := `<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE xtab PUBLIC "-//GOYY//DTD XTAB 1.0//EN" "http://goyy.org/dtd/xtab-1.0.dtd">
<configuration>
	<tables>
		<table id="menu" name="{{message "tmpl.tab.sys.menu.name"}}" extends="tree" generate="true" comment="{{message "tmpl.tab.sys.menu.memo"}}">
			<column id="href" name="{{message "tmpl.tab.sys.menu.href.name"}}" domain="memo" comment="{{message "tmpl.tab.sys.menu.href.memo"}}"/>
			<column id="target" name="{{message "tmpl.tab.sys.menu.target.name"}}" domain="code" comment="{{message "tmpl.tab.sys.menu.target.memo"}}"/>
			<column id="icon" name="{{message "tmpl.tab.sys.menu.icon.name"}}" domain="code" comment="{{message "tmpl.tab.sys.menu.icon.memo"}}"/>
			<column id="hidden" name="{{message "tmpl.tab.sys.menu.hidden.name"}}" domain="bool" comment="{{message "tmpl.tab.sys.menu.hidden.memo"}}" dict="{{message "tmpl.tab.sys.menu.hidden.dict"}}"/>
			<column id="permission" name="{{message "tmpl.tab.sys.menu.perm.name"}}" domain="memo" comment="{{message "tmpl.tab.sys.menu.perm.memo"}}"/>
			<column extends="genre" name="{{message "tmpl.tab.sys.menu.genre.name"}}" comment="{{message "tmpl.tab.sys.menu.genre.memo"}}" dict="{{message "tmpl.tab.sys.menu.genre.dict"}}"/>
		</table>
		<table id="post" name="{{message "tmpl.tab.sys.post.name"}}" extends="tree" generate="true" comment="{{message "tmpl.tab.sys.post.memo"}}">
			<column id="is_admin" name="{{message "tmpl.tab.sys.post.is_admin.name"}}" domain="bool" comment="{{message "tmpl.tab.sys.post.is_admin.memo"}}" dict="{{message "tmpl.tab.sys.post.is_admin.dict"}}" default="0"/>
			<column extends="genre" name="{{message "tmpl.tab.sys.post.genre.name"}}" comment="{{message "tmpl.tab.sys.post.genre.memo"}}" dict="{{message "tmpl.tab.sys.post.genre.dict"}}"/>
		</table>
		<table id="post_menu" name="{{message "tmpl.tab.sys.post_menu.name"}}" extends="sys" generate="true" menu="false" comment="{{message "tmpl.tab.sys.post_menu.memo"}}" master="post" slave="menu">
			<column extends="post_id"/>
			<column extends="menu_id"/>
		</table>
		<table id="role" name="{{message "tmpl.tab.sys.role.name"}}" extends="sys" generate="true" comment="{{message "tmpl.tab.sys.role.memo"}}">
			<column extends="name"/>
			<column extends="code"/>
			<column id="is_admin" name="{{message "tmpl.tab.sys.role.is_admin.name"}}" domain="bool" comment="{{message "tmpl.tab.sys.role.is_admin.memo"}}" dict="{{message "tmpl.tab.sys.role.is_admin.dict"}}" default="0"/>
			<column extends="genre" name="{{message "tmpl.tab.sys.role.genre.name"}}" comment="{{message "tmpl.tab.sys.role.genre.memo"}}" dict="{{message "tmpl.tab.sys.role.genre.dict"}}"/>
			<column extends="classify" name="{{message "tmpl.tab.sys.role.classify.name"}}" comment="{{message "tmpl.tab.sys.role.classify.memo"}}"/>
			<column extends="ordinal"/>
		</table>
		<table id="role_post" name="{{message "tmpl.tab.sys.role_post.name"}}" extends="sys" generate="true" menu="false" comment="{{message "tmpl.tab.sys.role_post.memo"}}" master="role" slave="post">
			<column extends="role_id"/>
			<column extends="post_id"/>
		</table>
		<table id="-" name="-" generate="false" menu="true" permissions="sys:area:view,sys:org:view,sys:user:view"/>
		<table id="area" name="{{message "tmpl.tab.sys.area.name"}}" extends="tree" generate="true" comment="{{message "tmpl.tab.sys.area.memo"}}">
		</table>
		<table id="org" name="{{message "tmpl.tab.sys.org.name"}}" extends="tree" generate="true" comment="{{message "tmpl.tab.sys.org.memo"}}">
			<column extends="area_id"/>
		</table>
		<table id="user" name="{{message "tmpl.tab.sys.user.name"}}" extends="sys" generate="true" comment="{{message "tmpl.tab.sys.user.memo"}}">
			<column extends="name" unique="true"/>
			<column extends="code"/>
			<column extends="passwd"/>
			<column extends="genre" dict="{{message "tmpl.tab.sys.user.genre.dict"}}"/>
			<column extends="email"/>
			<column extends="tel"/>
			<column extends="mobile"/>
			<column extends="area_id"/>
			<column extends="org_id"/>
			<column id="dimission" name="{{message "tmpl.tab.sys.user.dimission.name"}}" domain="bool" comment="{{message "tmpl.tab.sys.user.dimission.memo"}}" dict="{{message "tmpl.tab.sys.user.dimission.dict"}}" default="0"/>
			<column id="dimission_time" name="{{message "tmpl.tab.sys.user.dimission_time.name"}}" domain="long" comment="{{message "tmpl.tab.sys.user.dimission_time.memo"}}" default="-62135596800"/>
			<column id="freeze" name="{{message "tmpl.tab.sys.user.freeze.name"}}" domain="bool" comment="{{message "tmpl.tab.sys.user.freeze.memo"}}" dict="{{message "tmpl.tab.sys.user.freeze.dict"}}" default="0"/>
			<column id="freeze_time" name="{{message "tmpl.tab.sys.user.freeze_time.name"}}" domain="long" comment="{{message "tmpl.tab.sys.user.freeze_time.memo"}}" default="-62135596800"/>
			<column id="login_name" name="{{message "tmpl.tab.sys.user.login_name.name"}}" domain="code" comment="{{message "tmpl.tab.sys.user.login_name.memo"}}" unique="true"/>
		</table>
		<table id="user_role" name="{{message "tmpl.tab.sys.user_role.name"}}" extends="sys" generate="true" menu="false" comment="{{message "tmpl.tab.sys.user_role.memo"}}" master="user" slave="role">
			<column extends="user_id"/>
			<column extends="role_id"/>
		</table>
        <table id="user_login" name="{{message "tmpl.tab.sys.user_login.name"}}" extends="sys" generate="true" comment="{{message "tmpl.tab.sys.user_login.memo"}}">
            <column id="profiles" name="{{message "tmpl.tab.sys.user_login.profiles.name"}}" domain="code" comment="{{message "tmpl.tab.sys.user_login.profiles.memo"}}"/>
            <column extends="user_id"/>
			<column id="login_name" name="{{message "tmpl.tab.sys.user.login_name.name"}}" domain="code" comment="{{message "tmpl.tab.sys.user.login_name.memo"}}"/>
			<column id="login_ip" name="{{message "tmpl.tab.sys.user.login_ip.name"}}" domain="code" comment="{{message "tmpl.tab.sys.user.login_ip.memo"}}"/>
			<column id="login_time" name="{{message "tmpl.tab.sys.user.login_time.name"}}" domain="long" comment="{{message "tmpl.tab.sys.user.login_time.memo"}}"/>
        </table>
		<table id="-" name="-" generate="false" menu="true" permissions="sys:dict:view"/>
		<table id="dict" name="{{message "tmpl.tab.sys.dict.name"}}" extends="sys" generate="true" comment="{{message "tmpl.tab.sys.dict.memo"}}">
			<column extends="genre"/>
			<column extends="descr"/>
			<column extends="mkey"/>
			<column extends="mval"/>
			<column extends="filters"/>
			<column extends="ordinal"/>
		</table>
		<table id="-" name="-" generate="false" menu="true" permissions="sys:cache:view"/>
		<table id="cache" name="{{message "tmpl.tab.sys.cache"}}" generate="false" menu="true"/>
		<table id="blacklist" name="{{message "tmpl.tab.sys.blacklist.name"}}" extends="sys" generate="true" menu="false" comment="{{message "tmpl.tab.sys.blacklist.memo"}}">
			<column extends="name"/>
			<column extends="genre"/>
			<column extends="usable"/>
		</table>
		<table id="conf" name="{{message "tmpl.tab.sys.conf.name"}}" extends="sys" generate="true" menu="false" comment="{{message "tmpl.tab.sys.conf.name"}}">
			<column extends="name"/>
			<column extends="code"/>
			<column extends="content"/>
			<column extends="genre"/>
			<column extends="usable"/>
			<column extends="ordinal"/>
		</table>
	</tables>
</configuration>
`
	content := `<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE xtab PUBLIC "-//GOYY//DTD XTAB 1.0//EN" "http://goyy.org/dtd/xtab-1.0.dtd">
<configuration>
	<tables>
		<table id="table1" name="TABLE1" extends="sys" comment="table1">
			<column extends="name"/>
		</table>
	</tables>
</configuration>
`
	xconf := util.DecodeXML(xmodules)
	isExist := true
	for _, v := range xconf.Modules.Module {
		if v.ID == "example" {
			if util.InitFile("./conf/schema/tables-"+v.Project+"-example.xml", eg) == false {
				isExist = false
			}
		} else if v.ID == "sys" {
			if util.InitFile("./conf/schema/tables-"+v.Project+"-sys.xml", sys) == false {
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
