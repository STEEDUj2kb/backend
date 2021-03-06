// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/STEEDUj2kb/platform/ent/dailygaol"
	"github.com/STEEDUj2kb/platform/ent/predicate"
	"github.com/STEEDUj2kb/platform/ent/study"
	"github.com/STEEDUj2kb/platform/ent/weeklygaol"
)

// DailyGaolUpdate is the builder for updating DailyGaol entities.
type DailyGaolUpdate struct {
	config
	hooks    []Hook
	mutation *DailyGaolMutation
}

// Where appends a list predicates to the DailyGaolUpdate builder.
func (dgu *DailyGaolUpdate) Where(ps ...predicate.DailyGaol) *DailyGaolUpdate {
	dgu.mutation.Where(ps...)
	return dgu
}

// SetUpdatedAt sets the "updated_at" field.
func (dgu *DailyGaolUpdate) SetUpdatedAt(t time.Time) *DailyGaolUpdate {
	dgu.mutation.SetUpdatedAt(t)
	return dgu
}

// SetTodo sets the "todo" field.
func (dgu *DailyGaolUpdate) SetTodo(s string) *DailyGaolUpdate {
	dgu.mutation.SetTodo(s)
	return dgu
}

// SetDone sets the "done" field.
func (dgu *DailyGaolUpdate) SetDone(b bool) *DailyGaolUpdate {
	dgu.mutation.SetDone(b)
	return dgu
}

// SetNillableDone sets the "done" field if the given value is not nil.
func (dgu *DailyGaolUpdate) SetNillableDone(b *bool) *DailyGaolUpdate {
	if b != nil {
		dgu.SetDone(*b)
	}
	return dgu
}

// SetIsRemoved sets the "is_removed" field.
func (dgu *DailyGaolUpdate) SetIsRemoved(b bool) *DailyGaolUpdate {
	dgu.mutation.SetIsRemoved(b)
	return dgu
}

// SetNillableIsRemoved sets the "is_removed" field if the given value is not nil.
func (dgu *DailyGaolUpdate) SetNillableIsRemoved(b *bool) *DailyGaolUpdate {
	if b != nil {
		dgu.SetIsRemoved(*b)
	}
	return dgu
}

// SetStudyID sets the "study" edge to the Study entity by ID.
func (dgu *DailyGaolUpdate) SetStudyID(id int) *DailyGaolUpdate {
	dgu.mutation.SetStudyID(id)
	return dgu
}

// SetNillableStudyID sets the "study" edge to the Study entity by ID if the given value is not nil.
func (dgu *DailyGaolUpdate) SetNillableStudyID(id *int) *DailyGaolUpdate {
	if id != nil {
		dgu = dgu.SetStudyID(*id)
	}
	return dgu
}

// SetStudy sets the "study" edge to the Study entity.
func (dgu *DailyGaolUpdate) SetStudy(s *Study) *DailyGaolUpdate {
	return dgu.SetStudyID(s.ID)
}

// SetWgoalID sets the "wgoal" edge to the WeeklyGaol entity by ID.
func (dgu *DailyGaolUpdate) SetWgoalID(id int) *DailyGaolUpdate {
	dgu.mutation.SetWgoalID(id)
	return dgu
}

// SetNillableWgoalID sets the "wgoal" edge to the WeeklyGaol entity by ID if the given value is not nil.
func (dgu *DailyGaolUpdate) SetNillableWgoalID(id *int) *DailyGaolUpdate {
	if id != nil {
		dgu = dgu.SetWgoalID(*id)
	}
	return dgu
}

// SetWgoal sets the "wgoal" edge to the WeeklyGaol entity.
func (dgu *DailyGaolUpdate) SetWgoal(w *WeeklyGaol) *DailyGaolUpdate {
	return dgu.SetWgoalID(w.ID)
}

// Mutation returns the DailyGaolMutation object of the builder.
func (dgu *DailyGaolUpdate) Mutation() *DailyGaolMutation {
	return dgu.mutation
}

// ClearStudy clears the "study" edge to the Study entity.
func (dgu *DailyGaolUpdate) ClearStudy() *DailyGaolUpdate {
	dgu.mutation.ClearStudy()
	return dgu
}

// ClearWgoal clears the "wgoal" edge to the WeeklyGaol entity.
func (dgu *DailyGaolUpdate) ClearWgoal() *DailyGaolUpdate {
	dgu.mutation.ClearWgoal()
	return dgu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (dgu *DailyGaolUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	dgu.defaults()
	if len(dgu.hooks) == 0 {
		if err = dgu.check(); err != nil {
			return 0, err
		}
		affected, err = dgu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DailyGaolMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dgu.check(); err != nil {
				return 0, err
			}
			dgu.mutation = mutation
			affected, err = dgu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(dgu.hooks) - 1; i >= 0; i-- {
			if dgu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dgu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dgu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (dgu *DailyGaolUpdate) SaveX(ctx context.Context) int {
	affected, err := dgu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (dgu *DailyGaolUpdate) Exec(ctx context.Context) error {
	_, err := dgu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dgu *DailyGaolUpdate) ExecX(ctx context.Context) {
	if err := dgu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dgu *DailyGaolUpdate) defaults() {
	if _, ok := dgu.mutation.UpdatedAt(); !ok {
		v := dailygaol.UpdateDefaultUpdatedAt()
		dgu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dgu *DailyGaolUpdate) check() error {
	if v, ok := dgu.mutation.Todo(); ok {
		if err := dailygaol.TodoValidator(v); err != nil {
			return &ValidationError{Name: "todo", err: fmt.Errorf("ent: validator failed for field \"todo\": %w", err)}
		}
	}
	return nil
}

func (dgu *DailyGaolUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dailygaol.Table,
			Columns: dailygaol.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dailygaol.FieldID,
			},
		},
	}
	if ps := dgu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dgu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: dailygaol.FieldUpdatedAt,
		})
	}
	if value, ok := dgu.mutation.Todo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dailygaol.FieldTodo,
		})
	}
	if value, ok := dgu.mutation.Done(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: dailygaol.FieldDone,
		})
	}
	if value, ok := dgu.mutation.IsRemoved(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: dailygaol.FieldIsRemoved,
		})
	}
	if dgu.mutation.StudyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dailygaol.StudyTable,
			Columns: []string{dailygaol.StudyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: study.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dgu.mutation.StudyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dailygaol.StudyTable,
			Columns: []string{dailygaol.StudyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: study.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if dgu.mutation.WgoalCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dailygaol.WgoalTable,
			Columns: []string{dailygaol.WgoalColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: weeklygaol.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dgu.mutation.WgoalIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dailygaol.WgoalTable,
			Columns: []string{dailygaol.WgoalColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: weeklygaol.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, dgu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dailygaol.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// DailyGaolUpdateOne is the builder for updating a single DailyGaol entity.
type DailyGaolUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DailyGaolMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (dguo *DailyGaolUpdateOne) SetUpdatedAt(t time.Time) *DailyGaolUpdateOne {
	dguo.mutation.SetUpdatedAt(t)
	return dguo
}

// SetTodo sets the "todo" field.
func (dguo *DailyGaolUpdateOne) SetTodo(s string) *DailyGaolUpdateOne {
	dguo.mutation.SetTodo(s)
	return dguo
}

// SetDone sets the "done" field.
func (dguo *DailyGaolUpdateOne) SetDone(b bool) *DailyGaolUpdateOne {
	dguo.mutation.SetDone(b)
	return dguo
}

// SetNillableDone sets the "done" field if the given value is not nil.
func (dguo *DailyGaolUpdateOne) SetNillableDone(b *bool) *DailyGaolUpdateOne {
	if b != nil {
		dguo.SetDone(*b)
	}
	return dguo
}

// SetIsRemoved sets the "is_removed" field.
func (dguo *DailyGaolUpdateOne) SetIsRemoved(b bool) *DailyGaolUpdateOne {
	dguo.mutation.SetIsRemoved(b)
	return dguo
}

// SetNillableIsRemoved sets the "is_removed" field if the given value is not nil.
func (dguo *DailyGaolUpdateOne) SetNillableIsRemoved(b *bool) *DailyGaolUpdateOne {
	if b != nil {
		dguo.SetIsRemoved(*b)
	}
	return dguo
}

// SetStudyID sets the "study" edge to the Study entity by ID.
func (dguo *DailyGaolUpdateOne) SetStudyID(id int) *DailyGaolUpdateOne {
	dguo.mutation.SetStudyID(id)
	return dguo
}

// SetNillableStudyID sets the "study" edge to the Study entity by ID if the given value is not nil.
func (dguo *DailyGaolUpdateOne) SetNillableStudyID(id *int) *DailyGaolUpdateOne {
	if id != nil {
		dguo = dguo.SetStudyID(*id)
	}
	return dguo
}

// SetStudy sets the "study" edge to the Study entity.
func (dguo *DailyGaolUpdateOne) SetStudy(s *Study) *DailyGaolUpdateOne {
	return dguo.SetStudyID(s.ID)
}

// SetWgoalID sets the "wgoal" edge to the WeeklyGaol entity by ID.
func (dguo *DailyGaolUpdateOne) SetWgoalID(id int) *DailyGaolUpdateOne {
	dguo.mutation.SetWgoalID(id)
	return dguo
}

// SetNillableWgoalID sets the "wgoal" edge to the WeeklyGaol entity by ID if the given value is not nil.
func (dguo *DailyGaolUpdateOne) SetNillableWgoalID(id *int) *DailyGaolUpdateOne {
	if id != nil {
		dguo = dguo.SetWgoalID(*id)
	}
	return dguo
}

// SetWgoal sets the "wgoal" edge to the WeeklyGaol entity.
func (dguo *DailyGaolUpdateOne) SetWgoal(w *WeeklyGaol) *DailyGaolUpdateOne {
	return dguo.SetWgoalID(w.ID)
}

// Mutation returns the DailyGaolMutation object of the builder.
func (dguo *DailyGaolUpdateOne) Mutation() *DailyGaolMutation {
	return dguo.mutation
}

// ClearStudy clears the "study" edge to the Study entity.
func (dguo *DailyGaolUpdateOne) ClearStudy() *DailyGaolUpdateOne {
	dguo.mutation.ClearStudy()
	return dguo
}

// ClearWgoal clears the "wgoal" edge to the WeeklyGaol entity.
func (dguo *DailyGaolUpdateOne) ClearWgoal() *DailyGaolUpdateOne {
	dguo.mutation.ClearWgoal()
	return dguo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (dguo *DailyGaolUpdateOne) Select(field string, fields ...string) *DailyGaolUpdateOne {
	dguo.fields = append([]string{field}, fields...)
	return dguo
}

// Save executes the query and returns the updated DailyGaol entity.
func (dguo *DailyGaolUpdateOne) Save(ctx context.Context) (*DailyGaol, error) {
	var (
		err  error
		node *DailyGaol
	)
	dguo.defaults()
	if len(dguo.hooks) == 0 {
		if err = dguo.check(); err != nil {
			return nil, err
		}
		node, err = dguo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DailyGaolMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dguo.check(); err != nil {
				return nil, err
			}
			dguo.mutation = mutation
			node, err = dguo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dguo.hooks) - 1; i >= 0; i-- {
			if dguo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dguo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dguo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (dguo *DailyGaolUpdateOne) SaveX(ctx context.Context) *DailyGaol {
	node, err := dguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (dguo *DailyGaolUpdateOne) Exec(ctx context.Context) error {
	_, err := dguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dguo *DailyGaolUpdateOne) ExecX(ctx context.Context) {
	if err := dguo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dguo *DailyGaolUpdateOne) defaults() {
	if _, ok := dguo.mutation.UpdatedAt(); !ok {
		v := dailygaol.UpdateDefaultUpdatedAt()
		dguo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dguo *DailyGaolUpdateOne) check() error {
	if v, ok := dguo.mutation.Todo(); ok {
		if err := dailygaol.TodoValidator(v); err != nil {
			return &ValidationError{Name: "todo", err: fmt.Errorf("ent: validator failed for field \"todo\": %w", err)}
		}
	}
	return nil
}

func (dguo *DailyGaolUpdateOne) sqlSave(ctx context.Context) (_node *DailyGaol, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dailygaol.Table,
			Columns: dailygaol.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dailygaol.FieldID,
			},
		},
	}
	id, ok := dguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing DailyGaol.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := dguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dailygaol.FieldID)
		for _, f := range fields {
			if !dailygaol.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != dailygaol.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := dguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dguo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: dailygaol.FieldUpdatedAt,
		})
	}
	if value, ok := dguo.mutation.Todo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dailygaol.FieldTodo,
		})
	}
	if value, ok := dguo.mutation.Done(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: dailygaol.FieldDone,
		})
	}
	if value, ok := dguo.mutation.IsRemoved(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: dailygaol.FieldIsRemoved,
		})
	}
	if dguo.mutation.StudyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dailygaol.StudyTable,
			Columns: []string{dailygaol.StudyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: study.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dguo.mutation.StudyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dailygaol.StudyTable,
			Columns: []string{dailygaol.StudyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: study.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if dguo.mutation.WgoalCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dailygaol.WgoalTable,
			Columns: []string{dailygaol.WgoalColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: weeklygaol.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dguo.mutation.WgoalIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dailygaol.WgoalTable,
			Columns: []string{dailygaol.WgoalColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: weeklygaol.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &DailyGaol{config: dguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, dguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dailygaol.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
