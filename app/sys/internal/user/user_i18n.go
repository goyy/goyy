package user

import (
	"gopkg.in/goyy/goyy.v0/comm/i18n"
)

func SetLocale(locale string) {
	i18N.SetLocale(locale)
}

var i18N = i18n.New(locales)

var locales = map[string]map[string]string{
	i18n.LocaleZhCN: map[string]string{
		"pwd.get.err": "获取登录信息失败！",
		"pwd.nil.err": "密码不能为空！",
		"pwd.ne.err":  "新密码和确认新密码不相等！",
		"pwd.old.err": "旧密码错误！",
		"pwd.set.ok":  "密码修改成功！",
	},
	i18n.LocaleEnUS: map[string]string{
		"pwd.get.err": "Failed to get login information!",
		"pwd.nil.err": "Password can not be empty!",
		"pwd.ne.err":  "The new password and confirm the new password is not equal!",
		"pwd.old.err": "Old Password Error!",
		"pwd.set.ok":  "Password reset success!",
	},
}
