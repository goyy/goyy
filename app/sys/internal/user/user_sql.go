package user

const (
	listRoleIdByUserId     = `select role_id from sys_user_role where deletion = 0 and user_id = ?`
	repwdByUserId          = `update sys_user set passwd = ? where id = ?`
	deleteUserRoleByUserId = `delete from sys_user_role where user_id = ?`
)
