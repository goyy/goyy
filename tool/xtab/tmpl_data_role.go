// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplDataRole = `
insert into {{case "SYS_ROLE (ID, NAME, CODE, GENRE, ORDINAL, MEMO, CREATES, CREATER, CREATED, MODIFIER, MODIFIED, VERSION, DELETION, ARTIFICAL, HISTORY)"}}
values ('admin', '{{message "tmpl.data.role.admin"}}', '1900', '10', '99', '{{message "tmpl.data.role.admin"}}', null, null, ` + now + `, null, ` + now + `, 0, 0, 0, 0){{seperator}}
`
