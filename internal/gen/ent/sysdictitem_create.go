// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wanhello/omgind/internal/gen/ent/sysdictitem"
)

// SysDictItemCreate is the builder for creating a SysDictItem entity.
type SysDictItemCreate struct {
	config
	mutation *SysDictItemMutation
	hooks    []Hook
}

// SetIsDel sets the "is_del" field.
func (sdic *SysDictItemCreate) SetIsDel(b bool) *SysDictItemCreate {
	sdic.mutation.SetIsDel(b)
	return sdic
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (sdic *SysDictItemCreate) SetNillableIsDel(b *bool) *SysDictItemCreate {
	if b != nil {
		sdic.SetIsDel(*b)
	}
	return sdic
}

// SetMemo sets the "memo" field.
func (sdic *SysDictItemCreate) SetMemo(s string) *SysDictItemCreate {
	sdic.mutation.SetMemo(s)
	return sdic
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (sdic *SysDictItemCreate) SetNillableMemo(s *string) *SysDictItemCreate {
	if s != nil {
		sdic.SetMemo(*s)
	}
	return sdic
}

// SetSort sets the "sort" field.
func (sdic *SysDictItemCreate) SetSort(i int32) *SysDictItemCreate {
	sdic.mutation.SetSort(i)
	return sdic
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (sdic *SysDictItemCreate) SetNillableSort(i *int32) *SysDictItemCreate {
	if i != nil {
		sdic.SetSort(*i)
	}
	return sdic
}

// SetCreatedAt sets the "created_at" field.
func (sdic *SysDictItemCreate) SetCreatedAt(t time.Time) *SysDictItemCreate {
	sdic.mutation.SetCreatedAt(t)
	return sdic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sdic *SysDictItemCreate) SetNillableCreatedAt(t *time.Time) *SysDictItemCreate {
	if t != nil {
		sdic.SetCreatedAt(*t)
	}
	return sdic
}

// SetUpdatedAt sets the "updated_at" field.
func (sdic *SysDictItemCreate) SetUpdatedAt(t time.Time) *SysDictItemCreate {
	sdic.mutation.SetUpdatedAt(t)
	return sdic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sdic *SysDictItemCreate) SetNillableUpdatedAt(t *time.Time) *SysDictItemCreate {
	if t != nil {
		sdic.SetUpdatedAt(*t)
	}
	return sdic
}

// SetDeletedAt sets the "deleted_at" field.
func (sdic *SysDictItemCreate) SetDeletedAt(t time.Time) *SysDictItemCreate {
	sdic.mutation.SetDeletedAt(t)
	return sdic
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sdic *SysDictItemCreate) SetNillableDeletedAt(t *time.Time) *SysDictItemCreate {
	if t != nil {
		sdic.SetDeletedAt(*t)
	}
	return sdic
}

// SetLabel sets the "label" field.
func (sdic *SysDictItemCreate) SetLabel(s string) *SysDictItemCreate {
	sdic.mutation.SetLabel(s)
	return sdic
}

// SetValue sets the "value" field.
func (sdic *SysDictItemCreate) SetValue(i int) *SysDictItemCreate {
	sdic.mutation.SetValue(i)
	return sdic
}

// SetStatus sets the "status" field.
func (sdic *SysDictItemCreate) SetStatus(i int16) *SysDictItemCreate {
	sdic.mutation.SetStatus(i)
	return sdic
}

// SetDictID sets the "dict_id" field.
func (sdic *SysDictItemCreate) SetDictID(s string) *SysDictItemCreate {
	sdic.mutation.SetDictID(s)
	return sdic
}

// SetID sets the "id" field.
func (sdic *SysDictItemCreate) SetID(s string) *SysDictItemCreate {
	sdic.mutation.SetID(s)
	return sdic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sdic *SysDictItemCreate) SetNillableID(s *string) *SysDictItemCreate {
	if s != nil {
		sdic.SetID(*s)
	}
	return sdic
}

// Mutation returns the SysDictItemMutation object of the builder.
func (sdic *SysDictItemCreate) Mutation() *SysDictItemMutation {
	return sdic.mutation
}

// Save creates the SysDictItem in the database.
func (sdic *SysDictItemCreate) Save(ctx context.Context) (*SysDictItem, error) {
	var (
		err  error
		node *SysDictItem
	)
	sdic.defaults()
	if len(sdic.hooks) == 0 {
		if err = sdic.check(); err != nil {
			return nil, err
		}
		node, err = sdic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysDictItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sdic.check(); err != nil {
				return nil, err
			}
			sdic.mutation = mutation
			node, err = sdic.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sdic.hooks) - 1; i >= 0; i-- {
			mut = sdic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sdic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sdic *SysDictItemCreate) SaveX(ctx context.Context) *SysDictItem {
	v, err := sdic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (sdic *SysDictItemCreate) defaults() {
	if _, ok := sdic.mutation.IsDel(); !ok {
		v := sysdictitem.DefaultIsDel
		sdic.mutation.SetIsDel(v)
	}
	if _, ok := sdic.mutation.Memo(); !ok {
		v := sysdictitem.DefaultMemo
		sdic.mutation.SetMemo(v)
	}
	if _, ok := sdic.mutation.Sort(); !ok {
		v := sysdictitem.DefaultSort
		sdic.mutation.SetSort(v)
	}
	if _, ok := sdic.mutation.CreatedAt(); !ok {
		v := sysdictitem.DefaultCreatedAt()
		sdic.mutation.SetCreatedAt(v)
	}
	if _, ok := sdic.mutation.UpdatedAt(); !ok {
		v := sysdictitem.DefaultUpdatedAt()
		sdic.mutation.SetUpdatedAt(v)
	}
	if _, ok := sdic.mutation.ID(); !ok {
		v := sysdictitem.DefaultID()
		sdic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sdic *SysDictItemCreate) check() error {
	if _, ok := sdic.mutation.IsDel(); !ok {
		return &ValidationError{Name: "is_del", err: errors.New("ent: missing required field \"is_del\"")}
	}
	if _, ok := sdic.mutation.Memo(); !ok {
		return &ValidationError{Name: "memo", err: errors.New("ent: missing required field \"memo\"")}
	}
	if v, ok := sdic.mutation.Memo(); ok {
		if err := sysdictitem.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf("ent: validator failed for field \"memo\": %w", err)}
		}
	}
	if _, ok := sdic.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New("ent: missing required field \"sort\"")}
	}
	if _, ok := sdic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := sdic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("ent: missing required field \"updated_at\"")}
	}
	if _, ok := sdic.mutation.Label(); !ok {
		return &ValidationError{Name: "label", err: errors.New("ent: missing required field \"label\"")}
	}
	if v, ok := sdic.mutation.Label(); ok {
		if err := sysdictitem.LabelValidator(v); err != nil {
			return &ValidationError{Name: "label", err: fmt.Errorf("ent: validator failed for field \"label\": %w", err)}
		}
	}
	if _, ok := sdic.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New("ent: missing required field \"value\"")}
	}
	if _, ok := sdic.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New("ent: missing required field \"status\"")}
	}
	if _, ok := sdic.mutation.DictID(); !ok {
		return &ValidationError{Name: "dict_id", err: errors.New("ent: missing required field \"dict_id\"")}
	}
	if v, ok := sdic.mutation.DictID(); ok {
		if err := sysdictitem.DictIDValidator(v); err != nil {
			return &ValidationError{Name: "dict_id", err: fmt.Errorf("ent: validator failed for field \"dict_id\": %w", err)}
		}
	}
	if v, ok := sdic.mutation.ID(); ok {
		if err := sysdictitem.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf("ent: validator failed for field \"id\": %w", err)}
		}
	}
	return nil
}

func (sdic *SysDictItemCreate) sqlSave(ctx context.Context) (*SysDictItem, error) {
	_node, _spec := sdic.createSpec()
	if err := sqlgraph.CreateNode(ctx, sdic.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (sdic *SysDictItemCreate) createSpec() (*SysDictItem, *sqlgraph.CreateSpec) {
	var (
		_node = &SysDictItem{config: sdic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: sysdictitem.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: sysdictitem.FieldID,
			},
		}
	)
	if id, ok := sdic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sdic.mutation.IsDel(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysdictitem.FieldIsDel,
		})
		_node.IsDel = value
	}
	if value, ok := sdic.mutation.Memo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictitem.FieldMemo,
		})
		_node.Memo = value
	}
	if value, ok := sdic.mutation.Sort(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysdictitem.FieldSort,
		})
		_node.Sort = value
	}
	if value, ok := sdic.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysdictitem.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := sdic.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysdictitem.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := sdic.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysdictitem.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := sdic.mutation.Label(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictitem.FieldLabel,
		})
		_node.Label = value
	}
	if value, ok := sdic.mutation.Value(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: sysdictitem.FieldValue,
		})
		_node.Value = value
	}
	if value, ok := sdic.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  value,
			Column: sysdictitem.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := sdic.mutation.DictID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictitem.FieldDictID,
		})
		_node.DictID = value
	}
	return _node, _spec
}

// SysDictItemCreateBulk is the builder for creating many SysDictItem entities in bulk.
type SysDictItemCreateBulk struct {
	config
	builders []*SysDictItemCreate
}

// Save creates the SysDictItem entities in the database.
func (sdicb *SysDictItemCreateBulk) Save(ctx context.Context) ([]*SysDictItem, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sdicb.builders))
	nodes := make([]*SysDictItem, len(sdicb.builders))
	mutators := make([]Mutator, len(sdicb.builders))
	for i := range sdicb.builders {
		func(i int, root context.Context) {
			builder := sdicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SysDictItemMutation)
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
					_, err = mutators[i+1].Mutate(root, sdicb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sdicb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, sdicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sdicb *SysDictItemCreateBulk) SaveX(ctx context.Context) []*SysDictItem {
	v, err := sdicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
