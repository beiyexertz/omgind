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
	"github.com/wanhello/omgind/internal/gen/ent/xxxdemo"
)

// XxxDemoUpdate is the builder for updating XxxDemo entities.
type XxxDemoUpdate struct {
	config
	hooks    []Hook
	mutation *XxxDemoMutation
}

// Where adds a new predicate for the XxxDemoUpdate builder.
func (xdu *XxxDemoUpdate) Where(ps ...predicate.XxxDemo) *XxxDemoUpdate {
	xdu.mutation.predicates = append(xdu.mutation.predicates, ps...)
	return xdu
}

// SetMemo sets the "memo" field.
func (xdu *XxxDemoUpdate) SetMemo(s string) *XxxDemoUpdate {
	xdu.mutation.SetMemo(s)
	return xdu
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (xdu *XxxDemoUpdate) SetNillableMemo(s *string) *XxxDemoUpdate {
	if s != nil {
		xdu.SetMemo(*s)
	}
	return xdu
}

// SetSort sets the "sort" field.
func (xdu *XxxDemoUpdate) SetSort(i int32) *XxxDemoUpdate {
	xdu.mutation.ResetSort()
	xdu.mutation.SetSort(i)
	return xdu
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (xdu *XxxDemoUpdate) SetNillableSort(i *int32) *XxxDemoUpdate {
	if i != nil {
		xdu.SetSort(*i)
	}
	return xdu
}

// AddSort adds i to the "sort" field.
func (xdu *XxxDemoUpdate) AddSort(i int32) *XxxDemoUpdate {
	xdu.mutation.AddSort(i)
	return xdu
}

// SetUpdatedAt sets the "updated_at" field.
func (xdu *XxxDemoUpdate) SetUpdatedAt(t time.Time) *XxxDemoUpdate {
	xdu.mutation.SetUpdatedAt(t)
	return xdu
}

// SetDeletedAt sets the "deleted_at" field.
func (xdu *XxxDemoUpdate) SetDeletedAt(t time.Time) *XxxDemoUpdate {
	xdu.mutation.SetDeletedAt(t)
	return xdu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (xdu *XxxDemoUpdate) SetNillableDeletedAt(t *time.Time) *XxxDemoUpdate {
	if t != nil {
		xdu.SetDeletedAt(*t)
	}
	return xdu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (xdu *XxxDemoUpdate) ClearDeletedAt() *XxxDemoUpdate {
	xdu.mutation.ClearDeletedAt()
	return xdu
}

// SetCode sets the "code" field.
func (xdu *XxxDemoUpdate) SetCode(s string) *XxxDemoUpdate {
	xdu.mutation.SetCode(s)
	return xdu
}

// SetName sets the "name" field.
func (xdu *XxxDemoUpdate) SetName(s string) *XxxDemoUpdate {
	xdu.mutation.SetName(s)
	return xdu
}

// SetStatus sets the "status" field.
func (xdu *XxxDemoUpdate) SetStatus(i int) *XxxDemoUpdate {
	xdu.mutation.ResetStatus()
	xdu.mutation.SetStatus(i)
	return xdu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (xdu *XxxDemoUpdate) SetNillableStatus(i *int) *XxxDemoUpdate {
	if i != nil {
		xdu.SetStatus(*i)
	}
	return xdu
}

// AddStatus adds i to the "status" field.
func (xdu *XxxDemoUpdate) AddStatus(i int) *XxxDemoUpdate {
	xdu.mutation.AddStatus(i)
	return xdu
}

// Mutation returns the XxxDemoMutation object of the builder.
func (xdu *XxxDemoUpdate) Mutation() *XxxDemoMutation {
	return xdu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (xdu *XxxDemoUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	xdu.defaults()
	if len(xdu.hooks) == 0 {
		if err = xdu.check(); err != nil {
			return 0, err
		}
		affected, err = xdu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*XxxDemoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = xdu.check(); err != nil {
				return 0, err
			}
			xdu.mutation = mutation
			affected, err = xdu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(xdu.hooks) - 1; i >= 0; i-- {
			mut = xdu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, xdu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (xdu *XxxDemoUpdate) SaveX(ctx context.Context) int {
	affected, err := xdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (xdu *XxxDemoUpdate) Exec(ctx context.Context) error {
	_, err := xdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (xdu *XxxDemoUpdate) ExecX(ctx context.Context) {
	if err := xdu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (xdu *XxxDemoUpdate) defaults() {
	if _, ok := xdu.mutation.UpdatedAt(); !ok {
		v := xxxdemo.UpdateDefaultUpdatedAt()
		xdu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (xdu *XxxDemoUpdate) check() error {
	if v, ok := xdu.mutation.Memo(); ok {
		if err := xxxdemo.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf("ent: validator failed for field \"memo\": %w", err)}
		}
	}
	if v, ok := xdu.mutation.Code(); ok {
		if err := xxxdemo.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf("ent: validator failed for field \"code\": %w", err)}
		}
	}
	if v, ok := xdu.mutation.Name(); ok {
		if err := xxxdemo.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	return nil
}

func (xdu *XxxDemoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   xxxdemo.Table,
			Columns: xxxdemo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: xxxdemo.FieldID,
			},
		},
	}
	if ps := xdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := xdu.mutation.Memo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: xxxdemo.FieldMemo,
		})
	}
	if value, ok := xdu.mutation.Sort(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: xxxdemo.FieldSort,
		})
	}
	if value, ok := xdu.mutation.AddedSort(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: xxxdemo.FieldSort,
		})
	}
	if value, ok := xdu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: xxxdemo.FieldUpdatedAt,
		})
	}
	if value, ok := xdu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: xxxdemo.FieldDeletedAt,
		})
	}
	if xdu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: xxxdemo.FieldDeletedAt,
		})
	}
	if value, ok := xdu.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: xxxdemo.FieldCode,
		})
	}
	if value, ok := xdu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: xxxdemo.FieldName,
		})
	}
	if value, ok := xdu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: xxxdemo.FieldStatus,
		})
	}
	if value, ok := xdu.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: xxxdemo.FieldStatus,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, xdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{xxxdemo.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// XxxDemoUpdateOne is the builder for updating a single XxxDemo entity.
type XxxDemoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *XxxDemoMutation
}

// SetMemo sets the "memo" field.
func (xduo *XxxDemoUpdateOne) SetMemo(s string) *XxxDemoUpdateOne {
	xduo.mutation.SetMemo(s)
	return xduo
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (xduo *XxxDemoUpdateOne) SetNillableMemo(s *string) *XxxDemoUpdateOne {
	if s != nil {
		xduo.SetMemo(*s)
	}
	return xduo
}

// SetSort sets the "sort" field.
func (xduo *XxxDemoUpdateOne) SetSort(i int32) *XxxDemoUpdateOne {
	xduo.mutation.ResetSort()
	xduo.mutation.SetSort(i)
	return xduo
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (xduo *XxxDemoUpdateOne) SetNillableSort(i *int32) *XxxDemoUpdateOne {
	if i != nil {
		xduo.SetSort(*i)
	}
	return xduo
}

// AddSort adds i to the "sort" field.
func (xduo *XxxDemoUpdateOne) AddSort(i int32) *XxxDemoUpdateOne {
	xduo.mutation.AddSort(i)
	return xduo
}

// SetUpdatedAt sets the "updated_at" field.
func (xduo *XxxDemoUpdateOne) SetUpdatedAt(t time.Time) *XxxDemoUpdateOne {
	xduo.mutation.SetUpdatedAt(t)
	return xduo
}

// SetDeletedAt sets the "deleted_at" field.
func (xduo *XxxDemoUpdateOne) SetDeletedAt(t time.Time) *XxxDemoUpdateOne {
	xduo.mutation.SetDeletedAt(t)
	return xduo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (xduo *XxxDemoUpdateOne) SetNillableDeletedAt(t *time.Time) *XxxDemoUpdateOne {
	if t != nil {
		xduo.SetDeletedAt(*t)
	}
	return xduo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (xduo *XxxDemoUpdateOne) ClearDeletedAt() *XxxDemoUpdateOne {
	xduo.mutation.ClearDeletedAt()
	return xduo
}

// SetCode sets the "code" field.
func (xduo *XxxDemoUpdateOne) SetCode(s string) *XxxDemoUpdateOne {
	xduo.mutation.SetCode(s)
	return xduo
}

// SetName sets the "name" field.
func (xduo *XxxDemoUpdateOne) SetName(s string) *XxxDemoUpdateOne {
	xduo.mutation.SetName(s)
	return xduo
}

// SetStatus sets the "status" field.
func (xduo *XxxDemoUpdateOne) SetStatus(i int) *XxxDemoUpdateOne {
	xduo.mutation.ResetStatus()
	xduo.mutation.SetStatus(i)
	return xduo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (xduo *XxxDemoUpdateOne) SetNillableStatus(i *int) *XxxDemoUpdateOne {
	if i != nil {
		xduo.SetStatus(*i)
	}
	return xduo
}

// AddStatus adds i to the "status" field.
func (xduo *XxxDemoUpdateOne) AddStatus(i int) *XxxDemoUpdateOne {
	xduo.mutation.AddStatus(i)
	return xduo
}

// Mutation returns the XxxDemoMutation object of the builder.
func (xduo *XxxDemoUpdateOne) Mutation() *XxxDemoMutation {
	return xduo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (xduo *XxxDemoUpdateOne) Select(field string, fields ...string) *XxxDemoUpdateOne {
	xduo.fields = append([]string{field}, fields...)
	return xduo
}

// Save executes the query and returns the updated XxxDemo entity.
func (xduo *XxxDemoUpdateOne) Save(ctx context.Context) (*XxxDemo, error) {
	var (
		err  error
		node *XxxDemo
	)
	xduo.defaults()
	if len(xduo.hooks) == 0 {
		if err = xduo.check(); err != nil {
			return nil, err
		}
		node, err = xduo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*XxxDemoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = xduo.check(); err != nil {
				return nil, err
			}
			xduo.mutation = mutation
			node, err = xduo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(xduo.hooks) - 1; i >= 0; i-- {
			mut = xduo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, xduo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (xduo *XxxDemoUpdateOne) SaveX(ctx context.Context) *XxxDemo {
	node, err := xduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (xduo *XxxDemoUpdateOne) Exec(ctx context.Context) error {
	_, err := xduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (xduo *XxxDemoUpdateOne) ExecX(ctx context.Context) {
	if err := xduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (xduo *XxxDemoUpdateOne) defaults() {
	if _, ok := xduo.mutation.UpdatedAt(); !ok {
		v := xxxdemo.UpdateDefaultUpdatedAt()
		xduo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (xduo *XxxDemoUpdateOne) check() error {
	if v, ok := xduo.mutation.Memo(); ok {
		if err := xxxdemo.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf("ent: validator failed for field \"memo\": %w", err)}
		}
	}
	if v, ok := xduo.mutation.Code(); ok {
		if err := xxxdemo.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf("ent: validator failed for field \"code\": %w", err)}
		}
	}
	if v, ok := xduo.mutation.Name(); ok {
		if err := xxxdemo.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	return nil
}

func (xduo *XxxDemoUpdateOne) sqlSave(ctx context.Context) (_node *XxxDemo, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   xxxdemo.Table,
			Columns: xxxdemo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: xxxdemo.FieldID,
			},
		},
	}
	id, ok := xduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing XxxDemo.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := xduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, xxxdemo.FieldID)
		for _, f := range fields {
			if !xxxdemo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != xxxdemo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := xduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := xduo.mutation.Memo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: xxxdemo.FieldMemo,
		})
	}
	if value, ok := xduo.mutation.Sort(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: xxxdemo.FieldSort,
		})
	}
	if value, ok := xduo.mutation.AddedSort(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: xxxdemo.FieldSort,
		})
	}
	if value, ok := xduo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: xxxdemo.FieldUpdatedAt,
		})
	}
	if value, ok := xduo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: xxxdemo.FieldDeletedAt,
		})
	}
	if xduo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: xxxdemo.FieldDeletedAt,
		})
	}
	if value, ok := xduo.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: xxxdemo.FieldCode,
		})
	}
	if value, ok := xduo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: xxxdemo.FieldName,
		})
	}
	if value, ok := xduo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: xxxdemo.FieldStatus,
		})
	}
	if value, ok := xduo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: xxxdemo.FieldStatus,
		})
	}
	_node = &XxxDemo{config: xduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, xduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{xxxdemo.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
