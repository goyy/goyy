// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplDataDict = `
insert into {{case "SYS_DICT (ID, GENRE, DESCR, MKEY, MVAL, FILTERS, ORDINAL, MEMO, CREATES, CREATER, CREATED, MODIFIER, MODIFIED, VERSION, DELETION, ARTIFICAL, HISTORY)"}}
values ('53cc0ad27ad911e5b5a6902b34947e43', 'sys_menu.genre', '{{message "tmpl.data.dict.menu.genre"}}', '10', '{{message "tmpl.data.dict.menu.genre.10"}}', '', '10', '{{message "tmpl.data.dict.menu.genre"}}', '', '', ` + now + `, '', ` + now + `, 0, 0, 0, 0){{seperator}}

insert into {{case "SYS_DICT (ID, GENRE, DESCR, MKEY, MVAL, FILTERS, ORDINAL, MEMO, CREATES, CREATER, CREATED, MODIFIER, MODIFIED, VERSION, DELETION, ARTIFICAL, HISTORY)"}}
values ('53cc80057ad911e5b5a6902b34947e43', 'sys_menu.genre', '{{message "tmpl.data.dict.menu.genre"}}', '20', '{{message "tmpl.data.dict.menu.genre.20"}}', '', '20', '{{message "tmpl.data.dict.menu.genre"}}', '', '', ` + now + `, '', ` + now + `, 0, 0, 0, 0){{seperator}}

insert into {{case "SYS_DICT (ID, GENRE, DESCR, MKEY, MVAL, FILTERS, ORDINAL, MEMO, CREATES, CREATER, CREATED, MODIFIER, MODIFIED, VERSION, DELETION, ARTIFICAL, HISTORY)"}}
values ('53cd1c497ad911e5b5a6902b34947e43', 'sys_menu.genre', '{{message "tmpl.data.dict.menu.genre"}}', '30', '{{message "tmpl.data.dict.menu.genre.30"}}', '', '30', '{{message "tmpl.data.dict.menu.genre"}}', '', '', ` + now + `, '', ` + now + `, 0, 0, 0, 0){{seperator}}

insert into {{case "SYS_DICT (ID, GENRE, DESCR, MKEY, MVAL, FILTERS, ORDINAL, MEMO, CREATES, CREATER, CREATED, MODIFIER, MODIFIED, VERSION, DELETION, ARTIFICAL, HISTORY)"}}
values ('53cdb88d7ad911e5b5a6902b34947e43', 'sys_menu.hidden', '{{message "tmpl.data.dict.menu.hidden"}}', '0', '{{message "tmpl.data.dict.menu.hidden.0"}}', '', '10', '{{message "tmpl.data.dict.menu.hidden"}}', '', '', ` + now + `, '', ` + now + `, 0, 0, 0, 0){{seperator}}

insert into {{case "SYS_DICT (ID, GENRE, DESCR, MKEY, MVAL, FILTERS, ORDINAL, MEMO, CREATES, CREATER, CREATED, MODIFIER, MODIFIED, VERSION, DELETION, ARTIFICAL, HISTORY)"}}
values ('53ce2dc07ad911e5b5a6902b34947e43', 'sys_menu.hidden', '{{message "tmpl.data.dict.menu.hidden"}}', '1', '{{message "tmpl.data.dict.menu.hidden.1"}}', '', '20', '{{message "tmpl.data.dict.menu.hidden"}}', '', '', ` + now + `, '', ` + now + `, 0, 0, 0, 0){{seperator}}

insert into {{case "SYS_DICT (ID, GENRE, DESCR, MKEY, MVAL, FILTERS, ORDINAL, MEMO, CREATES, CREATER, CREATED, MODIFIER, MODIFIED, VERSION, DELETION, ARTIFICAL, HISTORY)"}}
values ('53ceca047ad911e5b5a6902b34947e43', 'sys_post.genre', '{{message "tmpl.data.dict.post.genre"}}', '10', '{{message "tmpl.data.dict.post.genre.10"}}', '', '10', '{{message "tmpl.data.dict.post.genre"}}', '', '', ` + now + `, '', ` + now + `, 0, 0, 0, 0){{seperator}}

insert into {{case "SYS_DICT (ID, GENRE, DESCR, MKEY, MVAL, FILTERS, ORDINAL, MEMO, CREATES, CREATER, CREATED, MODIFIER, MODIFIED, VERSION, DELETION, ARTIFICAL, HISTORY)"}}
values ('53cf66487ad911e5b5a6902b34947e43', 'sys_post.genre', '{{message "tmpl.data.dict.post.genre"}}', '20', '{{message "tmpl.data.dict.post.genre.20"}}', '', '20', '{{message "tmpl.data.dict.post.genre"}}', '', '', ` + now + `, '', ` + now + `, 0, 0, 0, 0){{seperator}}

insert into {{case "SYS_DICT (ID, GENRE, DESCR, MKEY, MVAL, FILTERS, ORDINAL, MEMO, CREATES, CREATER, CREATED, MODIFIER, MODIFIED, VERSION, DELETION, ARTIFICAL, HISTORY)"}}
values ('53cfdb7b7ad911e5b5a6902b34947e43', 'sys_post.genre', '{{message "tmpl.data.dict.post.genre"}}', '30', '{{message "tmpl.data.dict.post.genre.30"}}', '', '30', '{{message "tmpl.data.dict.post.genre"}}', '', '', ` + now + `, '', ` + now + `, 0, 0, 0, 0){{seperator}}

insert into {{case "SYS_DICT (ID, GENRE, DESCR, MKEY, MVAL, FILTERS, ORDINAL, MEMO, CREATES, CREATER, CREATED, MODIFIER, MODIFIED, VERSION, DELETION, ARTIFICAL, HISTORY)"}}
values ('53d077bf7ad911e5b5a6902b34947e43', 'sys_role.genre', '{{message "tmpl.data.dict.role.genre"}}', '10', '{{message "tmpl.data.dict.role.genre.10"}}', '', '10', '{{message "tmpl.data.dict.role.genre"}}', '', '', ` + now + `, '', ` + now + `, 0, 0, 0, 0){{seperator}}

insert into {{case "SYS_DICT (ID, GENRE, DESCR, MKEY, MVAL, FILTERS, ORDINAL, MEMO, CREATES, CREATER, CREATED, MODIFIER, MODIFIED, VERSION, DELETION, ARTIFICAL, HISTORY)"}}
values ('53d0ecf27ad911e5b5a6902b34947e43', 'sys_role.genre', '{{message "tmpl.data.dict.role.genre"}}', '20', '{{message "tmpl.data.dict.role.genre.20"}}', '', '20', '{{message "tmpl.data.dict.role.genre"}}', '', '', ` + now + `, '', ` + now + `, 0, 0, 0, 0){{seperator}}

insert into {{case "SYS_DICT (ID, GENRE, DESCR, MKEY, MVAL, FILTERS, ORDINAL, MEMO, CREATES, CREATER, CREATED, MODIFIER, MODIFIED, VERSION, DELETION, ARTIFICAL, HISTORY)"}}
values ('53d189367ad911e5b5a6902b34947e43', 'sys_user.genre', '{{message "tmpl.data.dict.user.genre"}}', '10', '{{message "tmpl.data.dict.user.genre.10"}}', '', '10', '{{message "tmpl.data.dict.user.genre"}}', '', '', ` + now + `, '', ` + now + `, 0, 0, 0, 0){{seperator}}

insert into {{case "SYS_DICT (ID, GENRE, DESCR, MKEY, MVAL, FILTERS, ORDINAL, MEMO, CREATES, CREATER, CREATED, MODIFIER, MODIFIED, VERSION, DELETION, ARTIFICAL, HISTORY)"}}
values ('53d1fe697ad911e5b5a6902b34947e43', 'sys_user.genre', '{{message "tmpl.data.dict.user.genre"}}', '20', '{{message "tmpl.data.dict.user.genre.20"}}', '', '20', '{{message "tmpl.data.dict.user.genre"}}', '', '', ` + now + `, '', ` + now + `, 0, 0, 0, 0){{seperator}}
`
