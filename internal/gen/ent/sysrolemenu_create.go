// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wanhello/omgind/internal/gen/ent/sysrolemenu"
)

// SysRoleMenuCreate is the builder for creating a SysRoleMenu entity.
type SysRoleMenuCreate struct {
	config
	mutation *SysRoleMenuMutation
	hooks    []Hook
}

// SetIsDel sets the "is_del" field.
func (srmc *SysRoleMenuCreate) SetIsDel(b bool) *SysRoleMenuCreate {
	srmc.mutation.SetIsDel(b)
	return srmc
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (srmc *SysRoleMenuCreate) SetNillableIsDel(b *bool) *SysRoleMenuCreate {
	if b != nil {
		srmc.SetIsDel(*b)
	}
	return srmc
}

// SetCreatedAt sets the "created_at" field.
func (srmc *SysRoleMenuCreate) SetCreatedAt(t time.Time) *SysRoleMenuCreate {
	srmc.mutation.SetCreatedAt(t)
	return srmc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (srmc *SysRoleMenuCreate) SetNillableCreatedAt(t *time.Time) *SysRoleMenuCreate {
	if t != nil {
		srmc.SetCreatedAt(*t)
	}
	return srmc
}

// SetUpdatedAt sets the "updated_at" field.
func (srmc *SysRoleMenuCreate) SetUpdatedAt(t time.Time) *SysRoleMenuCreate {
	srmc.mutation.SetUpdatedAt(t)
	return srmc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (srmc *SysRoleMenuCreate) SetNillableUpdatedAt(t *time.Time) *SysRoleMenuCreate {
	if t != nil {
		srmc.SetUpdatedAt(*t)
	}
	return srmc
}

// SetDeletedAt sets the "deleted_at" field.
func (srmc *SysRoleMenuCreate) SetDeletedAt(t time.Time) *SysRoleMenuCreate {
	srmc.mutation.SetDeletedAt(t)
	return srmc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (srmc *SysRoleMenuCreate) SetNillableDeletedAt(t *time.Time) *SysRoleMenuCreate {
	if t != nil {
		srmc.SetDeletedAt(*t)
	}
	return srmc
}

// SetRoleID sets the "role_id" field.
func (srmc *SysRoleMenuCreate) SetRoleID(s string) *SysRoleMenuCreate {
	srmc.mutation.SetRoleID(s)
	return srmc
}

// SetMenuID sets the "menu_id" field.
func (srmc *SysRoleMenuCreate) SetMenuID(s string) *SysRoleMenuCreate {
	srmc.mutation.SetMenuID(s)
	return srmc
}

// SetActionID sets the "action_id" field.
func (srmc *SysRoleMenuCreate) SetActionID(s string) *SysRoleMenuCreate {
	srmc.mutation.SetActionID(s)
	return srmc
}

// SetNillableActionID sets the "action_id" field if the given value is not nil.
func (srmc *SysRoleMenuCreate) SetNillableActionID(s *string) *SysRoleMenuCreate {
	if s != nil {
		srmc.SetActionID(*s)
	}
	return srmc
}

// SetID sets the "id" field.
func (srmc *SysRoleMenuCreate) SetID(s string) *SysRoleMenuCreate {
	srmc.mutation.SetID(s)
	return srmc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (srmc *SysRoleMenuCreate) SetNillableID(s *string) *SysRoleMenuCreate {
	if s != nil {
		srmc.SetID(*s)
	}
	return srmc
}

// Mutation returns the SysRoleMenuMutation object of the builder.
func (srmc *SysRoleMenuCreate) Mutation() *SysRoleMenuMutation {
	return srmc.mutation
}

// Save creates the SysRoleMenu in the database.
func (srmc *SysRoleMenuCreate) Save(ctx context.Context) (*SysRoleMenu, error) {
	var (
		err  error
		node *SysRoleMenu
	)
	srmc.defaults()
	if len(srmc.hooks) == 0 {
		if err = srmc.check(); err != nil {
			return nil, err
		}
		node, err = srmc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysRoleMenuMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = srmc.check(); err != nil {
				return nil, err
			}
			srmc.mutation = mutation
			node, err = srmc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(srmc.hooks) - 1; i >= 0; i-- {
			mut = srmc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, srmc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (srmc *SysRoleMenuCreate) SaveX(ctx context.Context) *SysRoleMenu {
	v, err := srmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (srmc *SysRoleMenuCreate) defaults() {
	if _, ok := srmc.mutation.IsDel(); !ok {
		v := sysrolemenu.DefaultIsDel
		srmc.mutation.SetIsDel(v)
	}
	if _, ok := srmc.mutation.CreatedAt(); !ok {
		v := sysrolemenu.DefaultCreatedAt()
		srmc.mutation.SetCreatedAt(v)
	}
	if _, ok := srmc.mutation.UpdatedAt(); !ok {
		v := sysrolemenu.DefaultUpdatedAt()
		srmc.mutation.SetUpdatedAt(v)
	}
	if _, ok := srmc.mutation.ID(); !ok {
		v := sysrolemenu.DefaultID()
		srmc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (srmc *SysRoleMenuCreate) check() error {
	if _, ok := srmc.mutation.IsDel(); !ok {
		return &ValidationError{Name: "is_del", err: errors.New("ent: missing required field \"is_del\"")}
	}
	if _, ok := srmc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := srmc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("ent: missing required field \"updated_at\"")}
	}
	if _, ok := srmc.mutation.RoleID(); !ok {
		return &ValidationError{Name: "role_id", err: errors.New("ent: missing required field \"role_id\"")}
	}
	if v, ok := srmc.mutation.RoleID(); ok {
		if err := sysrolemenu.RoleIDValidator(v); err != nil {
			return &ValidationError{Name: "role_id", err: fmt.Errorf("ent: validator failed for field \"role_id\": %w", err)}
		}
	}
	if _, ok := srmc.mutation.MenuID(); !ok {
		return &ValidationError{Name: "menu_id", err: errors.New("ent: missing required field \"menu_id\"")}
	}
	if v, ok := srmc.mutation.MenuID(); ok {
		if err := sysrolemenu.MenuIDValidator(v); err != nil {
			return &ValidationError{Name: "menu_id", err: fmt.Errorf("ent: validator failed for field \"menu_id\": %w", err)}
		}
	}
	if v, ok := srmc.mutation.ActionID(); ok {
		if err := sysrolemenu.ActionIDValidator(v); err != nil {
			return &ValidationError{Name: "action_id", err: fmt.Errorf("ent: validator failed for field \"action_id\": %w", err)}
		}
	}
	if v, ok := srmc.mutation.ID(); ok {
		if err := sysrolemenu.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf("ent: validator failed for field \"id\": %w", err)}
		}
	}
	return nil
}

func (srmc *SysRoleMenuCreate) sqlSave(ctx context.Context) (*SysRoleMenu, error) {
	_node, _spec := srmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, srmc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (srmc *SysRoleMenuCreate) createSpec() (*SysRoleMenu, *sqlgraph.CreateSpec) {
	var (
		_node = &SysRoleMenu{config: srmc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: sysrolemenu.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: sysrolemenu.FieldID,
			},
		}
	)
	if id, ok := srmc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := srmc.mutation.IsDel(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysrolemenu.FieldIsDel,
		})
		_node.IsDel = value
	}
	if value, ok := srmc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysrolemenu.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := srmc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysrolemenu.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := srmc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysrolemenu.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := srmc.mutation.RoleID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrolemenu.FieldRoleID,
		})
		_node.RoleID = value
	}
	if value, ok := srmc.mutation.MenuID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrolemenu.FieldMenuID,
		})
		_node.MenuID = value
	}
	if value, ok := srmc.mutation.ActionID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrolemenu.FieldActionID,
		})
		_node.ActionID = &value
	}
	return _node, _spec
}

// SysRoleMenuCreateBulk is the builder for creating many SysRoleMenu entities in bulk.
type SysRoleMenuCreateBulk struct {
	config
	builders []*SysRoleMenuCreate
}

// Save creates the SysRoleMenu entities in the database.
func (srmcb *SysRoleMenuCreateBulk) Save(ctx context.Context) ([]*SysRoleMenu, error) {
	specs := make([]*sqlgraph.CreateSpec, len(srmcb.builders))
	nodes := make([]*SysRoleMenu, len(srmcb.builders))
	mutators := make([]Mutator, len(srmcb.builders))
	for i := range srmcb.builders {
		func(i int, root context.Context) {
			builder := srmcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SysRoleMenuMutation)
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
					_, err = mutators[i+1].Mutate(root, srmcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, srmcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, srmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (srmcb *SysRoleMenuCreateBulk) SaveX(ctx context.Context) []*SysRoleMenu {
	v, err := srmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
