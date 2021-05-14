package entity

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/wanhello/omgind/internal/schema/mixin"
)

type XxxDemo struct {
	ent.Schema
}

func (xd XxxDemo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.MemoMixin{},
		mixin.SortMixin{},
		mixin.TimeMixin{},
	}
}

/*
	Code   string  `gorm:"column:code;size:50;index;default:'';not null;"`  // 编号
	Name   string  `gorm:"column:name;size:100;index;default:'';not null;"` // 名称
	Memo   *string `gorm:"column:memo;size:200;"`                           // 备注
	Status int     `gorm:"column:status;index;default:0;not null;"`         // 状态(1:启用 2:停用)
	IsDel  bool    `gorm:"column:is_del;default:false;index;not null;"`

*/

func (xd XxxDemo) Fields() []ent.Field {

	return []ent.Field{
		mixin.IdField("01"),

		field.String("code").StorageKey("code").
			MaxLen(128).StructTag(`json:"code,omitempty"`).Comment("编号"),
		field.String("name").StorageKey("name").
			MaxLen(128).StructTag(`json:"name,omitempty"`).Comment("名称"),

		field.Int("status").StorageKey("status").StructTag(`json:"status,omitempty"`).Default(1).Comment("状态"),
	}
}

func (xd XxxDemo) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.To("SysDictItems", SysDictItem.Type),
	}
}
