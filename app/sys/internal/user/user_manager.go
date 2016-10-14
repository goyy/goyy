package user

import (
	"gopkg.in/goyy/goyy.v0/comm/xtype"
	"gopkg.in/goyy/goyy.v0/data/domain"
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/web/secure"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func (me *Manager) Save(p xtype.Principal, e entity.Interface) error {
	tx, err := me.DB().Begin()
	if err != nil {
		return err
	}
	err = me.Manager.Save(p, e)
	if err != nil {
		tx.Rollback()
		return err
	}
	userId := e.Get("id").(string)
	_, err = RoleMgr.Exec(deleteUserRoleByUserId, userId)
	if err != nil {
		tx.Rollback()
		return err
	}
	roleIds := e.Get("role_ids").(string)
	for _, roleId := range strings.Split(roleIds, ",") {
		ur := NewRoleEntity()
		ur.SetUserId(userId)
		ur.SetRoleId(roleId)
		err = RoleMgr.SaveAndTx(p, ur)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

// Get names based on ID.
func (me *Manager) Name(id string) string {
	if strings.IsBlank(id) {
		return ""
	}
	e := NewEntity()
	e.SetId(id)
	if err := me.Get(e); err == nil {
		return e.Name()
	}
	return ""
}

// Verify that the password is correct.
func (me *Manager) CheckPasswd(loginName, passwd string) bool {
	if strings.IsNotBlank(loginName) && strings.IsNotBlank(passwd) {
		sDeletionEQ, _ := domain.NewSift("sDeletionEQ", "0")
		sLoginNameEQ, _ := domain.NewSift("sLoginNameEQ", loginName)
		sPasswdEQ, _ := domain.NewSift("sPasswdEQ", secure.EncryptPasswd(passwd))
		count, err := me.SelectCountBySift(sDeletionEQ, sLoginNameEQ, sPasswdEQ)
		if err != nil {
			logger.Error(err.Error())
			return false
		}
		if count == 1 {
			return true
		}
	}
	return false
}

// Modify password.
func (me *Manager) Repwd(c xhttp.Context, userId, passwd string) error {
	tx, _ := me.DB().Begin()
	_, err := me.DB().Exec(repwdByUserId, secure.EncryptPasswd(passwd), userId)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
