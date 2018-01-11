package wxalogin

import (
	"fmt"

	"gopkg.in/goyy/goyy.v0/data/service"
	"gopkg.in/goyy/goyy.v0/util/uuids"
)

func createUser(uid, code string) string {
	isql := fmt.Sprintf("INSERT INTO %s (%s, %s, %s, %s) VALUES (?, ?, ?, 'wxa')", User.TableName, User.Id, User.Uid, User.Code, User.Genre)
	id := uuids.New()
	_, err := service.DB.Exec(isql, id, uid, code)
	if err != nil {
		logger.Errorln("createUser err:", err)
		return ""
	}
	return id
}

func getUserId(uid string) (string, bool) {
	csql := fmt.Sprintf("SELECT count(1) FROM %s WHERE %s = ?", User.TableName, User.Uid)
	count, err := service.DB.Query(csql, uid).Int()
	if err != nil {
		logger.Errorln("getUserId.count err:", err)
		return "", false
	}
	if count == 1 {
		idsql := fmt.Sprintf("SELECT %v FROM %s WHERE %s = ?", User.Id, User.TableName, User.Uid)
		id, err := service.DB.Query(idsql, uid).Str()
		if err != nil {
			logger.Errorln("getUserId.id err:", err)
			return "", false
		}
		return id, true
	}
	return "", false
}
