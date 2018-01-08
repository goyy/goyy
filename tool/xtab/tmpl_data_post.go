// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplDataPost = `
insert into {{case "SYS_POST (ID, CODE, NAME, FULLNAME, GENRE, ORDINAL, PARENT_ID, PARENT_IDS, PARENT_CODES, PARENT_NAMES, LEAF, GRADE, MEMO, CREATES, CREATER, CREATED, MODIFIER, MODIFIED, VERSION, DELETION, ARTIFICAL, HISTORY)"}}
values ('root', '00', '{{message "tmpl.data.post.root"}}', '', '00', '00', '', '', '', '', 0, 1, '{{message "tmpl.data.post.root"}}', 'root', '', ` + now + `, '', ` + now + `, 0, 0, 0, 0){{seperator}}

insert into {{case "SYS_POST (ID, CODE, NAME, FULLNAME, GENRE, ORDINAL, PARENT_ID, PARENT_IDS, PARENT_CODES, PARENT_NAMES, LEAF, GRADE, MEMO, CREATES, CREATER, CREATED, MODIFIER, MODIFIED, VERSION, DELETION, ARTIFICAL, HISTORY)"}}
values ('admin', '10', '{{message "tmpl.data.post.admin"}}', '{{message "tmpl.data.post.admin"}}', '20', '10', 'root', 'root', '00', '{{message "tmpl.data.post.root"}}', 1, 2, '{{message "tmpl.data.post.admin"}}', 'root', '', ` + now + `, '', ` + now + `, 0, 0, 0, 0){{seperator}}

insert into {{case "SYS_POST (ID, CODE, NAME, FULLNAME, GENRE, ORDINAL, PARENT_ID, PARENT_IDS, PARENT_CODES, PARENT_NAMES, LEAF, GRADE, MEMO, CREATES, CREATER, CREATED, MODIFIER, MODIFIED, VERSION, DELETION, ARTIFICAL, HISTORY)"}}
values ('603ed6e46d9e11e5b888902b34947e43', '15', '{{message "tmpl.data.post.eg"}}', '{{message "tmpl.data.post.eg"}}', '20', '15', 'root', 'root', '00', '{{message "tmpl.data.post.root"}}', 0, 2, '{{message "tmpl.data.post.eg"}}', '', '029b7954e2fd11e3a512541fffec5618', ` + now + `, '029b7954e2fd11e3a512541fffec5618', ` + now + `, 0, 0, 0, 0){{seperator}}

insert into {{case "SYS_POST (ID, CODE, NAME, FULLNAME, GENRE, ORDINAL, PARENT_ID, PARENT_IDS, PARENT_CODES, PARENT_NAMES, LEAF, GRADE, MEMO, CREATES, CREATER, CREATED, MODIFIER, MODIFIED, VERSION, DELETION, ARTIFICAL, HISTORY)"}}
values ('2ca266227d4e11e58919005056bbc77f', '1510', '{{message "tmpl.data.post.eg.show"}}', '{{message "tmpl.data.post.eg"}} - {{message "tmpl.data.post.eg.show"}}', '30', '10', '603ed6e46d9e11e5b888902b34947e43', 'root,603ed6e46d9e11e5b888902b34947e43', '00,15', '{{message "tmpl.data.post.root"}},{{message "tmpl.data.post.eg"}}', 1, 3, '{{message "tmpl.data.post.eg.show"}}', '', '029b7954e2fd11e3a512541fffec5618', ` + now + `, '029b7954e2fd11e3a512541fffec5618', ` + now + `, 0, 0, 0, 0){{seperator}}

insert into {{case "SYS_POST (ID, CODE, NAME, FULLNAME, GENRE, ORDINAL, PARENT_ID, PARENT_IDS, PARENT_CODES, PARENT_NAMES, LEAF, GRADE, MEMO, CREATES, CREATER, CREATED, MODIFIER, MODIFIED, VERSION, DELETION, ARTIFICAL, HISTORY)"}}
values ('374ef7c87d4e11e58919005056bbc77f', '1520', '{{message "tmpl.data.post.eg.product"}}', '{{message "tmpl.data.post.eg"}} - {{message "tmpl.data.post.eg.product"}}', '30', '20', '603ed6e46d9e11e5b888902b34947e43', 'root,603ed6e46d9e11e5b888902b34947e43', '00,15', '{{message "tmpl.data.post.root"}},{{message "tmpl.data.post.eg"}}', 1, 3, '{{message "tmpl.data.post.eg.product"}}', '', '029b7954e2fd11e3a512541fffec5618', ` + now + `, '029b7954e2fd11e3a512541fffec5618', ` + now + `, 0, 0, 0, 0){{seperator}}

insert into {{case "SYS_POST (ID, CODE, NAME, FULLNAME, GENRE, ORDINAL, PARENT_ID, PARENT_IDS, PARENT_CODES, PARENT_NAMES, LEAF, GRADE, MEMO, CREATES, CREATER, CREATED, MODIFIER, MODIFIED, VERSION, DELETION, ARTIFICAL, HISTORY)"}}
values ('44fc4abe7d4e11e58919005056bbc77f', '1530', '{{message "tmpl.data.post.eg.order"}}', '{{message "tmpl.data.post.eg"}} - {{message "tmpl.data.post.eg.order"}}', '30', '30', '603ed6e46d9e11e5b888902b34947e43', 'root,603ed6e46d9e11e5b888902b34947e43', '00,15', '{{message "tmpl.data.post.root"}},{{message "tmpl.data.post.eg"}}', 1, 3, '{{message "tmpl.data.post.eg.order"}}', '', '029b7954e2fd11e3a512541fffec5618', ` + now + `, '029b7954e2fd11e3a512541fffec5618', ` + now + `, 0, 0, 0, 0){{seperator}}
`
