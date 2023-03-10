// Code generated by entc, DO NOT EDIT.

package sysuserrole

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/wanhello/omgind/internal/gen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
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
func IDNotIn(ids ...string) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
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
func IDGT(id string) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// IsDel applies equality check predicate on the "is_del" field. It's identical to IsDelEQ.
func IsDel(v bool) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsDel), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// RoleID applies equality check predicate on the "role_id" field. It's identical to RoleIDEQ.
func RoleID(v string) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRoleID), v))
	})
}

// IsDelEQ applies the EQ predicate on the "is_del" field.
func IsDelEQ(v bool) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsDel), v))
	})
}

// IsDelNEQ applies the NEQ predicate on the "is_del" field.
func IsDelNEQ(v bool) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsDel), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.SysUserRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysUserRole(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.SysUserRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysUserRole(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.SysUserRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysUserRole(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.SysUserRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysUserRole(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.SysUserRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysUserRole(func(s *sql.Selector) {
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
func DeletedAtNotIn(vs ...time.Time) predicate.SysUserRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysUserRole(func(s *sql.Selector) {
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
func DeletedAtGT(v time.Time) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.SysUserRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysUserRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.SysUserRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysUserRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	})
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	})
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	})
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	})
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUserID), v))
	})
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUserID), v))
	})
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUserID), v))
	})
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUserID), v))
	})
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUserID), v))
	})
}

// RoleIDEQ applies the EQ predicate on the "role_id" field.
func RoleIDEQ(v string) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRoleID), v))
	})
}

// RoleIDNEQ applies the NEQ predicate on the "role_id" field.
func RoleIDNEQ(v string) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRoleID), v))
	})
}

// RoleIDIn applies the In predicate on the "role_id" field.
func RoleIDIn(vs ...string) predicate.SysUserRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysUserRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRoleID), v...))
	})
}

// RoleIDNotIn applies the NotIn predicate on the "role_id" field.
func RoleIDNotIn(vs ...string) predicate.SysUserRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SysUserRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRoleID), v...))
	})
}

// RoleIDGT applies the GT predicate on the "role_id" field.
func RoleIDGT(v string) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRoleID), v))
	})
}

// RoleIDGTE applies the GTE predicate on the "role_id" field.
func RoleIDGTE(v string) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRoleID), v))
	})
}

// RoleIDLT applies the LT predicate on the "role_id" field.
func RoleIDLT(v string) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRoleID), v))
	})
}

// RoleIDLTE applies the LTE predicate on the "role_id" field.
func RoleIDLTE(v string) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRoleID), v))
	})
}

// RoleIDContains applies the Contains predicate on the "role_id" field.
func RoleIDContains(v string) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRoleID), v))
	})
}

// RoleIDHasPrefix applies the HasPrefix predicate on the "role_id" field.
func RoleIDHasPrefix(v string) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRoleID), v))
	})
}

// RoleIDHasSuffix applies the HasSuffix predicate on the "role_id" field.
func RoleIDHasSuffix(v string) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRoleID), v))
	})
}

// RoleIDEqualFold applies the EqualFold predicate on the "role_id" field.
func RoleIDEqualFold(v string) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRoleID), v))
	})
}

// RoleIDContainsFold applies the ContainsFold predicate on the "role_id" field.
func RoleIDContainsFold(v string) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRoleID), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SysUserRole) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SysUserRole) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
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
func Not(p predicate.SysUserRole) predicate.SysUserRole {
	return predicate.SysUserRole(func(s *sql.Selector) {
		p(s.Not())
	})
}
