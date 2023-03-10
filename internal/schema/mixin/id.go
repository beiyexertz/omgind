package mixin

import (
	"fmt"
	"math/rand"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/oklog/ulid/v2"
	"github.com/wanhello/omgind/pkg/helper/pulid"
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

		// Fixme:: make duplicate primary key
		//field.String("id").DefaultFunc(uid.MustString()).MaxLen(36).Immutable().NotEmpty().Comment("主键"),
		IdField(),
		field.Bool("is_del").Default(false).StructTag(`json:"is_del,omitempty"`).Comment("是否删除"),
	}
}

func (i IDMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
		index.Fields("is_del"),
	}
}

///////////////////////

// Helpers for ID generation for entities.
type timeNowFn func() time.Time

type idGenerator struct {
	prefix string
	timeFn timeNowFn
}

func (idg *idGenerator) newULID() pulid.ID {
	timeFn := idg.timeFn
	if idg.timeFn == nil {
		timeFn = time.Now
	}
	t := timeFn()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := pulid.ID{Prefix: idg.prefix}

	id.ULID = ulid.MustNew(ulid.Timestamp(t), entropy)

	fmt.Println(" 0000 ------- 0000 ", id.String())

	return id
}

func IdField1(prefix string) ent.Field {
	idGen := &idGenerator{prefix: prefix}

	return field.UUID("id", pulid.ID{}).Immutable().
		Unique().Default(idGen.newULID)
}

func IdField() ent.Field {
	// return field.UUID("id", uuid.UUID{}).Unique().Default(uuid.New)
	// return field.UUID("id", uuid.UUID{}).Default(uuid.New)

	//return field.String("id").DefaultFunc(uid.MustString()).MaxLen(36).Immutable().NotEmpty().Comment("主键")

	return field.String("id").MaxLen(36).NotEmpty().Immutable().DefaultFunc(func() string {

		seed := time.Now().UnixNano()
		source := rand.NewSource(seed)
		entropy := rand.New(source)

		return ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()
	})

}
