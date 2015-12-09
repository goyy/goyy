package internal

import (
	"bytes"
	"gopkg.in/goyy/goyy.v0/data/dialect"
	"gopkg.in/goyy/goyy.v0/util/errors"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

func (me *PermissionManager) SelectPermission(id string) (string, error) {
	if strings.IsNotBlank(id) {
		ps := NewPermissionEntities(200)
		var selectPermissionsByUserId string
		if me.Repository().Dialect().Type() == dialect.ORACLE {
			selectPermissionsByUserId = selectPermissionsByUserId_oracle
		} else {
			selectPermissionsByUserId = selectPermissionsByUserId_mysql
		}
		err := me.SelectList(ps, selectPermissionsByUserId, id)
		if err != nil {
			logger.Error(err.Error())
			return "", err
		}
		var b bytes.Buffer
		for i, p := range ps.Values() {
			if i > 0 {
				b.WriteString(",")
			}
			b.WriteString(p.Permission())
		}
		return b.String(), nil
	}
	return "", errors.NewNotBlank("loginName")
}
