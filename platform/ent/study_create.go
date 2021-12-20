// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/STEEDUj2kb/platform/ent/daily_gaol"
	"github.com/STEEDUj2kb/platform/ent/study"
	"github.com/STEEDUj2kb/platform/ent/user"
	"github.com/STEEDUj2kb/platform/ent/weekly_gaol"
)

// StudyCreate is the builder for creating a Study entity.
type StudyCreate struct {
	config
	mutation *StudyMutation
	hooks    []Hook
}

// SetStartedAt sets the "started_at" field.
func (sc *StudyCreate) SetStartedAt(t time.Time) *StudyCreate {
	sc.mutation.SetStartedAt(t)
	return sc
}

// SetNillableStartedAt sets the "started_at" field if the given value is not nil.
func (sc *StudyCreate) SetNillableStartedAt(t *time.Time) *StudyCreate {
	if t != nil {
		sc.SetStartedAt(*t)
	}
	return sc
}

// SetEndedAt sets the "ended_at" field.
func (sc *StudyCreate) SetEndedAt(t time.Time) *StudyCreate {
	sc.mutation.SetEndedAt(t)
	return sc
}

// SetNillableEndedAt sets the "ended_at" field if the given value is not nil.
func (sc *StudyCreate) SetNillableEndedAt(t *time.Time) *StudyCreate {
	if t != nil {
		sc.SetEndedAt(*t)
	}
	return sc
}

// SetContent sets the "content" field.
func (sc *StudyCreate) SetContent(s string) *StudyCreate {
	sc.mutation.SetContent(s)
	return sc
}

// SetPlannerID sets the "planner" edge to the User entity by ID.
func (sc *StudyCreate) SetPlannerID(id int) *StudyCreate {
	sc.mutation.SetPlannerID(id)
	return sc
}

// SetNillablePlannerID sets the "planner" edge to the User entity by ID if the given value is not nil.
func (sc *StudyCreate) SetNillablePlannerID(id *int) *StudyCreate {
	if id != nil {
		sc = sc.SetPlannerID(*id)
	}
	return sc
}

// SetPlanner sets the "planner" edge to the User entity.
func (sc *StudyCreate) SetPlanner(u *User) *StudyCreate {
	return sc.SetPlannerID(u.ID)
}

// AddDgoalIDs adds the "dgoals" edge to the Daily_Gaol entity by IDs.
func (sc *StudyCreate) AddDgoalIDs(ids ...int) *StudyCreate {
	sc.mutation.AddDgoalIDs(ids...)
	return sc
}

// AddDgoals adds the "dgoals" edges to the Daily_Gaol entity.
func (sc *StudyCreate) AddDgoals(d ...*Daily_Gaol) *StudyCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return sc.AddDgoalIDs(ids...)
}

// AddWgoalIDs adds the "wgoals" edge to the Weekly_Gaol entity by IDs.
func (sc *StudyCreate) AddWgoalIDs(ids ...int) *StudyCreate {
	sc.mutation.AddWgoalIDs(ids...)
	return sc
}

// AddWgoals adds the "wgoals" edges to the Weekly_Gaol entity.
func (sc *StudyCreate) AddWgoals(w ...*Weekly_Gaol) *StudyCreate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return sc.AddWgoalIDs(ids...)
}

// Mutation returns the StudyMutation object of the builder.
func (sc *StudyCreate) Mutation() *StudyMutation {
	return sc.mutation
}

// Save creates the Study in the database.
func (sc *StudyCreate) Save(ctx context.Context) (*Study, error) {
	var (
		err  error
		node *Study
	)
	sc.defaults()
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StudyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StudyCreate) SaveX(ctx context.Context) *Study {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *StudyCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *StudyCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *StudyCreate) defaults() {
	if _, ok := sc.mutation.StartedAt(); !ok {
		v := study.DefaultStartedAt()
		sc.mutation.SetStartedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *StudyCreate) check() error {
	if _, ok := sc.mutation.StartedAt(); !ok {
		return &ValidationError{Name: "started_at", err: errors.New(`ent: missing required field "started_at"`)}
	}
	if _, ok := sc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "content"`)}
	}
	if v, ok := sc.mutation.Content(); ok {
		if err := study.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "content": %w`, err)}
		}
	}
	return nil
}

func (sc *StudyCreate) sqlSave(ctx context.Context) (*Study, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (sc *StudyCreate) createSpec() (*Study, *sqlgraph.CreateSpec) {
	var (
		_node = &Study{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: study.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: study.FieldID,
			},
		}
	)
	if value, ok := sc.mutation.StartedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: study.FieldStartedAt,
		})
		_node.StartedAt = value
	}
	if value, ok := sc.mutation.EndedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: study.FieldEndedAt,
		})
		_node.EndedAt = value
	}
	if value, ok := sc.mutation.Content(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: study.FieldContent,
		})
		_node.Content = value
	}
	if nodes := sc.mutation.PlannerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   study.PlannerTable,
			Columns: []string{study.PlannerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_studies = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.DgoalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   study.DgoalsTable,
			Columns: []string{study.DgoalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: daily_gaol.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.WgoalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   study.WgoalsTable,
			Columns: []string{study.WgoalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: weekly_gaol.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// StudyCreateBulk is the builder for creating many Study entities in bulk.
type StudyCreateBulk struct {
	config
	builders []*StudyCreate
}

// Save creates the Study entities in the database.
func (scb *StudyCreateBulk) Save(ctx context.Context) ([]*Study, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Study, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StudyMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *StudyCreateBulk) SaveX(ctx context.Context) []*Study {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *StudyCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *StudyCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
