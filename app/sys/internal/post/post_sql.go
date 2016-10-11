package post

const (
	listMenuIdByPostId     = `select menu_id from sys_post_menu where deletion = 0 and post_id = ?`
	deletePostMenuByPostId = `delete from sys_post_menu where post_id = ?`
)
