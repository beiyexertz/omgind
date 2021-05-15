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
)

// SysDictUpdate is the builder for updating SysDict entities.
type SysDictUpdate struct {
	config
	hooks    []Hook
	mutation *SysDictMutation
}

// Where adds a new predicate for the SysDictUpdate builder.
func (sdu *SysDictUpdate) Where(ps ...predicate.SysDict) *SysDictUpdate {
	sdu.mutation.predicates = append(sdu.mutation.predicates, ps...)
	return sdu
}

// SetMemo sets the "memo" field.
func (sdu *SysDictUpdate) SetMemo(s string) *SysDictUpdate {
	sdu.mutation.SetMemo(s)
	return sdu
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (sdu *SysDictUpdate) SetNillableMemo(s *string) *SysDictUpdate {
	if s != nil {
		sdu.SetMemo(*s)
	}
	return sdu
}

// SetSort sets the "sort" field.
func (sdu *SysDictUpdate) SetSort(i int32) *SysDictUpdate {
	sdu.mutation.ResetSort()
	sdu.mutation.SetSort(i)
	return sdu
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (sdu *SysDictUpdate) SetNillableSort(i *int32) *SysDictUpdate {
	if i != nil {
		sdu.SetSort(*i)
	}
	return sdu
}

// AddSort adds i to the "sort" field.
func (sdu *SysDictUpdate) AddSort(i int32) *SysDictUpdate {
	sdu.mutation.AddSort(i)
	return sdu
}

// SetUpdatedAt sets the "updated_at" field.
func (sdu *SysDictUpdate) SetUpdatedAt(t time.Time) *SysDictUpdate {
	sdu.mutation.SetUpdatedAt(t)
	return sdu
}

// SetDeletedAt sets the "deleted_at" field.
func (sdu *SysDictUpdate) SetDeletedAt(t time.Time) *SysDictUpdate {
	sdu.mutation.SetDeletedAt(t)
	return sdu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sdu *SysDictUpdate) SetNillableDeletedAt(t *time.Time) *SysDictUpdate {
	if t != nil {
		sdu.SetDeletedAt(*t)
	}
	return sdu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (sdu *SysDictUpdate) ClearDeletedAt() *SysDictUpdate {
	sdu.mutation.ClearDeletedAt()
	return sdu
}

// SetNameCn sets the "name_cn" field.
func (sdu *SysDictUpdate) SetNameCn(s string) *SysDictUpdate {
	sdu.mutation.SetNameCn(s)
	return sdu
}

// SetNameEn sets the "name_en" field.
func (sdu *SysDictUpdate) SetNameEn(s string) *SysDictUpdate {
	sdu.mutation.SetNameEn(s)
	return sdu
}

// SetStatus sets the "status" field.
func (sdu *SysDictUpdate) SetStatus(b bool) *SysDictUpdate {
	sdu.mutation.SetStatus(b)
	return sdu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sdu *SysDictUpdate) SetNillableStatus(b *bool) *SysDictUpdate {
	if b != nil {
		sdu.SetStatus(*b)
	}
	return sdu
}

// Mutation returns the SysDictMutation object of the builder.
func (sdu *SysDictUpdate) Mutation() *SysDictMutation {
	return sdu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sdu *SysDictUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	sdu.defaults()
	if len(sdu.hooks) == 0 {
		if err = sdu.check(); err != nil {
			return 0, err
		}
		affected, err = sdu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysDictMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sdu.check(); err != nil {
				return 0, err
			}
			sdu.mutation = mutation
			affected, err = sdu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sdu.hooks) - 1; i >= 0; i-- {
			mut = sdu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sdu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (sdu *SysDictUpdate) SaveX(ctx context.Context) int {
	affected, err := sdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sdu *SysDictUpdate) Exec(ctx context.Context) error {
	_, err := sdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdu *SysDictUpdate) ExecX(ctx context.Context) {
	if err := sdu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sdu *SysDictUpdate) defaults() {
	if _, ok := sdu.mutation.UpdatedAt(); !ok {
		v := sysdict.UpdateDefaultUpdatedAt()
		sdu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sdu *SysDictUpdate) check() error {
	if v, ok := sdu.mutation.Memo(); ok {
		if err := sysdict.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf("ent: validator failed for field \"memo\": %w", err)}
		}
	}
	if v, ok := sdu.mutation.NameCn(); ok {
		if err := sysdict.NameCnValidator(v); err != nil {
			return &ValidationError{Name: "name_cn", err: fmt.Errorf("ent: validator failed for field \"name_cn\": %w", err)}
		}
	}
	if v, ok := sdu.mutation.NameEn(); ok {
		if err := sysdict.NameEnValidator(v); err != nil {
			return &ValidationError{Name: "name_en", err: fmt.Errorf("ent: validator failed for field \"name_en\": %w", err)}
		}
	}
	return nil
}

func (sdu *SysDictUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysdict.Table,
			Columns: sysdict.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: sysdict.FieldID,
			},
		},
	}
	if ps := sdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sdu.mutation.Memo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdict.FieldMemo,
		})
	}
	if value, ok := sdu.mutation.Sort(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysdict.FieldSort,
		})
	}
	if value, ok := sdu.mutation.AddedSort(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysdict.FieldSort,
		})
	}
	if value, ok := sdu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysdict.FieldUpdatedAt,
		})
	}
	if value, ok := sdu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysdict.FieldDeletedAt,
		})
	}
	if sdu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: sysdict.FieldDeletedAt,
		})
	}
	if value, ok := sdu.mutation.NameCn(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdict.FieldNameCn,
		})
	}
	if value, ok := sdu.mutation.NameEn(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdict.FieldNameEn,
		})
	}
	if value, ok := sdu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysdict.FieldStatus,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, sdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysdict.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// SysDictUpdateOne is the builder for updating a single SysDict entity.
type SysDictUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SysDictMutation
}

// SetMemo sets the "memo" field.
func (sduo *SysDictUpdateOne) SetMemo(s string) *SysDictUpdateOne {
	sduo.mutation.SetMemo(s)
	return sduo
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (sduo *SysDictUpdateOne) SetNillableMemo(s *string) *SysDictUpdateOne {
	if s != nil {
		sduo.SetMemo(*s)
	}
	return sduo
}

// SetSort sets the "sort" field.
func (sduo *SysDictUpdateOne) SetSort(i int32) *SysDictUpdateOne {
	sduo.mutation.ResetSort()
	sduo.mutation.SetSort(i)
	return sduo
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (sduo *SysDictUpdateOne) SetNillableSort(i *int32) *SysDictUpdateOne {
	if i != nil {
		sduo.SetSort(*i)
	}
	return sduo
}

// AddSort adds i to the "sort" field.
func (sduo *SysDictUpdateOne) AddSort(i int32) *SysDictUpdateOne {
	sduo.mutation.AddSort(i)
	return sduo
}

// SetUpdatedAt sets the "updated_at" field.
func (sduo *SysDictUpdateOne) SetUpdatedAt(t time.Time) *SysDictUpdateOne {
	sduo.mutation.SetUpdatedAt(t)
	return sduo
}

// SetDeletedAt sets the "deleted_at" field.
func (sduo *SysDictUpdateOne) SetDeletedAt(t time.Time) *SysDictUpdateOne {
	sduo.mutation.SetDeletedAt(t)
	return sduo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sduo *SysDictUpdateOne) SetNillableDeletedAt(t *time.Time) *SysDictUpdateOne {
	if t != nil {
		sduo.SetDeletedAt(*t)
	}
	return sduo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (sduo *SysDictUpdateOne) ClearDeletedAt() *SysDictUpdateOne {
	sduo.mutation.ClearDeletedAt()
	return sduo
}

// SetNameCn sets the "name_cn" field.
func (sduo *SysDictUpdateOne) SetNameCn(s string) *SysDictUpdateOne {
	sduo.mutation.SetNameCn(s)
	return sduo
}

// SetNameEn sets the "name_en" field.
func (sduo *SysDictUpdateOne) SetNameEn(s string) *SysDictUpdateOne {
	sduo.mutation.SetNameEn(s)
	return sduo
}

// SetStatus sets the "status" field.
func (sduo *SysDictUpdateOne) SetStatus(b bool) *SysDictUpdateOne {
	sduo.mutation.SetStatus(b)
	return sduo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sduo *SysDictUpdateOne) SetNillableStatus(b *bool) *SysDictUpdateOne {
	if b != nil {
		sduo.SetStatus(*b)
	}
	return sduo
}

// Mutation returns the SysDictMutation object of the builder.
func (sduo *SysDictUpdateOne) Mutation() *SysDictMutation {
	return sduo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sduo *SysDictUpdateOne) Select(field string, fields ...string) *SysDictUpdateOne {
	sduo.fields = append([]string{field}, fields...)
	return sduo
}

// Save executes the query and returns the updated SysDict entity.
func (sduo *SysDictUpdateOne) Save(ctx context.Context) (*SysDict, error) {
	var (
		err  error
		node *SysDict
	)
	sduo.defaults()
	if len(sduo.hooks) == 0 {
		if err = sduo.check(); err != nil {
			return nil, err
		}
		node, err = sduo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysDictMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sduo.check(); err != nil {
				return nil, err
			}
			sduo.mutation = mutation
			node, err = sduo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sduo.hooks) - 1; i >= 0; i-- {
			mut = sduo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sduo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (sduo *SysDictUpdateOne) SaveX(ctx context.Context) *SysDict {
	node, err := sduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sduo *SysDictUpdateOne) Exec(ctx context.Context) error {
	_, err := sduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sduo *SysDictUpdateOne) ExecX(ctx context.Context) {
	if err := sduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sduo *SysDictUpdateOne) defaults() {
	if _, ok := sduo.mutation.UpdatedAt(); !ok {
		v := sysdict.UpdateDefaultUpdatedAt()
		sduo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sduo *SysDictUpdateOne) check() error {
	if v, ok := sduo.mutation.Memo(); ok {
		if err := sysdict.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf("ent: validator failed for field \"memo\": %w", err)}
		}
	}
	if v, ok := sduo.mutation.NameCn(); ok {
		if err := sysdict.NameCnValidator(v); err != nil {
			return &ValidationError{Name: "name_cn", err: fmt.Errorf("ent: validator failed for field \"name_cn\": %w", err)}
		}
	}
	if v, ok := sduo.mutation.NameEn(); ok {
		if err := sysdict.NameEnValidator(v); err != nil {
			return &ValidationError{Name: "name_en", err: fmt.Errorf("ent: validator failed for field \"name_en\": %w", err)}
		}
	}
	return nil
}

func (sduo *SysDictUpdateOne) sqlSave(ctx context.Context) (_node *SysDict, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysdict.Table,
			Columns: sysdict.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: sysdict.FieldID,
			},
		},
	}
	id, ok := sduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing SysDict.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := sduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysdict.FieldID)
		for _, f := range fields {
			if !sysdict.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != sysdict.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sduo.mutation.Memo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdict.FieldMemo,
		})
	}
	if value, ok := sduo.mutation.Sort(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysdict.FieldSort,
		})
	}
	if value, ok := sduo.mutation.AddedSort(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysdict.FieldSort,
		})
	}
	if value, ok := sduo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysdict.FieldUpdatedAt,
		})
	}
	if value, ok := sduo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysdict.FieldDeletedAt,
		})
	}
	if sduo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: sysdict.FieldDeletedAt,
		})
	}
	if value, ok := sduo.mutation.NameCn(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdict.FieldNameCn,
		})
	}
	if value, ok := sduo.mutation.NameEn(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdict.FieldNameEn,
		})
	}
	if value, ok := sduo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysdict.FieldStatus,
		})
	}
	_node = &SysDict{config: sduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysdict.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
