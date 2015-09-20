// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of me source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package internal

const (
	repwdByUserId             = `update sys_user set passwd = ? where id = ?`
	countByLogin              = `select count(*) from sys_user where deletion = 0 and login_name = ? and passwd = ?`
	selectUserByLoginName     = `select id, name, login_name from sys_user where deletion = 0 and login_name = ?`
	selectPermissionsByUserId = `
select distinct m.permission
  from sys_menu m, sys_post_menu pm, sys_role_post rp, sys_user_role ur
 where m.deletion = 0
   and m.genre = '30'
   and m.id = pm.menu_id
   and pm.post_id = rp.post_id
   and rp.role_id = ur.role_id
   and ur.user_id = ?`
)
