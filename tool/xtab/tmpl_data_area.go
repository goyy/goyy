// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplDataArea = `
insert into {{case "SYS_AREA (ID, CODE, NAME, FULLNAME, GENRE, ORDINAL, PARENT_ID, PARENT_IDS, PARENT_CODES, PARENT_NAMES, LEAF, GRADE, MEMO, CREATES, CREATER, CREATED, MODIFIER, MODIFIED, VERSION, DELETION, ARTIFICAL, HISTORY)"}}
values ('root', '00', '{{message "tmpl.data.area"}}', null, null, '00', null, null, null, null, 0, 1, '{{message "tmpl.data.area"}}', null, null, ` + now + `, null, ` + now + `, 0, 0, 0, 0){{seperator}}
`
