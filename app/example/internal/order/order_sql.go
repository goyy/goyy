package order

const (
	getCountByPrice_SQL = `select count(1) from eg_order where price > ? and deletion = 0`
	findByPrice_SQL     = `select * from eg_order where price > ? and deletion = 0 order by num desc`
	findBy_SQL          = `
select *
  from eg_order
 where 1 = 1
   {{if exist . "sPriceGT"}} and price > #{sPriceGT}{{end}}
   {{if exist . "sNumGT"}} and price > #{sNumGT}{{end}}
   and deletion = 0
 order by created desc
`
)
