// generated by xgen -- DO NOT EDIT
package menu

import (
	"bytes"
	"fmt"
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

var (
	ENTITY              = schema.TABLE("sys_menu")
	ENTITY_ID           = ENTITY.PRIMARY("id")
	ENTITY_CODE         = ENTITY.COLUMN("code")
	ENTITY_NAME         = ENTITY.COLUMN("name")
	ENTITY_FULLNAME     = ENTITY.COLUMN("fullname")
	ENTITY_GENRE        = ENTITY.COLUMN("genre")
	ENTITY_LEAF         = ENTITY.COLUMN("leaf")
	ENTITY_GRADE        = ENTITY.COLUMN("grade")
	ENTITY_ORDINAL      = ENTITY.COLUMN("ordinal")
	ENTITY_PARENT_ID    = ENTITY.COLUMN("parent_id")
	ENTITY_PARENT_IDS   = ENTITY.COLUMN("parent_ids")
	ENTITY_PARENT_CODES = ENTITY.COLUMN("parent_codes")
	ENTITY_PARENT_NAMES = ENTITY.COLUMN("parent_names")
	ENTITY_MEMO         = ENTITY.COLUMN("memo")
	ENTITY_CREATES      = ENTITY.COLUMN("creates")
	ENTITY_CREATER      = ENTITY.CREATER("creater")
	ENTITY_CREATED      = ENTITY.CREATED("created")
	ENTITY_MODIFIER     = ENTITY.MODIFIER("modifier")
	ENTITY_MODIFIED     = ENTITY.MODIFIED("modified")
	ENTITY_VERSION      = ENTITY.VERSION("version")
	ENTITY_DELETION     = ENTITY.DELETION("deletion")
	ENTITY_ARTIFICAL    = ENTITY.COLUMN("artifical")
	ENTITY_HISTORY      = ENTITY.COLUMN("history")
	ENTITY_HREF         = ENTITY.COLUMN("href")
	ENTITY_TARGET       = ENTITY.COLUMN("target")
	ENTITY_ICON         = ENTITY.COLUMN("icon")
	ENTITY_HIDDEN       = ENTITY.COLUMN("hidden")
	ENTITY_PERMISSION   = ENTITY.COLUMN("permission")
)

func NewEntity() *Entity {
	e := &Entity{}
	e.init()
	return e
}

func (me *Entity) Href() string {
	return me.href.Value()
}

func (me *Entity) SetHref(v string) {
	me.href.SetValue(v)
}

func (me *Entity) Target() string {
	return me.target.Value()
}

func (me *Entity) SetTarget(v string) {
	me.target.SetValue(v)
}

func (me *Entity) Icon() string {
	return me.icon.Value()
}

func (me *Entity) SetIcon(v string) {
	me.icon.SetValue(v)
}

func (me *Entity) Hidden() bool {
	return me.hidden.Value()
}

func (me *Entity) SetHidden(v bool) {
	me.hidden.SetValue(v)
}

func (me *Entity) Permission() string {
	return me.permission.Value()
}

func (me *Entity) SetPermission(v string) {
	me.permission.SetValue(v)
}

func (me *Entity) init() {
	me.table = ENTITY

	if t, ok := me.Tree.Type("id"); ok {
		t.SetColumn(ENTITY_ID)
	}
	if t, ok := me.Tree.Type("code"); ok {
		t.SetColumn(ENTITY_CODE)
	}
	if t, ok := me.Tree.Type("name"); ok {
		t.SetColumn(ENTITY_NAME)
	}
	if t, ok := me.Tree.Type("fullname"); ok {
		t.SetColumn(ENTITY_FULLNAME)
	}
	if t, ok := me.Tree.Type("genre"); ok {
		t.SetColumn(ENTITY_GENRE)
	}
	if t, ok := me.Tree.Type("leaf"); ok {
		t.SetColumn(ENTITY_LEAF)
	}
	if t, ok := me.Tree.Type("grade"); ok {
		t.SetColumn(ENTITY_GRADE)
	}
	if t, ok := me.Tree.Type("ordinal"); ok {
		t.SetColumn(ENTITY_ORDINAL)
	}
	if t, ok := me.Tree.Type("parent_id"); ok {
		t.SetColumn(ENTITY_PARENT_ID)
	}
	if t, ok := me.Tree.Type("parent_ids"); ok {
		t.SetColumn(ENTITY_PARENT_IDS)
	}
	if t, ok := me.Tree.Type("parent_codes"); ok {
		t.SetColumn(ENTITY_PARENT_CODES)
	}
	if t, ok := me.Tree.Type("parent_names"); ok {
		t.SetColumn(ENTITY_PARENT_NAMES)
	}
	if t, ok := me.Tree.Type("memo"); ok {
		t.SetColumn(ENTITY_MEMO)
	}
	if t, ok := me.Tree.Type("creates"); ok {
		t.SetColumn(ENTITY_CREATES)
	}
	if t, ok := me.Tree.Type("creater"); ok {
		t.SetColumn(ENTITY_CREATER)
	}
	if t, ok := me.Tree.Type("created"); ok {
		t.SetColumn(ENTITY_CREATED)
	}
	if t, ok := me.Tree.Type("modifier"); ok {
		t.SetColumn(ENTITY_MODIFIER)
	}
	if t, ok := me.Tree.Type("modified"); ok {
		t.SetColumn(ENTITY_MODIFIED)
	}
	if t, ok := me.Tree.Type("version"); ok {
		t.SetColumn(ENTITY_VERSION)
	}
	if t, ok := me.Tree.Type("deletion"); ok {
		t.SetColumn(ENTITY_DELETION)
	}
	if t, ok := me.Tree.Type("artifical"); ok {
		t.SetColumn(ENTITY_ARTIFICAL)
	}
	if t, ok := me.Tree.Type("history"); ok {
		t.SetColumn(ENTITY_HISTORY)
	}
	me.href.SetColumn(ENTITY_HREF)
	me.target.SetColumn(ENTITY_TARGET)
	me.icon.SetColumn(ENTITY_ICON)
	me.hidden.SetColumn(ENTITY_HIDDEN)
	me.permission.SetColumn(ENTITY_PERMISSION)

	if t, ok := me.Tree.Type("created"); ok {
		t.SetDefault("-62135596800")
	}
	if t, ok := me.Tree.Type("modified"); ok {
		t.SetDefault("-62135596800")
	}

	columns := []string{"id", "code", "name", "fullname", "genre", "leaf", "grade",
		"ordinal", "parent_id", "parent_ids", "parent_codes", "parent_names",
		"memo", "creates", "creater", "created", "modifier", "modified",
		"version", "deletion", "artifical", "history"}
	for _, c := range columns {
		if t, ok := me.Tree.Type(c); ok {
			t.SetField(entity.DefaultField())
		}
	}
	me.href.SetField(entity.DefaultField())
	me.target.SetField(entity.DefaultField())
	me.icon.SetField(entity.DefaultField())
	me.hidden.SetField(entity.DefaultField())
	me.permission.SetField(entity.DefaultField())
}

func (me Entity) New() entity.Interface {
	return NewEntity()
}

func (me *Entity) Get(column string) interface{} {
	switch column {
	case ENTITY_HREF.Name():
		return me.href.Value()
	case ENTITY_TARGET.Name():
		return me.target.Value()
	case ENTITY_ICON.Name():
		return me.icon.Value()
	case ENTITY_HIDDEN.Name():
		return me.hidden.Value()
	case ENTITY_PERMISSION.Name():
		return me.permission.Value()
	}
	return me.Tree.Get(column)
}

func (me *Entity) GetPtr(column string) interface{} {
	switch column {
	case ENTITY_HREF.Name():
		return me.href.ValuePtr()
	case ENTITY_TARGET.Name():
		return me.target.ValuePtr()
	case ENTITY_ICON.Name():
		return me.icon.ValuePtr()
	case ENTITY_HIDDEN.Name():
		return me.hidden.ValuePtr()
	case ENTITY_PERMISSION.Name():
		return me.permission.ValuePtr()
	}
	return me.Tree.GetPtr(column)
}

func (me *Entity) Table() schema.Table {
	return me.table
}

func (me *Entity) Type(column string) (entity.Type, bool) {
	switch column {
	case ENTITY_HREF.Name():
		return &me.href, true
	case ENTITY_TARGET.Name():
		return &me.target, true
	case ENTITY_ICON.Name():
		return &me.icon, true
	case ENTITY_HIDDEN.Name():
		return &me.hidden, true
	case ENTITY_PERMISSION.Name():
		return &me.permission, true
	}
	return me.Tree.Type(column)
}

func (me *Entity) Column(field string) (schema.Column, bool) {
	switch strings.ToLowerFirst(field) {
	case "href":
		return ENTITY_HREF, true
	case "target":
		return ENTITY_TARGET, true
	case "icon":
		return ENTITY_ICON, true
	case "hidden":
		return ENTITY_HIDDEN, true
	case "permission":
		return ENTITY_PERMISSION, true
	}
	return me.Tree.Column(field)
}

func (me *Entity) Columns() []schema.Column {
	return []schema.Column{
		ENTITY_ID,
		ENTITY_CODE,
		ENTITY_NAME,
		ENTITY_FULLNAME,
		ENTITY_GENRE,
		ENTITY_LEAF,
		ENTITY_GRADE,
		ENTITY_ORDINAL,
		ENTITY_PARENT_ID,
		ENTITY_PARENT_IDS,
		ENTITY_PARENT_CODES,
		ENTITY_PARENT_NAMES,
		ENTITY_MEMO,
		ENTITY_CREATES,
		ENTITY_CREATER,
		ENTITY_CREATED,
		ENTITY_MODIFIER,
		ENTITY_MODIFIED,
		ENTITY_VERSION,
		ENTITY_DELETION,
		ENTITY_ARTIFICAL,
		ENTITY_HISTORY,
		ENTITY_HREF,
		ENTITY_TARGET,
		ENTITY_ICON,
		ENTITY_HIDDEN,
		ENTITY_PERMISSION,
	}
}

func (me *Entity) Names() []string {
	return []string{
		"id",
		"code",
		"name",
		"fullname",
		"genre",
		"leaf",
		"grade",
		"ordinal",
		"parentId",
		"parentIds",
		"parentCodes",
		"parentNames",
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
		"href",
		"target",
		"icon",
		"hidden",
		"permission",
	}
}

func (me *Entity) Value() *Entity {
	return me
}

func (me *Entity) SetString(field, value string) error {
	switch strings.ToLowerFirst(field) {
	case "href":
		return me.href.SetString(value)
	case "target":
		return me.target.SetString(value)
	case "icon":
		return me.icon.SetString(value)
	case "hidden":
		return me.hidden.SetString(value)
	case "permission":
		return me.permission.SetString(value)
	}
	return me.Tree.SetString(field, value)
}

func (me *Entity) Validate() error {
	return nil
}

func (me *Entity) JSON() string {
	var b bytes.Buffer
	b.WriteString("{")
	b.WriteString(fmt.Sprintf(`"id":%q`, me.Tree.Sys.Pk.Id()))
	b.WriteString(fmt.Sprintf(`,"memo":%q`, me.Tree.Sys.Memo()))
	b.WriteString(fmt.Sprintf(`,"creates":%q`, me.Tree.Sys.Creates()))
	b.WriteString(fmt.Sprintf(`,"creater":%q`, me.Tree.Sys.Creater()))
	b.WriteString(fmt.Sprintf(`,"created":%d`, me.Tree.Sys.Created()))
	b.WriteString(fmt.Sprintf(`,"modifier":%q`, me.Tree.Sys.Modifier()))
	b.WriteString(fmt.Sprintf(`,"modified":%d`, me.Tree.Sys.Modified()))
	b.WriteString(fmt.Sprintf(`,"version":%d`, me.Tree.Sys.Version()))
	b.WriteString(fmt.Sprintf(`,"deletion":%d`, me.Tree.Sys.Deletion()))
	b.WriteString(fmt.Sprintf(`,"artifical":%d`, me.Tree.Sys.Artifical()))
	b.WriteString(fmt.Sprintf(`,"history":%d`, me.Tree.Sys.History()))
	b.WriteString(fmt.Sprintf(`,"code":%q`, me.Tree.Code()))
	b.WriteString(fmt.Sprintf(`,"name":%q`, me.Tree.Name()))
	b.WriteString(fmt.Sprintf(`,"fullname":%q`, me.Tree.Fullname()))
	b.WriteString(fmt.Sprintf(`,"genre":%q`, me.Tree.Genre()))
	b.WriteString(fmt.Sprintf(`,"leaf":%d`, me.Tree.Leaf()))
	b.WriteString(fmt.Sprintf(`,"grade":%d`, me.Tree.Grade()))
	b.WriteString(fmt.Sprintf(`,"ordinal":%q`, me.Tree.Ordinal()))
	b.WriteString(fmt.Sprintf(`,"parentId":%q`, me.Tree.ParentId()))
	b.WriteString(fmt.Sprintf(`,"parentIds":%q`, me.Tree.ParentIds()))
	b.WriteString(fmt.Sprintf(`,"parentCodes":%q`, me.Tree.ParentCodes()))
	b.WriteString(fmt.Sprintf(`,"parentNames":%q`, me.Tree.ParentNames()))
	b.WriteString(fmt.Sprintf(`,"href":%q`, me.href.String()))
	b.WriteString(fmt.Sprintf(`,"target":%q`, me.target.String()))
	b.WriteString(fmt.Sprintf(`,"icon":%q`, me.icon.String()))
	b.WriteString(fmt.Sprintf(`,"hidden":%q`, me.hidden.String()))
	b.WriteString(fmt.Sprintf(`,"permission":%q`, me.permission.String()))
	b.WriteString("}")
	return b.String()
}