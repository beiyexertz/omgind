// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wanhello/omgind/internal/gen/ent/predicate"
	"github.com/wanhello/omgind/internal/gen/ent/sysjwtblock"
)

// SysJwtBlockDelete is the builder for deleting a SysJwtBlock entity.
type SysJwtBlockDelete struct {
	config
	hooks    []Hook
	mutation *SysJwtBlockMutation
}

// Where adds a new predicate to the SysJwtBlockDelete builder.
func (sjbd *SysJwtBlockDelete) Where(ps ...predicate.SysJwtBlock) *SysJwtBlockDelete {
	sjbd.mutation.predicates = append(sjbd.mutation.predicates, ps...)
	return sjbd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sjbd *SysJwtBlockDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sjbd.hooks) == 0 {
		affected, err = sjbd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysJwtBlockMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sjbd.mutation = mutation
			affected, err = sjbd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sjbd.hooks) - 1; i >= 0; i-- {
			mut = sjbd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sjbd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (sjbd *SysJwtBlockDelete) ExecX(ctx context.Context) int {
	n, err := sjbd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sjbd *SysJwtBlockDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: sysjwtblock.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: sysjwtblock.FieldID,
			},
		},
	}
	if ps := sjbd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, sjbd.driver, _spec)
}

// SysJwtBlockDeleteOne is the builder for deleting a single SysJwtBlock entity.
type SysJwtBlockDeleteOne struct {
	sjbd *SysJwtBlockDelete
}

// Exec executes the deletion query.
func (sjbdo *SysJwtBlockDeleteOne) Exec(ctx context.Context) error {
	n, err := sjbdo.sjbd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{sysjwtblock.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sjbdo *SysJwtBlockDeleteOne) ExecX(ctx context.Context) {
	sjbdo.sjbd.ExecX(ctx)
}
