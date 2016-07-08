// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"gopkg.in/goyy/goyy.v0/comm/i18n"
)

var i18N = i18n.NewByEnv(locales)

var locales = map[string]map[string]string{
	i18n.Locale_zh_CN: map[string]string{
		"col.comment.id":           "标识",
		"col.comment.memo":         "备注",
		"col.comment.creates":      "创建机构",
		"col.comment.creater":      "创建人员",
		"col.comment.created":      "创建时间",
		"col.comment.modifier":     "更新人员",
		"col.comment.modified":     "更新时间",
		"col.comment.version":      "乐观锁",
		"col.comment.deletion":     "删除标志",
		"col.comment.artifical":    "人造数据",
		"col.comment.history":      "历史数据",
		"col.comment.code":         "编号",
		"col.comment.name":         "名称",
		"col.comment.fullname":     "全称",
		"col.comment.genre":        "类型",
		"col.comment.leaf":         "是否叶子",
		"col.comment.grade":        "节点等级",
		"col.comment.ordinal":      "排序",
		"col.comment.parent_id":    "父表主键",
		"col.comment.parent_ids":   "父ID集",
		"col.comment.parent_codes": "父编号集",
		"col.comment.parent_names": "父名称集",
	},
	i18n.Locale_en_US: map[string]string{
		"col.comment.id":           "ID",
		"col.comment.memo":         "MEMO",
		"col.comment.creates":      "CREATES",
		"col.comment.creater":      "CREATER",
		"col.comment.created":      "CREATED",
		"col.comment.modifier":     "MODIFIER",
		"col.comment.modified":     "MODIFIED",
		"col.comment.version":      "VERSION",
		"col.comment.deletion":     "DELETION",
		"col.comment.artifical":    "ARTIFICAL",
		"col.comment.history":      "HISTORY",
		"col.comment.code":         "CODE",
		"col.comment.name":         "NAME",
		"col.comment.fullname":     "FULLNAME",
		"col.comment.genre":        "GENRE",
		"col.comment.leaf":         "LEAF",
		"col.comment.grade":        "GRADE",
		"col.comment.ordinal":      "ORDINAL",
		"col.comment.parent_id":    "PARENT_ID",
		"col.comment.parent_ids":   "PARENT_IDS",
		"col.comment.parent_codes": "PARENT_CODES",
		"col.comment.parent_names": "PARENT_NAMES",
	},
}
