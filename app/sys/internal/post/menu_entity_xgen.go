// generated by xgen -- DO NOT EDIT
package post

import (
	"bytes"

	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
	"gopkg.in/goyy/goyy.v0/util/jsons"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

var (
	MENU_ENTITY           = schema.TABLE("sys_post_menu", "POST MENU")
	MENU_ENTITY_ID        = MENU_ENTITY.PRIMARY("id", "ID")
	MENU_ENTITY_MEMO      = MENU_ENTITY.COLUMN("memo", "MEMO")
	MENU_ENTITY_CREATES   = MENU_ENTITY.COLUMN("creates", "CREATES")
	MENU_ENTITY_CREATER   = MENU_ENTITY.CREATER("creater", "CREATER")
	MENU_ENTITY_CREATED   = MENU_ENTITY.CREATED("created", "CREATED")
	MENU_ENTITY_MODIFIER  = MENU_ENTITY.MODIFIER("modifier", "MODIFIER")
	MENU_ENTITY_MODIFIED  = MENU_ENTITY.MODIFIED("modified", "MODIFIED")
	MENU_ENTITY_VERSION   = MENU_ENTITY.VERSION("version", "VERSION")
	MENU_ENTITY_DELETION  = MENU_ENTITY.DELETION("deletion", "DELETION")
	MENU_ENTITY_ARTIFICAL = MENU_ENTITY.COLUMN("artifical", "ARTIFICAL")
	MENU_ENTITY_HISTORY   = MENU_ENTITY.COLUMN("history", "HISTORY")
	MENU_ENTITY_POST_ID   = MENU_ENTITY.COLUMN("post_id", "POST_ID")
	MENU_ENTITY_MENU_ID   = MENU_ENTITY.COLUMN("menu_id", "MENU_ID")
)

func NewMenuEntity() *MenuEntity {
	e := &MenuEntity{}
	e.init()
	return e
}

func (me *MenuEntity) PostId() string {
	return me.postId.Value()
}

func (me *MenuEntity) SetPostId(v string) {
	me.postId.SetValue(v)
}

func (me *MenuEntity) MenuId() string {
	return me.menuId.Value()
}

func (me *MenuEntity) SetMenuId(v string) {
	me.menuId.SetValue(v)
}

func (me *MenuEntity) init() {
	me.table = MENU_ENTITY
	me.initSetDict()
	me.initSetColumn()
	me.initSetDefault()
	me.initSetField()
	me.initSetExcel()
	me.initSetJson()
	me.initSetXml()
}

func (me *MenuEntity) initSetDict() {
}

func (me *MenuEntity) initSetColumn() {
	if t, ok := me.Sys.Type("id"); ok {
		t.SetColumn(MENU_ENTITY_ID)
	}
	if t, ok := me.Sys.Type("memo"); ok {
		t.SetColumn(MENU_ENTITY_MEMO)
	}
	if t, ok := me.Sys.Type("creates"); ok {
		t.SetColumn(MENU_ENTITY_CREATES)
	}
	if t, ok := me.Sys.Type("creater"); ok {
		t.SetColumn(MENU_ENTITY_CREATER)
	}
	if t, ok := me.Sys.Type("created"); ok {
		t.SetColumn(MENU_ENTITY_CREATED)
	}
	if t, ok := me.Sys.Type("modifier"); ok {
		t.SetColumn(MENU_ENTITY_MODIFIER)
	}
	if t, ok := me.Sys.Type("modified"); ok {
		t.SetColumn(MENU_ENTITY_MODIFIED)
	}
	if t, ok := me.Sys.Type("version"); ok {
		t.SetColumn(MENU_ENTITY_VERSION)
	}
	if t, ok := me.Sys.Type("deletion"); ok {
		t.SetColumn(MENU_ENTITY_DELETION)
	}
	if t, ok := me.Sys.Type("artifical"); ok {
		t.SetColumn(MENU_ENTITY_ARTIFICAL)
	}
	if t, ok := me.Sys.Type("history"); ok {
		t.SetColumn(MENU_ENTITY_HISTORY)
	}
	me.postId.SetColumn(MENU_ENTITY_POST_ID)
	me.menuId.SetColumn(MENU_ENTITY_MENU_ID)
}

func (me *MenuEntity) initSetDefault() {
	if t, ok := me.Sys.Type("created"); ok {
		t.SetDefault("-62135596800")
	}
	if t, ok := me.Sys.Type("modified"); ok {
		t.SetDefault("-62135596800")
	}
}

func (me *MenuEntity) initSetField() {
	for _, c := range entity.SysColumns {
		if t, ok := me.Sys.Type(c); ok {
			t.SetField(entity.DefaultField())
		}
	}
	me.postId.SetField(entity.DefaultField())
	me.menuId.SetField(entity.DefaultField())
}

func (me *MenuEntity) initSetExcel() {
}

func (me *MenuEntity) initSetJson() {
	for _, c := range entity.SysColumns {
		if t, ok := me.Sys.Type(c); ok {
			t.Field().SetJson(entity.NewJsonBy(c))
		}
	}
	me.postId.Field().SetJson(entity.NewJsonBy("postId"))
	me.menuId.Field().SetJson(entity.NewJsonBy("menuId"))
}

func (me *MenuEntity) initSetXml() {
	for _, c := range entity.SysColumns {
		if t, ok := me.Sys.Type(c); ok {
			t.Field().SetXml(entity.NewXmlBy(c))
		}
	}
	me.postId.Field().SetXml(entity.NewXmlBy("postId"))
	me.menuId.Field().SetXml(entity.NewXmlBy("menuId"))
}

func (me MenuEntity) New() entity.Interface {
	return NewMenuEntity()
}

func (me *MenuEntity) Get(column string) interface{} {
	switch column {
	case MENU_ENTITY_POST_ID.Name():
		return me.postId.Value()
	case MENU_ENTITY_MENU_ID.Name():
		return me.menuId.Value()
	}
	return me.Sys.Get(column)
}

func (me *MenuEntity) GetPtr(column string) interface{} {
	switch column {
	case MENU_ENTITY_POST_ID.Name():
		return me.postId.ValuePtr()
	case MENU_ENTITY_MENU_ID.Name():
		return me.menuId.ValuePtr()
	}
	return me.Sys.GetPtr(column)
}

func (me *MenuEntity) GetString(field string) string {
	switch strings.ToLowerFirst(field) {
	case "postId":
		return me.postId.String()
	case "menuId":
		return me.menuId.String()
	}
	return me.Sys.GetString(field)
}

func (me *MenuEntity) SetString(field, value string) error {
	switch strings.ToLowerFirst(field) {
	case "postId":
		return me.postId.SetString(value)
	case "menuId":
		return me.menuId.SetString(value)
	}
	return me.Sys.SetString(field, value)
}

func (me *MenuEntity) Table() schema.Table {
	return me.table
}

func (me *MenuEntity) Type(column string) (entity.Type, bool) {
	switch column {
	case MENU_ENTITY_POST_ID.Name():
		return &me.postId, true
	case MENU_ENTITY_MENU_ID.Name():
		return &me.menuId, true
	}
	return me.Sys.Type(column)
}

func (me *MenuEntity) Column(field string) (schema.Column, bool) {
	switch strings.ToLowerFirst(field) {
	case "postId":
		return MENU_ENTITY_POST_ID, true
	case "menuId":
		return MENU_ENTITY_MENU_ID, true
	}
	return me.Sys.Column(field)
}

func (me *MenuEntity) Columns() []schema.Column {
	return []schema.Column{
		MENU_ENTITY_ID,
		MENU_ENTITY_MEMO,
		MENU_ENTITY_CREATES,
		MENU_ENTITY_CREATER,
		MENU_ENTITY_CREATED,
		MENU_ENTITY_MODIFIER,
		MENU_ENTITY_MODIFIED,
		MENU_ENTITY_VERSION,
		MENU_ENTITY_DELETION,
		MENU_ENTITY_ARTIFICAL,
		MENU_ENTITY_HISTORY,
		MENU_ENTITY_POST_ID,
		MENU_ENTITY_MENU_ID,
	}
}

func (me *MenuEntity) Names() []string {
	return []string{
		"id",
		"memo",
		"creates",
		"creater",
		"created",
		"modifier",
		"modified",
		"version",
		"deletion",
		"artifical",
		"history",
		"postId",
		"menuId",
	}
}

func (me *MenuEntity) Value() *MenuEntity {
	return me
}

func (me *MenuEntity) Validate() error {
	return nil
}

func (me *MenuEntity) JSON() string {
	var b bytes.Buffer
	b.WriteString("{")
	b.WriteString(`"id":"` + jsons.Format(me.GetString("id")) + `"`)
	b.WriteString(`,"memo":"` + jsons.Format(me.GetString("memo")) + `"`)
	b.WriteString(`,"creates":"` + jsons.Format(me.GetString("creates")) + `"`)
	b.WriteString(`,"creater":"` + jsons.Format(me.GetString("creater")) + `"`)
	b.WriteString(`,"created":` + me.GetString("created"))
	b.WriteString(`,"modifier":"` + jsons.Format(me.GetString("modifier")) + `"`)
	b.WriteString(`,"modified":` + me.GetString("modified"))
	b.WriteString(`,"version":` + me.GetString("version"))
	b.WriteString(`,"deletion":` + me.GetString("deletion"))
	b.WriteString(`,"artifical":` + me.GetString("artifical"))
	b.WriteString(`,"history":` + me.GetString("history"))
	b.WriteString(`,"postId":"` + jsons.Format(me.GetString("postId")) + `"`)
	b.WriteString(`,"menuId":"` + jsons.Format(me.GetString("menuId")) + `"`)
	b.WriteString("}")
	return b.String()
}

func (me *MenuEntity) ExcelColumns() []string {
	return nil
}
