// generated by xgen -- DO NOT EDIT
package menu

import (
	"bytes"

	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
	"gopkg.in/goyy/goyy.v0/util/jsons"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

var (
	ENTITY              = schema.TABLE("sys_menu", "MENU")
	ENTITY_ID           = ENTITY.PRIMARY("id", "ID")
	ENTITY_CODE         = ENTITY.COLUMN("code", "CODE")
	ENTITY_NAME         = ENTITY.COLUMN("name", "NAME")
	ENTITY_FULLNAME     = ENTITY.COLUMN("fullname", "FULLNAME")
	ENTITY_GENRE        = ENTITY.COLUMN("genre", "GENRE")
	ENTITY_LEAF         = ENTITY.COLUMN("leaf", "LEAF")
	ENTITY_GRADE        = ENTITY.COLUMN("grade", "GRADE")
	ENTITY_ORDINAL      = ENTITY.COLUMN("ordinal", "ORDINAL")
	ENTITY_PARENT_ID    = ENTITY.COLUMN("parent_id", "PARENT_ID")
	ENTITY_PARENT_IDS   = ENTITY.COLUMN("parent_ids", "PARENT_IDS")
	ENTITY_PARENT_CODES = ENTITY.COLUMN("parent_codes", "PARENT_CODES")
	ENTITY_PARENT_NAMES = ENTITY.COLUMN("parent_names", "PARENT_NAMES")
	ENTITY_MEMO         = ENTITY.COLUMN("memo", "MEMO")
	ENTITY_CREATES      = ENTITY.COLUMN("creates", "CREATES")
	ENTITY_CREATER      = ENTITY.CREATER("creater", "CREATER")
	ENTITY_CREATED      = ENTITY.CREATED("created", "CREATED")
	ENTITY_MODIFIER     = ENTITY.MODIFIER("modifier", "MODIFIER")
	ENTITY_MODIFIED     = ENTITY.MODIFIED("modified", "MODIFIED")
	ENTITY_VERSION      = ENTITY.VERSION("version", "VERSION")
	ENTITY_DELETION     = ENTITY.DELETION("deletion", "DELETION")
	ENTITY_ARTIFICAL    = ENTITY.COLUMN("artifical", "ARTIFICAL")
	ENTITY_HISTORY      = ENTITY.COLUMN("history", "HISTORY")
	ENTITY_HREF         = ENTITY.COLUMN("href", "HREF")
	ENTITY_TARGET       = ENTITY.COLUMN("target", "TARGET")
	ENTITY_ICON         = ENTITY.COLUMN("icon", "ICON")
	ENTITY_HIDDEN       = ENTITY.COLUMN("hidden", "HIDDEN")
	ENTITY_PERMISSION   = ENTITY.COLUMN("permission", "PERMISSION")
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
	me.initSetDict()
	me.initSetColumn()
	me.initSetDefault()
	me.initSetField()
	me.initSetExcel()
	me.initSetJson()
	me.initSetXml()
}

func (me *Entity) initSetDict() {
}

func (me *Entity) initSetColumn() {
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
}

func (me *Entity) initSetDefault() {
	if t, ok := me.Tree.Type("created"); ok {
		t.SetDefault("-62135596800")
	}
	if t, ok := me.Tree.Type("modified"); ok {
		t.SetDefault("-62135596800")
	}
}

func (me *Entity) initSetField() {
	for _, c := range entity.TreeColumns {
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

func (me *Entity) initSetExcel() {
}

func (me *Entity) initSetJson() {
	for _, c := range entity.TreeColumns {
		if t, ok := me.Tree.Type(c); ok {
			t.Field().SetJson(entity.NewJsonBy(c))
		}
	}
	me.href.Field().SetJson(entity.NewJsonBy("href"))
	me.target.Field().SetJson(entity.NewJsonBy("target"))
	me.icon.Field().SetJson(entity.NewJsonBy("icon"))
	me.hidden.Field().SetJson(entity.NewJsonBy("hidden"))
	me.permission.Field().SetJson(entity.NewJsonBy("permission"))
}

func (me *Entity) initSetXml() {
	for _, c := range entity.TreeColumns {
		if t, ok := me.Tree.Type(c); ok {
			t.Field().SetXml(entity.NewXmlBy(c))
		}
	}
	me.href.Field().SetXml(entity.NewXmlBy("href"))
	me.target.Field().SetXml(entity.NewXmlBy("target"))
	me.icon.Field().SetXml(entity.NewXmlBy("icon"))
	me.hidden.Field().SetXml(entity.NewXmlBy("hidden"))
	me.permission.Field().SetXml(entity.NewXmlBy("permission"))
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

func (me *Entity) GetString(field string) string {
	switch strings.ToLowerFirst(field) {
	case "href":
		return me.href.String()
	case "target":
		return me.target.String()
	case "icon":
		return me.icon.String()
	case "hidden":
		return me.hidden.String()
	case "permission":
		return me.permission.String()
	}
	return me.Tree.GetString(field)
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

func (me *Entity) Validate() error {
	return nil
}

func (me *Entity) JSON() string {
	var b bytes.Buffer
	b.WriteString("{")
	b.WriteString(`"id":"` + jsons.Format(me.GetString("id")) + `"`)
	b.WriteString(`,"code":"` + jsons.Format(me.GetString("code")) + `"`)
	b.WriteString(`,"name":"` + jsons.Format(me.GetString("name")) + `"`)
	b.WriteString(`,"fullname":"` + jsons.Format(me.GetString("fullname")) + `"`)
	b.WriteString(`,"genre":"` + jsons.Format(me.GetString("genre")) + `"`)
	b.WriteString(`,"leaf":` + me.GetString("leaf"))
	b.WriteString(`,"grade":` + me.GetString("grade"))
	b.WriteString(`,"ordinal":"` + jsons.Format(me.GetString("ordinal")) + `"`)
	b.WriteString(`,"parent_id":"` + jsons.Format(me.GetString("parent_id")) + `"`)
	b.WriteString(`,"parent_ids":"` + jsons.Format(me.GetString("parent_ids")) + `"`)
	b.WriteString(`,"parent_codes":"` + jsons.Format(me.GetString("parent_codes")) + `"`)
	b.WriteString(`,"parent_names":"` + jsons.Format(me.GetString("parent_names")) + `"`)
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
	b.WriteString(`,"href":"` + jsons.Format(me.GetString("href")) + `"`)
	b.WriteString(`,"target":"` + jsons.Format(me.GetString("target")) + `"`)
	b.WriteString(`,"icon":"` + jsons.Format(me.GetString("icon")) + `"`)
	b.WriteString(`,"hidden":` + me.GetString("hidden"))
	b.WriteString(`,"permission":"` + jsons.Format(me.GetString("permission")) + `"`)
	b.WriteString("}")
	return b.String()
}

func (me *Entity) ExcelColumns() []string {
	return nil
}