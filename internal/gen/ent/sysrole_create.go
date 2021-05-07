// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wanhello/omgind/internal/gen/ent/sysrole"
	"github.com/wanhello/omgind/internal/gen/ent/sysuserrole"
)

// SysRoleCreate is the builder for creating a SysRole entity.
type SysRoleCreate struct {
	config
	mutation *SysRoleMutation
	hooks    []Hook
}

// SetIsDel sets the "is_del" field.
func (src *SysRoleCreate) SetIsDel(b bool) *SysRoleCreate {
	src.mutation.SetIsDel(b)
	return src
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (src *SysRoleCreate) SetNillableIsDel(b *bool) *SysRoleCreate {
	if b != nil {
		src.SetIsDel(*b)
	}
	return src
}

// SetStatus sets the "status" field.
func (src *SysRoleCreate) SetStatus(i int32) *SysRoleCreate {
	src.mutation.SetStatus(i)
	return src
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (src *SysRoleCreate) SetNillableStatus(i *int32) *SysRoleCreate {
	if i != nil {
		src.SetStatus(*i)
	}
	return src
}

// SetSort sets the "sort" field.
func (src *SysRoleCreate) SetSort(i int32) *SysRoleCreate {
	src.mutation.SetSort(i)
	return src
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (src *SysRoleCreate) SetNillableSort(i *int32) *SysRoleCreate {
	if i != nil {
		src.SetSort(*i)
	}
	return src
}

// SetMemo sets the "memo" field.
func (src *SysRoleCreate) SetMemo(s string) *SysRoleCreate {
	src.mutation.SetMemo(s)
	return src
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (src *SysRoleCreate) SetNillableMemo(s *string) *SysRoleCreate {
	if s != nil {
		src.SetMemo(*s)
	}
	return src
}

// SetCreatedAt sets the "created_at" field.
func (src *SysRoleCreate) SetCreatedAt(t time.Time) *SysRoleCreate {
	src.mutation.SetCreatedAt(t)
	return src
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (src *SysRoleCreate) SetNillableCreatedAt(t *time.Time) *SysRoleCreate {
	if t != nil {
		src.SetCreatedAt(*t)
	}
	return src
}

// SetUpdatedAt sets the "updated_at" field.
func (src *SysRoleCreate) SetUpdatedAt(t time.Time) *SysRoleCreate {
	src.mutation.SetUpdatedAt(t)
	return src
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (src *SysRoleCreate) SetNillableUpdatedAt(t *time.Time) *SysRoleCreate {
	if t != nil {
		src.SetUpdatedAt(*t)
	}
	return src
}

// SetDeletedAt sets the "deleted_at" field.
func (src *SysRoleCreate) SetDeletedAt(t time.Time) *SysRoleCreate {
	src.mutation.SetDeletedAt(t)
	return src
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (src *SysRoleCreate) SetNillableDeletedAt(t *time.Time) *SysRoleCreate {
	if t != nil {
		src.SetDeletedAt(*t)
	}
	return src
}

// SetName sets the "name" field.
func (src *SysRoleCreate) SetName(s string) *SysRoleCreate {
	src.mutation.SetName(s)
	return src
}

// SetID sets the "id" field.
func (src *SysRoleCreate) SetID(s string) *SysRoleCreate {
	src.mutation.SetID(s)
	return src
}

// SetNillableID sets the "id" field if the given value is not nil.
func (src *SysRoleCreate) SetNillableID(s *string) *SysRoleCreate {
	if s != nil {
		src.SetID(*s)
	}
	return src
}

// AddUserRoleIDs adds the "userRoles" edge to the SysUserRole entity by IDs.
func (src *SysRoleCreate) AddUserRoleIDs(ids ...string) *SysRoleCreate {
	src.mutation.AddUserRoleIDs(ids...)
	return src
}

// AddUserRoles adds the "userRoles" edges to the SysUserRole entity.
func (src *SysRoleCreate) AddUserRoles(s ...*SysUserRole) *SysRoleCreate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return src.AddUserRoleIDs(ids...)
}

// Mutation returns the SysRoleMutation object of the builder.
func (src *SysRoleCreate) Mutation() *SysRoleMutation {
	return src.mutation
}

// Save creates the SysRole in the database.
func (src *SysRoleCreate) Save(ctx context.Context) (*SysRole, error) {
	var (
		err  error
		node *SysRole
	)
	src.defaults()
	if len(src.hooks) == 0 {
		if err = src.check(); err != nil {
			return nil, err
		}
		node, err = src.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = src.check(); err != nil {
				return nil, err
			}
			src.mutation = mutation
			node, err = src.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(src.hooks) - 1; i >= 0; i-- {
			mut = src.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, src.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (src *SysRoleCreate) SaveX(ctx context.Context) *SysRole {
	v, err := src.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (src *SysRoleCreate) defaults() {
	if _, ok := src.mutation.IsDel(); !ok {
		v := sysrole.DefaultIsDel
		src.mutation.SetIsDel(v)
	}
	if _, ok := src.mutation.Status(); !ok {
		v := sysrole.DefaultStatus
		src.mutation.SetStatus(v)
	}
	if _, ok := src.mutation.Sort(); !ok {
		v := sysrole.DefaultSort
		src.mutation.SetSort(v)
	}
	if _, ok := src.mutation.Memo(); !ok {
		v := sysrole.DefaultMemo
		src.mutation.SetMemo(v)
	}
	if _, ok := src.mutation.CreatedAt(); !ok {
		v := sysrole.DefaultCreatedAt()
		src.mutation.SetCreatedAt(v)
	}
	if _, ok := src.mutation.UpdatedAt(); !ok {
		v := sysrole.DefaultUpdatedAt()
		src.mutation.SetUpdatedAt(v)
	}
	if _, ok := src.mutation.ID(); !ok {
		v := sysrole.DefaultID
		src.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (src *SysRoleCreate) check() error {
	if _, ok := src.mutation.IsDel(); !ok {
		return &ValidationError{Name: "is_del", err: errors.New("ent: missing required field \"is_del\"")}
	}
	if _, ok := src.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New("ent: missing required field \"status\"")}
	}
	if _, ok := src.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New("ent: missing required field \"sort\"")}
	}
	if _, ok := src.mutation.Memo(); !ok {
		return &ValidationError{Name: "memo", err: errors.New("ent: missing required field \"memo\"")}
	}
	if v, ok := src.mutation.Memo(); ok {
		if err := sysrole.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf("ent: validator failed for field \"memo\": %w", err)}
		}
	}
	if _, ok := src.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := src.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("ent: missing required field \"updated_at\"")}
	}
	if _, ok := src.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := src.mutation.Name(); ok {
		if err := sysrole.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := src.mutation.ID(); ok {
		if err := sysrole.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf("ent: validator failed for field \"id\": %w", err)}
		}
	}
	return nil
}

func (src *SysRoleCreate) sqlSave(ctx context.Context) (*SysRole, error) {
	_node, _spec := src.createSpec()
	if err := sqlgraph.CreateNode(ctx, src.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (src *SysRoleCreate) createSpec() (*SysRole, *sqlgraph.CreateSpec) {
	var (
		_node = &SysRole{config: src.config}
		_spec = &sqlgraph.CreateSpec{
			Table: sysrole.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: sysrole.FieldID,
			},
		}
	)
	if id, ok := src.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := src.mutation.IsDel(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysrole.FieldIsDel,
		})
		_node.IsDel = value
	}
	if value, ok := src.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysrole.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := src.mutation.Sort(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysrole.FieldSort,
		})
		_node.Sort = value
	}
	if value, ok := src.mutation.Memo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrole.FieldMemo,
		})
		_node.Memo = value
	}
	if value, ok := src.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysrole.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := src.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysrole.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := src.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysrole.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := src.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrole.FieldName,
		})
		_node.Name = value
	}
	if nodes := src.mutation.UserRolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   sysrole.UserRolesTable,
			Columns: []string{sysrole.UserRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: sysuserrole.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SysRoleCreateBulk is the builder for creating many SysRole entities in bulk.
type SysRoleCreateBulk struct {
	config
	builders []*SysRoleCreate
}

// Save creates the SysRole entities in the database.
func (srcb *SysRoleCreateBulk) Save(ctx context.Context) ([]*SysRole, error) {
	specs := make([]*sqlgraph.CreateSpec, len(srcb.builders))
	nodes := make([]*SysRole, len(srcb.builders))
	mutators := make([]Mutator, len(srcb.builders))
	for i := range srcb.builders {
		func(i int, root context.Context) {
			builder := srcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SysRoleMutation)
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
					_, err = mutators[i+1].Mutate(root, srcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, srcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, srcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (srcb *SysRoleCreateBulk) SaveX(ctx context.Context) []*SysRole {
	v, err := srcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
