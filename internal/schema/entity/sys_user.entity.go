package entity

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/wanhello/omgind/internal/schema/mixin"
)

// SysUser holds the schema definition for the SysUser entity.
type SysUser struct {
	ent.Schema
}

func (su SysUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{},
		mixin.SortMixin{},
		mixin.TimeMixin{},
	}
}

// Fields of the SysUser.
func (SysUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("UserName").StorageKey("user_name").
			MinLen(5).MaxLen(128).NotEmpty().StructTag(`json:"user_name,omitempty"`).Comment("用户名"),
		field.String("RealName").StorageKey("real_name").MaxLen(64),
		field.String("FirstName").StorageKey("first_name").
			MaxLen(31).Nillable().Optional().StructTag(`json:"first_name,omitempty"`).Comment("名"),
		field.String("LastName").StorageKey("last_name").
			MaxLen(31).Nillable().Optional().StructTag(`json:"last_name,omitempty"`).Comment("姓"),
		field.String("Password").StorageKey("passwd").
			MaxLen(256).Sensitive().Comment("密码"),
		field.String("Email").StorageKey("email").
			MaxLen(256).StructTag(`json:"email,omitempty"`).Comment("电子邮箱"),
		field.String("Phone").StorageKey("phone").
			MaxLen(20).StructTag(`json:"phone,omitempty"`).Comment("电话号码"),
	}
}

func (su SysUser) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("UserName").Unique(),
	}
}

// Edges of the SysUser.
func (SysUser) Edges() []ent.Edge {
	return nil
}
