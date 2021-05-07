package entity

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/wanhello/omgind/internal/schema/mixin"
)

// SysMenu holds the schema definition for the SysMenu entity.
type SysMenu struct {
	ent.Schema
}

func (sm SysMenu) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{},
		mixin.MemoMixin{},
		mixin.SortMixin{},
		mixin.TimeMixin{},
	}
}

// Fields of the SysMenu.
func (SysMenu) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").StorageKey("name").
			MaxLen(64).StructTag(`json:"name,omitempty"`).Comment("菜单名称"),
		field.String("icon").StorageKey("icon").
			MaxLen(256).StructTag(`json:"icon,omitempty"`).Comment("菜单图标"),
	}
}

// Edges of the SysMenu.
func (SysMenu) Edges() []ent.Edge {
	return nil
}
