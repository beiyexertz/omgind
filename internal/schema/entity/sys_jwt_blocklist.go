package entity

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/wanhello/omgind/internal/schema/mixin"
)

type SysJwtBlock struct {
	ent.Schema
}

func (sjb SysJwtBlock) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.MemoMixin{},
		mixin.TimeMixin{},
		mixin.StatusMixin{},
	}
}

func (sjb SysJwtBlock) Fields() []ent.Field {
	return []ent.Field{
		mixin.IdField("01"),

		field.Text("jwt").StorageKey("jwt").NotEmpty().Comment("jwt"),
	}
}
