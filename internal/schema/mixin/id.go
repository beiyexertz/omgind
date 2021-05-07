package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"

	uid "github.com/wanhello/omgind/pkg/helper/uid/ulid"
)

/////////////////////////
// 使用ulid 作为主键
/////////////////////////
type IDMixin struct {
	mixin.Schema
}

func (i IDMixin) Fields() []ent.Field {
	return []ent.Field{
		//field.String("id").DefaultFunc(func() string {
		//	now := time.Now()
		//	seed := now.UnixNano()
		//	source := rand.NewSource(seed)
		//	entropy := rand.New(source)
		//	uid := ulid.MustNew(ulid.Timestamp(now), entropy).String()
		//	return uid
		//}).MaxLen(36).Unique().Immutable().NotEmpty().Comment("主键"),

		field.String("id").DefaultFunc(uid.MustString()).MaxLen(36).Immutable().NotEmpty().Comment("主键"),
		field.Bool("is_del").Default(false).StructTag(`json:"is_del,omitempty"`).Comment("是否删除"),
	}
}

func (i IDMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
		index.Fields("is_del"),
	}
}
