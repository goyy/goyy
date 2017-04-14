package order

import (
	"gopkg.in/goyy/goyy.v0/data/domain"
)

func (me *Manager) GetCountByPrice() int {
	sPriceGT, _ := domain.NewSift("sPriceGT", "1000")    // 价格大于1000
	sDeletionEQ, _ := domain.NewSift("sDeletionEQ", "0") // 未删除
	out, err := Mgr.SelectCountBySift(sPriceGT, sDeletionEQ)
	if err != nil {
		logger.Error(err)
		return 0
	}
	return out
}

func (me *Manager) GetOrdersByPrice() (*Entities, error) {
	sPriceGT, _ := domain.NewSift("sPriceGT", "1000")    // 价格大于1000
	sDeletionEQ, _ := domain.NewSift("sDeletionEQ", "0") // 未删除
	sNumOD, _ := domain.NewSift("sNumOD", "10")          // 按购买数量倒序
	orders := NewEntities(10)
	err := Mgr.SelectListBySift(orders, sPriceGT, sDeletionEQ, sNumOD)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return orders, nil
}

func (me *Manager) FindBy(price float64, num int) (*Entities, error) {
	args := map[string]interface{}{
		"sPriceGT": price,
		"sNumGT":   num,
	}
	orders := NewEntities(10)
	err := Mgr.SelectListByNamed(orders, findBy_SQL, args)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return orders, nil
}
