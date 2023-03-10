package entity

import (
	"entgo.io/ent"
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
		field.Int16("status").StorageKey("status").StructTag(`json:"status,omitempty"`).Comment("启用状态"),

		field.String("dict_id").MaxLen(36).NotEmpty().Comment("sys_dict.id"),
	}
}

//func (sdd SysDictItem) Edges() []ent.Edge {
//	return []ent.Edge{
//		//edge.From("SysDict", SysDict.Type).Field("dict_id").Ref("SysDictItems").Unique().Required(),
//	}
//}

func (SysDictItem) Indexes() []ent.Index {
	return []ent.Index{}
}
