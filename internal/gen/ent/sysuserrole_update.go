// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wanhello/omgind/internal/gen/ent/predicate"
	"github.com/wanhello/omgind/internal/gen/ent/sysuserrole"
)

// SysUserRoleUpdate is the builder for updating SysUserRole entities.
type SysUserRoleUpdate struct {
	config
	hooks    []Hook
	mutation *SysUserRoleMutation
}

// Where adds a new predicate for the SysUserRoleUpdate builder.
func (suru *SysUserRoleUpdate) Where(ps ...predicate.SysUserRole) *SysUserRoleUpdate {
	suru.mutation.predicates = append(suru.mutation.predicates, ps...)
	return suru
}

// Mutation returns the SysUserRoleMutation object of the builder.
func (suru *SysUserRoleUpdate) Mutation() *SysUserRoleMutation {
	return suru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (suru *SysUserRoleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(suru.hooks) == 0 {
		affected, err = suru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysUserRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suru.mutation = mutation
			affected, err = suru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(suru.hooks) - 1; i >= 0; i-- {
			mut = suru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (suru *SysUserRoleUpdate) SaveX(ctx context.Context) int {
	affected, err := suru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (suru *SysUserRoleUpdate) Exec(ctx context.Context) error {
	_, err := suru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suru *SysUserRoleUpdate) ExecX(ctx context.Context) {
	if err := suru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suru *SysUserRoleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysuserrole.Table,
			Columns: sysuserrole.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: sysuserrole.FieldID,
			},
		},
	}
	if ps := suru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, suru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysuserrole.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// SysUserRoleUpdateOne is the builder for updating a single SysUserRole entity.
type SysUserRoleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SysUserRoleMutation
}

// Mutation returns the SysUserRoleMutation object of the builder.
func (suruo *SysUserRoleUpdateOne) Mutation() *SysUserRoleMutation {
	return suruo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suruo *SysUserRoleUpdateOne) Select(field string, fields ...string) *SysUserRoleUpdateOne {
	suruo.fields = append([]string{field}, fields...)
	return suruo
}

// Save executes the query and returns the updated SysUserRole entity.
func (suruo *SysUserRoleUpdateOne) Save(ctx context.Context) (*SysUserRole, error) {
	var (
		err  error
		node *SysUserRole
	)
	if len(suruo.hooks) == 0 {
		node, err = suruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysUserRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suruo.mutation = mutation
			node, err = suruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suruo.hooks) - 1; i >= 0; i-- {
			mut = suruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suruo *SysUserRoleUpdateOne) SaveX(ctx context.Context) *SysUserRole {
	node, err := suruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suruo *SysUserRoleUpdateOne) Exec(ctx context.Context) error {
	_, err := suruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suruo *SysUserRoleUpdateOne) ExecX(ctx context.Context) {
	if err := suruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suruo *SysUserRoleUpdateOne) sqlSave(ctx context.Context) (_node *SysUserRole, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysuserrole.Table,
			Columns: sysuserrole.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: sysuserrole.FieldID,
			},
		},
	}
	id, ok := suruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing SysUserRole.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := suruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysuserrole.FieldID)
		for _, f := range fields {
			if !sysuserrole.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != sysuserrole.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_node = &SysUserRole{config: suruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysuserrole.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
