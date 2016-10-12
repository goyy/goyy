package dict

const (
	ListByGenre = `select distinct genre from sys_dict where deletion = 0 order by genre`
)
