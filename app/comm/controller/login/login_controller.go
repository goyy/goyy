package login

import (
	"gopkg.in/goyy/goyy.v0/comm/captcha"
	"gopkg.in/goyy/goyy.v0/util/crypto/aes"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/web/controller"
	"gopkg.in/goyy/goyy.v0/web/secure"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func init() {
	xhttp.GET("/login", ctl.login)
	xhttp.POST("/signin", ctl.signin)
	xhttp.GET("/logout", ctl.logout)
}

func (me *Controller) login(c xhttp.Context) {
	c.Redirect(xhttp.Conf.Secure.LoginUrl)
}

func (me *Controller) logout(c xhttp.Context) {
	err := secure.Logout(c)
	if err != nil {
		logger.Errorln("logout:", err)
		c.Redirect(xhttp.Conf.Err.Err500)
		return
	}
	c.Redirect(xhttp.Conf.Secure.LoginUrl)
}

func (me *Controller) signin(c xhttp.Context) {
	var err error
	if !captcha.Verify(c, true) {
		msg := i18N.Message("login.captcha.err")
		err = c.JSON(xhttp.StatusOK, me.FaultStatusMsg(c, "-1", msg, ""))
		if err != nil {
			logger.Errorln("response err:", err)
		}
		return
	}
	err = secure.Login(c, c.Param("loginName"), c.Param("passwd"))
	if err != nil {
		msg := i18N.Message("login.user.err")
		err = c.JSON(xhttp.StatusOK, me.FaultStatusMsg(c, "-100", msg, ""))
		if err != nil {
			logger.Errorln("response err:", err)
		}
		return
	}
	redirect := c.Param("redirect")
	if strings.IsNotBlank(redirect) {
		if url, err := aes.DecryptHex(redirect, aes.DefaultKey); err == nil {
			redirect = url
		} else {
			redirect = "/"
		}
	} else {
		redirect = "/"
	}
	msg := i18N.Message("login.user.ok")
	err = c.JSON(xhttp.StatusOK, me.SuccessMsg(c, msg, redirect))
	if err != nil {
		logger.Errorln("response err:", err)
	}
}

type Controller struct {
	controller.JSONController
}

var ctl = &Controller{}
