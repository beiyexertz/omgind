// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wanhello/omgind/internal/gen/ent/predicate"
	"github.com/wanhello/omgind/internal/gen/ent/sysrolemenu"
)

// SysRoleMenuDelete is the builder for deleting a SysRoleMenu entity.
type SysRoleMenuDelete struct {
	config
	hooks    []Hook
	mutation *SysRoleMenuMutation
}

// Where adds a new predicate to the SysRoleMenuDelete builder.
func (srmd *SysRoleMenuDelete) Where(ps ...predicate.SysRoleMenu) *SysRoleMenuDelete {
	srmd.mutation.predicates = append(srmd.mutation.predicates, ps...)
	return srmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (srmd *SysRoleMenuDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(srmd.hooks) == 0 {
		affected, err = srmd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysRoleMenuMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			srmd.mutation = mutation
			affected, err = srmd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(srmd.hooks) - 1; i >= 0; i-- {
			mut = srmd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, srmd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (srmd *SysRoleMenuDelete) ExecX(ctx context.Context) int {
	n, err := srmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (srmd *SysRoleMenuDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: sysrolemenu.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: sysrolemenu.FieldID,
			},
		},
	}
	if ps := srmd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, srmd.driver, _spec)
}

// SysRoleMenuDeleteOne is the builder for deleting a single SysRoleMenu entity.
type SysRoleMenuDeleteOne struct {
	srmd *SysRoleMenuDelete
}

// Exec executes the deletion query.
func (srmdo *SysRoleMenuDeleteOne) Exec(ctx context.Context) error {
	n, err := srmdo.srmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{sysrolemenu.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (srmdo *SysRoleMenuDeleteOne) ExecX(ctx context.Context) {
	srmdo.srmd.ExecX(ctx)
}
