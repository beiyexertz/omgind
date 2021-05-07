package mixin

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

type TimeMixin struct {
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			StorageKey("crtd_at").
			//StructTag(`json:"created_at,omitempty" sql:"crtd_at"`).
			StructTag(`json:"created_at,omitempty"`).
			Immutable().
			Default(time.Now).
			Comment("创建时间,由程序自动生成"),

		field.Time("updated_at").
			StorageKey("uptd_at").
			//StructTag(`json:"updated_at,omitempty" sql:"uptd_at"`).
			StructTag(`json:"updated_at,omitempty"`).
			Default(time.Now).
			UpdateDefault(time.Now).
			//Immutable().
			Comment("更新时间,由程序自动生成"),

		field.Time("deleted_at").
			StorageKey("dltd_at").Nillable().Optional().
			StructTag(`json:"deleted_at,omitempty"`).
			Comment("删除时间,"),
	}
}

func (m TimeMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at"),
		index.Fields("updated_at"),
		index.Fields("deleted_at"),
	}
}
