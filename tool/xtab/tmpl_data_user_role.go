// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplDataUserRole = `
insert into {{case "SYS_USER_ROLE (ID, USER_ID, ROLE_ID, MEMO, CREATES, CREATER, CREATED, MODIFIER, MODIFIED, VERSION, DELETION, ARTIFICAL, HISTORY)"}}
values ('8199a34aa07511e4a96d005056a88ea8', 'admin', 'admin', '', '', '', ` + now + `, '', ` + now + `, 0, 0, 0, 0){{seperator}}
`
