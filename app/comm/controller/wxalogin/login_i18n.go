package wxalogin

import (
	"gopkg.in/goyy/goyy.v0/comm/i18n"
)

func SetLocale(locale string) {
	i18N.SetLocale(locale)
}

var i18N = i18n.New(locales)

var locales = map[string]map[string]string{
	i18n.LocaleZhCN: map[string]string{
		"login.captcha.err": "验证码有误，请重新输入",
		"login.user.err":    "用户名或密码不正确",
		"login.user.ok":     "登录成功",
	},
	i18n.LocaleEnUS: map[string]string{
		"login.captcha.err": "Captcha is incorrect. Please re-enter it.",
		"login.user.err":    "Username or password incorrect",
		"login.user.ok":     "Login success",
	},
}
