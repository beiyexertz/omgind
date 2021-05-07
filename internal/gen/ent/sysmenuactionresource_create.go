// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wanhello/omgind/internal/gen/ent/sysmenuaction"
	"github.com/wanhello/omgind/internal/gen/ent/sysmenuactionresource"
)

// SysMenuActionResourceCreate is the builder for creating a SysMenuActionResource entity.
type SysMenuActionResourceCreate struct {
	config
	mutation *SysMenuActionResourceMutation
	hooks    []Hook
}

// SetIsDel sets the "is_del" field.
func (smarc *SysMenuActionResourceCreate) SetIsDel(b bool) *SysMenuActionResourceCreate {
	smarc.mutation.SetIsDel(b)
	return smarc
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (smarc *SysMenuActionResourceCreate) SetNillableIsDel(b *bool) *SysMenuActionResourceCreate {
	if b != nil {
		smarc.SetIsDel(*b)
	}
	return smarc
}

// SetSort sets the "sort" field.
func (smarc *SysMenuActionResourceCreate) SetSort(i int32) *SysMenuActionResourceCreate {
	smarc.mutation.SetSort(i)
	return smarc
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (smarc *SysMenuActionResourceCreate) SetNillableSort(i *int32) *SysMenuActionResourceCreate {
	if i != nil {
		smarc.SetSort(*i)
	}
	return smarc
}

// SetMemo sets the "memo" field.
func (smarc *SysMenuActionResourceCreate) SetMemo(s string) *SysMenuActionResourceCreate {
	smarc.mutation.SetMemo(s)
	return smarc
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (smarc *SysMenuActionResourceCreate) SetNillableMemo(s *string) *SysMenuActionResourceCreate {
	if s != nil {
		smarc.SetMemo(*s)
	}
	return smarc
}

// SetCreatedAt sets the "created_at" field.
func (smarc *SysMenuActionResourceCreate) SetCreatedAt(t time.Time) *SysMenuActionResourceCreate {
	smarc.mutation.SetCreatedAt(t)
	return smarc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (smarc *SysMenuActionResourceCreate) SetNillableCreatedAt(t *time.Time) *SysMenuActionResourceCreate {
	if t != nil {
		smarc.SetCreatedAt(*t)
	}
	return smarc
}

// SetUpdatedAt sets the "updated_at" field.
func (smarc *SysMenuActionResourceCreate) SetUpdatedAt(t time.Time) *SysMenuActionResourceCreate {
	smarc.mutation.SetUpdatedAt(t)
	return smarc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (smarc *SysMenuActionResourceCreate) SetNillableUpdatedAt(t *time.Time) *SysMenuActionResourceCreate {
	if t != nil {
		smarc.SetUpdatedAt(*t)
	}
	return smarc
}

// SetDeletedAt sets the "deleted_at" field.
func (smarc *SysMenuActionResourceCreate) SetDeletedAt(t time.Time) *SysMenuActionResourceCreate {
	smarc.mutation.SetDeletedAt(t)
	return smarc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (smarc *SysMenuActionResourceCreate) SetNillableDeletedAt(t *time.Time) *SysMenuActionResourceCreate {
	if t != nil {
		smarc.SetDeletedAt(*t)
	}
	return smarc
}

// SetStatus sets the "status" field.
func (smarc *SysMenuActionResourceCreate) SetStatus(i int32) *SysMenuActionResourceCreate {
	smarc.mutation.SetStatus(i)
	return smarc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (smarc *SysMenuActionResourceCreate) SetNillableStatus(i *int32) *SysMenuActionResourceCreate {
	if i != nil {
		smarc.SetStatus(*i)
	}
	return smarc
}

// SetMethod sets the "method" field.
func (smarc *SysMenuActionResourceCreate) SetMethod(s string) *SysMenuActionResourceCreate {
	smarc.mutation.SetMethod(s)
	return smarc
}

// SetPath sets the "path" field.
func (smarc *SysMenuActionResourceCreate) SetPath(s string) *SysMenuActionResourceCreate {
	smarc.mutation.SetPath(s)
	return smarc
}

// SetActionID sets the "action_id" field.
func (smarc *SysMenuActionResourceCreate) SetActionID(s string) *SysMenuActionResourceCreate {
	smarc.mutation.SetActionID(s)
	return smarc
}

// SetID sets the "id" field.
func (smarc *SysMenuActionResourceCreate) SetID(s string) *SysMenuActionResourceCreate {
	smarc.mutation.SetID(s)
	return smarc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (smarc *SysMenuActionResourceCreate) SetNillableID(s *string) *SysMenuActionResourceCreate {
	if s != nil {
		smarc.SetID(*s)
	}
	return smarc
}

// SetAction sets the "action" edge to the SysMenuAction entity.
func (smarc *SysMenuActionResourceCreate) SetAction(s *SysMenuAction) *SysMenuActionResourceCreate {
	return smarc.SetActionID(s.ID)
}

// Mutation returns the SysMenuActionResourceMutation object of the builder.
func (smarc *SysMenuActionResourceCreate) Mutation() *SysMenuActionResourceMutation {
	return smarc.mutation
}

// Save creates the SysMenuActionResource in the database.
func (smarc *SysMenuActionResourceCreate) Save(ctx context.Context) (*SysMenuActionResource, error) {
	var (
		err  error
		node *SysMenuActionResource
	)
	smarc.defaults()
	if len(smarc.hooks) == 0 {
		if err = smarc.check(); err != nil {
			return nil, err
		}
		node, err = smarc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysMenuActionResourceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = smarc.check(); err != nil {
				return nil, err
			}
			smarc.mutation = mutation
			node, err = smarc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(smarc.hooks) - 1; i >= 0; i-- {
			mut = smarc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, smarc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (smarc *SysMenuActionResourceCreate) SaveX(ctx context.Context) *SysMenuActionResource {
	v, err := smarc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (smarc *SysMenuActionResourceCreate) defaults() {
	if _, ok := smarc.mutation.IsDel(); !ok {
		v := sysmenuactionresource.DefaultIsDel
		smarc.mutation.SetIsDel(v)
	}
	if _, ok := smarc.mutation.Sort(); !ok {
		v := sysmenuactionresource.DefaultSort
		smarc.mutation.SetSort(v)
	}
	if _, ok := smarc.mutation.Memo(); !ok {
		v := sysmenuactionresource.DefaultMemo
		smarc.mutation.SetMemo(v)
	}
	if _, ok := smarc.mutation.CreatedAt(); !ok {
		v := sysmenuactionresource.DefaultCreatedAt()
		smarc.mutation.SetCreatedAt(v)
	}
	if _, ok := smarc.mutation.UpdatedAt(); !ok {
		v := sysmenuactionresource.DefaultUpdatedAt()
		smarc.mutation.SetUpdatedAt(v)
	}
	if _, ok := smarc.mutation.Status(); !ok {
		v := sysmenuactionresource.DefaultStatus
		smarc.mutation.SetStatus(v)
	}
	if _, ok := smarc.mutation.ID(); !ok {
		v := sysmenuactionresource.DefaultID
		smarc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (smarc *SysMenuActionResourceCreate) check() error {
	if _, ok := smarc.mutation.IsDel(); !ok {
		return &ValidationError{Name: "is_del", err: errors.New("ent: missing required field \"is_del\"")}
	}
	if _, ok := smarc.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New("ent: missing required field \"sort\"")}
	}
	if _, ok := smarc.mutation.Memo(); !ok {
		return &ValidationError{Name: "memo", err: errors.New("ent: missing required field \"memo\"")}
	}
	if v, ok := smarc.mutation.Memo(); ok {
		if err := sysmenuactionresource.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf("ent: validator failed for field \"memo\": %w", err)}
		}
	}
	if _, ok := smarc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := smarc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("ent: missing required field \"updated_at\"")}
	}
	if _, ok := smarc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New("ent: missing required field \"status\"")}
	}
	if _, ok := smarc.mutation.Method(); !ok {
		return &ValidationError{Name: "method", err: errors.New("ent: missing required field \"method\"")}
	}
	if v, ok := smarc.mutation.Method(); ok {
		if err := sysmenuactionresource.MethodValidator(v); err != nil {
			return &ValidationError{Name: "method", err: fmt.Errorf("ent: validator failed for field \"method\": %w", err)}
		}
	}
	if _, ok := smarc.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New("ent: missing required field \"path\"")}
	}
	if v, ok := smarc.mutation.Path(); ok {
		if err := sysmenuactionresource.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf("ent: validator failed for field \"path\": %w", err)}
		}
	}
	if _, ok := smarc.mutation.ActionID(); !ok {
		return &ValidationError{Name: "action_id", err: errors.New("ent: missing required field \"action_id\"")}
	}
	if v, ok := smarc.mutation.ActionID(); ok {
		if err := sysmenuactionresource.ActionIDValidator(v); err != nil {
			return &ValidationError{Name: "action_id", err: fmt.Errorf("ent: validator failed for field \"action_id\": %w", err)}
		}
	}
	if v, ok := smarc.mutation.ID(); ok {
		if err := sysmenuactionresource.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf("ent: validator failed for field \"id\": %w", err)}
		}
	}
	if _, ok := smarc.mutation.ActionID(); !ok {
		return &ValidationError{Name: "action", err: errors.New("ent: missing required edge \"action\"")}
	}
	return nil
}

func (smarc *SysMenuActionResourceCreate) sqlSave(ctx context.Context) (*SysMenuActionResource, error) {
	_node, _spec := smarc.createSpec()
	if err := sqlgraph.CreateNode(ctx, smarc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (smarc *SysMenuActionResourceCreate) createSpec() (*SysMenuActionResource, *sqlgraph.CreateSpec) {
	var (
		_node = &SysMenuActionResource{config: smarc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: sysmenuactionresource.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: sysmenuactionresource.FieldID,
			},
		}
	)
	if id, ok := smarc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := smarc.mutation.IsDel(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysmenuactionresource.FieldIsDel,
		})
		_node.IsDel = value
	}
	if value, ok := smarc.mutation.Sort(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysmenuactionresource.FieldSort,
		})
		_node.Sort = value
	}
	if value, ok := smarc.mutation.Memo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenuactionresource.FieldMemo,
		})
		_node.Memo = value
	}
	if value, ok := smarc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysmenuactionresource.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := smarc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysmenuactionresource.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := smarc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysmenuactionresource.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := smarc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysmenuactionresource.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := smarc.mutation.Method(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenuactionresource.FieldMethod,
		})
		_node.Method = value
	}
	if value, ok := smarc.mutation.Path(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenuactionresource.FieldPath,
		})
		_node.Path = value
	}
	if nodes := smarc.mutation.ActionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sysmenuactionresource.ActionTable,
			Columns: []string{sysmenuactionresource.ActionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: sysmenuaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ActionID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SysMenuActionResourceCreateBulk is the builder for creating many SysMenuActionResource entities in bulk.
type SysMenuActionResourceCreateBulk struct {
	config
	builders []*SysMenuActionResourceCreate
}

// Save creates the SysMenuActionResource entities in the database.
func (smarcb *SysMenuActionResourceCreateBulk) Save(ctx context.Context) ([]*SysMenuActionResource, error) {
	specs := make([]*sqlgraph.CreateSpec, len(smarcb.builders))
	nodes := make([]*SysMenuActionResource, len(smarcb.builders))
	mutators := make([]Mutator, len(smarcb.builders))
	for i := range smarcb.builders {
		func(i int, root context.Context) {
			builder := smarcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SysMenuActionResourceMutation)
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
					_, err = mutators[i+1].Mutate(root, smarcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, smarcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, smarcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (smarcb *SysMenuActionResourceCreateBulk) SaveX(ctx context.Context) []*SysMenuActionResource {
	v, err := smarcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
