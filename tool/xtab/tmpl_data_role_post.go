// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplDataRolePost = `
insert into {{case "SYS_ROLE_POST (ID, ROLE_ID, POST_ID, MEMO, CREATES, CREATER, CREATED, MODIFIER, MODIFIED, VERSION, DELETION, ARTIFICAL, HISTORY)"}}
values ('6188defdc62f11e5bb26005056bbc77f', 'admin', 'admin', null, null, null, ` + now + `, null, ` + now + `, 0, 0, 0, 0){{seperator}}
`
