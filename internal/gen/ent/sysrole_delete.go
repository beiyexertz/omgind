// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wanhello/omgind/internal/gen/ent/predicate"
	"github.com/wanhello/omgind/internal/gen/ent/sysrole"
)

// SysRoleDelete is the builder for deleting a SysRole entity.
type SysRoleDelete struct {
	config
	hooks    []Hook
	mutation *SysRoleMutation
}

// Where adds a new predicate to the SysRoleDelete builder.
func (srd *SysRoleDelete) Where(ps ...predicate.SysRole) *SysRoleDelete {
	srd.mutation.predicates = append(srd.mutation.predicates, ps...)
	return srd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (srd *SysRoleDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(srd.hooks) == 0 {
		affected, err = srd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			srd.mutation = mutation
			affected, err = srd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(srd.hooks) - 1; i >= 0; i-- {
			mut = srd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, srd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (srd *SysRoleDelete) ExecX(ctx context.Context) int {
	n, err := srd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (srd *SysRoleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: sysrole.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: sysrole.FieldID,
			},
		},
	}
	if ps := srd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, srd.driver, _spec)
}

// SysRoleDeleteOne is the builder for deleting a single SysRole entity.
type SysRoleDeleteOne struct {
	srd *SysRoleDelete
}

// Exec executes the deletion query.
func (srdo *SysRoleDeleteOne) Exec(ctx context.Context) error {
	n, err := srdo.srd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{sysrole.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (srdo *SysRoleDeleteOne) ExecX(ctx context.Context) {
	srdo.srd.ExecX(ctx)
}
