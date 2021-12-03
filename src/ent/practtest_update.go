// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/0B1t322/CP-Rosseti-Back/ent/practtest"
	"github.com/0B1t322/CP-Rosseti-Back/ent/predicate"
	"github.com/0B1t322/CP-Rosseti-Back/ent/schema"
	"github.com/0B1t322/CP-Rosseti-Back/ent/submoduletest"
)

// PractTestUpdate is the builder for updating PractTest entities.
type PractTestUpdate struct {
	config
	hooks    []Hook
	mutation *PractTestMutation
}

// Where appends a list predicates to the PractTestUpdate builder.
func (ptu *PractTestUpdate) Where(ps ...predicate.PractTest) *PractTestUpdate {
	ptu.mutation.Where(ps...)
	return ptu
}

// SetSubmoduletestID sets the "submoduletest_id" field.
func (ptu *PractTestUpdate) SetSubmoduletestID(i int) *PractTestUpdate {
	ptu.mutation.SetSubmoduletestID(i)
	return ptu
}

// SetConfig sets the "config" field.
func (ptu *PractTestUpdate) SetConfig(so schema.JSONObject) *PractTestUpdate {
	ptu.mutation.SetConfig(so)
	return ptu
}

// SetSubModuleTestID sets the "SubModuleTest" edge to the SubModuleTest entity by ID.
func (ptu *PractTestUpdate) SetSubModuleTestID(id int) *PractTestUpdate {
	ptu.mutation.SetSubModuleTestID(id)
	return ptu
}

// SetSubModuleTest sets the "SubModuleTest" edge to the SubModuleTest entity.
func (ptu *PractTestUpdate) SetSubModuleTest(s *SubModuleTest) *PractTestUpdate {
	return ptu.SetSubModuleTestID(s.ID)
}

// Mutation returns the PractTestMutation object of the builder.
func (ptu *PractTestUpdate) Mutation() *PractTestMutation {
	return ptu.mutation
}

// ClearSubModuleTest clears the "SubModuleTest" edge to the SubModuleTest entity.
func (ptu *PractTestUpdate) ClearSubModuleTest() *PractTestUpdate {
	ptu.mutation.ClearSubModuleTest()
	return ptu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ptu *PractTestUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ptu.hooks) == 0 {
		if err = ptu.check(); err != nil {
			return 0, err
		}
		affected, err = ptu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PractTestMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ptu.check(); err != nil {
				return 0, err
			}
			ptu.mutation = mutation
			affected, err = ptu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ptu.hooks) - 1; i >= 0; i-- {
			if ptu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ptu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ptu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ptu *PractTestUpdate) SaveX(ctx context.Context) int {
	affected, err := ptu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ptu *PractTestUpdate) Exec(ctx context.Context) error {
	_, err := ptu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptu *PractTestUpdate) ExecX(ctx context.Context) {
	if err := ptu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ptu *PractTestUpdate) check() error {
	if _, ok := ptu.mutation.SubModuleTestID(); ptu.mutation.SubModuleTestCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"SubModuleTest\"")
	}
	return nil
}

func (ptu *PractTestUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   practtest.Table,
			Columns: practtest.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: practtest.FieldID,
			},
		},
	}
	if ps := ptu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ptu.mutation.Config(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: practtest.FieldConfig,
		})
	}
	if ptu.mutation.SubModuleTestCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   practtest.SubModuleTestTable,
			Columns: []string{practtest.SubModuleTestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: submoduletest.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptu.mutation.SubModuleTestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   practtest.SubModuleTestTable,
			Columns: []string{practtest.SubModuleTestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: submoduletest.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ptu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{practtest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// PractTestUpdateOne is the builder for updating a single PractTest entity.
type PractTestUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PractTestMutation
}

// SetSubmoduletestID sets the "submoduletest_id" field.
func (ptuo *PractTestUpdateOne) SetSubmoduletestID(i int) *PractTestUpdateOne {
	ptuo.mutation.SetSubmoduletestID(i)
	return ptuo
}

// SetConfig sets the "config" field.
func (ptuo *PractTestUpdateOne) SetConfig(so schema.JSONObject) *PractTestUpdateOne {
	ptuo.mutation.SetConfig(so)
	return ptuo
}

// SetSubModuleTestID sets the "SubModuleTest" edge to the SubModuleTest entity by ID.
func (ptuo *PractTestUpdateOne) SetSubModuleTestID(id int) *PractTestUpdateOne {
	ptuo.mutation.SetSubModuleTestID(id)
	return ptuo
}

// SetSubModuleTest sets the "SubModuleTest" edge to the SubModuleTest entity.
func (ptuo *PractTestUpdateOne) SetSubModuleTest(s *SubModuleTest) *PractTestUpdateOne {
	return ptuo.SetSubModuleTestID(s.ID)
}

// Mutation returns the PractTestMutation object of the builder.
func (ptuo *PractTestUpdateOne) Mutation() *PractTestMutation {
	return ptuo.mutation
}

// ClearSubModuleTest clears the "SubModuleTest" edge to the SubModuleTest entity.
func (ptuo *PractTestUpdateOne) ClearSubModuleTest() *PractTestUpdateOne {
	ptuo.mutation.ClearSubModuleTest()
	return ptuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ptuo *PractTestUpdateOne) Select(field string, fields ...string) *PractTestUpdateOne {
	ptuo.fields = append([]string{field}, fields...)
	return ptuo
}

// Save executes the query and returns the updated PractTest entity.
func (ptuo *PractTestUpdateOne) Save(ctx context.Context) (*PractTest, error) {
	var (
		err  error
		node *PractTest
	)
	if len(ptuo.hooks) == 0 {
		if err = ptuo.check(); err != nil {
			return nil, err
		}
		node, err = ptuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PractTestMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ptuo.check(); err != nil {
				return nil, err
			}
			ptuo.mutation = mutation
			node, err = ptuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ptuo.hooks) - 1; i >= 0; i-- {
			if ptuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ptuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ptuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ptuo *PractTestUpdateOne) SaveX(ctx context.Context) *PractTest {
	node, err := ptuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ptuo *PractTestUpdateOne) Exec(ctx context.Context) error {
	_, err := ptuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptuo *PractTestUpdateOne) ExecX(ctx context.Context) {
	if err := ptuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ptuo *PractTestUpdateOne) check() error {
	if _, ok := ptuo.mutation.SubModuleTestID(); ptuo.mutation.SubModuleTestCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"SubModuleTest\"")
	}
	return nil
}

func (ptuo *PractTestUpdateOne) sqlSave(ctx context.Context) (_node *PractTest, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   practtest.Table,
			Columns: practtest.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: practtest.FieldID,
			},
		},
	}
	id, ok := ptuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing PractTest.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ptuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, practtest.FieldID)
		for _, f := range fields {
			if !practtest.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != practtest.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ptuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ptuo.mutation.Config(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: practtest.FieldConfig,
		})
	}
	if ptuo.mutation.SubModuleTestCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   practtest.SubModuleTestTable,
			Columns: []string{practtest.SubModuleTestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: submoduletest.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptuo.mutation.SubModuleTestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   practtest.SubModuleTestTable,
			Columns: []string{practtest.SubModuleTestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: submoduletest.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PractTest{config: ptuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ptuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{practtest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}