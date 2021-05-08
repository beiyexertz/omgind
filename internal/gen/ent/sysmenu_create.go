// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wanhello/omgind/internal/gen/ent/sysmenu"
)

// SysMenuCreate is the builder for creating a SysMenu entity.
type SysMenuCreate struct {
	config
	mutation *SysMenuMutation
	hooks    []Hook
}

// SetIsDel sets the "is_del" field.
func (smc *SysMenuCreate) SetIsDel(b bool) *SysMenuCreate {
	smc.mutation.SetIsDel(b)
	return smc
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (smc *SysMenuCreate) SetNillableIsDel(b *bool) *SysMenuCreate {
	if b != nil {
		smc.SetIsDel(*b)
	}
	return smc
}

// SetMemo sets the "memo" field.
func (smc *SysMenuCreate) SetMemo(s string) *SysMenuCreate {
	smc.mutation.SetMemo(s)
	return smc
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (smc *SysMenuCreate) SetNillableMemo(s *string) *SysMenuCreate {
	if s != nil {
		smc.SetMemo(*s)
	}
	return smc
}

// SetSort sets the "sort" field.
func (smc *SysMenuCreate) SetSort(i int32) *SysMenuCreate {
	smc.mutation.SetSort(i)
	return smc
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (smc *SysMenuCreate) SetNillableSort(i *int32) *SysMenuCreate {
	if i != nil {
		smc.SetSort(*i)
	}
	return smc
}

// SetCreatedAt sets the "created_at" field.
func (smc *SysMenuCreate) SetCreatedAt(t time.Time) *SysMenuCreate {
	smc.mutation.SetCreatedAt(t)
	return smc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (smc *SysMenuCreate) SetNillableCreatedAt(t *time.Time) *SysMenuCreate {
	if t != nil {
		smc.SetCreatedAt(*t)
	}
	return smc
}

// SetUpdatedAt sets the "updated_at" field.
func (smc *SysMenuCreate) SetUpdatedAt(t time.Time) *SysMenuCreate {
	smc.mutation.SetUpdatedAt(t)
	return smc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (smc *SysMenuCreate) SetNillableUpdatedAt(t *time.Time) *SysMenuCreate {
	if t != nil {
		smc.SetUpdatedAt(*t)
	}
	return smc
}

// SetDeletedAt sets the "deleted_at" field.
func (smc *SysMenuCreate) SetDeletedAt(t time.Time) *SysMenuCreate {
	smc.mutation.SetDeletedAt(t)
	return smc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (smc *SysMenuCreate) SetNillableDeletedAt(t *time.Time) *SysMenuCreate {
	if t != nil {
		smc.SetDeletedAt(*t)
	}
	return smc
}

// SetStatus sets the "status" field.
func (smc *SysMenuCreate) SetStatus(i int32) *SysMenuCreate {
	smc.mutation.SetStatus(i)
	return smc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (smc *SysMenuCreate) SetNillableStatus(i *int32) *SysMenuCreate {
	if i != nil {
		smc.SetStatus(*i)
	}
	return smc
}

// SetName sets the "name" field.
func (smc *SysMenuCreate) SetName(s string) *SysMenuCreate {
	smc.mutation.SetName(s)
	return smc
}

// SetIcon sets the "icon" field.
func (smc *SysMenuCreate) SetIcon(s string) *SysMenuCreate {
	smc.mutation.SetIcon(s)
	return smc
}

// SetRouter sets the "router" field.
func (smc *SysMenuCreate) SetRouter(s string) *SysMenuCreate {
	smc.mutation.SetRouter(s)
	return smc
}

// SetIsShow sets the "is_show" field.
func (smc *SysMenuCreate) SetIsShow(b bool) *SysMenuCreate {
	smc.mutation.SetIsShow(b)
	return smc
}

// SetNillableIsShow sets the "is_show" field if the given value is not nil.
func (smc *SysMenuCreate) SetNillableIsShow(b *bool) *SysMenuCreate {
	if b != nil {
		smc.SetIsShow(*b)
	}
	return smc
}

// SetParentID sets the "parent_id" field.
func (smc *SysMenuCreate) SetParentID(s string) *SysMenuCreate {
	smc.mutation.SetParentID(s)
	return smc
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (smc *SysMenuCreate) SetNillableParentID(s *string) *SysMenuCreate {
	if s != nil {
		smc.SetParentID(*s)
	}
	return smc
}

// SetParentPath sets the "parent_path" field.
func (smc *SysMenuCreate) SetParentPath(s string) *SysMenuCreate {
	smc.mutation.SetParentPath(s)
	return smc
}

// SetNillableParentPath sets the "parent_path" field if the given value is not nil.
func (smc *SysMenuCreate) SetNillableParentPath(s *string) *SysMenuCreate {
	if s != nil {
		smc.SetParentPath(*s)
	}
	return smc
}

// SetID sets the "id" field.
func (smc *SysMenuCreate) SetID(s string) *SysMenuCreate {
	smc.mutation.SetID(s)
	return smc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (smc *SysMenuCreate) SetNillableID(s *string) *SysMenuCreate {
	if s != nil {
		smc.SetID(*s)
	}
	return smc
}

// Mutation returns the SysMenuMutation object of the builder.
func (smc *SysMenuCreate) Mutation() *SysMenuMutation {
	return smc.mutation
}

// Save creates the SysMenu in the database.
func (smc *SysMenuCreate) Save(ctx context.Context) (*SysMenu, error) {
	var (
		err  error
		node *SysMenu
	)
	smc.defaults()
	if len(smc.hooks) == 0 {
		if err = smc.check(); err != nil {
			return nil, err
		}
		node, err = smc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysMenuMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = smc.check(); err != nil {
				return nil, err
			}
			smc.mutation = mutation
			node, err = smc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(smc.hooks) - 1; i >= 0; i-- {
			mut = smc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, smc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (smc *SysMenuCreate) SaveX(ctx context.Context) *SysMenu {
	v, err := smc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (smc *SysMenuCreate) defaults() {
	if _, ok := smc.mutation.IsDel(); !ok {
		v := sysmenu.DefaultIsDel
		smc.mutation.SetIsDel(v)
	}
	if _, ok := smc.mutation.Memo(); !ok {
		v := sysmenu.DefaultMemo
		smc.mutation.SetMemo(v)
	}
	if _, ok := smc.mutation.Sort(); !ok {
		v := sysmenu.DefaultSort
		smc.mutation.SetSort(v)
	}
	if _, ok := smc.mutation.CreatedAt(); !ok {
		v := sysmenu.DefaultCreatedAt()
		smc.mutation.SetCreatedAt(v)
	}
	if _, ok := smc.mutation.UpdatedAt(); !ok {
		v := sysmenu.DefaultUpdatedAt()
		smc.mutation.SetUpdatedAt(v)
	}
	if _, ok := smc.mutation.Status(); !ok {
		v := sysmenu.DefaultStatus
		smc.mutation.SetStatus(v)
	}
	if _, ok := smc.mutation.IsShow(); !ok {
		v := sysmenu.DefaultIsShow
		smc.mutation.SetIsShow(v)
	}
	if _, ok := smc.mutation.ID(); !ok {
		v := sysmenu.DefaultID
		smc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (smc *SysMenuCreate) check() error {
	if _, ok := smc.mutation.IsDel(); !ok {
		return &ValidationError{Name: "is_del", err: errors.New("ent: missing required field \"is_del\"")}
	}
	if _, ok := smc.mutation.Memo(); !ok {
		return &ValidationError{Name: "memo", err: errors.New("ent: missing required field \"memo\"")}
	}
	if v, ok := smc.mutation.Memo(); ok {
		if err := sysmenu.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf("ent: validator failed for field \"memo\": %w", err)}
		}
	}
	if _, ok := smc.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New("ent: missing required field \"sort\"")}
	}
	if _, ok := smc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := smc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("ent: missing required field \"updated_at\"")}
	}
	if _, ok := smc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New("ent: missing required field \"status\"")}
	}
	if _, ok := smc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := smc.mutation.Name(); ok {
		if err := sysmenu.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := smc.mutation.Icon(); !ok {
		return &ValidationError{Name: "icon", err: errors.New("ent: missing required field \"icon\"")}
	}
	if v, ok := smc.mutation.Icon(); ok {
		if err := sysmenu.IconValidator(v); err != nil {
			return &ValidationError{Name: "icon", err: fmt.Errorf("ent: validator failed for field \"icon\": %w", err)}
		}
	}
	if _, ok := smc.mutation.Router(); !ok {
		return &ValidationError{Name: "router", err: errors.New("ent: missing required field \"router\"")}
	}
	if v, ok := smc.mutation.Router(); ok {
		if err := sysmenu.RouterValidator(v); err != nil {
			return &ValidationError{Name: "router", err: fmt.Errorf("ent: validator failed for field \"router\": %w", err)}
		}
	}
	if _, ok := smc.mutation.IsShow(); !ok {
		return &ValidationError{Name: "is_show", err: errors.New("ent: missing required field \"is_show\"")}
	}
	if v, ok := smc.mutation.ParentID(); ok {
		if err := sysmenu.ParentIDValidator(v); err != nil {
			return &ValidationError{Name: "parent_id", err: fmt.Errorf("ent: validator failed for field \"parent_id\": %w", err)}
		}
	}
	if v, ok := smc.mutation.ParentPath(); ok {
		if err := sysmenu.ParentPathValidator(v); err != nil {
			return &ValidationError{Name: "parent_path", err: fmt.Errorf("ent: validator failed for field \"parent_path\": %w", err)}
		}
	}
	if v, ok := smc.mutation.ID(); ok {
		if err := sysmenu.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf("ent: validator failed for field \"id\": %w", err)}
		}
	}
	return nil
}

func (smc *SysMenuCreate) sqlSave(ctx context.Context) (*SysMenu, error) {
	_node, _spec := smc.createSpec()
	if err := sqlgraph.CreateNode(ctx, smc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (smc *SysMenuCreate) createSpec() (*SysMenu, *sqlgraph.CreateSpec) {
	var (
		_node = &SysMenu{config: smc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: sysmenu.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: sysmenu.FieldID,
			},
		}
	)
	if id, ok := smc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := smc.mutation.IsDel(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysmenu.FieldIsDel,
		})
		_node.IsDel = value
	}
	if value, ok := smc.mutation.Memo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenu.FieldMemo,
		})
		_node.Memo = value
	}
	if value, ok := smc.mutation.Sort(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysmenu.FieldSort,
		})
		_node.Sort = value
	}
	if value, ok := smc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysmenu.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := smc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysmenu.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := smc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysmenu.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := smc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysmenu.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := smc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenu.FieldName,
		})
		_node.Name = value
	}
	if value, ok := smc.mutation.Icon(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenu.FieldIcon,
		})
		_node.Icon = value
	}
	if value, ok := smc.mutation.Router(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenu.FieldRouter,
		})
		_node.Router = value
	}
	if value, ok := smc.mutation.IsShow(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysmenu.FieldIsShow,
		})
		_node.IsShow = value
	}
	if value, ok := smc.mutation.ParentID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenu.FieldParentID,
		})
		_node.ParentID = &value
	}
	if value, ok := smc.mutation.ParentPath(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenu.FieldParentPath,
		})
		_node.ParentPath = &value
	}
	return _node, _spec
}

// SysMenuCreateBulk is the builder for creating many SysMenu entities in bulk.
type SysMenuCreateBulk struct {
	config
	builders []*SysMenuCreate
}

// Save creates the SysMenu entities in the database.
func (smcb *SysMenuCreateBulk) Save(ctx context.Context) ([]*SysMenu, error) {
	specs := make([]*sqlgraph.CreateSpec, len(smcb.builders))
	nodes := make([]*SysMenu, len(smcb.builders))
	mutators := make([]Mutator, len(smcb.builders))
	for i := range smcb.builders {
		func(i int, root context.Context) {
			builder := smcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SysMenuMutation)
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
					_, err = mutators[i+1].Mutate(root, smcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, smcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, smcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (smcb *SysMenuCreateBulk) SaveX(ctx context.Context) []*SysMenu {
	v, err := smcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}