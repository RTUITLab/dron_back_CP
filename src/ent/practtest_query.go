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
	"github.com/0B1t322/CP-Rosseti-Back/ent/practtest"
	"github.com/0B1t322/CP-Rosseti-Back/ent/predicate"
	"github.com/0B1t322/CP-Rosseti-Back/ent/submoduletest"
)

// PractTestQuery is the builder for querying PractTest entities.
type PractTestQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.PractTest
	// eager-loading edges.
	withSubModuleTest *SubModuleTestQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PractTestQuery builder.
func (ptq *PractTestQuery) Where(ps ...predicate.PractTest) *PractTestQuery {
	ptq.predicates = append(ptq.predicates, ps...)
	return ptq
}

// Limit adds a limit step to the query.
func (ptq *PractTestQuery) Limit(limit int) *PractTestQuery {
	ptq.limit = &limit
	return ptq
}

// Offset adds an offset step to the query.
func (ptq *PractTestQuery) Offset(offset int) *PractTestQuery {
	ptq.offset = &offset
	return ptq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ptq *PractTestQuery) Unique(unique bool) *PractTestQuery {
	ptq.unique = &unique
	return ptq
}

// Order adds an order step to the query.
func (ptq *PractTestQuery) Order(o ...OrderFunc) *PractTestQuery {
	ptq.order = append(ptq.order, o...)
	return ptq
}

// QuerySubModuleTest chains the current query on the "SubModuleTest" edge.
func (ptq *PractTestQuery) QuerySubModuleTest() *SubModuleTestQuery {
	query := &SubModuleTestQuery{config: ptq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ptq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ptq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(practtest.Table, practtest.FieldID, selector),
			sqlgraph.To(submoduletest.Table, submoduletest.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, practtest.SubModuleTestTable, practtest.SubModuleTestColumn),
		)
		fromU = sqlgraph.SetNeighbors(ptq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PractTest entity from the query.
// Returns a *NotFoundError when no PractTest was found.
func (ptq *PractTestQuery) First(ctx context.Context) (*PractTest, error) {
	nodes, err := ptq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{practtest.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ptq *PractTestQuery) FirstX(ctx context.Context) *PractTest {
	node, err := ptq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PractTest ID from the query.
// Returns a *NotFoundError when no PractTest ID was found.
func (ptq *PractTestQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ptq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{practtest.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ptq *PractTestQuery) FirstIDX(ctx context.Context) int {
	id, err := ptq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PractTest entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one PractTest entity is not found.
// Returns a *NotFoundError when no PractTest entities are found.
func (ptq *PractTestQuery) Only(ctx context.Context) (*PractTest, error) {
	nodes, err := ptq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{practtest.Label}
	default:
		return nil, &NotSingularError{practtest.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ptq *PractTestQuery) OnlyX(ctx context.Context) *PractTest {
	node, err := ptq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PractTest ID in the query.
// Returns a *NotSingularError when exactly one PractTest ID is not found.
// Returns a *NotFoundError when no entities are found.
func (ptq *PractTestQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ptq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{practtest.Label}
	default:
		err = &NotSingularError{practtest.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ptq *PractTestQuery) OnlyIDX(ctx context.Context) int {
	id, err := ptq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PractTests.
func (ptq *PractTestQuery) All(ctx context.Context) ([]*PractTest, error) {
	if err := ptq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ptq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ptq *PractTestQuery) AllX(ctx context.Context) []*PractTest {
	nodes, err := ptq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PractTest IDs.
func (ptq *PractTestQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ptq.Select(practtest.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ptq *PractTestQuery) IDsX(ctx context.Context) []int {
	ids, err := ptq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ptq *PractTestQuery) Count(ctx context.Context) (int, error) {
	if err := ptq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ptq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ptq *PractTestQuery) CountX(ctx context.Context) int {
	count, err := ptq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ptq *PractTestQuery) Exist(ctx context.Context) (bool, error) {
	if err := ptq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ptq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ptq *PractTestQuery) ExistX(ctx context.Context) bool {
	exist, err := ptq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PractTestQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ptq *PractTestQuery) Clone() *PractTestQuery {
	if ptq == nil {
		return nil
	}
	return &PractTestQuery{
		config:            ptq.config,
		limit:             ptq.limit,
		offset:            ptq.offset,
		order:             append([]OrderFunc{}, ptq.order...),
		predicates:        append([]predicate.PractTest{}, ptq.predicates...),
		withSubModuleTest: ptq.withSubModuleTest.Clone(),
		// clone intermediate query.
		sql:  ptq.sql.Clone(),
		path: ptq.path,
	}
}

// WithSubModuleTest tells the query-builder to eager-load the nodes that are connected to
// the "SubModuleTest" edge. The optional arguments are used to configure the query builder of the edge.
func (ptq *PractTestQuery) WithSubModuleTest(opts ...func(*SubModuleTestQuery)) *PractTestQuery {
	query := &SubModuleTestQuery{config: ptq.config}
	for _, opt := range opts {
		opt(query)
	}
	ptq.withSubModuleTest = query
	return ptq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		SubmoduletestID int `json:"submoduletest_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PractTest.Query().
//		GroupBy(practtest.FieldSubmoduletestID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ptq *PractTestQuery) GroupBy(field string, fields ...string) *PractTestGroupBy {
	group := &PractTestGroupBy{config: ptq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ptq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ptq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		SubmoduletestID int `json:"submoduletest_id,omitempty"`
//	}
//
//	client.PractTest.Query().
//		Select(practtest.FieldSubmoduletestID).
//		Scan(ctx, &v)
//
func (ptq *PractTestQuery) Select(fields ...string) *PractTestSelect {
	ptq.fields = append(ptq.fields, fields...)
	return &PractTestSelect{PractTestQuery: ptq}
}

func (ptq *PractTestQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ptq.fields {
		if !practtest.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ptq.path != nil {
		prev, err := ptq.path(ctx)
		if err != nil {
			return err
		}
		ptq.sql = prev
	}
	return nil
}

func (ptq *PractTestQuery) sqlAll(ctx context.Context) ([]*PractTest, error) {
	var (
		nodes       = []*PractTest{}
		_spec       = ptq.querySpec()
		loadedTypes = [1]bool{
			ptq.withSubModuleTest != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &PractTest{config: ptq.config}
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
	if err := sqlgraph.QueryNodes(ctx, ptq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ptq.withSubModuleTest; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*PractTest)
		for i := range nodes {
			fk := nodes[i].SubmoduletestID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(submoduletest.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "submoduletest_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.SubModuleTest = n
			}
		}
	}

	return nodes, nil
}

func (ptq *PractTestQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ptq.querySpec()
	return sqlgraph.CountNodes(ctx, ptq.driver, _spec)
}

func (ptq *PractTestQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ptq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ptq *PractTestQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   practtest.Table,
			Columns: practtest.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: practtest.FieldID,
			},
		},
		From:   ptq.sql,
		Unique: true,
	}
	if unique := ptq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ptq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, practtest.FieldID)
		for i := range fields {
			if fields[i] != practtest.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ptq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ptq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ptq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ptq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ptq *PractTestQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ptq.driver.Dialect())
	t1 := builder.Table(practtest.Table)
	columns := ptq.fields
	if len(columns) == 0 {
		columns = practtest.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ptq.sql != nil {
		selector = ptq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range ptq.predicates {
		p(selector)
	}
	for _, p := range ptq.order {
		p(selector)
	}
	if offset := ptq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ptq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PractTestGroupBy is the group-by builder for PractTest entities.
type PractTestGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ptgb *PractTestGroupBy) Aggregate(fns ...AggregateFunc) *PractTestGroupBy {
	ptgb.fns = append(ptgb.fns, fns...)
	return ptgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ptgb *PractTestGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ptgb.path(ctx)
	if err != nil {
		return err
	}
	ptgb.sql = query
	return ptgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ptgb *PractTestGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ptgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ptgb *PractTestGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ptgb.fields) > 1 {
		return nil, errors.New("ent: PractTestGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ptgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ptgb *PractTestGroupBy) StringsX(ctx context.Context) []string {
	v, err := ptgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ptgb *PractTestGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ptgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{practtest.Label}
	default:
		err = fmt.Errorf("ent: PractTestGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ptgb *PractTestGroupBy) StringX(ctx context.Context) string {
	v, err := ptgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ptgb *PractTestGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ptgb.fields) > 1 {
		return nil, errors.New("ent: PractTestGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ptgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ptgb *PractTestGroupBy) IntsX(ctx context.Context) []int {
	v, err := ptgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ptgb *PractTestGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ptgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{practtest.Label}
	default:
		err = fmt.Errorf("ent: PractTestGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ptgb *PractTestGroupBy) IntX(ctx context.Context) int {
	v, err := ptgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ptgb *PractTestGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ptgb.fields) > 1 {
		return nil, errors.New("ent: PractTestGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ptgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ptgb *PractTestGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ptgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ptgb *PractTestGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ptgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{practtest.Label}
	default:
		err = fmt.Errorf("ent: PractTestGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ptgb *PractTestGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ptgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ptgb *PractTestGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ptgb.fields) > 1 {
		return nil, errors.New("ent: PractTestGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ptgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ptgb *PractTestGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ptgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ptgb *PractTestGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ptgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{practtest.Label}
	default:
		err = fmt.Errorf("ent: PractTestGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ptgb *PractTestGroupBy) BoolX(ctx context.Context) bool {
	v, err := ptgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ptgb *PractTestGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ptgb.fields {
		if !practtest.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ptgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ptgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ptgb *PractTestGroupBy) sqlQuery() *sql.Selector {
	selector := ptgb.sql.Select()
	aggregation := make([]string, 0, len(ptgb.fns))
	for _, fn := range ptgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ptgb.fields)+len(ptgb.fns))
		for _, f := range ptgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ptgb.fields...)...)
}

// PractTestSelect is the builder for selecting fields of PractTest entities.
type PractTestSelect struct {
	*PractTestQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (pts *PractTestSelect) Scan(ctx context.Context, v interface{}) error {
	if err := pts.prepareQuery(ctx); err != nil {
		return err
	}
	pts.sql = pts.PractTestQuery.sqlQuery(ctx)
	return pts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pts *PractTestSelect) ScanX(ctx context.Context, v interface{}) {
	if err := pts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (pts *PractTestSelect) Strings(ctx context.Context) ([]string, error) {
	if len(pts.fields) > 1 {
		return nil, errors.New("ent: PractTestSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := pts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pts *PractTestSelect) StringsX(ctx context.Context) []string {
	v, err := pts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (pts *PractTestSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pts.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{practtest.Label}
	default:
		err = fmt.Errorf("ent: PractTestSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pts *PractTestSelect) StringX(ctx context.Context) string {
	v, err := pts.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (pts *PractTestSelect) Ints(ctx context.Context) ([]int, error) {
	if len(pts.fields) > 1 {
		return nil, errors.New("ent: PractTestSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := pts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pts *PractTestSelect) IntsX(ctx context.Context) []int {
	v, err := pts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (pts *PractTestSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pts.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{practtest.Label}
	default:
		err = fmt.Errorf("ent: PractTestSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pts *PractTestSelect) IntX(ctx context.Context) int {
	v, err := pts.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (pts *PractTestSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(pts.fields) > 1 {
		return nil, errors.New("ent: PractTestSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := pts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pts *PractTestSelect) Float64sX(ctx context.Context) []float64 {
	v, err := pts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (pts *PractTestSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pts.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{practtest.Label}
	default:
		err = fmt.Errorf("ent: PractTestSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pts *PractTestSelect) Float64X(ctx context.Context) float64 {
	v, err := pts.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (pts *PractTestSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(pts.fields) > 1 {
		return nil, errors.New("ent: PractTestSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := pts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pts *PractTestSelect) BoolsX(ctx context.Context) []bool {
	v, err := pts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (pts *PractTestSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pts.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{practtest.Label}
	default:
		err = fmt.Errorf("ent: PractTestSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pts *PractTestSelect) BoolX(ctx context.Context) bool {
	v, err := pts.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pts *PractTestSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pts.sql.Query()
	if err := pts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
