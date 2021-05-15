package entity

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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
		mixin.IdField(),

		field.Text("jwt").StorageKey("jwt").NotEmpty().Comment("jwt"),
	}
}

func (SysJwtBlock) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
	}
}
