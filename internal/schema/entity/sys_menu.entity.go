package entity

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/wanhello/omgind/internal/schema/mixin"
)

// SysMenu holds the schema definition for the SysMenu entity.
type SysMenu struct {
	ent.Schema
}

func (sm SysMenu) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.MemoMixin{},
		mixin.SortMixin{},
		mixin.TimeMixin{},
		mixin.StatusMixin{},
	}
}

// Fields of the SysMenu.
func (SysMenu) Fields() []ent.Field {
	return []ent.Field{
		mixin.IdField(),

		field.String("name").StorageKey("name").MaxLen(64).NotEmpty().
			StructTag(`json:"name,omitempty"`).Comment("菜单名称"),
		field.String("icon").StorageKey("icon").MaxLen(256).
			StructTag(`json:"icon,omitempty"`).Comment("菜单图标"),
		field.String("router").StorageKey("router").MaxLen(1024).NotEmpty().
			StructTag(`json:"router,omitempty"`).Comment("前端路由"),
		field.Bool("is_show").Default(true).StorageKey("is_show").StructTag(`json:"is_show,omitempty"`).Comment("是否显示"),

		field.String("parent_id").MaxLen(36).
			Nillable().Optional().StorageKey("pid").
			StructTag(`json:"parent_id,omitempty"`).Comment("父级id"),
		field.String("parent_path").MaxLen(40 * 4).
			Nillable().Optional().StorageKey("ppath").
			StructTag(`json:"parent_path,omitempty"`).Comment("父级路径: 1/2/3"),
	}
}

//// Edges of the SysMenu.
//func (SysMenu) Edges() []ent.Edge {
//	return []ent.Edge{
//		//edge.To("actions", SysMenuAction.Type),
//	}
//}

func (sm SysMenu) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("parent_id"),
		index.Fields("parent_id", "name"),
	}
}
