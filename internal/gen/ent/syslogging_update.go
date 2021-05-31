// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wanhello/omgind/internal/gen/ent/predicate"
	"github.com/wanhello/omgind/internal/gen/ent/syslogging"
)

// SysLoggingUpdate is the builder for updating SysLogging entities.
type SysLoggingUpdate struct {
	config
	hooks    []Hook
	mutation *SysLoggingMutation
}

// Where adds a new predicate for the SysLoggingUpdate builder.
func (slu *SysLoggingUpdate) Where(ps ...predicate.SysLogging) *SysLoggingUpdate {
	slu.mutation.predicates = append(slu.mutation.predicates, ps...)
	return slu
}

// SetIsDel sets the "is_del" field.
func (slu *SysLoggingUpdate) SetIsDel(b bool) *SysLoggingUpdate {
	slu.mutation.SetIsDel(b)
	return slu
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (slu *SysLoggingUpdate) SetNillableIsDel(b *bool) *SysLoggingUpdate {
	if b != nil {
		slu.SetIsDel(*b)
	}
	return slu
}

// SetMemo sets the "memo" field.
func (slu *SysLoggingUpdate) SetMemo(s string) *SysLoggingUpdate {
	slu.mutation.SetMemo(s)
	return slu
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (slu *SysLoggingUpdate) SetNillableMemo(s *string) *SysLoggingUpdate {
	if s != nil {
		slu.SetMemo(*s)
	}
	return slu
}

// SetLevel sets the "level" field.
func (slu *SysLoggingUpdate) SetLevel(s string) *SysLoggingUpdate {
	slu.mutation.SetLevel(s)
	return slu
}

// SetTraceID sets the "trace_id" field.
func (slu *SysLoggingUpdate) SetTraceID(s string) *SysLoggingUpdate {
	slu.mutation.SetTraceID(s)
	return slu
}

// SetUserID sets the "user_id" field.
func (slu *SysLoggingUpdate) SetUserID(s string) *SysLoggingUpdate {
	slu.mutation.SetUserID(s)
	return slu
}

// SetTag sets the "tag" field.
func (slu *SysLoggingUpdate) SetTag(s string) *SysLoggingUpdate {
	slu.mutation.SetTag(s)
	return slu
}

// SetVersion sets the "version" field.
func (slu *SysLoggingUpdate) SetVersion(s string) *SysLoggingUpdate {
	slu.mutation.SetVersion(s)
	return slu
}

// SetMessage sets the "message" field.
func (slu *SysLoggingUpdate) SetMessage(s string) *SysLoggingUpdate {
	slu.mutation.SetMessage(s)
	return slu
}

// SetData sets the "data" field.
func (slu *SysLoggingUpdate) SetData(s string) *SysLoggingUpdate {
	slu.mutation.SetData(s)
	return slu
}

// SetErrorStack sets the "error_stack" field.
func (slu *SysLoggingUpdate) SetErrorStack(s string) *SysLoggingUpdate {
	slu.mutation.SetErrorStack(s)
	return slu
}

// Mutation returns the SysLoggingMutation object of the builder.
func (slu *SysLoggingUpdate) Mutation() *SysLoggingMutation {
	return slu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (slu *SysLoggingUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(slu.hooks) == 0 {
		if err = slu.check(); err != nil {
			return 0, err
		}
		affected, err = slu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysLoggingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = slu.check(); err != nil {
				return 0, err
			}
			slu.mutation = mutation
			affected, err = slu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(slu.hooks) - 1; i >= 0; i-- {
			mut = slu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, slu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (slu *SysLoggingUpdate) SaveX(ctx context.Context) int {
	affected, err := slu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (slu *SysLoggingUpdate) Exec(ctx context.Context) error {
	_, err := slu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (slu *SysLoggingUpdate) ExecX(ctx context.Context) {
	if err := slu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (slu *SysLoggingUpdate) check() error {
	if v, ok := slu.mutation.Memo(); ok {
		if err := syslogging.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf("ent: validator failed for field \"memo\": %w", err)}
		}
	}
	if v, ok := slu.mutation.Level(); ok {
		if err := syslogging.LevelValidator(v); err != nil {
			return &ValidationError{Name: "level", err: fmt.Errorf("ent: validator failed for field \"level\": %w", err)}
		}
	}
	if v, ok := slu.mutation.TraceID(); ok {
		if err := syslogging.TraceIDValidator(v); err != nil {
			return &ValidationError{Name: "trace_id", err: fmt.Errorf("ent: validator failed for field \"trace_id\": %w", err)}
		}
	}
	if v, ok := slu.mutation.UserID(); ok {
		if err := syslogging.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf("ent: validator failed for field \"user_id\": %w", err)}
		}
	}
	if v, ok := slu.mutation.Tag(); ok {
		if err := syslogging.TagValidator(v); err != nil {
			return &ValidationError{Name: "tag", err: fmt.Errorf("ent: validator failed for field \"tag\": %w", err)}
		}
	}
	if v, ok := slu.mutation.Version(); ok {
		if err := syslogging.VersionValidator(v); err != nil {
			return &ValidationError{Name: "version", err: fmt.Errorf("ent: validator failed for field \"version\": %w", err)}
		}
	}
	return nil
}

func (slu *SysLoggingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   syslogging.Table,
			Columns: syslogging.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: syslogging.FieldID,
			},
		},
	}
	if ps := slu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := slu.mutation.IsDel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: syslogging.FieldIsDel,
		})
	}
	if value, ok := slu.mutation.Memo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syslogging.FieldMemo,
		})
	}
	if value, ok := slu.mutation.Level(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syslogging.FieldLevel,
		})
	}
	if value, ok := slu.mutation.TraceID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syslogging.FieldTraceID,
		})
	}
	if value, ok := slu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syslogging.FieldUserID,
		})
	}
	if value, ok := slu.mutation.Tag(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syslogging.FieldTag,
		})
	}
	if value, ok := slu.mutation.Version(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syslogging.FieldVersion,
		})
	}
	if value, ok := slu.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syslogging.FieldMessage,
		})
	}
	if value, ok := slu.mutation.Data(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syslogging.FieldData,
		})
	}
	if value, ok := slu.mutation.ErrorStack(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syslogging.FieldErrorStack,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, slu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{syslogging.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// SysLoggingUpdateOne is the builder for updating a single SysLogging entity.
type SysLoggingUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SysLoggingMutation
}

// SetIsDel sets the "is_del" field.
func (sluo *SysLoggingUpdateOne) SetIsDel(b bool) *SysLoggingUpdateOne {
	sluo.mutation.SetIsDel(b)
	return sluo
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (sluo *SysLoggingUpdateOne) SetNillableIsDel(b *bool) *SysLoggingUpdateOne {
	if b != nil {
		sluo.SetIsDel(*b)
	}
	return sluo
}

// SetMemo sets the "memo" field.
func (sluo *SysLoggingUpdateOne) SetMemo(s string) *SysLoggingUpdateOne {
	sluo.mutation.SetMemo(s)
	return sluo
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (sluo *SysLoggingUpdateOne) SetNillableMemo(s *string) *SysLoggingUpdateOne {
	if s != nil {
		sluo.SetMemo(*s)
	}
	return sluo
}

// SetLevel sets the "level" field.
func (sluo *SysLoggingUpdateOne) SetLevel(s string) *SysLoggingUpdateOne {
	sluo.mutation.SetLevel(s)
	return sluo
}

// SetTraceID sets the "trace_id" field.
func (sluo *SysLoggingUpdateOne) SetTraceID(s string) *SysLoggingUpdateOne {
	sluo.mutation.SetTraceID(s)
	return sluo
}

// SetUserID sets the "user_id" field.
func (sluo *SysLoggingUpdateOne) SetUserID(s string) *SysLoggingUpdateOne {
	sluo.mutation.SetUserID(s)
	return sluo
}

// SetTag sets the "tag" field.
func (sluo *SysLoggingUpdateOne) SetTag(s string) *SysLoggingUpdateOne {
	sluo.mutation.SetTag(s)
	return sluo
}

// SetVersion sets the "version" field.
func (sluo *SysLoggingUpdateOne) SetVersion(s string) *SysLoggingUpdateOne {
	sluo.mutation.SetVersion(s)
	return sluo
}

// SetMessage sets the "message" field.
func (sluo *SysLoggingUpdateOne) SetMessage(s string) *SysLoggingUpdateOne {
	sluo.mutation.SetMessage(s)
	return sluo
}

// SetData sets the "data" field.
func (sluo *SysLoggingUpdateOne) SetData(s string) *SysLoggingUpdateOne {
	sluo.mutation.SetData(s)
	return sluo
}

// SetErrorStack sets the "error_stack" field.
func (sluo *SysLoggingUpdateOne) SetErrorStack(s string) *SysLoggingUpdateOne {
	sluo.mutation.SetErrorStack(s)
	return sluo
}

// Mutation returns the SysLoggingMutation object of the builder.
func (sluo *SysLoggingUpdateOne) Mutation() *SysLoggingMutation {
	return sluo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sluo *SysLoggingUpdateOne) Select(field string, fields ...string) *SysLoggingUpdateOne {
	sluo.fields = append([]string{field}, fields...)
	return sluo
}

// Save executes the query and returns the updated SysLogging entity.
func (sluo *SysLoggingUpdateOne) Save(ctx context.Context) (*SysLogging, error) {
	var (
		err  error
		node *SysLogging
	)
	if len(sluo.hooks) == 0 {
		if err = sluo.check(); err != nil {
			return nil, err
		}
		node, err = sluo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysLoggingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sluo.check(); err != nil {
				return nil, err
			}
			sluo.mutation = mutation
			node, err = sluo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sluo.hooks) - 1; i >= 0; i-- {
			mut = sluo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sluo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (sluo *SysLoggingUpdateOne) SaveX(ctx context.Context) *SysLogging {
	node, err := sluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sluo *SysLoggingUpdateOne) Exec(ctx context.Context) error {
	_, err := sluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sluo *SysLoggingUpdateOne) ExecX(ctx context.Context) {
	if err := sluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sluo *SysLoggingUpdateOne) check() error {
	if v, ok := sluo.mutation.Memo(); ok {
		if err := syslogging.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf("ent: validator failed for field \"memo\": %w", err)}
		}
	}
	if v, ok := sluo.mutation.Level(); ok {
		if err := syslogging.LevelValidator(v); err != nil {
			return &ValidationError{Name: "level", err: fmt.Errorf("ent: validator failed for field \"level\": %w", err)}
		}
	}
	if v, ok := sluo.mutation.TraceID(); ok {
		if err := syslogging.TraceIDValidator(v); err != nil {
			return &ValidationError{Name: "trace_id", err: fmt.Errorf("ent: validator failed for field \"trace_id\": %w", err)}
		}
	}
	if v, ok := sluo.mutation.UserID(); ok {
		if err := syslogging.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf("ent: validator failed for field \"user_id\": %w", err)}
		}
	}
	if v, ok := sluo.mutation.Tag(); ok {
		if err := syslogging.TagValidator(v); err != nil {
			return &ValidationError{Name: "tag", err: fmt.Errorf("ent: validator failed for field \"tag\": %w", err)}
		}
	}
	if v, ok := sluo.mutation.Version(); ok {
		if err := syslogging.VersionValidator(v); err != nil {
			return &ValidationError{Name: "version", err: fmt.Errorf("ent: validator failed for field \"version\": %w", err)}
		}
	}
	return nil
}

func (sluo *SysLoggingUpdateOne) sqlSave(ctx context.Context) (_node *SysLogging, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   syslogging.Table,
			Columns: syslogging.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: syslogging.FieldID,
			},
		},
	}
	id, ok := sluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing SysLogging.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := sluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, syslogging.FieldID)
		for _, f := range fields {
			if !syslogging.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != syslogging.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sluo.mutation.IsDel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: syslogging.FieldIsDel,
		})
	}
	if value, ok := sluo.mutation.Memo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syslogging.FieldMemo,
		})
	}
	if value, ok := sluo.mutation.Level(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syslogging.FieldLevel,
		})
	}
	if value, ok := sluo.mutation.TraceID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syslogging.FieldTraceID,
		})
	}
	if value, ok := sluo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syslogging.FieldUserID,
		})
	}
	if value, ok := sluo.mutation.Tag(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syslogging.FieldTag,
		})
	}
	if value, ok := sluo.mutation.Version(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syslogging.FieldVersion,
		})
	}
	if value, ok := sluo.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syslogging.FieldMessage,
		})
	}
	if value, ok := sluo.mutation.Data(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syslogging.FieldData,
		})
	}
	if value, ok := sluo.mutation.ErrorStack(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syslogging.FieldErrorStack,
		})
	}
	_node = &SysLogging{config: sluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{syslogging.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
