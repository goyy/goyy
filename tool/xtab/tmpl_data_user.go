// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplDataUser = `
insert into {{case "SYS_USER (ID, NAME, CODE, PASSWD, GENRE, EMAIL, TEL, MOBILE, AREA_ID, ORG_ID, DIMISSION_TIME, FREEZE_TIME, LOGIN_NAME, MEMO, CREATES, CREATER, CREATED, MODIFIER, MODIFIED, VERSION, DELETION, ARTIFICAL, HISTORY)"}}
values ('admin', '{{message "tmpl.data.user.admin"}}', null, '92d55a4a6b07', '10', 'admin@goyy.org', null, null, 'root', 'root', -62135596800, -62135596800, 'admin', null, null, null, ` + now + `, null, ` + now + `, 0, 0, 0, 0){{seperator}}
`
