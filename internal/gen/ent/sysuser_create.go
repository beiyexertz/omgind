// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wanhello/omgind/internal/gen/ent/sysuser"
)

// SysUserCreate is the builder for creating a SysUser entity.
type SysUserCreate struct {
	config
	mutation *SysUserMutation
	hooks    []Hook
}

// SetSort sets the "sort" field.
func (suc *SysUserCreate) SetSort(i int32) *SysUserCreate {
	suc.mutation.SetSort(i)
	return suc
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableSort(i *int32) *SysUserCreate {
	if i != nil {
		suc.SetSort(*i)
	}
	return suc
}

// SetCreatedAt sets the "created_at" field.
func (suc *SysUserCreate) SetCreatedAt(t time.Time) *SysUserCreate {
	suc.mutation.SetCreatedAt(t)
	return suc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableCreatedAt(t *time.Time) *SysUserCreate {
	if t != nil {
		suc.SetCreatedAt(*t)
	}
	return suc
}

// SetUpdatedAt sets the "updated_at" field.
func (suc *SysUserCreate) SetUpdatedAt(t time.Time) *SysUserCreate {
	suc.mutation.SetUpdatedAt(t)
	return suc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableUpdatedAt(t *time.Time) *SysUserCreate {
	if t != nil {
		suc.SetUpdatedAt(*t)
	}
	return suc
}

// SetDeletedAt sets the "deleted_at" field.
func (suc *SysUserCreate) SetDeletedAt(t time.Time) *SysUserCreate {
	suc.mutation.SetDeletedAt(t)
	return suc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableDeletedAt(t *time.Time) *SysUserCreate {
	if t != nil {
		suc.SetDeletedAt(*t)
	}
	return suc
}

// SetStatus sets the "status" field.
func (suc *SysUserCreate) SetStatus(i int32) *SysUserCreate {
	suc.mutation.SetStatus(i)
	return suc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableStatus(i *int32) *SysUserCreate {
	if i != nil {
		suc.SetStatus(*i)
	}
	return suc
}

// SetUserName sets the "user_name" field.
func (suc *SysUserCreate) SetUserName(s string) *SysUserCreate {
	suc.mutation.SetUserName(s)
	return suc
}

// SetRealName sets the "real_name" field.
func (suc *SysUserCreate) SetRealName(s string) *SysUserCreate {
	suc.mutation.SetRealName(s)
	return suc
}

// SetNillableRealName sets the "real_name" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableRealName(s *string) *SysUserCreate {
	if s != nil {
		suc.SetRealName(*s)
	}
	return suc
}

// SetFirstName sets the "first_name" field.
func (suc *SysUserCreate) SetFirstName(s string) *SysUserCreate {
	suc.mutation.SetFirstName(s)
	return suc
}

// SetNillableFirstName sets the "first_name" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableFirstName(s *string) *SysUserCreate {
	if s != nil {
		suc.SetFirstName(*s)
	}
	return suc
}

// SetLastName sets the "last_name" field.
func (suc *SysUserCreate) SetLastName(s string) *SysUserCreate {
	suc.mutation.SetLastName(s)
	return suc
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableLastName(s *string) *SysUserCreate {
	if s != nil {
		suc.SetLastName(*s)
	}
	return suc
}

// SetPassword sets the "Password" field.
func (suc *SysUserCreate) SetPassword(s string) *SysUserCreate {
	suc.mutation.SetPassword(s)
	return suc
}

// SetEmail sets the "Email" field.
func (suc *SysUserCreate) SetEmail(s string) *SysUserCreate {
	suc.mutation.SetEmail(s)
	return suc
}

// SetPhone sets the "Phone" field.
func (suc *SysUserCreate) SetPhone(s string) *SysUserCreate {
	suc.mutation.SetPhone(s)
	return suc
}

// SetSalt sets the "salt" field.
func (suc *SysUserCreate) SetSalt(s string) *SysUserCreate {
	suc.mutation.SetSalt(s)
	return suc
}

// SetNillableSalt sets the "salt" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableSalt(s *string) *SysUserCreate {
	if s != nil {
		suc.SetSalt(*s)
	}
	return suc
}

// SetID sets the "id" field.
func (suc *SysUserCreate) SetID(s string) *SysUserCreate {
	suc.mutation.SetID(s)
	return suc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableID(s *string) *SysUserCreate {
	if s != nil {
		suc.SetID(*s)
	}
	return suc
}

// Mutation returns the SysUserMutation object of the builder.
func (suc *SysUserCreate) Mutation() *SysUserMutation {
	return suc.mutation
}

// Save creates the SysUser in the database.
func (suc *SysUserCreate) Save(ctx context.Context) (*SysUser, error) {
	var (
		err  error
		node *SysUser
	)
	suc.defaults()
	if len(suc.hooks) == 0 {
		if err = suc.check(); err != nil {
			return nil, err
		}
		node, err = suc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = suc.check(); err != nil {
				return nil, err
			}
			suc.mutation = mutation
			node, err = suc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suc.hooks) - 1; i >= 0; i-- {
			mut = suc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (suc *SysUserCreate) SaveX(ctx context.Context) *SysUser {
	v, err := suc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (suc *SysUserCreate) defaults() {
	if _, ok := suc.mutation.Sort(); !ok {
		v := sysuser.DefaultSort
		suc.mutation.SetSort(v)
	}
	if _, ok := suc.mutation.CreatedAt(); !ok {
		v := sysuser.DefaultCreatedAt()
		suc.mutation.SetCreatedAt(v)
	}
	if _, ok := suc.mutation.UpdatedAt(); !ok {
		v := sysuser.DefaultUpdatedAt()
		suc.mutation.SetUpdatedAt(v)
	}
	if _, ok := suc.mutation.Status(); !ok {
		v := sysuser.DefaultStatus
		suc.mutation.SetStatus(v)
	}
	if _, ok := suc.mutation.Salt(); !ok {
		v := sysuser.DefaultSalt()
		suc.mutation.SetSalt(v)
	}
	if _, ok := suc.mutation.ID(); !ok {
		v := sysuser.DefaultID
		suc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suc *SysUserCreate) check() error {
	if _, ok := suc.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New("ent: missing required field \"sort\"")}
	}
	if _, ok := suc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := suc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("ent: missing required field \"updated_at\"")}
	}
	if _, ok := suc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New("ent: missing required field \"status\"")}
	}
	if _, ok := suc.mutation.UserName(); !ok {
		return &ValidationError{Name: "user_name", err: errors.New("ent: missing required field \"user_name\"")}
	}
	if v, ok := suc.mutation.UserName(); ok {
		if err := sysuser.UserNameValidator(v); err != nil {
			return &ValidationError{Name: "user_name", err: fmt.Errorf("ent: validator failed for field \"user_name\": %w", err)}
		}
	}
	if v, ok := suc.mutation.RealName(); ok {
		if err := sysuser.RealNameValidator(v); err != nil {
			return &ValidationError{Name: "real_name", err: fmt.Errorf("ent: validator failed for field \"real_name\": %w", err)}
		}
	}
	if v, ok := suc.mutation.FirstName(); ok {
		if err := sysuser.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "first_name", err: fmt.Errorf("ent: validator failed for field \"first_name\": %w", err)}
		}
	}
	if v, ok := suc.mutation.LastName(); ok {
		if err := sysuser.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf("ent: validator failed for field \"last_name\": %w", err)}
		}
	}
	if _, ok := suc.mutation.Password(); !ok {
		return &ValidationError{Name: "Password", err: errors.New("ent: missing required field \"Password\"")}
	}
	if v, ok := suc.mutation.Password(); ok {
		if err := sysuser.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "Password", err: fmt.Errorf("ent: validator failed for field \"Password\": %w", err)}
		}
	}
	if _, ok := suc.mutation.Email(); !ok {
		return &ValidationError{Name: "Email", err: errors.New("ent: missing required field \"Email\"")}
	}
	if v, ok := suc.mutation.Email(); ok {
		if err := sysuser.EmailValidator(v); err != nil {
			return &ValidationError{Name: "Email", err: fmt.Errorf("ent: validator failed for field \"Email\": %w", err)}
		}
	}
	if _, ok := suc.mutation.Phone(); !ok {
		return &ValidationError{Name: "Phone", err: errors.New("ent: missing required field \"Phone\"")}
	}
	if v, ok := suc.mutation.Phone(); ok {
		if err := sysuser.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "Phone", err: fmt.Errorf("ent: validator failed for field \"Phone\": %w", err)}
		}
	}
	if _, ok := suc.mutation.Salt(); !ok {
		return &ValidationError{Name: "salt", err: errors.New("ent: missing required field \"salt\"")}
	}
	if v, ok := suc.mutation.ID(); ok {
		if err := sysuser.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf("ent: validator failed for field \"id\": %w", err)}
		}
	}
	return nil
}

func (suc *SysUserCreate) sqlSave(ctx context.Context) (*SysUser, error) {
	_node, _spec := suc.createSpec()
	if err := sqlgraph.CreateNode(ctx, suc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (suc *SysUserCreate) createSpec() (*SysUser, *sqlgraph.CreateSpec) {
	var (
		_node = &SysUser{config: suc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: sysuser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: sysuser.FieldID,
			},
		}
	)
	if id, ok := suc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := suc.mutation.Sort(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysuser.FieldSort,
		})
		_node.Sort = value
	}
	if value, ok := suc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysuser.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := suc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysuser.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := suc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysuser.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := suc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysuser.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := suc.mutation.UserName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysuser.FieldUserName,
		})
		_node.UserName = value
	}
	if value, ok := suc.mutation.RealName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysuser.FieldRealName,
		})
		_node.RealName = &value
	}
	if value, ok := suc.mutation.FirstName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysuser.FieldFirstName,
		})
		_node.FirstName = &value
	}
	if value, ok := suc.mutation.LastName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysuser.FieldLastName,
		})
		_node.LastName = &value
	}
	if value, ok := suc.mutation.Password(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysuser.FieldPassword,
		})
		_node.Password = value
	}
	if value, ok := suc.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysuser.FieldEmail,
		})
		_node.Email = value
	}
	if value, ok := suc.mutation.Phone(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysuser.FieldPhone,
		})
		_node.Phone = value
	}
	if value, ok := suc.mutation.Salt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysuser.FieldSalt,
		})
		_node.Salt = value
	}
	return _node, _spec
}

// SysUserCreateBulk is the builder for creating many SysUser entities in bulk.
type SysUserCreateBulk struct {
	config
	builders []*SysUserCreate
}

// Save creates the SysUser entities in the database.
func (sucb *SysUserCreateBulk) Save(ctx context.Context) ([]*SysUser, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sucb.builders))
	nodes := make([]*SysUser, len(sucb.builders))
	mutators := make([]Mutator, len(sucb.builders))
	for i := range sucb.builders {
		func(i int, root context.Context) {
			builder := sucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SysUserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, sucb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sucb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, sucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sucb *SysUserCreateBulk) SaveX(ctx context.Context) []*SysUser {
	v, err := sucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
