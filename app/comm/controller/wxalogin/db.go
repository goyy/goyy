package wxalogin

import (
	"fmt"

	"gopkg.in/goyy/goyy.v0/data/service"
	"gopkg.in/goyy/goyy.v0/util/uuids"
)

func createUser(code, loginName, passwd string) string {
	isql := fmt.Sprintf("INSERT INTO %s (%s, %s, %s, %s) VALUES (?, ?, ?, ?)", User.TableName, User.Id, User.Code, User.Passwd, User.LoginName)
	id := uuids.New()
	_, err := service.DB.Exec(isql, id, code, loginName, passwd)
	if err != nil {
		logger.Errorln("createUser err:", err)
		return ""
	}
	return id
}

func getUserId(loginName string) (string, bool) {
	csql := fmt.Sprintf("SELECT count(1) FROM %s WHERE %s = ?", User.TableName, User.LoginName)
	count, err := service.DB.Query(csql, loginName).Int()
	if err != nil {
		logger.Errorln("getUserId.count err:", err)
		return "", false
	}
	if count == 1 {
		idsql := fmt.Sprintf("SELECT %v FROM %s WHERE %s = ?", User.Id, User.TableName, User.LoginName)
		id, err := service.DB.Query(idsql, loginName).Str()
		if err != nil {
			logger.Errorln("getUserId.id err:", err)
			return "", false
		}
		return id, true
	}
	return "", false
}
