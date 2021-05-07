package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

//////////////////////////////////

type StatusMixin struct {
	mixin.Schema
}

func (s StatusMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("status").
			Default(0).
			StructTag(`json:"sort,omitempty"`).
			Comment("状态, "),
	}
}

func (s StatusMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("status"),
	}
}
