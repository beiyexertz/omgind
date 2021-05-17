// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wanhello/omgind/internal/gen/ent/sysjwtblock"
)

// SysJwtBlockCreate is the builder for creating a SysJwtBlock entity.
type SysJwtBlockCreate struct {
	config
	mutation *SysJwtBlockMutation
	hooks    []Hook
}

// SetMemo sets the "memo" field.
func (sjbc *SysJwtBlockCreate) SetMemo(s string) *SysJwtBlockCreate {
	sjbc.mutation.SetMemo(s)
	return sjbc
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (sjbc *SysJwtBlockCreate) SetNillableMemo(s *string) *SysJwtBlockCreate {
	if s != nil {
		sjbc.SetMemo(*s)
	}
	return sjbc
}

// SetCreatedAt sets the "created_at" field.
func (sjbc *SysJwtBlockCreate) SetCreatedAt(t time.Time) *SysJwtBlockCreate {
	sjbc.mutation.SetCreatedAt(t)
	return sjbc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sjbc *SysJwtBlockCreate) SetNillableCreatedAt(t *time.Time) *SysJwtBlockCreate {
	if t != nil {
		sjbc.SetCreatedAt(*t)
	}
	return sjbc
}

// SetUpdatedAt sets the "updated_at" field.
func (sjbc *SysJwtBlockCreate) SetUpdatedAt(t time.Time) *SysJwtBlockCreate {
	sjbc.mutation.SetUpdatedAt(t)
	return sjbc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sjbc *SysJwtBlockCreate) SetNillableUpdatedAt(t *time.Time) *SysJwtBlockCreate {
	if t != nil {
		sjbc.SetUpdatedAt(*t)
	}
	return sjbc
}

// SetDeletedAt sets the "deleted_at" field.
func (sjbc *SysJwtBlockCreate) SetDeletedAt(t time.Time) *SysJwtBlockCreate {
	sjbc.mutation.SetDeletedAt(t)
	return sjbc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sjbc *SysJwtBlockCreate) SetNillableDeletedAt(t *time.Time) *SysJwtBlockCreate {
	if t != nil {
		sjbc.SetDeletedAt(*t)
	}
	return sjbc
}

// SetStatus sets the "status" field.
func (sjbc *SysJwtBlockCreate) SetStatus(i int32) *SysJwtBlockCreate {
	sjbc.mutation.SetStatus(i)
	return sjbc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sjbc *SysJwtBlockCreate) SetNillableStatus(i *int32) *SysJwtBlockCreate {
	if i != nil {
		sjbc.SetStatus(*i)
	}
	return sjbc
}

// SetJwt sets the "jwt" field.
func (sjbc *SysJwtBlockCreate) SetJwt(s string) *SysJwtBlockCreate {
	sjbc.mutation.SetJwt(s)
	return sjbc
}

// SetID sets the "id" field.
func (sjbc *SysJwtBlockCreate) SetID(s string) *SysJwtBlockCreate {
	sjbc.mutation.SetID(s)
	return sjbc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sjbc *SysJwtBlockCreate) SetNillableID(s *string) *SysJwtBlockCreate {
	if s != nil {
		sjbc.SetID(*s)
	}
	return sjbc
}

// Mutation returns the SysJwtBlockMutation object of the builder.
func (sjbc *SysJwtBlockCreate) Mutation() *SysJwtBlockMutation {
	return sjbc.mutation
}

// Save creates the SysJwtBlock in the database.
func (sjbc *SysJwtBlockCreate) Save(ctx context.Context) (*SysJwtBlock, error) {
	var (
		err  error
		node *SysJwtBlock
	)
	sjbc.defaults()
	if len(sjbc.hooks) == 0 {
		if err = sjbc.check(); err != nil {
			return nil, err
		}
		node, err = sjbc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysJwtBlockMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sjbc.check(); err != nil {
				return nil, err
			}
			sjbc.mutation = mutation
			node, err = sjbc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sjbc.hooks) - 1; i >= 0; i-- {
			mut = sjbc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sjbc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sjbc *SysJwtBlockCreate) SaveX(ctx context.Context) *SysJwtBlock {
	v, err := sjbc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (sjbc *SysJwtBlockCreate) defaults() {
	if _, ok := sjbc.mutation.Memo(); !ok {
		v := sysjwtblock.DefaultMemo
		sjbc.mutation.SetMemo(v)
	}
	if _, ok := sjbc.mutation.CreatedAt(); !ok {
		v := sysjwtblock.DefaultCreatedAt()
		sjbc.mutation.SetCreatedAt(v)
	}
	if _, ok := sjbc.mutation.UpdatedAt(); !ok {
		v := sysjwtblock.DefaultUpdatedAt()
		sjbc.mutation.SetUpdatedAt(v)
	}
	if _, ok := sjbc.mutation.Status(); !ok {
		v := sysjwtblock.DefaultStatus
		sjbc.mutation.SetStatus(v)
	}
	if _, ok := sjbc.mutation.ID(); !ok {
		v := sysjwtblock.DefaultID
		sjbc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sjbc *SysJwtBlockCreate) check() error {
	if _, ok := sjbc.mutation.Memo(); !ok {
		return &ValidationError{Name: "memo", err: errors.New("ent: missing required field \"memo\"")}
	}
	if v, ok := sjbc.mutation.Memo(); ok {
		if err := sysjwtblock.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf("ent: validator failed for field \"memo\": %w", err)}
		}
	}
	if _, ok := sjbc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := sjbc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("ent: missing required field \"updated_at\"")}
	}
	if _, ok := sjbc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New("ent: missing required field \"status\"")}
	}
	if _, ok := sjbc.mutation.Jwt(); !ok {
		return &ValidationError{Name: "jwt", err: errors.New("ent: missing required field \"jwt\"")}
	}
	if v, ok := sjbc.mutation.Jwt(); ok {
		if err := sysjwtblock.JwtValidator(v); err != nil {
			return &ValidationError{Name: "jwt", err: fmt.Errorf("ent: validator failed for field \"jwt\": %w", err)}
		}
	}
	if v, ok := sjbc.mutation.ID(); ok {
		if err := sysjwtblock.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf("ent: validator failed for field \"id\": %w", err)}
		}
	}
	return nil
}

func (sjbc *SysJwtBlockCreate) sqlSave(ctx context.Context) (*SysJwtBlock, error) {
	_node, _spec := sjbc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sjbc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (sjbc *SysJwtBlockCreate) createSpec() (*SysJwtBlock, *sqlgraph.CreateSpec) {
	var (
		_node = &SysJwtBlock{config: sjbc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: sysjwtblock.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: sysjwtblock.FieldID,
			},
		}
	)
	if id, ok := sjbc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sjbc.mutation.Memo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysjwtblock.FieldMemo,
		})
		_node.Memo = value
	}
	if value, ok := sjbc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysjwtblock.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := sjbc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysjwtblock.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := sjbc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysjwtblock.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := sjbc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysjwtblock.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := sjbc.mutation.Jwt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysjwtblock.FieldJwt,
		})
		_node.Jwt = value
	}
	return _node, _spec
}

// SysJwtBlockCreateBulk is the builder for creating many SysJwtBlock entities in bulk.
type SysJwtBlockCreateBulk struct {
	config
	builders []*SysJwtBlockCreate
}

// Save creates the SysJwtBlock entities in the database.
func (sjbcb *SysJwtBlockCreateBulk) Save(ctx context.Context) ([]*SysJwtBlock, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sjbcb.builders))
	nodes := make([]*SysJwtBlock, len(sjbcb.builders))
	mutators := make([]Mutator, len(sjbcb.builders))
	for i := range sjbcb.builders {
		func(i int, root context.Context) {
			builder := sjbcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SysJwtBlockMutation)
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
					_, err = mutators[i+1].Mutate(root, sjbcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sjbcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, sjbcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sjbcb *SysJwtBlockCreateBulk) SaveX(ctx context.Context) []*SysJwtBlock {
	v, err := sjbcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
