package entity

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/wanhello/omgind/internal/schema/mixin"
)

// SysRole holds the schema definition for the SysRole entity.
type SysRole struct {
	ent.Schema
}

func (sr SysRole) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.StatusMixin{},
		mixin.SortMixin{},
		mixin.MemoMixin{},
		mixin.TimeMixin{},
	}
}

// Fields of the SysRole.
func (SysRole) Fields() []ent.Field {
	return []ent.Field{
		mixin.IdField(),

		field.String("name").MaxLen(64).MinLen(2).NotEmpty().Comment("角色名称")}
}

//// Edges of the SysRole.
//func (SysRole) Edges() []ent.Edge {
//	return []ent.Edge{
//		//edge.To("userRoles", SysUserRole.Type).Comment("userroles"),
//	}
//}

func (SysRole) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
	}
}
