// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/0B1t322/CP-Rosseti-Back/ent/predicate"
	"github.com/0B1t322/CP-Rosseti-Back/ent/question"
	"github.com/0B1t322/CP-Rosseti-Back/ent/test"
	"github.com/0B1t322/CP-Rosseti-Back/ent/theoreticaltest"
)

// TheoreticalTestQuery is the builder for querying TheoreticalTest entities.
type TheoreticalTestQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.TheoreticalTest
	// eager-loading edges.
	withTest     *TestQuery
	withQuestion *QuestionQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TheoreticalTestQuery builder.
func (ttq *TheoreticalTestQuery) Where(ps ...predicate.TheoreticalTest) *TheoreticalTestQuery {
	ttq.predicates = append(ttq.predicates, ps...)
	return ttq
}

// Limit adds a limit step to the query.
func (ttq *TheoreticalTestQuery) Limit(limit int) *TheoreticalTestQuery {
	ttq.limit = &limit
	return ttq
}

// Offset adds an offset step to the query.
func (ttq *TheoreticalTestQuery) Offset(offset int) *TheoreticalTestQuery {
	ttq.offset = &offset
	return ttq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ttq *TheoreticalTestQuery) Unique(unique bool) *TheoreticalTestQuery {
	ttq.unique = &unique
	return ttq
}

// Order adds an order step to the query.
func (ttq *TheoreticalTestQuery) Order(o ...OrderFunc) *TheoreticalTestQuery {
	ttq.order = append(ttq.order, o...)
	return ttq
}

// QueryTest chains the current query on the "Test" edge.
func (ttq *TheoreticalTestQuery) QueryTest() *TestQuery {
	query := &TestQuery{config: ttq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ttq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ttq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(theoreticaltest.Table, theoreticaltest.FieldID, selector),
			sqlgraph.To(test.Table, test.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, theoreticaltest.TestTable, theoreticaltest.TestColumn),
		)
		fromU = sqlgraph.SetNeighbors(ttq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryQuestion chains the current query on the "Question" edge.
func (ttq *TheoreticalTestQuery) QueryQuestion() *QuestionQuery {
	query := &QuestionQuery{config: ttq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ttq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ttq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(theoreticaltest.Table, theoreticaltest.FieldID, selector),
			sqlgraph.To(question.Table, question.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, theoreticaltest.QuestionTable, theoreticaltest.QuestionColumn),
		)
		fromU = sqlgraph.SetNeighbors(ttq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TheoreticalTest entity from the query.
// Returns a *NotFoundError when no TheoreticalTest was found.
func (ttq *TheoreticalTestQuery) First(ctx context.Context) (*TheoreticalTest, error) {
	nodes, err := ttq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{theoreticaltest.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ttq *TheoreticalTestQuery) FirstX(ctx context.Context) *TheoreticalTest {
	node, err := ttq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TheoreticalTest ID from the query.
// Returns a *NotFoundError when no TheoreticalTest ID was found.
func (ttq *TheoreticalTestQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ttq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{theoreticaltest.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ttq *TheoreticalTestQuery) FirstIDX(ctx context.Context) int {
	id, err := ttq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TheoreticalTest entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one TheoreticalTest entity is not found.
// Returns a *NotFoundError when no TheoreticalTest entities are found.
func (ttq *TheoreticalTestQuery) Only(ctx context.Context) (*TheoreticalTest, error) {
	nodes, err := ttq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{theoreticaltest.Label}
	default:
		return nil, &NotSingularError{theoreticaltest.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ttq *TheoreticalTestQuery) OnlyX(ctx context.Context) *TheoreticalTest {
	node, err := ttq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TheoreticalTest ID in the query.
// Returns a *NotSingularError when exactly one TheoreticalTest ID is not found.
// Returns a *NotFoundError when no entities are found.
func (ttq *TheoreticalTestQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ttq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{theoreticaltest.Label}
	default:
		err = &NotSingularError{theoreticaltest.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ttq *TheoreticalTestQuery) OnlyIDX(ctx context.Context) int {
	id, err := ttq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TheoreticalTests.
func (ttq *TheoreticalTestQuery) All(ctx context.Context) ([]*TheoreticalTest, error) {
	if err := ttq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ttq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ttq *TheoreticalTestQuery) AllX(ctx context.Context) []*TheoreticalTest {
	nodes, err := ttq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TheoreticalTest IDs.
func (ttq *TheoreticalTestQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ttq.Select(theoreticaltest.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ttq *TheoreticalTestQuery) IDsX(ctx context.Context) []int {
	ids, err := ttq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ttq *TheoreticalTestQuery) Count(ctx context.Context) (int, error) {
	if err := ttq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ttq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ttq *TheoreticalTestQuery) CountX(ctx context.Context) int {
	count, err := ttq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ttq *TheoreticalTestQuery) Exist(ctx context.Context) (bool, error) {
	if err := ttq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ttq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ttq *TheoreticalTestQuery) ExistX(ctx context.Context) bool {
	exist, err := ttq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TheoreticalTestQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ttq *TheoreticalTestQuery) Clone() *TheoreticalTestQuery {
	if ttq == nil {
		return nil
	}
	return &TheoreticalTestQuery{
		config:       ttq.config,
		limit:        ttq.limit,
		offset:       ttq.offset,
		order:        append([]OrderFunc{}, ttq.order...),
		predicates:   append([]predicate.TheoreticalTest{}, ttq.predicates...),
		withTest:     ttq.withTest.Clone(),
		withQuestion: ttq.withQuestion.Clone(),
		// clone intermediate query.
		sql:  ttq.sql.Clone(),
		path: ttq.path,
	}
}

// WithTest tells the query-builder to eager-load the nodes that are connected to
// the "Test" edge. The optional arguments are used to configure the query builder of the edge.
func (ttq *TheoreticalTestQuery) WithTest(opts ...func(*TestQuery)) *TheoreticalTestQuery {
	query := &TestQuery{config: ttq.config}
	for _, opt := range opts {
		opt(query)
	}
	ttq.withTest = query
	return ttq
}

// WithQuestion tells the query-builder to eager-load the nodes that are connected to
// the "Question" edge. The optional arguments are used to configure the query builder of the edge.
func (ttq *TheoreticalTestQuery) WithQuestion(opts ...func(*QuestionQuery)) *TheoreticalTestQuery {
	query := &QuestionQuery{config: ttq.config}
	for _, opt := range opts {
		opt(query)
	}
	ttq.withQuestion = query
	return ttq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		TestID int `json:"test_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TheoreticalTest.Query().
//		GroupBy(theoreticaltest.FieldTestID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ttq *TheoreticalTestQuery) GroupBy(field string, fields ...string) *TheoreticalTestGroupBy {
	group := &TheoreticalTestGroupBy{config: ttq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ttq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ttq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		TestID int `json:"test_id,omitempty"`
//	}
//
//	client.TheoreticalTest.Query().
//		Select(theoreticaltest.FieldTestID).
//		Scan(ctx, &v)
//
func (ttq *TheoreticalTestQuery) Select(fields ...string) *TheoreticalTestSelect {
	ttq.fields = append(ttq.fields, fields...)
	return &TheoreticalTestSelect{TheoreticalTestQuery: ttq}
}

func (ttq *TheoreticalTestQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ttq.fields {
		if !theoreticaltest.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ttq.path != nil {
		prev, err := ttq.path(ctx)
		if err != nil {
			return err
		}
		ttq.sql = prev
	}
	return nil
}

func (ttq *TheoreticalTestQuery) sqlAll(ctx context.Context) ([]*TheoreticalTest, error) {
	var (
		nodes       = []*TheoreticalTest{}
		_spec       = ttq.querySpec()
		loadedTypes = [2]bool{
			ttq.withTest != nil,
			ttq.withQuestion != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &TheoreticalTest{config: ttq.config}
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
	if err := sqlgraph.QueryNodes(ctx, ttq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ttq.withTest; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*TheoreticalTest)
		for i := range nodes {
			fk := nodes[i].TestID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(test.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "test_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Test = n
			}
		}
	}

	if query := ttq.withQuestion; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*TheoreticalTest)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Question = []*Question{}
		}
		query.Where(predicate.Question(func(s *sql.Selector) {
			s.Where(sql.InValues(theoreticaltest.QuestionColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.TheoricalTestID
			node, ok := nodeids[fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "theorical_test_id" returned %v for node %v`, fk, n.ID)
			}
			node.Edges.Question = append(node.Edges.Question, n)
		}
	}

	return nodes, nil
}

func (ttq *TheoreticalTestQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ttq.querySpec()
	return sqlgraph.CountNodes(ctx, ttq.driver, _spec)
}

func (ttq *TheoreticalTestQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ttq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ttq *TheoreticalTestQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   theoreticaltest.Table,
			Columns: theoreticaltest.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: theoreticaltest.FieldID,
			},
		},
		From:   ttq.sql,
		Unique: true,
	}
	if unique := ttq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ttq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, theoreticaltest.FieldID)
		for i := range fields {
			if fields[i] != theoreticaltest.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ttq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ttq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ttq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ttq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ttq *TheoreticalTestQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ttq.driver.Dialect())
	t1 := builder.Table(theoreticaltest.Table)
	columns := ttq.fields
	if len(columns) == 0 {
		columns = theoreticaltest.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ttq.sql != nil {
		selector = ttq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range ttq.predicates {
		p(selector)
	}
	for _, p := range ttq.order {
		p(selector)
	}
	if offset := ttq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ttq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TheoreticalTestGroupBy is the group-by builder for TheoreticalTest entities.
type TheoreticalTestGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ttgb *TheoreticalTestGroupBy) Aggregate(fns ...AggregateFunc) *TheoreticalTestGroupBy {
	ttgb.fns = append(ttgb.fns, fns...)
	return ttgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ttgb *TheoreticalTestGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ttgb.path(ctx)
	if err != nil {
		return err
	}
	ttgb.sql = query
	return ttgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ttgb *TheoreticalTestGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ttgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ttgb *TheoreticalTestGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ttgb.fields) > 1 {
		return nil, errors.New("ent: TheoreticalTestGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ttgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ttgb *TheoreticalTestGroupBy) StringsX(ctx context.Context) []string {
	v, err := ttgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ttgb *TheoreticalTestGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ttgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{theoreticaltest.Label}
	default:
		err = fmt.Errorf("ent: TheoreticalTestGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ttgb *TheoreticalTestGroupBy) StringX(ctx context.Context) string {
	v, err := ttgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ttgb *TheoreticalTestGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ttgb.fields) > 1 {
		return nil, errors.New("ent: TheoreticalTestGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ttgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ttgb *TheoreticalTestGroupBy) IntsX(ctx context.Context) []int {
	v, err := ttgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ttgb *TheoreticalTestGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ttgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{theoreticaltest.Label}
	default:
		err = fmt.Errorf("ent: TheoreticalTestGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ttgb *TheoreticalTestGroupBy) IntX(ctx context.Context) int {
	v, err := ttgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ttgb *TheoreticalTestGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ttgb.fields) > 1 {
		return nil, errors.New("ent: TheoreticalTestGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ttgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ttgb *TheoreticalTestGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ttgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ttgb *TheoreticalTestGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ttgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{theoreticaltest.Label}
	default:
		err = fmt.Errorf("ent: TheoreticalTestGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ttgb *TheoreticalTestGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ttgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ttgb *TheoreticalTestGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ttgb.fields) > 1 {
		return nil, errors.New("ent: TheoreticalTestGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ttgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ttgb *TheoreticalTestGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ttgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ttgb *TheoreticalTestGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ttgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{theoreticaltest.Label}
	default:
		err = fmt.Errorf("ent: TheoreticalTestGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ttgb *TheoreticalTestGroupBy) BoolX(ctx context.Context) bool {
	v, err := ttgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ttgb *TheoreticalTestGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ttgb.fields {
		if !theoreticaltest.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ttgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ttgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ttgb *TheoreticalTestGroupBy) sqlQuery() *sql.Selector {
	selector := ttgb.sql.Select()
	aggregation := make([]string, 0, len(ttgb.fns))
	for _, fn := range ttgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ttgb.fields)+len(ttgb.fns))
		for _, f := range ttgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ttgb.fields...)...)
}

// TheoreticalTestSelect is the builder for selecting fields of TheoreticalTest entities.
type TheoreticalTestSelect struct {
	*TheoreticalTestQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (tts *TheoreticalTestSelect) Scan(ctx context.Context, v interface{}) error {
	if err := tts.prepareQuery(ctx); err != nil {
		return err
	}
	tts.sql = tts.TheoreticalTestQuery.sqlQuery(ctx)
	return tts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tts *TheoreticalTestSelect) ScanX(ctx context.Context, v interface{}) {
	if err := tts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (tts *TheoreticalTestSelect) Strings(ctx context.Context) ([]string, error) {
	if len(tts.fields) > 1 {
		return nil, errors.New("ent: TheoreticalTestSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := tts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tts *TheoreticalTestSelect) StringsX(ctx context.Context) []string {
	v, err := tts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (tts *TheoreticalTestSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tts.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{theoreticaltest.Label}
	default:
		err = fmt.Errorf("ent: TheoreticalTestSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tts *TheoreticalTestSelect) StringX(ctx context.Context) string {
	v, err := tts.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (tts *TheoreticalTestSelect) Ints(ctx context.Context) ([]int, error) {
	if len(tts.fields) > 1 {
		return nil, errors.New("ent: TheoreticalTestSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := tts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tts *TheoreticalTestSelect) IntsX(ctx context.Context) []int {
	v, err := tts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (tts *TheoreticalTestSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tts.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{theoreticaltest.Label}
	default:
		err = fmt.Errorf("ent: TheoreticalTestSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tts *TheoreticalTestSelect) IntX(ctx context.Context) int {
	v, err := tts.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (tts *TheoreticalTestSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(tts.fields) > 1 {
		return nil, errors.New("ent: TheoreticalTestSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := tts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tts *TheoreticalTestSelect) Float64sX(ctx context.Context) []float64 {
	v, err := tts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (tts *TheoreticalTestSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tts.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{theoreticaltest.Label}
	default:
		err = fmt.Errorf("ent: TheoreticalTestSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tts *TheoreticalTestSelect) Float64X(ctx context.Context) float64 {
	v, err := tts.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (tts *TheoreticalTestSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(tts.fields) > 1 {
		return nil, errors.New("ent: TheoreticalTestSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := tts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tts *TheoreticalTestSelect) BoolsX(ctx context.Context) []bool {
	v, err := tts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (tts *TheoreticalTestSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tts.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{theoreticaltest.Label}
	default:
		err = fmt.Errorf("ent: TheoreticalTestSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tts *TheoreticalTestSelect) BoolX(ctx context.Context) bool {
	v, err := tts.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tts *TheoreticalTestSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := tts.sql.Query()
	if err := tts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
