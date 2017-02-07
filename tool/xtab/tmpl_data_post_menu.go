// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplDataPostMenu = `
-- MySQL
-- delete from {{case "SYS_POST_MENU"}} where {{case "POST_ID"}} = 'admin'{{seperator}}
insert into {{case "SYS_POST_MENU(ID,POST_ID,MENU_ID,MEMO,CREATES,CREATER,CREATED,MODIFIER,MODIFIED)"}}
select replace(uuid(), '-', ''),'admin',t.{{case "ID"}},t.{{case "MEMO"}},t.{{case "CREATES"}},t.{{case "CREATER"}},t.{{case "CREATED"}},t.{{case "MODIFIER"}},t.{{case "MODIFIED"}}
from {{case "SYS_MENU"}} t{{seperator}}

-- Oracle
-- delete from {{case "SYS_POST_MENU"}} where {{case "POST_ID"}} = 'admin'{{seperator}}
-- insert into {{case "SYS_POST_MENU(ID,POST_ID,MENU_ID,MEMO,CREATES,CREATER,CREATED,MODIFIER,MODIFIED)"}}
-- select sys_guid(),'admin',t.{{case "ID"}},t.{{case "MEMO"}},t.{{case "CREATES"}},t.{{case "CREATER"}},t.{{case "CREATED"}},t.{{case "MODIFIER"}},t.{{case "MODIFIED"}}
-- from {{case "SYS_MENU"}} t{{seperator}}
`
