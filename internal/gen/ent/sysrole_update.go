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
	"github.com/wanhello/omgind/internal/gen/ent/sysrole"
)

// SysRoleUpdate is the builder for updating SysRole entities.
type SysRoleUpdate struct {
	config
	hooks    []Hook
	mutation *SysRoleMutation
}

// Where adds a new predicate for the SysRoleUpdate builder.
func (sru *SysRoleUpdate) Where(ps ...predicate.SysRole) *SysRoleUpdate {
	sru.mutation.predicates = append(sru.mutation.predicates, ps...)
	return sru
}

// SetIsDel sets the "is_del" field.
func (sru *SysRoleUpdate) SetIsDel(b bool) *SysRoleUpdate {
	sru.mutation.SetIsDel(b)
	return sru
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (sru *SysRoleUpdate) SetNillableIsDel(b *bool) *SysRoleUpdate {
	if b != nil {
		sru.SetIsDel(*b)
	}
	return sru
}

// SetStatus sets the "status" field.
func (sru *SysRoleUpdate) SetStatus(i int) *SysRoleUpdate {
	sru.mutation.ResetStatus()
	sru.mutation.SetStatus(i)
	return sru
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sru *SysRoleUpdate) SetNillableStatus(i *int) *SysRoleUpdate {
	if i != nil {
		sru.SetStatus(*i)
	}
	return sru
}

// AddStatus adds i to the "status" field.
func (sru *SysRoleUpdate) AddStatus(i int) *SysRoleUpdate {
	sru.mutation.AddStatus(i)
	return sru
}

// SetSort sets the "sort" field.
func (sru *SysRoleUpdate) SetSort(i int) *SysRoleUpdate {
	sru.mutation.ResetSort()
	sru.mutation.SetSort(i)
	return sru
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (sru *SysRoleUpdate) SetNillableSort(i *int) *SysRoleUpdate {
	if i != nil {
		sru.SetSort(*i)
	}
	return sru
}

// AddSort adds i to the "sort" field.
func (sru *SysRoleUpdate) AddSort(i int) *SysRoleUpdate {
	sru.mutation.AddSort(i)
	return sru
}

// SetMemo sets the "memo" field.
func (sru *SysRoleUpdate) SetMemo(s string) *SysRoleUpdate {
	sru.mutation.SetMemo(s)
	return sru
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (sru *SysRoleUpdate) SetNillableMemo(s *string) *SysRoleUpdate {
	if s != nil {
		sru.SetMemo(*s)
	}
	return sru
}

// SetUpdatedAt sets the "updated_at" field.
func (sru *SysRoleUpdate) SetUpdatedAt(t time.Time) *SysRoleUpdate {
	sru.mutation.SetUpdatedAt(t)
	return sru
}

// SetDeletedAt sets the "deleted_at" field.
func (sru *SysRoleUpdate) SetDeletedAt(t time.Time) *SysRoleUpdate {
	sru.mutation.SetDeletedAt(t)
	return sru
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sru *SysRoleUpdate) SetNillableDeletedAt(t *time.Time) *SysRoleUpdate {
	if t != nil {
		sru.SetDeletedAt(*t)
	}
	return sru
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (sru *SysRoleUpdate) ClearDeletedAt() *SysRoleUpdate {
	sru.mutation.ClearDeletedAt()
	return sru
}

// SetName sets the "name" field.
func (sru *SysRoleUpdate) SetName(s string) *SysRoleUpdate {
	sru.mutation.SetName(s)
	return sru
}

// Mutation returns the SysRoleMutation object of the builder.
func (sru *SysRoleUpdate) Mutation() *SysRoleMutation {
	return sru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sru *SysRoleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	sru.defaults()
	if len(sru.hooks) == 0 {
		if err = sru.check(); err != nil {
			return 0, err
		}
		affected, err = sru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sru.check(); err != nil {
				return 0, err
			}
			sru.mutation = mutation
			affected, err = sru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sru.hooks) - 1; i >= 0; i-- {
			mut = sru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (sru *SysRoleUpdate) SaveX(ctx context.Context) int {
	affected, err := sru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sru *SysRoleUpdate) Exec(ctx context.Context) error {
	_, err := sru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sru *SysRoleUpdate) ExecX(ctx context.Context) {
	if err := sru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sru *SysRoleUpdate) defaults() {
	if _, ok := sru.mutation.UpdatedAt(); !ok {
		v := sysrole.UpdateDefaultUpdatedAt()
		sru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sru *SysRoleUpdate) check() error {
	if v, ok := sru.mutation.Memo(); ok {
		if err := sysrole.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf("ent: validator failed for field \"memo\": %w", err)}
		}
	}
	if v, ok := sru.mutation.Name(); ok {
		if err := sysrole.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	return nil
}

func (sru *SysRoleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysrole.Table,
			Columns: sysrole.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: sysrole.FieldID,
			},
		},
	}
	if ps := sru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sru.mutation.IsDel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysrole.FieldIsDel,
		})
	}
	if value, ok := sru.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: sysrole.FieldStatus,
		})
	}
	if value, ok := sru.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: sysrole.FieldStatus,
		})
	}
	if value, ok := sru.mutation.Sort(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: sysrole.FieldSort,
		})
	}
	if value, ok := sru.mutation.AddedSort(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: sysrole.FieldSort,
		})
	}
	if value, ok := sru.mutation.Memo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrole.FieldMemo,
		})
	}
	if value, ok := sru.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysrole.FieldUpdatedAt,
		})
	}
	if value, ok := sru.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysrole.FieldDeletedAt,
		})
	}
	if sru.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: sysrole.FieldDeletedAt,
		})
	}
	if value, ok := sru.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrole.FieldName,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, sru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysrole.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// SysRoleUpdateOne is the builder for updating a single SysRole entity.
type SysRoleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SysRoleMutation
}

// SetIsDel sets the "is_del" field.
func (sruo *SysRoleUpdateOne) SetIsDel(b bool) *SysRoleUpdateOne {
	sruo.mutation.SetIsDel(b)
	return sruo
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (sruo *SysRoleUpdateOne) SetNillableIsDel(b *bool) *SysRoleUpdateOne {
	if b != nil {
		sruo.SetIsDel(*b)
	}
	return sruo
}

// SetStatus sets the "status" field.
func (sruo *SysRoleUpdateOne) SetStatus(i int) *SysRoleUpdateOne {
	sruo.mutation.ResetStatus()
	sruo.mutation.SetStatus(i)
	return sruo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sruo *SysRoleUpdateOne) SetNillableStatus(i *int) *SysRoleUpdateOne {
	if i != nil {
		sruo.SetStatus(*i)
	}
	return sruo
}

// AddStatus adds i to the "status" field.
func (sruo *SysRoleUpdateOne) AddStatus(i int) *SysRoleUpdateOne {
	sruo.mutation.AddStatus(i)
	return sruo
}

// SetSort sets the "sort" field.
func (sruo *SysRoleUpdateOne) SetSort(i int) *SysRoleUpdateOne {
	sruo.mutation.ResetSort()
	sruo.mutation.SetSort(i)
	return sruo
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (sruo *SysRoleUpdateOne) SetNillableSort(i *int) *SysRoleUpdateOne {
	if i != nil {
		sruo.SetSort(*i)
	}
	return sruo
}

// AddSort adds i to the "sort" field.
func (sruo *SysRoleUpdateOne) AddSort(i int) *SysRoleUpdateOne {
	sruo.mutation.AddSort(i)
	return sruo
}

// SetMemo sets the "memo" field.
func (sruo *SysRoleUpdateOne) SetMemo(s string) *SysRoleUpdateOne {
	sruo.mutation.SetMemo(s)
	return sruo
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (sruo *SysRoleUpdateOne) SetNillableMemo(s *string) *SysRoleUpdateOne {
	if s != nil {
		sruo.SetMemo(*s)
	}
	return sruo
}

// SetUpdatedAt sets the "updated_at" field.
func (sruo *SysRoleUpdateOne) SetUpdatedAt(t time.Time) *SysRoleUpdateOne {
	sruo.mutation.SetUpdatedAt(t)
	return sruo
}

// SetDeletedAt sets the "deleted_at" field.
func (sruo *SysRoleUpdateOne) SetDeletedAt(t time.Time) *SysRoleUpdateOne {
	sruo.mutation.SetDeletedAt(t)
	return sruo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sruo *SysRoleUpdateOne) SetNillableDeletedAt(t *time.Time) *SysRoleUpdateOne {
	if t != nil {
		sruo.SetDeletedAt(*t)
	}
	return sruo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (sruo *SysRoleUpdateOne) ClearDeletedAt() *SysRoleUpdateOne {
	sruo.mutation.ClearDeletedAt()
	return sruo
}

// SetName sets the "name" field.
func (sruo *SysRoleUpdateOne) SetName(s string) *SysRoleUpdateOne {
	sruo.mutation.SetName(s)
	return sruo
}

// Mutation returns the SysRoleMutation object of the builder.
func (sruo *SysRoleUpdateOne) Mutation() *SysRoleMutation {
	return sruo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sruo *SysRoleUpdateOne) Select(field string, fields ...string) *SysRoleUpdateOne {
	sruo.fields = append([]string{field}, fields...)
	return sruo
}

// Save executes the query and returns the updated SysRole entity.
func (sruo *SysRoleUpdateOne) Save(ctx context.Context) (*SysRole, error) {
	var (
		err  error
		node *SysRole
	)
	sruo.defaults()
	if len(sruo.hooks) == 0 {
		if err = sruo.check(); err != nil {
			return nil, err
		}
		node, err = sruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sruo.check(); err != nil {
				return nil, err
			}
			sruo.mutation = mutation
			node, err = sruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sruo.hooks) - 1; i >= 0; i-- {
			mut = sruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (sruo *SysRoleUpdateOne) SaveX(ctx context.Context) *SysRole {
	node, err := sruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sruo *SysRoleUpdateOne) Exec(ctx context.Context) error {
	_, err := sruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sruo *SysRoleUpdateOne) ExecX(ctx context.Context) {
	if err := sruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sruo *SysRoleUpdateOne) defaults() {
	if _, ok := sruo.mutation.UpdatedAt(); !ok {
		v := sysrole.UpdateDefaultUpdatedAt()
		sruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sruo *SysRoleUpdateOne) check() error {
	if v, ok := sruo.mutation.Memo(); ok {
		if err := sysrole.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf("ent: validator failed for field \"memo\": %w", err)}
		}
	}
	if v, ok := sruo.mutation.Name(); ok {
		if err := sysrole.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	return nil
}

func (sruo *SysRoleUpdateOne) sqlSave(ctx context.Context) (_node *SysRole, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysrole.Table,
			Columns: sysrole.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: sysrole.FieldID,
			},
		},
	}
	id, ok := sruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing SysRole.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := sruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysrole.FieldID)
		for _, f := range fields {
			if !sysrole.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != sysrole.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sruo.mutation.IsDel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysrole.FieldIsDel,
		})
	}
	if value, ok := sruo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: sysrole.FieldStatus,
		})
	}
	if value, ok := sruo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: sysrole.FieldStatus,
		})
	}
	if value, ok := sruo.mutation.Sort(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: sysrole.FieldSort,
		})
	}
	if value, ok := sruo.mutation.AddedSort(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: sysrole.FieldSort,
		})
	}
	if value, ok := sruo.mutation.Memo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrole.FieldMemo,
		})
	}
	if value, ok := sruo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysrole.FieldUpdatedAt,
		})
	}
	if value, ok := sruo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysrole.FieldDeletedAt,
		})
	}
	if sruo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: sysrole.FieldDeletedAt,
		})
	}
	if value, ok := sruo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrole.FieldName,
		})
	}
	_node = &SysRole{config: sruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysrole.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
