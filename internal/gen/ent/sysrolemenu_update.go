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
	"github.com/wanhello/omgind/internal/gen/ent/sysrolemenu"
)

// SysRoleMenuUpdate is the builder for updating SysRoleMenu entities.
type SysRoleMenuUpdate struct {
	config
	hooks    []Hook
	mutation *SysRoleMenuMutation
}

// Where adds a new predicate for the SysRoleMenuUpdate builder.
func (srmu *SysRoleMenuUpdate) Where(ps ...predicate.SysRoleMenu) *SysRoleMenuUpdate {
	srmu.mutation.predicates = append(srmu.mutation.predicates, ps...)
	return srmu
}

// SetIsDel sets the "is_del" field.
func (srmu *SysRoleMenuUpdate) SetIsDel(b bool) *SysRoleMenuUpdate {
	srmu.mutation.SetIsDel(b)
	return srmu
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (srmu *SysRoleMenuUpdate) SetNillableIsDel(b *bool) *SysRoleMenuUpdate {
	if b != nil {
		srmu.SetIsDel(*b)
	}
	return srmu
}

// SetUpdatedAt sets the "updated_at" field.
func (srmu *SysRoleMenuUpdate) SetUpdatedAt(t time.Time) *SysRoleMenuUpdate {
	srmu.mutation.SetUpdatedAt(t)
	return srmu
}

// SetDeletedAt sets the "deleted_at" field.
func (srmu *SysRoleMenuUpdate) SetDeletedAt(t time.Time) *SysRoleMenuUpdate {
	srmu.mutation.SetDeletedAt(t)
	return srmu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (srmu *SysRoleMenuUpdate) SetNillableDeletedAt(t *time.Time) *SysRoleMenuUpdate {
	if t != nil {
		srmu.SetDeletedAt(*t)
	}
	return srmu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (srmu *SysRoleMenuUpdate) ClearDeletedAt() *SysRoleMenuUpdate {
	srmu.mutation.ClearDeletedAt()
	return srmu
}

// SetRoleID sets the "role_id" field.
func (srmu *SysRoleMenuUpdate) SetRoleID(s string) *SysRoleMenuUpdate {
	srmu.mutation.SetRoleID(s)
	return srmu
}

// SetMenuID sets the "menu_id" field.
func (srmu *SysRoleMenuUpdate) SetMenuID(s string) *SysRoleMenuUpdate {
	srmu.mutation.SetMenuID(s)
	return srmu
}

// SetActionID sets the "action_id" field.
func (srmu *SysRoleMenuUpdate) SetActionID(s string) *SysRoleMenuUpdate {
	srmu.mutation.SetActionID(s)
	return srmu
}

// SetNillableActionID sets the "action_id" field if the given value is not nil.
func (srmu *SysRoleMenuUpdate) SetNillableActionID(s *string) *SysRoleMenuUpdate {
	if s != nil {
		srmu.SetActionID(*s)
	}
	return srmu
}

// ClearActionID clears the value of the "action_id" field.
func (srmu *SysRoleMenuUpdate) ClearActionID() *SysRoleMenuUpdate {
	srmu.mutation.ClearActionID()
	return srmu
}

// Mutation returns the SysRoleMenuMutation object of the builder.
func (srmu *SysRoleMenuUpdate) Mutation() *SysRoleMenuMutation {
	return srmu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (srmu *SysRoleMenuUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	srmu.defaults()
	if len(srmu.hooks) == 0 {
		if err = srmu.check(); err != nil {
			return 0, err
		}
		affected, err = srmu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysRoleMenuMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = srmu.check(); err != nil {
				return 0, err
			}
			srmu.mutation = mutation
			affected, err = srmu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(srmu.hooks) - 1; i >= 0; i-- {
			mut = srmu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, srmu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (srmu *SysRoleMenuUpdate) SaveX(ctx context.Context) int {
	affected, err := srmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (srmu *SysRoleMenuUpdate) Exec(ctx context.Context) error {
	_, err := srmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (srmu *SysRoleMenuUpdate) ExecX(ctx context.Context) {
	if err := srmu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (srmu *SysRoleMenuUpdate) defaults() {
	if _, ok := srmu.mutation.UpdatedAt(); !ok {
		v := sysrolemenu.UpdateDefaultUpdatedAt()
		srmu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (srmu *SysRoleMenuUpdate) check() error {
	if v, ok := srmu.mutation.RoleID(); ok {
		if err := sysrolemenu.RoleIDValidator(v); err != nil {
			return &ValidationError{Name: "role_id", err: fmt.Errorf("ent: validator failed for field \"role_id\": %w", err)}
		}
	}
	if v, ok := srmu.mutation.MenuID(); ok {
		if err := sysrolemenu.MenuIDValidator(v); err != nil {
			return &ValidationError{Name: "menu_id", err: fmt.Errorf("ent: validator failed for field \"menu_id\": %w", err)}
		}
	}
	if v, ok := srmu.mutation.ActionID(); ok {
		if err := sysrolemenu.ActionIDValidator(v); err != nil {
			return &ValidationError{Name: "action_id", err: fmt.Errorf("ent: validator failed for field \"action_id\": %w", err)}
		}
	}
	return nil
}

func (srmu *SysRoleMenuUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysrolemenu.Table,
			Columns: sysrolemenu.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: sysrolemenu.FieldID,
			},
		},
	}
	if ps := srmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := srmu.mutation.IsDel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysrolemenu.FieldIsDel,
		})
	}
	if value, ok := srmu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysrolemenu.FieldUpdatedAt,
		})
	}
	if value, ok := srmu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysrolemenu.FieldDeletedAt,
		})
	}
	if srmu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: sysrolemenu.FieldDeletedAt,
		})
	}
	if value, ok := srmu.mutation.RoleID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrolemenu.FieldRoleID,
		})
	}
	if value, ok := srmu.mutation.MenuID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrolemenu.FieldMenuID,
		})
	}
	if value, ok := srmu.mutation.ActionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrolemenu.FieldActionID,
		})
	}
	if srmu.mutation.ActionIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysrolemenu.FieldActionID,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, srmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysrolemenu.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// SysRoleMenuUpdateOne is the builder for updating a single SysRoleMenu entity.
type SysRoleMenuUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SysRoleMenuMutation
}

// SetIsDel sets the "is_del" field.
func (srmuo *SysRoleMenuUpdateOne) SetIsDel(b bool) *SysRoleMenuUpdateOne {
	srmuo.mutation.SetIsDel(b)
	return srmuo
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (srmuo *SysRoleMenuUpdateOne) SetNillableIsDel(b *bool) *SysRoleMenuUpdateOne {
	if b != nil {
		srmuo.SetIsDel(*b)
	}
	return srmuo
}

// SetUpdatedAt sets the "updated_at" field.
func (srmuo *SysRoleMenuUpdateOne) SetUpdatedAt(t time.Time) *SysRoleMenuUpdateOne {
	srmuo.mutation.SetUpdatedAt(t)
	return srmuo
}

// SetDeletedAt sets the "deleted_at" field.
func (srmuo *SysRoleMenuUpdateOne) SetDeletedAt(t time.Time) *SysRoleMenuUpdateOne {
	srmuo.mutation.SetDeletedAt(t)
	return srmuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (srmuo *SysRoleMenuUpdateOne) SetNillableDeletedAt(t *time.Time) *SysRoleMenuUpdateOne {
	if t != nil {
		srmuo.SetDeletedAt(*t)
	}
	return srmuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (srmuo *SysRoleMenuUpdateOne) ClearDeletedAt() *SysRoleMenuUpdateOne {
	srmuo.mutation.ClearDeletedAt()
	return srmuo
}

// SetRoleID sets the "role_id" field.
func (srmuo *SysRoleMenuUpdateOne) SetRoleID(s string) *SysRoleMenuUpdateOne {
	srmuo.mutation.SetRoleID(s)
	return srmuo
}

// SetMenuID sets the "menu_id" field.
func (srmuo *SysRoleMenuUpdateOne) SetMenuID(s string) *SysRoleMenuUpdateOne {
	srmuo.mutation.SetMenuID(s)
	return srmuo
}

// SetActionID sets the "action_id" field.
func (srmuo *SysRoleMenuUpdateOne) SetActionID(s string) *SysRoleMenuUpdateOne {
	srmuo.mutation.SetActionID(s)
	return srmuo
}

// SetNillableActionID sets the "action_id" field if the given value is not nil.
func (srmuo *SysRoleMenuUpdateOne) SetNillableActionID(s *string) *SysRoleMenuUpdateOne {
	if s != nil {
		srmuo.SetActionID(*s)
	}
	return srmuo
}

// ClearActionID clears the value of the "action_id" field.
func (srmuo *SysRoleMenuUpdateOne) ClearActionID() *SysRoleMenuUpdateOne {
	srmuo.mutation.ClearActionID()
	return srmuo
}

// Mutation returns the SysRoleMenuMutation object of the builder.
func (srmuo *SysRoleMenuUpdateOne) Mutation() *SysRoleMenuMutation {
	return srmuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (srmuo *SysRoleMenuUpdateOne) Select(field string, fields ...string) *SysRoleMenuUpdateOne {
	srmuo.fields = append([]string{field}, fields...)
	return srmuo
}

// Save executes the query and returns the updated SysRoleMenu entity.
func (srmuo *SysRoleMenuUpdateOne) Save(ctx context.Context) (*SysRoleMenu, error) {
	var (
		err  error
		node *SysRoleMenu
	)
	srmuo.defaults()
	if len(srmuo.hooks) == 0 {
		if err = srmuo.check(); err != nil {
			return nil, err
		}
		node, err = srmuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysRoleMenuMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = srmuo.check(); err != nil {
				return nil, err
			}
			srmuo.mutation = mutation
			node, err = srmuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(srmuo.hooks) - 1; i >= 0; i-- {
			mut = srmuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, srmuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (srmuo *SysRoleMenuUpdateOne) SaveX(ctx context.Context) *SysRoleMenu {
	node, err := srmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (srmuo *SysRoleMenuUpdateOne) Exec(ctx context.Context) error {
	_, err := srmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (srmuo *SysRoleMenuUpdateOne) ExecX(ctx context.Context) {
	if err := srmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (srmuo *SysRoleMenuUpdateOne) defaults() {
	if _, ok := srmuo.mutation.UpdatedAt(); !ok {
		v := sysrolemenu.UpdateDefaultUpdatedAt()
		srmuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (srmuo *SysRoleMenuUpdateOne) check() error {
	if v, ok := srmuo.mutation.RoleID(); ok {
		if err := sysrolemenu.RoleIDValidator(v); err != nil {
			return &ValidationError{Name: "role_id", err: fmt.Errorf("ent: validator failed for field \"role_id\": %w", err)}
		}
	}
	if v, ok := srmuo.mutation.MenuID(); ok {
		if err := sysrolemenu.MenuIDValidator(v); err != nil {
			return &ValidationError{Name: "menu_id", err: fmt.Errorf("ent: validator failed for field \"menu_id\": %w", err)}
		}
	}
	if v, ok := srmuo.mutation.ActionID(); ok {
		if err := sysrolemenu.ActionIDValidator(v); err != nil {
			return &ValidationError{Name: "action_id", err: fmt.Errorf("ent: validator failed for field \"action_id\": %w", err)}
		}
	}
	return nil
}

func (srmuo *SysRoleMenuUpdateOne) sqlSave(ctx context.Context) (_node *SysRoleMenu, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysrolemenu.Table,
			Columns: sysrolemenu.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: sysrolemenu.FieldID,
			},
		},
	}
	id, ok := srmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing SysRoleMenu.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := srmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysrolemenu.FieldID)
		for _, f := range fields {
			if !sysrolemenu.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != sysrolemenu.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := srmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := srmuo.mutation.IsDel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysrolemenu.FieldIsDel,
		})
	}
	if value, ok := srmuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysrolemenu.FieldUpdatedAt,
		})
	}
	if value, ok := srmuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysrolemenu.FieldDeletedAt,
		})
	}
	if srmuo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: sysrolemenu.FieldDeletedAt,
		})
	}
	if value, ok := srmuo.mutation.RoleID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrolemenu.FieldRoleID,
		})
	}
	if value, ok := srmuo.mutation.MenuID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrolemenu.FieldMenuID,
		})
	}
	if value, ok := srmuo.mutation.ActionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrolemenu.FieldActionID,
		})
	}
	if srmuo.mutation.ActionIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysrolemenu.FieldActionID,
		})
	}
	_node = &SysRoleMenu{config: srmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, srmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysrolemenu.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
