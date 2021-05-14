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
	"github.com/wanhello/omgind/pkg/helper/pulid"
)

// SysMenuActionCreate is the builder for creating a SysMenuAction entity.
type SysMenuActionCreate struct {
	config
	mutation *SysMenuActionMutation
	hooks    []Hook
}

// SetSort sets the "sort" field.
func (smac *SysMenuActionCreate) SetSort(i int32) *SysMenuActionCreate {
	smac.mutation.SetSort(i)
	return smac
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (smac *SysMenuActionCreate) SetNillableSort(i *int32) *SysMenuActionCreate {
	if i != nil {
		smac.SetSort(*i)
	}
	return smac
}

// SetStatus sets the "status" field.
func (smac *SysMenuActionCreate) SetStatus(i int32) *SysMenuActionCreate {
	smac.mutation.SetStatus(i)
	return smac
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (smac *SysMenuActionCreate) SetNillableStatus(i *int32) *SysMenuActionCreate {
	if i != nil {
		smac.SetStatus(*i)
	}
	return smac
}

// SetMemo sets the "memo" field.
func (smac *SysMenuActionCreate) SetMemo(s string) *SysMenuActionCreate {
	smac.mutation.SetMemo(s)
	return smac
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (smac *SysMenuActionCreate) SetNillableMemo(s *string) *SysMenuActionCreate {
	if s != nil {
		smac.SetMemo(*s)
	}
	return smac
}

// SetCreatedAt sets the "created_at" field.
func (smac *SysMenuActionCreate) SetCreatedAt(t time.Time) *SysMenuActionCreate {
	smac.mutation.SetCreatedAt(t)
	return smac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (smac *SysMenuActionCreate) SetNillableCreatedAt(t *time.Time) *SysMenuActionCreate {
	if t != nil {
		smac.SetCreatedAt(*t)
	}
	return smac
}

// SetUpdatedAt sets the "updated_at" field.
func (smac *SysMenuActionCreate) SetUpdatedAt(t time.Time) *SysMenuActionCreate {
	smac.mutation.SetUpdatedAt(t)
	return smac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (smac *SysMenuActionCreate) SetNillableUpdatedAt(t *time.Time) *SysMenuActionCreate {
	if t != nil {
		smac.SetUpdatedAt(*t)
	}
	return smac
}

// SetDeletedAt sets the "deleted_at" field.
func (smac *SysMenuActionCreate) SetDeletedAt(t time.Time) *SysMenuActionCreate {
	smac.mutation.SetDeletedAt(t)
	return smac
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (smac *SysMenuActionCreate) SetNillableDeletedAt(t *time.Time) *SysMenuActionCreate {
	if t != nil {
		smac.SetDeletedAt(*t)
	}
	return smac
}

// SetMenuID sets the "menu_id" field.
func (smac *SysMenuActionCreate) SetMenuID(s string) *SysMenuActionCreate {
	smac.mutation.SetMenuID(s)
	return smac
}

// SetCode sets the "code" field.
func (smac *SysMenuActionCreate) SetCode(s string) *SysMenuActionCreate {
	smac.mutation.SetCode(s)
	return smac
}

// SetName sets the "name" field.
func (smac *SysMenuActionCreate) SetName(s string) *SysMenuActionCreate {
	smac.mutation.SetName(s)
	return smac
}

// SetID sets the "id" field.
func (smac *SysMenuActionCreate) SetID(pu pulid.ID) *SysMenuActionCreate {
	smac.mutation.SetID(pu)
	return smac
}

// Mutation returns the SysMenuActionMutation object of the builder.
func (smac *SysMenuActionCreate) Mutation() *SysMenuActionMutation {
	return smac.mutation
}

// Save creates the SysMenuAction in the database.
func (smac *SysMenuActionCreate) Save(ctx context.Context) (*SysMenuAction, error) {
	var (
		err  error
		node *SysMenuAction
	)
	smac.defaults()
	if len(smac.hooks) == 0 {
		if err = smac.check(); err != nil {
			return nil, err
		}
		node, err = smac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysMenuActionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = smac.check(); err != nil {
				return nil, err
			}
			smac.mutation = mutation
			node, err = smac.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(smac.hooks) - 1; i >= 0; i-- {
			mut = smac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, smac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (smac *SysMenuActionCreate) SaveX(ctx context.Context) *SysMenuAction {
	v, err := smac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (smac *SysMenuActionCreate) defaults() {
	if _, ok := smac.mutation.Sort(); !ok {
		v := sysmenuaction.DefaultSort
		smac.mutation.SetSort(v)
	}
	if _, ok := smac.mutation.Status(); !ok {
		v := sysmenuaction.DefaultStatus
		smac.mutation.SetStatus(v)
	}
	if _, ok := smac.mutation.Memo(); !ok {
		v := sysmenuaction.DefaultMemo
		smac.mutation.SetMemo(v)
	}
	if _, ok := smac.mutation.CreatedAt(); !ok {
		v := sysmenuaction.DefaultCreatedAt()
		smac.mutation.SetCreatedAt(v)
	}
	if _, ok := smac.mutation.UpdatedAt(); !ok {
		v := sysmenuaction.DefaultUpdatedAt()
		smac.mutation.SetUpdatedAt(v)
	}
	if _, ok := smac.mutation.ID(); !ok {
		v := sysmenuaction.DefaultID()
		smac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (smac *SysMenuActionCreate) check() error {
	if _, ok := smac.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New("ent: missing required field \"sort\"")}
	}
	if _, ok := smac.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New("ent: missing required field \"status\"")}
	}
	if _, ok := smac.mutation.Memo(); !ok {
		return &ValidationError{Name: "memo", err: errors.New("ent: missing required field \"memo\"")}
	}
	if v, ok := smac.mutation.Memo(); ok {
		if err := sysmenuaction.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf("ent: validator failed for field \"memo\": %w", err)}
		}
	}
	if _, ok := smac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := smac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("ent: missing required field \"updated_at\"")}
	}
	if _, ok := smac.mutation.MenuID(); !ok {
		return &ValidationError{Name: "menu_id", err: errors.New("ent: missing required field \"menu_id\"")}
	}
	if v, ok := smac.mutation.MenuID(); ok {
		if err := sysmenuaction.MenuIDValidator(v); err != nil {
			return &ValidationError{Name: "menu_id", err: fmt.Errorf("ent: validator failed for field \"menu_id\": %w", err)}
		}
	}
	if _, ok := smac.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New("ent: missing required field \"code\"")}
	}
	if v, ok := smac.mutation.Code(); ok {
		if err := sysmenuaction.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf("ent: validator failed for field \"code\": %w", err)}
		}
	}
	if _, ok := smac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := smac.mutation.Name(); ok {
		if err := sysmenuaction.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	return nil
}

func (smac *SysMenuActionCreate) sqlSave(ctx context.Context) (*SysMenuAction, error) {
	_node, _spec := smac.createSpec()
	if err := sqlgraph.CreateNode(ctx, smac.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (smac *SysMenuActionCreate) createSpec() (*SysMenuAction, *sqlgraph.CreateSpec) {
	var (
		_node = &SysMenuAction{config: smac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: sysmenuaction.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: sysmenuaction.FieldID,
			},
		}
	)
	if id, ok := smac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := smac.mutation.Sort(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysmenuaction.FieldSort,
		})
		_node.Sort = value
	}
	if value, ok := smac.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysmenuaction.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := smac.mutation.Memo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenuaction.FieldMemo,
		})
		_node.Memo = value
	}
	if value, ok := smac.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysmenuaction.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := smac.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysmenuaction.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := smac.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysmenuaction.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := smac.mutation.MenuID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenuaction.FieldMenuID,
		})
		_node.MenuID = value
	}
	if value, ok := smac.mutation.Code(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenuaction.FieldCode,
		})
		_node.Code = value
	}
	if value, ok := smac.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenuaction.FieldName,
		})
		_node.Name = value
	}
	return _node, _spec
}

// SysMenuActionCreateBulk is the builder for creating many SysMenuAction entities in bulk.
type SysMenuActionCreateBulk struct {
	config
	builders []*SysMenuActionCreate
}

// Save creates the SysMenuAction entities in the database.
func (smacb *SysMenuActionCreateBulk) Save(ctx context.Context) ([]*SysMenuAction, error) {
	specs := make([]*sqlgraph.CreateSpec, len(smacb.builders))
	nodes := make([]*SysMenuAction, len(smacb.builders))
	mutators := make([]Mutator, len(smacb.builders))
	for i := range smacb.builders {
		func(i int, root context.Context) {
			builder := smacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SysMenuActionMutation)
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
					_, err = mutators[i+1].Mutate(root, smacb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, smacb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, smacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (smacb *SysMenuActionCreateBulk) SaveX(ctx context.Context) []*SysMenuAction {
	v, err := smacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
