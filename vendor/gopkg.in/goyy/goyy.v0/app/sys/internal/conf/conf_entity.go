package conf

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
)

//go:generate xgen -entity=$GOFILE -scaffold -clidir=../../../bms -clipath=gopkg.in/goyy/goyy.v0/app/bms -apipath=gopkg.in/goyy/goyy.v0/app/sys

// CONF Entity.
// @entity(project:"sys")
type Entity struct {
	entity.Sys
	table   schema.Table  `db:"table=sys_conf&comment=CONF"`
	name    entity.String `db:"column=name"`
	code    entity.String `db:"column=code"`
	content entity.String `db:"column=content"`
	genre   entity.String `db:"column=genre"`
	usable  entity.String `db:"column=usable"`
	ordinal entity.String `db:"column=ordinal"`
}
