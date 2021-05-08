// Code generated by entc, DO NOT EDIT.

package sysmenuaction

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/wanhello/omgind/internal/gen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// IsDel applies equality check predicate on the "is_del" field. It's identical to IsDelEQ.
func IsDel(v bool) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsDel), v))
	})
}

// Sort applies equality check predicate on the "sort" field. It's identical to SortEQ.
func Sort(v int32) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSort), v))
	})
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int32) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// Memo applies equality check predicate on the "memo" field. It's identical to MemoEQ.
func Memo(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMemo), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// MenuID applies equality check predicate on the "menu_id" field. It's identical to MenuIDEQ.
func MenuID(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMenuID), v))
	})
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCode), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// IsDelEQ applies the EQ predicate on the "is_del" field.
func IsDelEQ(v bool) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsDel), v))
	})
}

// IsDelNEQ applies the NEQ predicate on the "is_del" field.
func IsDelNEQ(v bool) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsDel), v))
	})
}

// SortEQ applies the EQ predicate on the "sort" field.
func SortEQ(v int32) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSort), v))
	})
}

// SortNEQ applies the NEQ predicate on the "sort" field.
func SortNEQ(v int32) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSort), v))
	})
}

// SortIn applies the In predicate on the "sort" field.
func SortIn(vs ...int32) predicate.SysMenuAction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenuAction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSort), v...))
	})
}

// SortNotIn applies the NotIn predicate on the "sort" field.
func SortNotIn(vs ...int32) predicate.SysMenuAction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenuAction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSort), v...))
	})
}

// SortGT applies the GT predicate on the "sort" field.
func SortGT(v int32) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSort), v))
	})
}

// SortGTE applies the GTE predicate on the "sort" field.
func SortGTE(v int32) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSort), v))
	})
}

// SortLT applies the LT predicate on the "sort" field.
func SortLT(v int32) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSort), v))
	})
}

// SortLTE applies the LTE predicate on the "sort" field.
func SortLTE(v int32) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSort), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int32) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int32) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int32) predicate.SysMenuAction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenuAction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int32) predicate.SysMenuAction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenuAction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int32) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatus), v))
	})
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int32) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatus), v))
	})
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int32) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatus), v))
	})
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int32) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatus), v))
	})
}

// MemoEQ applies the EQ predicate on the "memo" field.
func MemoEQ(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMemo), v))
	})
}

// MemoNEQ applies the NEQ predicate on the "memo" field.
func MemoNEQ(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMemo), v))
	})
}

// MemoIn applies the In predicate on the "memo" field.
func MemoIn(vs ...string) predicate.SysMenuAction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenuAction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMemo), v...))
	})
}

// MemoNotIn applies the NotIn predicate on the "memo" field.
func MemoNotIn(vs ...string) predicate.SysMenuAction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenuAction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMemo), v...))
	})
}

// MemoGT applies the GT predicate on the "memo" field.
func MemoGT(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMemo), v))
	})
}

// MemoGTE applies the GTE predicate on the "memo" field.
func MemoGTE(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMemo), v))
	})
}

// MemoLT applies the LT predicate on the "memo" field.
func MemoLT(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMemo), v))
	})
}

// MemoLTE applies the LTE predicate on the "memo" field.
func MemoLTE(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMemo), v))
	})
}

// MemoContains applies the Contains predicate on the "memo" field.
func MemoContains(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMemo), v))
	})
}

// MemoHasPrefix applies the HasPrefix predicate on the "memo" field.
func MemoHasPrefix(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMemo), v))
	})
}

// MemoHasSuffix applies the HasSuffix predicate on the "memo" field.
func MemoHasSuffix(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMemo), v))
	})
}

// MemoEqualFold applies the EqualFold predicate on the "memo" field.
func MemoEqualFold(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMemo), v))
	})
}

// MemoContainsFold applies the ContainsFold predicate on the "memo" field.
func MemoContainsFold(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMemo), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.SysMenuAction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenuAction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.SysMenuAction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenuAction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.SysMenuAction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenuAction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.SysMenuAction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenuAction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.SysMenuAction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenuAction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.SysMenuAction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenuAction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// MenuIDEQ applies the EQ predicate on the "menu_id" field.
func MenuIDEQ(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMenuID), v))
	})
}

// MenuIDNEQ applies the NEQ predicate on the "menu_id" field.
func MenuIDNEQ(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMenuID), v))
	})
}

// MenuIDIn applies the In predicate on the "menu_id" field.
func MenuIDIn(vs ...string) predicate.SysMenuAction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenuAction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMenuID), v...))
	})
}

// MenuIDNotIn applies the NotIn predicate on the "menu_id" field.
func MenuIDNotIn(vs ...string) predicate.SysMenuAction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenuAction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMenuID), v...))
	})
}

// MenuIDGT applies the GT predicate on the "menu_id" field.
func MenuIDGT(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMenuID), v))
	})
}

// MenuIDGTE applies the GTE predicate on the "menu_id" field.
func MenuIDGTE(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMenuID), v))
	})
}

// MenuIDLT applies the LT predicate on the "menu_id" field.
func MenuIDLT(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMenuID), v))
	})
}

// MenuIDLTE applies the LTE predicate on the "menu_id" field.
func MenuIDLTE(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMenuID), v))
	})
}

// MenuIDContains applies the Contains predicate on the "menu_id" field.
func MenuIDContains(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMenuID), v))
	})
}

// MenuIDHasPrefix applies the HasPrefix predicate on the "menu_id" field.
func MenuIDHasPrefix(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMenuID), v))
	})
}

// MenuIDHasSuffix applies the HasSuffix predicate on the "menu_id" field.
func MenuIDHasSuffix(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMenuID), v))
	})
}

// MenuIDEqualFold applies the EqualFold predicate on the "menu_id" field.
func MenuIDEqualFold(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMenuID), v))
	})
}

// MenuIDContainsFold applies the ContainsFold predicate on the "menu_id" field.
func MenuIDContainsFold(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMenuID), v))
	})
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCode), v))
	})
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCode), v))
	})
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.SysMenuAction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenuAction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCode), v...))
	})
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.SysMenuAction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenuAction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCode), v...))
	})
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCode), v))
	})
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCode), v))
	})
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCode), v))
	})
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCode), v))
	})
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCode), v))
	})
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCode), v))
	})
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCode), v))
	})
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCode), v))
	})
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCode), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.SysMenuAction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenuAction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.SysMenuAction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysMenuAction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// HasResources applies the HasEdge predicate on the "resources" edge.
func HasResources() predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ResourcesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ResourcesTable, ResourcesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasResourcesWith applies the HasEdge predicate on the "resources" edge with a given conditions (other predicates).
func HasResourcesWith(preds ...predicate.SysMenuActionResource) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ResourcesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ResourcesTable, ResourcesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMenu applies the HasEdge predicate on the "menu" edge.
func HasMenu() predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MenuTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MenuTable, MenuColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMenuWith applies the HasEdge predicate on the "menu" edge with a given conditions (other predicates).
func HasMenuWith(preds ...predicate.SysMenu) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MenuInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MenuTable, MenuColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SysMenuAction) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SysMenuAction) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SysMenuAction) predicate.SysMenuAction {
	return predicate.SysMenuAction(func(s *sql.Selector) {
		p(s.Not())
	})
}
