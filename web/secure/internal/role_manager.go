package internal

import (
	"gopkg.in/goyy/goyy.v0/data/dialect"
	"gopkg.in/goyy/goyy.v0/util/errors"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

func (me *RoleManager) SelectRole(userId string) ([]*Role, error) {
	if strings.IsNotBlank(userId) {
		out := NewRoleEntities(10)
		var selectRoleByUserId string
		if me.DB().Dialect().Type() == dialect.ORACLE {
			selectRoleByUserId = selectRoleByUserId_oracle
		} else {
			selectRoleByUserId = selectRoleByUserId_mysql
		}
		err := me.SelectList(out, selectRoleByUserId, userId)
		if err != nil {
			logger.Error(err.Error())
			return nil, err
		}
		return out.Values(), nil
	}
	return nil, errors.NewNotBlank("userId")
}
