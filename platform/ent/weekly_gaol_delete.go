// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/STEEDUj2kb/platform/ent/predicate"
	"github.com/STEEDUj2kb/platform/ent/weekly_gaol"
)

// WeeklyGaolDelete is the builder for deleting a Weekly_Gaol entity.
type WeeklyGaolDelete struct {
	config
	hooks    []Hook
	mutation *WeeklyGaolMutation
}

// Where appends a list predicates to the WeeklyGaolDelete builder.
func (wgd *WeeklyGaolDelete) Where(ps ...predicate.Weekly_Gaol) *WeeklyGaolDelete {
	wgd.mutation.Where(ps...)
	return wgd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wgd *WeeklyGaolDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wgd.hooks) == 0 {
		affected, err = wgd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WeeklyGaolMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wgd.mutation = mutation
			affected, err = wgd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wgd.hooks) - 1; i >= 0; i-- {
			if wgd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wgd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wgd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (wgd *WeeklyGaolDelete) ExecX(ctx context.Context) int {
	n, err := wgd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wgd *WeeklyGaolDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: weekly_gaol.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: weekly_gaol.FieldID,
			},
		},
	}
	if ps := wgd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, wgd.driver, _spec)
}

// WeeklyGaolDeleteOne is the builder for deleting a single Weekly_Gaol entity.
type WeeklyGaolDeleteOne struct {
	wgd *WeeklyGaolDelete
}

// Exec executes the deletion query.
func (wgdo *WeeklyGaolDeleteOne) Exec(ctx context.Context) error {
	n, err := wgdo.wgd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{weekly_gaol.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wgdo *WeeklyGaolDeleteOne) ExecX(ctx context.Context) {
	wgdo.wgd.ExecX(ctx)
}
