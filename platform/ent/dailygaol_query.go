// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/STEEDUj2kb/platform/ent/dailygaol"
	"github.com/STEEDUj2kb/platform/ent/predicate"
	"github.com/STEEDUj2kb/platform/ent/study"
	"github.com/STEEDUj2kb/platform/ent/weeklygaol"
)

// DailyGaolQuery is the builder for querying DailyGaol entities.
type DailyGaolQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.DailyGaol
	// eager-loading edges.
	withStudy *StudyQuery
	withWgoal *WeeklyGaolQuery
	withFKs   bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DailyGaolQuery builder.
func (dgq *DailyGaolQuery) Where(ps ...predicate.DailyGaol) *DailyGaolQuery {
	dgq.predicates = append(dgq.predicates, ps...)
	return dgq
}

// Limit adds a limit step to the query.
func (dgq *DailyGaolQuery) Limit(limit int) *DailyGaolQuery {
	dgq.limit = &limit
	return dgq
}

// Offset adds an offset step to the query.
func (dgq *DailyGaolQuery) Offset(offset int) *DailyGaolQuery {
	dgq.offset = &offset
	return dgq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dgq *DailyGaolQuery) Unique(unique bool) *DailyGaolQuery {
	dgq.unique = &unique
	return dgq
}

// Order adds an order step to the query.
func (dgq *DailyGaolQuery) Order(o ...OrderFunc) *DailyGaolQuery {
	dgq.order = append(dgq.order, o...)
	return dgq
}

// QueryStudy chains the current query on the "study" edge.
func (dgq *DailyGaolQuery) QueryStudy() *StudyQuery {
	query := &StudyQuery{config: dgq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dgq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dailygaol.Table, dailygaol.FieldID, selector),
			sqlgraph.To(study.Table, study.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, dailygaol.StudyTable, dailygaol.StudyColumn),
		)
		fromU = sqlgraph.SetNeighbors(dgq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWgoal chains the current query on the "wgoal" edge.
func (dgq *DailyGaolQuery) QueryWgoal() *WeeklyGaolQuery {
	query := &WeeklyGaolQuery{config: dgq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dgq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dailygaol.Table, dailygaol.FieldID, selector),
			sqlgraph.To(weeklygaol.Table, weeklygaol.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, dailygaol.WgoalTable, dailygaol.WgoalColumn),
		)
		fromU = sqlgraph.SetNeighbors(dgq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DailyGaol entity from the query.
// Returns a *NotFoundError when no DailyGaol was found.
func (dgq *DailyGaolQuery) First(ctx context.Context) (*DailyGaol, error) {
	nodes, err := dgq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{dailygaol.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dgq *DailyGaolQuery) FirstX(ctx context.Context) *DailyGaol {
	node, err := dgq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DailyGaol ID from the query.
// Returns a *NotFoundError when no DailyGaol ID was found.
func (dgq *DailyGaolQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dgq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{dailygaol.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dgq *DailyGaolQuery) FirstIDX(ctx context.Context) int {
	id, err := dgq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DailyGaol entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one DailyGaol entity is not found.
// Returns a *NotFoundError when no DailyGaol entities are found.
func (dgq *DailyGaolQuery) Only(ctx context.Context) (*DailyGaol, error) {
	nodes, err := dgq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{dailygaol.Label}
	default:
		return nil, &NotSingularError{dailygaol.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dgq *DailyGaolQuery) OnlyX(ctx context.Context) *DailyGaol {
	node, err := dgq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DailyGaol ID in the query.
// Returns a *NotSingularError when exactly one DailyGaol ID is not found.
// Returns a *NotFoundError when no entities are found.
func (dgq *DailyGaolQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dgq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{dailygaol.Label}
	default:
		err = &NotSingularError{dailygaol.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dgq *DailyGaolQuery) OnlyIDX(ctx context.Context) int {
	id, err := dgq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DailyGaols.
func (dgq *DailyGaolQuery) All(ctx context.Context) ([]*DailyGaol, error) {
	if err := dgq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return dgq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (dgq *DailyGaolQuery) AllX(ctx context.Context) []*DailyGaol {
	nodes, err := dgq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DailyGaol IDs.
func (dgq *DailyGaolQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := dgq.Select(dailygaol.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dgq *DailyGaolQuery) IDsX(ctx context.Context) []int {
	ids, err := dgq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dgq *DailyGaolQuery) Count(ctx context.Context) (int, error) {
	if err := dgq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return dgq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (dgq *DailyGaolQuery) CountX(ctx context.Context) int {
	count, err := dgq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dgq *DailyGaolQuery) Exist(ctx context.Context) (bool, error) {
	if err := dgq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return dgq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (dgq *DailyGaolQuery) ExistX(ctx context.Context) bool {
	exist, err := dgq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DailyGaolQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dgq *DailyGaolQuery) Clone() *DailyGaolQuery {
	if dgq == nil {
		return nil
	}
	return &DailyGaolQuery{
		config:     dgq.config,
		limit:      dgq.limit,
		offset:     dgq.offset,
		order:      append([]OrderFunc{}, dgq.order...),
		predicates: append([]predicate.DailyGaol{}, dgq.predicates...),
		withStudy:  dgq.withStudy.Clone(),
		withWgoal:  dgq.withWgoal.Clone(),
		// clone intermediate query.
		sql:  dgq.sql.Clone(),
		path: dgq.path,
	}
}

// WithStudy tells the query-builder to eager-load the nodes that are connected to
// the "study" edge. The optional arguments are used to configure the query builder of the edge.
func (dgq *DailyGaolQuery) WithStudy(opts ...func(*StudyQuery)) *DailyGaolQuery {
	query := &StudyQuery{config: dgq.config}
	for _, opt := range opts {
		opt(query)
	}
	dgq.withStudy = query
	return dgq
}

// WithWgoal tells the query-builder to eager-load the nodes that are connected to
// the "wgoal" edge. The optional arguments are used to configure the query builder of the edge.
func (dgq *DailyGaolQuery) WithWgoal(opts ...func(*WeeklyGaolQuery)) *DailyGaolQuery {
	query := &WeeklyGaolQuery{config: dgq.config}
	for _, opt := range opts {
		opt(query)
	}
	dgq.withWgoal = query
	return dgq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DailyGaol.Query().
//		GroupBy(dailygaol.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (dgq *DailyGaolQuery) GroupBy(field string, fields ...string) *DailyGaolGroupBy {
	group := &DailyGaolGroupBy{config: dgq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dgq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.DailyGaol.Query().
//		Select(dailygaol.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (dgq *DailyGaolQuery) Select(fields ...string) *DailyGaolSelect {
	dgq.fields = append(dgq.fields, fields...)
	return &DailyGaolSelect{DailyGaolQuery: dgq}
}

func (dgq *DailyGaolQuery) prepareQuery(ctx context.Context) error {
	for _, f := range dgq.fields {
		if !dailygaol.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dgq.path != nil {
		prev, err := dgq.path(ctx)
		if err != nil {
			return err
		}
		dgq.sql = prev
	}
	return nil
}

func (dgq *DailyGaolQuery) sqlAll(ctx context.Context) ([]*DailyGaol, error) {
	var (
		nodes       = []*DailyGaol{}
		withFKs     = dgq.withFKs
		_spec       = dgq.querySpec()
		loadedTypes = [2]bool{
			dgq.withStudy != nil,
			dgq.withWgoal != nil,
		}
	)
	if dgq.withStudy != nil || dgq.withWgoal != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, dailygaol.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &DailyGaol{config: dgq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, dgq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := dgq.withStudy; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*DailyGaol)
		for i := range nodes {
			if nodes[i].study_dgoals == nil {
				continue
			}
			fk := *nodes[i].study_dgoals
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(study.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "study_dgoals" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Study = n
			}
		}
	}

	if query := dgq.withWgoal; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*DailyGaol)
		for i := range nodes {
			if nodes[i].weekly_gaol_dgaols == nil {
				continue
			}
			fk := *nodes[i].weekly_gaol_dgaols
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(weeklygaol.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "weekly_gaol_dgaols" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Wgoal = n
			}
		}
	}

	return nodes, nil
}

func (dgq *DailyGaolQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dgq.querySpec()
	return sqlgraph.CountNodes(ctx, dgq.driver, _spec)
}

func (dgq *DailyGaolQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := dgq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (dgq *DailyGaolQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dailygaol.Table,
			Columns: dailygaol.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dailygaol.FieldID,
			},
		},
		From:   dgq.sql,
		Unique: true,
	}
	if unique := dgq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := dgq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dailygaol.FieldID)
		for i := range fields {
			if fields[i] != dailygaol.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dgq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dgq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dgq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dgq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dgq *DailyGaolQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dgq.driver.Dialect())
	t1 := builder.Table(dailygaol.Table)
	columns := dgq.fields
	if len(columns) == 0 {
		columns = dailygaol.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dgq.sql != nil {
		selector = dgq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range dgq.predicates {
		p(selector)
	}
	for _, p := range dgq.order {
		p(selector)
	}
	if offset := dgq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dgq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DailyGaolGroupBy is the group-by builder for DailyGaol entities.
type DailyGaolGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dggb *DailyGaolGroupBy) Aggregate(fns ...AggregateFunc) *DailyGaolGroupBy {
	dggb.fns = append(dggb.fns, fns...)
	return dggb
}

// Scan applies the group-by query and scans the result into the given value.
func (dggb *DailyGaolGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := dggb.path(ctx)
	if err != nil {
		return err
	}
	dggb.sql = query
	return dggb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (dggb *DailyGaolGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := dggb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (dggb *DailyGaolGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(dggb.fields) > 1 {
		return nil, errors.New("ent: DailyGaolGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := dggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (dggb *DailyGaolGroupBy) StringsX(ctx context.Context) []string {
	v, err := dggb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dggb *DailyGaolGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = dggb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dailygaol.Label}
	default:
		err = fmt.Errorf("ent: DailyGaolGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (dggb *DailyGaolGroupBy) StringX(ctx context.Context) string {
	v, err := dggb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (dggb *DailyGaolGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(dggb.fields) > 1 {
		return nil, errors.New("ent: DailyGaolGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := dggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (dggb *DailyGaolGroupBy) IntsX(ctx context.Context) []int {
	v, err := dggb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dggb *DailyGaolGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = dggb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dailygaol.Label}
	default:
		err = fmt.Errorf("ent: DailyGaolGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (dggb *DailyGaolGroupBy) IntX(ctx context.Context) int {
	v, err := dggb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (dggb *DailyGaolGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(dggb.fields) > 1 {
		return nil, errors.New("ent: DailyGaolGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := dggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (dggb *DailyGaolGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := dggb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dggb *DailyGaolGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = dggb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dailygaol.Label}
	default:
		err = fmt.Errorf("ent: DailyGaolGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (dggb *DailyGaolGroupBy) Float64X(ctx context.Context) float64 {
	v, err := dggb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (dggb *DailyGaolGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(dggb.fields) > 1 {
		return nil, errors.New("ent: DailyGaolGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := dggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (dggb *DailyGaolGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := dggb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dggb *DailyGaolGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = dggb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dailygaol.Label}
	default:
		err = fmt.Errorf("ent: DailyGaolGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (dggb *DailyGaolGroupBy) BoolX(ctx context.Context) bool {
	v, err := dggb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dggb *DailyGaolGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range dggb.fields {
		if !dailygaol.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := dggb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dggb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dggb *DailyGaolGroupBy) sqlQuery() *sql.Selector {
	selector := dggb.sql.Select()
	aggregation := make([]string, 0, len(dggb.fns))
	for _, fn := range dggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(dggb.fields)+len(dggb.fns))
		for _, f := range dggb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(dggb.fields...)...)
}

// DailyGaolSelect is the builder for selecting fields of DailyGaol entities.
type DailyGaolSelect struct {
	*DailyGaolQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (dgs *DailyGaolSelect) Scan(ctx context.Context, v interface{}) error {
	if err := dgs.prepareQuery(ctx); err != nil {
		return err
	}
	dgs.sql = dgs.DailyGaolQuery.sqlQuery(ctx)
	return dgs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (dgs *DailyGaolSelect) ScanX(ctx context.Context, v interface{}) {
	if err := dgs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (dgs *DailyGaolSelect) Strings(ctx context.Context) ([]string, error) {
	if len(dgs.fields) > 1 {
		return nil, errors.New("ent: DailyGaolSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := dgs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (dgs *DailyGaolSelect) StringsX(ctx context.Context) []string {
	v, err := dgs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (dgs *DailyGaolSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = dgs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dailygaol.Label}
	default:
		err = fmt.Errorf("ent: DailyGaolSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (dgs *DailyGaolSelect) StringX(ctx context.Context) string {
	v, err := dgs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (dgs *DailyGaolSelect) Ints(ctx context.Context) ([]int, error) {
	if len(dgs.fields) > 1 {
		return nil, errors.New("ent: DailyGaolSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := dgs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (dgs *DailyGaolSelect) IntsX(ctx context.Context) []int {
	v, err := dgs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (dgs *DailyGaolSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = dgs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dailygaol.Label}
	default:
		err = fmt.Errorf("ent: DailyGaolSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (dgs *DailyGaolSelect) IntX(ctx context.Context) int {
	v, err := dgs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (dgs *DailyGaolSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(dgs.fields) > 1 {
		return nil, errors.New("ent: DailyGaolSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := dgs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (dgs *DailyGaolSelect) Float64sX(ctx context.Context) []float64 {
	v, err := dgs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (dgs *DailyGaolSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = dgs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dailygaol.Label}
	default:
		err = fmt.Errorf("ent: DailyGaolSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (dgs *DailyGaolSelect) Float64X(ctx context.Context) float64 {
	v, err := dgs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (dgs *DailyGaolSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(dgs.fields) > 1 {
		return nil, errors.New("ent: DailyGaolSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := dgs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (dgs *DailyGaolSelect) BoolsX(ctx context.Context) []bool {
	v, err := dgs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (dgs *DailyGaolSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = dgs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dailygaol.Label}
	default:
		err = fmt.Errorf("ent: DailyGaolSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (dgs *DailyGaolSelect) BoolX(ctx context.Context) bool {
	v, err := dgs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dgs *DailyGaolSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := dgs.sql.Query()
	if err := dgs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
