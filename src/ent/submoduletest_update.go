// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/0B1t322/CP-Rosseti-Back/ent/predicate"
	"github.com/0B1t322/CP-Rosseti-Back/ent/submodule"
	"github.com/0B1t322/CP-Rosseti-Back/ent/submoduletest"
	"github.com/0B1t322/CP-Rosseti-Back/ent/test"
)

// SubModuleTestUpdate is the builder for updating SubModuleTest entities.
type SubModuleTestUpdate struct {
	config
	hooks    []Hook
	mutation *SubModuleTestMutation
}

// Where appends a list predicates to the SubModuleTestUpdate builder.
func (smtu *SubModuleTestUpdate) Where(ps ...predicate.SubModuleTest) *SubModuleTestUpdate {
	smtu.mutation.Where(ps...)
	return smtu
}

// SetSubmoduleID sets the "submodule_id" field.
func (smtu *SubModuleTestUpdate) SetSubmoduleID(i int) *SubModuleTestUpdate {
	smtu.mutation.SetSubmoduleID(i)
	return smtu
}

// SetTestID sets the "test_id" field.
func (smtu *SubModuleTestUpdate) SetTestID(i int) *SubModuleTestUpdate {
	smtu.mutation.SetTestID(i)
	return smtu
}

// SetSubModuleID sets the "SubModule" edge to the SubModule entity by ID.
func (smtu *SubModuleTestUpdate) SetSubModuleID(id int) *SubModuleTestUpdate {
	smtu.mutation.SetSubModuleID(id)
	return smtu
}

// SetSubModule sets the "SubModule" edge to the SubModule entity.
func (smtu *SubModuleTestUpdate) SetSubModule(s *SubModule) *SubModuleTestUpdate {
	return smtu.SetSubModuleID(s.ID)
}

// SetTest sets the "Test" edge to the Test entity.
func (smtu *SubModuleTestUpdate) SetTest(t *Test) *SubModuleTestUpdate {
	return smtu.SetTestID(t.ID)
}

// Mutation returns the SubModuleTestMutation object of the builder.
func (smtu *SubModuleTestUpdate) Mutation() *SubModuleTestMutation {
	return smtu.mutation
}

// ClearSubModule clears the "SubModule" edge to the SubModule entity.
func (smtu *SubModuleTestUpdate) ClearSubModule() *SubModuleTestUpdate {
	smtu.mutation.ClearSubModule()
	return smtu
}

// ClearTest clears the "Test" edge to the Test entity.
func (smtu *SubModuleTestUpdate) ClearTest() *SubModuleTestUpdate {
	smtu.mutation.ClearTest()
	return smtu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (smtu *SubModuleTestUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(smtu.hooks) == 0 {
		if err = smtu.check(); err != nil {
			return 0, err
		}
		affected, err = smtu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SubModuleTestMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = smtu.check(); err != nil {
				return 0, err
			}
			smtu.mutation = mutation
			affected, err = smtu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(smtu.hooks) - 1; i >= 0; i-- {
			if smtu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = smtu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, smtu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (smtu *SubModuleTestUpdate) SaveX(ctx context.Context) int {
	affected, err := smtu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (smtu *SubModuleTestUpdate) Exec(ctx context.Context) error {
	_, err := smtu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smtu *SubModuleTestUpdate) ExecX(ctx context.Context) {
	if err := smtu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (smtu *SubModuleTestUpdate) check() error {
	if _, ok := smtu.mutation.SubModuleID(); smtu.mutation.SubModuleCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"SubModule\"")
	}
	if _, ok := smtu.mutation.TestID(); smtu.mutation.TestCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"Test\"")
	}
	return nil
}

func (smtu *SubModuleTestUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   submoduletest.Table,
			Columns: submoduletest.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: submoduletest.FieldID,
			},
		},
	}
	if ps := smtu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if smtu.mutation.SubModuleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   submoduletest.SubModuleTable,
			Columns: []string{submoduletest.SubModuleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: submodule.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := smtu.mutation.SubModuleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   submoduletest.SubModuleTable,
			Columns: []string{submoduletest.SubModuleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: submodule.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if smtu.mutation.TestCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   submoduletest.TestTable,
			Columns: []string{submoduletest.TestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: test.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := smtu.mutation.TestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   submoduletest.TestTable,
			Columns: []string{submoduletest.TestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: test.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, smtu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{submoduletest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// SubModuleTestUpdateOne is the builder for updating a single SubModuleTest entity.
type SubModuleTestUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SubModuleTestMutation
}

// SetSubmoduleID sets the "submodule_id" field.
func (smtuo *SubModuleTestUpdateOne) SetSubmoduleID(i int) *SubModuleTestUpdateOne {
	smtuo.mutation.SetSubmoduleID(i)
	return smtuo
}

// SetTestID sets the "test_id" field.
func (smtuo *SubModuleTestUpdateOne) SetTestID(i int) *SubModuleTestUpdateOne {
	smtuo.mutation.SetTestID(i)
	return smtuo
}

// SetSubModuleID sets the "SubModule" edge to the SubModule entity by ID.
func (smtuo *SubModuleTestUpdateOne) SetSubModuleID(id int) *SubModuleTestUpdateOne {
	smtuo.mutation.SetSubModuleID(id)
	return smtuo
}

// SetSubModule sets the "SubModule" edge to the SubModule entity.
func (smtuo *SubModuleTestUpdateOne) SetSubModule(s *SubModule) *SubModuleTestUpdateOne {
	return smtuo.SetSubModuleID(s.ID)
}

// SetTest sets the "Test" edge to the Test entity.
func (smtuo *SubModuleTestUpdateOne) SetTest(t *Test) *SubModuleTestUpdateOne {
	return smtuo.SetTestID(t.ID)
}

// Mutation returns the SubModuleTestMutation object of the builder.
func (smtuo *SubModuleTestUpdateOne) Mutation() *SubModuleTestMutation {
	return smtuo.mutation
}

// ClearSubModule clears the "SubModule" edge to the SubModule entity.
func (smtuo *SubModuleTestUpdateOne) ClearSubModule() *SubModuleTestUpdateOne {
	smtuo.mutation.ClearSubModule()
	return smtuo
}

// ClearTest clears the "Test" edge to the Test entity.
func (smtuo *SubModuleTestUpdateOne) ClearTest() *SubModuleTestUpdateOne {
	smtuo.mutation.ClearTest()
	return smtuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (smtuo *SubModuleTestUpdateOne) Select(field string, fields ...string) *SubModuleTestUpdateOne {
	smtuo.fields = append([]string{field}, fields...)
	return smtuo
}

// Save executes the query and returns the updated SubModuleTest entity.
func (smtuo *SubModuleTestUpdateOne) Save(ctx context.Context) (*SubModuleTest, error) {
	var (
		err  error
		node *SubModuleTest
	)
	if len(smtuo.hooks) == 0 {
		if err = smtuo.check(); err != nil {
			return nil, err
		}
		node, err = smtuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SubModuleTestMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = smtuo.check(); err != nil {
				return nil, err
			}
			smtuo.mutation = mutation
			node, err = smtuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(smtuo.hooks) - 1; i >= 0; i-- {
			if smtuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = smtuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, smtuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (smtuo *SubModuleTestUpdateOne) SaveX(ctx context.Context) *SubModuleTest {
	node, err := smtuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (smtuo *SubModuleTestUpdateOne) Exec(ctx context.Context) error {
	_, err := smtuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smtuo *SubModuleTestUpdateOne) ExecX(ctx context.Context) {
	if err := smtuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (smtuo *SubModuleTestUpdateOne) check() error {
	if _, ok := smtuo.mutation.SubModuleID(); smtuo.mutation.SubModuleCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"SubModule\"")
	}
	if _, ok := smtuo.mutation.TestID(); smtuo.mutation.TestCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"Test\"")
	}
	return nil
}

func (smtuo *SubModuleTestUpdateOne) sqlSave(ctx context.Context) (_node *SubModuleTest, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   submoduletest.Table,
			Columns: submoduletest.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: submoduletest.FieldID,
			},
		},
	}
	id, ok := smtuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing SubModuleTest.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := smtuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, submoduletest.FieldID)
		for _, f := range fields {
			if !submoduletest.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != submoduletest.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := smtuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if smtuo.mutation.SubModuleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   submoduletest.SubModuleTable,
			Columns: []string{submoduletest.SubModuleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: submodule.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := smtuo.mutation.SubModuleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   submoduletest.SubModuleTable,
			Columns: []string{submoduletest.SubModuleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: submodule.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if smtuo.mutation.TestCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   submoduletest.TestTable,
			Columns: []string{submoduletest.TestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: test.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := smtuo.mutation.TestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   submoduletest.TestTable,
			Columns: []string{submoduletest.TestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: test.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &SubModuleTest{config: smtuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, smtuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{submoduletest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
