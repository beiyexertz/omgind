// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wanhello/omgind/internal/gen/ent/predicate"
	"github.com/wanhello/omgind/internal/gen/ent/sysdict"
	"github.com/wanhello/omgind/internal/gen/ent/sysdictitem"
)

// SysDictItemUpdate is the builder for updating SysDictItem entities.
type SysDictItemUpdate struct {
	config
	hooks    []Hook
	mutation *SysDictItemMutation
}

// Where adds a new predicate for the SysDictItemUpdate builder.
func (sdiu *SysDictItemUpdate) Where(ps ...predicate.SysDictItem) *SysDictItemUpdate {
	sdiu.mutation.predicates = append(sdiu.mutation.predicates, ps...)
	return sdiu
}

// SetIsDel sets the "is_del" field.
func (sdiu *SysDictItemUpdate) SetIsDel(b bool) *SysDictItemUpdate {
	sdiu.mutation.SetIsDel(b)
	return sdiu
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (sdiu *SysDictItemUpdate) SetNillableIsDel(b *bool) *SysDictItemUpdate {
	if b != nil {
		sdiu.SetIsDel(*b)
	}
	return sdiu
}

// SetMemo sets the "memo" field.
func (sdiu *SysDictItemUpdate) SetMemo(s string) *SysDictItemUpdate {
	sdiu.mutation.SetMemo(s)
	return sdiu
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (sdiu *SysDictItemUpdate) SetNillableMemo(s *string) *SysDictItemUpdate {
	if s != nil {
		sdiu.SetMemo(*s)
	}
	return sdiu
}

// SetSort sets the "sort" field.
func (sdiu *SysDictItemUpdate) SetSort(i int32) *SysDictItemUpdate {
	sdiu.mutation.ResetSort()
	sdiu.mutation.SetSort(i)
	return sdiu
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (sdiu *SysDictItemUpdate) SetNillableSort(i *int32) *SysDictItemUpdate {
	if i != nil {
		sdiu.SetSort(*i)
	}
	return sdiu
}

// AddSort adds i to the "sort" field.
func (sdiu *SysDictItemUpdate) AddSort(i int32) *SysDictItemUpdate {
	sdiu.mutation.AddSort(i)
	return sdiu
}

// SetUpdatedAt sets the "updated_at" field.
func (sdiu *SysDictItemUpdate) SetUpdatedAt(t time.Time) *SysDictItemUpdate {
	sdiu.mutation.SetUpdatedAt(t)
	return sdiu
}

// SetDeletedAt sets the "deleted_at" field.
func (sdiu *SysDictItemUpdate) SetDeletedAt(t time.Time) *SysDictItemUpdate {
	sdiu.mutation.SetDeletedAt(t)
	return sdiu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sdiu *SysDictItemUpdate) SetNillableDeletedAt(t *time.Time) *SysDictItemUpdate {
	if t != nil {
		sdiu.SetDeletedAt(*t)
	}
	return sdiu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (sdiu *SysDictItemUpdate) ClearDeletedAt() *SysDictItemUpdate {
	sdiu.mutation.ClearDeletedAt()
	return sdiu
}

// SetLabel sets the "label" field.
func (sdiu *SysDictItemUpdate) SetLabel(s string) *SysDictItemUpdate {
	sdiu.mutation.SetLabel(s)
	return sdiu
}

// SetValue sets the "value" field.
func (sdiu *SysDictItemUpdate) SetValue(i int) *SysDictItemUpdate {
	sdiu.mutation.ResetValue()
	sdiu.mutation.SetValue(i)
	return sdiu
}

// AddValue adds i to the "value" field.
func (sdiu *SysDictItemUpdate) AddValue(i int) *SysDictItemUpdate {
	sdiu.mutation.AddValue(i)
	return sdiu
}

// SetStatus sets the "status" field.
func (sdiu *SysDictItemUpdate) SetStatus(b bool) *SysDictItemUpdate {
	sdiu.mutation.SetStatus(b)
	return sdiu
}

// SetSysDictID sets the "SysDict" edge to the SysDict entity by ID.
func (sdiu *SysDictItemUpdate) SetSysDictID(id string) *SysDictItemUpdate {
	sdiu.mutation.SetSysDictID(id)
	return sdiu
}

// SetNillableSysDictID sets the "SysDict" edge to the SysDict entity by ID if the given value is not nil.
func (sdiu *SysDictItemUpdate) SetNillableSysDictID(id *string) *SysDictItemUpdate {
	if id != nil {
		sdiu = sdiu.SetSysDictID(*id)
	}
	return sdiu
}

// SetSysDict sets the "SysDict" edge to the SysDict entity.
func (sdiu *SysDictItemUpdate) SetSysDict(s *SysDict) *SysDictItemUpdate {
	return sdiu.SetSysDictID(s.ID)
}

// Mutation returns the SysDictItemMutation object of the builder.
func (sdiu *SysDictItemUpdate) Mutation() *SysDictItemMutation {
	return sdiu.mutation
}

// ClearSysDict clears the "SysDict" edge to the SysDict entity.
func (sdiu *SysDictItemUpdate) ClearSysDict() *SysDictItemUpdate {
	sdiu.mutation.ClearSysDict()
	return sdiu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sdiu *SysDictItemUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	sdiu.defaults()
	if len(sdiu.hooks) == 0 {
		if err = sdiu.check(); err != nil {
			return 0, err
		}
		affected, err = sdiu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysDictItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sdiu.check(); err != nil {
				return 0, err
			}
			sdiu.mutation = mutation
			affected, err = sdiu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sdiu.hooks) - 1; i >= 0; i-- {
			mut = sdiu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sdiu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (sdiu *SysDictItemUpdate) SaveX(ctx context.Context) int {
	affected, err := sdiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sdiu *SysDictItemUpdate) Exec(ctx context.Context) error {
	_, err := sdiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdiu *SysDictItemUpdate) ExecX(ctx context.Context) {
	if err := sdiu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sdiu *SysDictItemUpdate) defaults() {
	if _, ok := sdiu.mutation.UpdatedAt(); !ok {
		v := sysdictitem.UpdateDefaultUpdatedAt()
		sdiu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sdiu *SysDictItemUpdate) check() error {
	if v, ok := sdiu.mutation.Memo(); ok {
		if err := sysdictitem.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf("ent: validator failed for field \"memo\": %w", err)}
		}
	}
	if v, ok := sdiu.mutation.Label(); ok {
		if err := sysdictitem.LabelValidator(v); err != nil {
			return &ValidationError{Name: "label", err: fmt.Errorf("ent: validator failed for field \"label\": %w", err)}
		}
	}
	return nil
}

func (sdiu *SysDictItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysdictitem.Table,
			Columns: sysdictitem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: sysdictitem.FieldID,
			},
		},
	}
	if ps := sdiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sdiu.mutation.IsDel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysdictitem.FieldIsDel,
		})
	}
	if value, ok := sdiu.mutation.Memo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictitem.FieldMemo,
		})
	}
	if value, ok := sdiu.mutation.Sort(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysdictitem.FieldSort,
		})
	}
	if value, ok := sdiu.mutation.AddedSort(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysdictitem.FieldSort,
		})
	}
	if value, ok := sdiu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysdictitem.FieldUpdatedAt,
		})
	}
	if value, ok := sdiu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysdictitem.FieldDeletedAt,
		})
	}
	if sdiu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: sysdictitem.FieldDeletedAt,
		})
	}
	if value, ok := sdiu.mutation.Label(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictitem.FieldLabel,
		})
	}
	if value, ok := sdiu.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: sysdictitem.FieldValue,
		})
	}
	if value, ok := sdiu.mutation.AddedValue(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: sysdictitem.FieldValue,
		})
	}
	if value, ok := sdiu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysdictitem.FieldStatus,
		})
	}
	if sdiu.mutation.SysDictCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sysdictitem.SysDictTable,
			Columns: []string{sysdictitem.SysDictColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: sysdict.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sdiu.mutation.SysDictIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sysdictitem.SysDictTable,
			Columns: []string{sysdictitem.SysDictColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: sysdict.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, sdiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysdictitem.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// SysDictItemUpdateOne is the builder for updating a single SysDictItem entity.
type SysDictItemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SysDictItemMutation
}

// SetIsDel sets the "is_del" field.
func (sdiuo *SysDictItemUpdateOne) SetIsDel(b bool) *SysDictItemUpdateOne {
	sdiuo.mutation.SetIsDel(b)
	return sdiuo
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (sdiuo *SysDictItemUpdateOne) SetNillableIsDel(b *bool) *SysDictItemUpdateOne {
	if b != nil {
		sdiuo.SetIsDel(*b)
	}
	return sdiuo
}

// SetMemo sets the "memo" field.
func (sdiuo *SysDictItemUpdateOne) SetMemo(s string) *SysDictItemUpdateOne {
	sdiuo.mutation.SetMemo(s)
	return sdiuo
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (sdiuo *SysDictItemUpdateOne) SetNillableMemo(s *string) *SysDictItemUpdateOne {
	if s != nil {
		sdiuo.SetMemo(*s)
	}
	return sdiuo
}

// SetSort sets the "sort" field.
func (sdiuo *SysDictItemUpdateOne) SetSort(i int32) *SysDictItemUpdateOne {
	sdiuo.mutation.ResetSort()
	sdiuo.mutation.SetSort(i)
	return sdiuo
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (sdiuo *SysDictItemUpdateOne) SetNillableSort(i *int32) *SysDictItemUpdateOne {
	if i != nil {
		sdiuo.SetSort(*i)
	}
	return sdiuo
}

// AddSort adds i to the "sort" field.
func (sdiuo *SysDictItemUpdateOne) AddSort(i int32) *SysDictItemUpdateOne {
	sdiuo.mutation.AddSort(i)
	return sdiuo
}

// SetUpdatedAt sets the "updated_at" field.
func (sdiuo *SysDictItemUpdateOne) SetUpdatedAt(t time.Time) *SysDictItemUpdateOne {
	sdiuo.mutation.SetUpdatedAt(t)
	return sdiuo
}

// SetDeletedAt sets the "deleted_at" field.
func (sdiuo *SysDictItemUpdateOne) SetDeletedAt(t time.Time) *SysDictItemUpdateOne {
	sdiuo.mutation.SetDeletedAt(t)
	return sdiuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sdiuo *SysDictItemUpdateOne) SetNillableDeletedAt(t *time.Time) *SysDictItemUpdateOne {
	if t != nil {
		sdiuo.SetDeletedAt(*t)
	}
	return sdiuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (sdiuo *SysDictItemUpdateOne) ClearDeletedAt() *SysDictItemUpdateOne {
	sdiuo.mutation.ClearDeletedAt()
	return sdiuo
}

// SetLabel sets the "label" field.
func (sdiuo *SysDictItemUpdateOne) SetLabel(s string) *SysDictItemUpdateOne {
	sdiuo.mutation.SetLabel(s)
	return sdiuo
}

// SetValue sets the "value" field.
func (sdiuo *SysDictItemUpdateOne) SetValue(i int) *SysDictItemUpdateOne {
	sdiuo.mutation.ResetValue()
	sdiuo.mutation.SetValue(i)
	return sdiuo
}

// AddValue adds i to the "value" field.
func (sdiuo *SysDictItemUpdateOne) AddValue(i int) *SysDictItemUpdateOne {
	sdiuo.mutation.AddValue(i)
	return sdiuo
}

// SetStatus sets the "status" field.
func (sdiuo *SysDictItemUpdateOne) SetStatus(b bool) *SysDictItemUpdateOne {
	sdiuo.mutation.SetStatus(b)
	return sdiuo
}

// SetSysDictID sets the "SysDict" edge to the SysDict entity by ID.
func (sdiuo *SysDictItemUpdateOne) SetSysDictID(id string) *SysDictItemUpdateOne {
	sdiuo.mutation.SetSysDictID(id)
	return sdiuo
}

// SetNillableSysDictID sets the "SysDict" edge to the SysDict entity by ID if the given value is not nil.
func (sdiuo *SysDictItemUpdateOne) SetNillableSysDictID(id *string) *SysDictItemUpdateOne {
	if id != nil {
		sdiuo = sdiuo.SetSysDictID(*id)
	}
	return sdiuo
}

// SetSysDict sets the "SysDict" edge to the SysDict entity.
func (sdiuo *SysDictItemUpdateOne) SetSysDict(s *SysDict) *SysDictItemUpdateOne {
	return sdiuo.SetSysDictID(s.ID)
}

// Mutation returns the SysDictItemMutation object of the builder.
func (sdiuo *SysDictItemUpdateOne) Mutation() *SysDictItemMutation {
	return sdiuo.mutation
}

// ClearSysDict clears the "SysDict" edge to the SysDict entity.
func (sdiuo *SysDictItemUpdateOne) ClearSysDict() *SysDictItemUpdateOne {
	sdiuo.mutation.ClearSysDict()
	return sdiuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sdiuo *SysDictItemUpdateOne) Select(field string, fields ...string) *SysDictItemUpdateOne {
	sdiuo.fields = append([]string{field}, fields...)
	return sdiuo
}

// Save executes the query and returns the updated SysDictItem entity.
func (sdiuo *SysDictItemUpdateOne) Save(ctx context.Context) (*SysDictItem, error) {
	var (
		err  error
		node *SysDictItem
	)
	sdiuo.defaults()
	if len(sdiuo.hooks) == 0 {
		if err = sdiuo.check(); err != nil {
			return nil, err
		}
		node, err = sdiuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysDictItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sdiuo.check(); err != nil {
				return nil, err
			}
			sdiuo.mutation = mutation
			node, err = sdiuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sdiuo.hooks) - 1; i >= 0; i-- {
			mut = sdiuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sdiuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (sdiuo *SysDictItemUpdateOne) SaveX(ctx context.Context) *SysDictItem {
	node, err := sdiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sdiuo *SysDictItemUpdateOne) Exec(ctx context.Context) error {
	_, err := sdiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdiuo *SysDictItemUpdateOne) ExecX(ctx context.Context) {
	if err := sdiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sdiuo *SysDictItemUpdateOne) defaults() {
	if _, ok := sdiuo.mutation.UpdatedAt(); !ok {
		v := sysdictitem.UpdateDefaultUpdatedAt()
		sdiuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sdiuo *SysDictItemUpdateOne) check() error {
	if v, ok := sdiuo.mutation.Memo(); ok {
		if err := sysdictitem.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf("ent: validator failed for field \"memo\": %w", err)}
		}
	}
	if v, ok := sdiuo.mutation.Label(); ok {
		if err := sysdictitem.LabelValidator(v); err != nil {
			return &ValidationError{Name: "label", err: fmt.Errorf("ent: validator failed for field \"label\": %w", err)}
		}
	}
	return nil
}

func (sdiuo *SysDictItemUpdateOne) sqlSave(ctx context.Context) (_node *SysDictItem, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysdictitem.Table,
			Columns: sysdictitem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: sysdictitem.FieldID,
			},
		},
	}
	id, ok := sdiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing SysDictItem.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := sdiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysdictitem.FieldID)
		for _, f := range fields {
			if !sysdictitem.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != sysdictitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sdiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sdiuo.mutation.IsDel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysdictitem.FieldIsDel,
		})
	}
	if value, ok := sdiuo.mutation.Memo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictitem.FieldMemo,
		})
	}
	if value, ok := sdiuo.mutation.Sort(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysdictitem.FieldSort,
		})
	}
	if value, ok := sdiuo.mutation.AddedSort(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysdictitem.FieldSort,
		})
	}
	if value, ok := sdiuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysdictitem.FieldUpdatedAt,
		})
	}
	if value, ok := sdiuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysdictitem.FieldDeletedAt,
		})
	}
	if sdiuo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: sysdictitem.FieldDeletedAt,
		})
	}
	if value, ok := sdiuo.mutation.Label(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictitem.FieldLabel,
		})
	}
	if value, ok := sdiuo.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: sysdictitem.FieldValue,
		})
	}
	if value, ok := sdiuo.mutation.AddedValue(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: sysdictitem.FieldValue,
		})
	}
	if value, ok := sdiuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysdictitem.FieldStatus,
		})
	}
	if sdiuo.mutation.SysDictCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sysdictitem.SysDictTable,
			Columns: []string{sysdictitem.SysDictColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: sysdict.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sdiuo.mutation.SysDictIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sysdictitem.SysDictTable,
			Columns: []string{sysdictitem.SysDictColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: sysdict.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &SysDictItem{config: sdiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sdiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysdictitem.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
