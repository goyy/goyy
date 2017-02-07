// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplDataPost = `
insert into {{case "SYS_POST (ID, CODE, NAME, FULLNAME, GENRE, ORDINAL, PARENT_ID, PARENT_IDS, PARENT_CODES, PARENT_NAMES, LEAF, GRADE, MEMO, CREATES, CREATER, CREATED, MODIFIER, MODIFIED, VERSION, DELETION, ARTIFICAL, HISTORY)"}}
values ('root', '00', '{{message "tmpl.data.post.root"}}', null, '00', '00', null, null, null, null, 0, 1, '{{message "tmpl.data.post.root"}}', 'root', null, ` + now + `, null, ` + now + `, 0, 0, 0, 0){{seperator}}

insert into {{case "SYS_POST (ID, CODE, NAME, FULLNAME, GENRE, ORDINAL, PARENT_ID, PARENT_IDS, PARENT_CODES, PARENT_NAMES, LEAF, GRADE, MEMO, CREATES, CREATER, CREATED, MODIFIER, MODIFIED, VERSION, DELETION, ARTIFICAL, HISTORY)"}}
values ('admin', '10', '{{message "tmpl.data.post.admin"}}', '{{message "tmpl.data.post.admin"}}', '20', '10', 'root', 'root', '00', '{{message "tmpl.data.post.root"}}', 1, 2, '{{message "tmpl.data.post.admin"}}', 'root', null, ` + now + `, null, ` + now + `, 0, 0, 0, 0){{seperator}}
`
