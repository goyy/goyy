package role

const (
	listPostIdByRoleId     = `select post_id from sys_role_post where deletion = 0 and role_id = ?`
	deleteRolePostByRoleId = `delete from sys_role_post where role_id = ?`
)
