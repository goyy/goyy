// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package validate

import (
	"gopkg.in/goyy/goyy.v0/comm/i18n"
)

func SetLocale(locale string) {
	i18N.SetLocale(locale)
}

var i18N = i18n.New(locales)

// Validation errors.
var locales = map[string]map[string]string{
	i18n.Locale_zh_CN: map[string]string{
		typRequired:    "非空.",
		typMin:         "不能小于 %d.",
		typMax:         "不能大于 %d.",
		typRange:       "必须介于 %d 和 %d 之间.",
		typMinf:        "不能小于 %g.",
		typMaxf:        "不能大于 %g.",
		typRangef:      "必须介于 %g 和 %g 之间.",
		typMinlen:      "长度最小是 %d 的字符串（汉字算一个字符）.",
		typMaxlen:      "长度最大是 %d 的字符串（汉字算一个字符）.",
		typRangelen:    "长度必须介于 %d 和 %d 之间的字符串（汉字算一个字符）.",
		typFloat:       "必须是浮点类型.",
		typInteger:     "必须是整数类型.",
		typAlpha:       "只能包含大小写字母.",
		typAlrod:       "只能包含大小写字母或横杠（-_）.",
		typAlnum:       "只能包含大小写字母或数字.",
		typAlnumrod:    "只能包含大小写字母或数字或横杠（-_）.",
		typAlnumhan:    "只能包含大小写字母或数字或汉字.",
		typAlnumhanrod: "只能包含大小写字母或数字或汉字或横杠（-_）.",
		typAlhan:       "只能包含大小写字母或汉字.",
		typAlhanrod:    "只能包含大小写字母或汉字或横杠（-_）.",
		typHan:         "只能包含汉字.",
		typHanrod:      "只能包含汉字或横杠（-_）.",
		typMatch:       "必须正则匹配 s%.",
		typNomatch:     "必须正则不匹配 s%.",
		typEmail:       "必须是正确格式的电子邮件.",
		typURL:         "必须是正确格式的网址.",
		typIP:          "必须是正确格式的IP地址.",
		typMobile:      "必须是有效的手机号.",
		typTel:         "必须是有效的固定电话号.",
		typPhone:       "必须是有效的手机号或固定电话号.",
		typZipcode:     "必须是有效的邮政编码.",
	},
	i18n.Locale_en_US: map[string]string{
		typRequired:    "Expecting a non empty value.",
		typMin:         "Expecting a value greater than or equal to %d.",
		typMax:         "Expecting a value less than or equal to %d.",
		typRange:       "Expecting a value between %d and %d.",
		typMinf:        "Expecting a value greater than or equal to %g.",
		typMaxf:        "Expecting a value less than or equal to %g.",
		typRangef:      "Expecting a value between %g and %g.",
		typMinlen:      "Expecting at least %d characters.",
		typMaxlen:      "Expecting no more than %d characters.",
		typRangelen:    "Expecting a value between %d and %d characters long.",
		typFloat:       "Expecting a floating point number.",
		typInteger:     "Expecting an integer.",
		typAlpha:       "Expecting an alphabetic string.",
		typAlrod:       "Expecting an alphabetic or rod string.",
		typAlnum:       "Expecting an alphanumeric string.",
		typAlnumrod:    "Expecting an alphanumeric or rod string.",
		typAlnumhan:    "Expecting an alphanumeric or chinese string.",
		typAlnumhanrod: "Expecting an alphanumeric or chinese or rod string.",
		typAlhan:       "Expecting an alphabetic or chinese string.",
		typAlhanrod:    "Expecting an alphabetic or chinese or rod string.",
		typHan:         "Expecting an chinese string.",
		typHanrod:      "Expecting an chinese or rod string.",
		typMatch:       "Expecting a value to match s%.",
		typNomatch:     "Expecting a value that does not match s%.",
		typEmail:       "Expecting a valid email address.",
		typURL:         "Expecting a valid URL.",
		typIP:          "Expecting a valid ip address.",
		typMobile:      "Expecting a valid mobile number.",
		typTel:         "Expecting a valid telephone number.",
		typPhone:       "Expecting a valid telephone or mobile phone number.",
		typZipcode:     "Expecting a valid zipcode.",
	},
}
