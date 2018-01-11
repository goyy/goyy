package user

import (
	"bytes"

	"gopkg.in/goyy/goyy.v0/data/result"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/web/secure"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func init() {
	xhttp.GET(ctl.ApiIndex(), ctl.Index, ctl.PermitView())
	xhttp.POST(ctl.ApiIndex(), ctl.Index, ctl.PermitView())
	xhttp.GET(ctl.ApiShow(), ctl.Show, ctl.PermitView())
	xhttp.POST(ctl.ApiShow(), ctl.Show, ctl.PermitView())
	xhttp.POST(ctl.ApiAdd(), ctl.Add, ctl.PermitAdd())
	xhttp.POST(ctl.ApiEdit(), ctl.Edit, ctl.PermitEdit())
	xhttp.POST(ctl.ApiSave(), ctl.Save, ctl.PermitAdd(), ctl.PermitEdit())
	xhttp.POST(ctl.ApiDisable(), ctl.DisableAndTx, ctl.PermitDisable())
	xhttp.POST(ctl.ApiBy("repwd"), ctl.repwd)
	xhttp.GET(ctl.ApiBy("name"), ctl.name)
	// Gets the current login user has information
	xhttp.GET(ctl.ApiBy("principal"), ctl.principal)

	ctl.PreAdd = preSetRoleIds
	ctl.PreEdit = preSetRoleIds
	ctl.PreShow = preSetRoleIds
	ctl.PreSave = preSave
	ctl.PostSave = refresh
	ctl.PostDisable = refresh
}

// Gets the current login user has information.
func (me *Controller) principal(c xhttp.Context) {
	user, err := c.Session().Principal()
	if err != nil {
		c.JSON(xhttp.StatusOK, me.Fault(c, i18N.Message("pwd.get.err")))
		return
	}
	mp := make(map[string]string, 1)
	mp["loginName"] = user.LoginName
	mp["name"] = user.Name
	mp["loginTime"] = user.LoginTime
	c.JSON(xhttp.StatusOK, me.Success(c, mp))
	return
}

// Modify user login password.
func (me *Controller) repwd(c xhttp.Context) {
	oldpwd := c.Param("oldPasswd")
	newpwd := c.Param("newPasswd")
	okpwd := c.Param("okPasswd")
	if strings.IsBlank(oldpwd) || strings.IsBlank(newpwd) || strings.IsBlank(okpwd) {
		c.JSON(xhttp.StatusOK, me.Fault(c, i18N.Message("pwd.nil.err")))
		return
	}
	if newpwd != okpwd {
		c.JSON(xhttp.StatusOK, me.Fault(c, i18N.Message("pwd.ne.err")))
		return
	}
	p, err := c.Session().Principal()
	if err != nil {
		logger.Error(err.Error())
		me.Error(c, err)
		return
	}
	checked := Mgr.CheckPasswd(p.LoginName, oldpwd)
	if !checked {
		c.JSON(xhttp.StatusOK, me.Fault(c, i18N.Message("pwd.old.err")))
		return
	}
	err = Mgr.Repwd(c, p.Id, newpwd)
	if err != nil {
		me.Error(c, err)
		return
	}
	c.JSON(xhttp.StatusOK, me.Success(c, i18N.Message("pwd.set.ok")))
	return
}

// Get names based on ID.
func (me *Controller) name(c xhttp.Context) {
	val := Mgr.Name(c.Param("id"))
	c.JSON(xhttp.StatusOK, me.Success(c, val))
}

func preSetRoleIds(c xhttp.Context) error {
	userId := c.Param("sIdTR")
	var b bytes.Buffer
	userRoles := NewRoleEntities(0)
	err := RoleMgr.SelectList(userRoles, listRoleIdByUserId, userId)
	if err != nil {
		return err
	}
	for i, data := range userRoles.Values() {
		if i > 0 {
			b.WriteString(",")
		}
		b.WriteString(data.RoleId())
	}
	c.Params().Set("roleIds", b.String())
	return nil
}

// Encryption user password.
func preSave(c xhttp.Context) error {
	// If the password is empty or ●●●●●●, that does not modify the password
	pwdkey := "passwd"
	passwd := c.Param(pwdkey)
	if strings.IsBlank(passwd) || passwd == "●●●●●●" {
		c.Params().Del(pwdkey)
		return nil
	} else {
		passwd = secure.EncryptPasswd(passwd)
		c.Params().Set(pwdkey, passwd)
	}
	return nil
}

// Refresh the cache data for the user dictionary.
func refresh(c xhttp.Context, r *result.Entity) error {
	refreshCache()
	return nil
}
