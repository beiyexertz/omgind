package entity

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/wanhello/omgind/internal/schema/mixin"
)

type SysDictItem struct {
	ent.Schema
}

func (sdd SysDictItem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{},
		mixin.MemoMixin{},
		mixin.SortMixin{},
		mixin.TimeMixin{},
	}
}

func (sdd SysDictItem) Fields() []ent.Field {
	return []ent.Field{
		field.String("label").StorageKey("label").
			MaxLen(128).StructTag(`json:"label,omitempty"`).Comment("显示值"),
		field.Int("value").StorageKey("val").StructTag(`json:"value,omitempty"`).Comment("字典值"),
		field.Bool("status").StorageKey("status").StructTag(`json:"status,omitempty"`).Comment("启用状态"),
	}
}

func (sdd SysDictItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("SysDict", SysDict.Type).
			Ref("SysDictItems").Unique(),
	}
}
