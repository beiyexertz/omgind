// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wanhello/omgind/internal/gen/ent/sysdict"
)

// SysDictCreate is the builder for creating a SysDict entity.
type SysDictCreate struct {
	config
	mutation *SysDictMutation
	hooks    []Hook
}

// SetIsDel sets the "is_del" field.
func (sdc *SysDictCreate) SetIsDel(b bool) *SysDictCreate {
	sdc.mutation.SetIsDel(b)
	return sdc
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (sdc *SysDictCreate) SetNillableIsDel(b *bool) *SysDictCreate {
	if b != nil {
		sdc.SetIsDel(*b)
	}
	return sdc
}

// SetMemo sets the "memo" field.
func (sdc *SysDictCreate) SetMemo(s string) *SysDictCreate {
	sdc.mutation.SetMemo(s)
	return sdc
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (sdc *SysDictCreate) SetNillableMemo(s *string) *SysDictCreate {
	if s != nil {
		sdc.SetMemo(*s)
	}
	return sdc
}

// SetSort sets the "sort" field.
func (sdc *SysDictCreate) SetSort(i int32) *SysDictCreate {
	sdc.mutation.SetSort(i)
	return sdc
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (sdc *SysDictCreate) SetNillableSort(i *int32) *SysDictCreate {
	if i != nil {
		sdc.SetSort(*i)
	}
	return sdc
}

// SetCreatedAt sets the "created_at" field.
func (sdc *SysDictCreate) SetCreatedAt(t time.Time) *SysDictCreate {
	sdc.mutation.SetCreatedAt(t)
	return sdc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sdc *SysDictCreate) SetNillableCreatedAt(t *time.Time) *SysDictCreate {
	if t != nil {
		sdc.SetCreatedAt(*t)
	}
	return sdc
}

// SetUpdatedAt sets the "updated_at" field.
func (sdc *SysDictCreate) SetUpdatedAt(t time.Time) *SysDictCreate {
	sdc.mutation.SetUpdatedAt(t)
	return sdc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sdc *SysDictCreate) SetNillableUpdatedAt(t *time.Time) *SysDictCreate {
	if t != nil {
		sdc.SetUpdatedAt(*t)
	}
	return sdc
}

// SetDeletedAt sets the "deleted_at" field.
func (sdc *SysDictCreate) SetDeletedAt(t time.Time) *SysDictCreate {
	sdc.mutation.SetDeletedAt(t)
	return sdc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sdc *SysDictCreate) SetNillableDeletedAt(t *time.Time) *SysDictCreate {
	if t != nil {
		sdc.SetDeletedAt(*t)
	}
	return sdc
}

// SetStatus sets the "status" field.
func (sdc *SysDictCreate) SetStatus(i int16) *SysDictCreate {
	sdc.mutation.SetStatus(i)
	return sdc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sdc *SysDictCreate) SetNillableStatus(i *int16) *SysDictCreate {
	if i != nil {
		sdc.SetStatus(*i)
	}
	return sdc
}

// SetNameCn sets the "name_cn" field.
func (sdc *SysDictCreate) SetNameCn(s string) *SysDictCreate {
	sdc.mutation.SetNameCn(s)
	return sdc
}

// SetNameEn sets the "name_en" field.
func (sdc *SysDictCreate) SetNameEn(s string) *SysDictCreate {
	sdc.mutation.SetNameEn(s)
	return sdc
}

// SetID sets the "id" field.
func (sdc *SysDictCreate) SetID(s string) *SysDictCreate {
	sdc.mutation.SetID(s)
	return sdc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sdc *SysDictCreate) SetNillableID(s *string) *SysDictCreate {
	if s != nil {
		sdc.SetID(*s)
	}
	return sdc
}

// Mutation returns the SysDictMutation object of the builder.
func (sdc *SysDictCreate) Mutation() *SysDictMutation {
	return sdc.mutation
}

// Save creates the SysDict in the database.
func (sdc *SysDictCreate) Save(ctx context.Context) (*SysDict, error) {
	var (
		err  error
		node *SysDict
	)
	sdc.defaults()
	if len(sdc.hooks) == 0 {
		if err = sdc.check(); err != nil {
			return nil, err
		}
		node, err = sdc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysDictMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sdc.check(); err != nil {
				return nil, err
			}
			sdc.mutation = mutation
			node, err = sdc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sdc.hooks) - 1; i >= 0; i-- {
			mut = sdc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sdc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sdc *SysDictCreate) SaveX(ctx context.Context) *SysDict {
	v, err := sdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (sdc *SysDictCreate) defaults() {
	if _, ok := sdc.mutation.IsDel(); !ok {
		v := sysdict.DefaultIsDel
		sdc.mutation.SetIsDel(v)
	}
	if _, ok := sdc.mutation.Memo(); !ok {
		v := sysdict.DefaultMemo
		sdc.mutation.SetMemo(v)
	}
	if _, ok := sdc.mutation.Sort(); !ok {
		v := sysdict.DefaultSort
		sdc.mutation.SetSort(v)
	}
	if _, ok := sdc.mutation.CreatedAt(); !ok {
		v := sysdict.DefaultCreatedAt()
		sdc.mutation.SetCreatedAt(v)
	}
	if _, ok := sdc.mutation.UpdatedAt(); !ok {
		v := sysdict.DefaultUpdatedAt()
		sdc.mutation.SetUpdatedAt(v)
	}
	if _, ok := sdc.mutation.Status(); !ok {
		v := sysdict.DefaultStatus
		sdc.mutation.SetStatus(v)
	}
	if _, ok := sdc.mutation.ID(); !ok {
		v := sysdict.DefaultID()
		sdc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sdc *SysDictCreate) check() error {
	if _, ok := sdc.mutation.IsDel(); !ok {
		return &ValidationError{Name: "is_del", err: errors.New("ent: missing required field \"is_del\"")}
	}
	if _, ok := sdc.mutation.Memo(); !ok {
		return &ValidationError{Name: "memo", err: errors.New("ent: missing required field \"memo\"")}
	}
	if v, ok := sdc.mutation.Memo(); ok {
		if err := sysdict.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf("ent: validator failed for field \"memo\": %w", err)}
		}
	}
	if _, ok := sdc.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New("ent: missing required field \"sort\"")}
	}
	if _, ok := sdc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := sdc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("ent: missing required field \"updated_at\"")}
	}
	if _, ok := sdc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New("ent: missing required field \"status\"")}
	}
	if _, ok := sdc.mutation.NameCn(); !ok {
		return &ValidationError{Name: "name_cn", err: errors.New("ent: missing required field \"name_cn\"")}
	}
	if v, ok := sdc.mutation.NameCn(); ok {
		if err := sysdict.NameCnValidator(v); err != nil {
			return &ValidationError{Name: "name_cn", err: fmt.Errorf("ent: validator failed for field \"name_cn\": %w", err)}
		}
	}
	if _, ok := sdc.mutation.NameEn(); !ok {
		return &ValidationError{Name: "name_en", err: errors.New("ent: missing required field \"name_en\"")}
	}
	if v, ok := sdc.mutation.NameEn(); ok {
		if err := sysdict.NameEnValidator(v); err != nil {
			return &ValidationError{Name: "name_en", err: fmt.Errorf("ent: validator failed for field \"name_en\": %w", err)}
		}
	}
	if v, ok := sdc.mutation.ID(); ok {
		if err := sysdict.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf("ent: validator failed for field \"id\": %w", err)}
		}
	}
	return nil
}

func (sdc *SysDictCreate) sqlSave(ctx context.Context) (*SysDict, error) {
	_node, _spec := sdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sdc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (sdc *SysDictCreate) createSpec() (*SysDict, *sqlgraph.CreateSpec) {
	var (
		_node = &SysDict{config: sdc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: sysdict.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: sysdict.FieldID,
			},
		}
	)
	if id, ok := sdc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sdc.mutation.IsDel(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysdict.FieldIsDel,
		})
		_node.IsDel = value
	}
	if value, ok := sdc.mutation.Memo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdict.FieldMemo,
		})
		_node.Memo = value
	}
	if value, ok := sdc.mutation.Sort(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysdict.FieldSort,
		})
		_node.Sort = value
	}
	if value, ok := sdc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysdict.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := sdc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysdict.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := sdc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysdict.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := sdc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  value,
			Column: sysdict.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := sdc.mutation.NameCn(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdict.FieldNameCn,
		})
		_node.NameCn = value
	}
	if value, ok := sdc.mutation.NameEn(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdict.FieldNameEn,
		})
		_node.NameEn = value
	}
	return _node, _spec
}

// SysDictCreateBulk is the builder for creating many SysDict entities in bulk.
type SysDictCreateBulk struct {
	config
	builders []*SysDictCreate
}

// Save creates the SysDict entities in the database.
func (sdcb *SysDictCreateBulk) Save(ctx context.Context) ([]*SysDict, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sdcb.builders))
	nodes := make([]*SysDict, len(sdcb.builders))
	mutators := make([]Mutator, len(sdcb.builders))
	for i := range sdcb.builders {
		func(i int, root context.Context) {
			builder := sdcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SysDictMutation)
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
					_, err = mutators[i+1].Mutate(root, sdcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sdcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, sdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sdcb *SysDictCreateBulk) SaveX(ctx context.Context) []*SysDict {
	v, err := sdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
