// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/0B1t322/CP-Rosseti-Back/ent/answer"
	"github.com/0B1t322/CP-Rosseti-Back/ent/predicate"
	"github.com/0B1t322/CP-Rosseti-Back/ent/question"
	"github.com/0B1t322/CP-Rosseti-Back/ent/theoreticaltest"
)

// QuestionUpdate is the builder for updating Question entities.
type QuestionUpdate struct {
	config
	hooks    []Hook
	mutation *QuestionMutation
}

// Where appends a list predicates to the QuestionUpdate builder.
func (qu *QuestionUpdate) Where(ps ...predicate.Question) *QuestionUpdate {
	qu.mutation.Where(ps...)
	return qu
}

// SetTheoricalTestID sets the "theorical_test_id" field.
func (qu *QuestionUpdate) SetTheoricalTestID(i int) *QuestionUpdate {
	qu.mutation.SetTheoricalTestID(i)
	return qu
}

// SetQuestion sets the "question" field.
func (qu *QuestionUpdate) SetQuestion(s string) *QuestionUpdate {
	qu.mutation.SetQuestion(s)
	return qu
}

// SetTheoreticalTestID sets the "TheoreticalTest" edge to the TheoreticalTest entity by ID.
func (qu *QuestionUpdate) SetTheoreticalTestID(id int) *QuestionUpdate {
	qu.mutation.SetTheoreticalTestID(id)
	return qu
}

// SetTheoreticalTest sets the "TheoreticalTest" edge to the TheoreticalTest entity.
func (qu *QuestionUpdate) SetTheoreticalTest(t *TheoreticalTest) *QuestionUpdate {
	return qu.SetTheoreticalTestID(t.ID)
}

// AddAnswerIDs adds the "Answer" edge to the Answer entity by IDs.
func (qu *QuestionUpdate) AddAnswerIDs(ids ...int) *QuestionUpdate {
	qu.mutation.AddAnswerIDs(ids...)
	return qu
}

// AddAnswer adds the "Answer" edges to the Answer entity.
func (qu *QuestionUpdate) AddAnswer(a ...*Answer) *QuestionUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return qu.AddAnswerIDs(ids...)
}

// Mutation returns the QuestionMutation object of the builder.
func (qu *QuestionUpdate) Mutation() *QuestionMutation {
	return qu.mutation
}

// ClearTheoreticalTest clears the "TheoreticalTest" edge to the TheoreticalTest entity.
func (qu *QuestionUpdate) ClearTheoreticalTest() *QuestionUpdate {
	qu.mutation.ClearTheoreticalTest()
	return qu
}

// ClearAnswer clears all "Answer" edges to the Answer entity.
func (qu *QuestionUpdate) ClearAnswer() *QuestionUpdate {
	qu.mutation.ClearAnswer()
	return qu
}

// RemoveAnswerIDs removes the "Answer" edge to Answer entities by IDs.
func (qu *QuestionUpdate) RemoveAnswerIDs(ids ...int) *QuestionUpdate {
	qu.mutation.RemoveAnswerIDs(ids...)
	return qu
}

// RemoveAnswer removes "Answer" edges to Answer entities.
func (qu *QuestionUpdate) RemoveAnswer(a ...*Answer) *QuestionUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return qu.RemoveAnswerIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (qu *QuestionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(qu.hooks) == 0 {
		if err = qu.check(); err != nil {
			return 0, err
		}
		affected, err = qu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*QuestionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = qu.check(); err != nil {
				return 0, err
			}
			qu.mutation = mutation
			affected, err = qu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(qu.hooks) - 1; i >= 0; i-- {
			if qu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = qu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, qu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (qu *QuestionUpdate) SaveX(ctx context.Context) int {
	affected, err := qu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (qu *QuestionUpdate) Exec(ctx context.Context) error {
	_, err := qu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qu *QuestionUpdate) ExecX(ctx context.Context) {
	if err := qu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (qu *QuestionUpdate) check() error {
	if _, ok := qu.mutation.TheoreticalTestID(); qu.mutation.TheoreticalTestCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"TheoreticalTest\"")
	}
	return nil
}

func (qu *QuestionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   question.Table,
			Columns: question.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: question.FieldID,
			},
		},
	}
	if ps := qu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := qu.mutation.Question(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: question.FieldQuestion,
		})
	}
	if qu.mutation.TheoreticalTestCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   question.TheoreticalTestTable,
			Columns: []string{question.TheoreticalTestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: theoreticaltest.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qu.mutation.TheoreticalTestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   question.TheoreticalTestTable,
			Columns: []string{question.TheoreticalTestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: theoreticaltest.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if qu.mutation.AnswerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   question.AnswerTable,
			Columns: []string{question.AnswerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: answer.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qu.mutation.RemovedAnswerIDs(); len(nodes) > 0 && !qu.mutation.AnswerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   question.AnswerTable,
			Columns: []string{question.AnswerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: answer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qu.mutation.AnswerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   question.AnswerTable,
			Columns: []string{question.AnswerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: answer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, qu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{question.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// QuestionUpdateOne is the builder for updating a single Question entity.
type QuestionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *QuestionMutation
}

// SetTheoricalTestID sets the "theorical_test_id" field.
func (quo *QuestionUpdateOne) SetTheoricalTestID(i int) *QuestionUpdateOne {
	quo.mutation.SetTheoricalTestID(i)
	return quo
}

// SetQuestion sets the "question" field.
func (quo *QuestionUpdateOne) SetQuestion(s string) *QuestionUpdateOne {
	quo.mutation.SetQuestion(s)
	return quo
}

// SetTheoreticalTestID sets the "TheoreticalTest" edge to the TheoreticalTest entity by ID.
func (quo *QuestionUpdateOne) SetTheoreticalTestID(id int) *QuestionUpdateOne {
	quo.mutation.SetTheoreticalTestID(id)
	return quo
}

// SetTheoreticalTest sets the "TheoreticalTest" edge to the TheoreticalTest entity.
func (quo *QuestionUpdateOne) SetTheoreticalTest(t *TheoreticalTest) *QuestionUpdateOne {
	return quo.SetTheoreticalTestID(t.ID)
}

// AddAnswerIDs adds the "Answer" edge to the Answer entity by IDs.
func (quo *QuestionUpdateOne) AddAnswerIDs(ids ...int) *QuestionUpdateOne {
	quo.mutation.AddAnswerIDs(ids...)
	return quo
}

// AddAnswer adds the "Answer" edges to the Answer entity.
func (quo *QuestionUpdateOne) AddAnswer(a ...*Answer) *QuestionUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return quo.AddAnswerIDs(ids...)
}

// Mutation returns the QuestionMutation object of the builder.
func (quo *QuestionUpdateOne) Mutation() *QuestionMutation {
	return quo.mutation
}

// ClearTheoreticalTest clears the "TheoreticalTest" edge to the TheoreticalTest entity.
func (quo *QuestionUpdateOne) ClearTheoreticalTest() *QuestionUpdateOne {
	quo.mutation.ClearTheoreticalTest()
	return quo
}

// ClearAnswer clears all "Answer" edges to the Answer entity.
func (quo *QuestionUpdateOne) ClearAnswer() *QuestionUpdateOne {
	quo.mutation.ClearAnswer()
	return quo
}

// RemoveAnswerIDs removes the "Answer" edge to Answer entities by IDs.
func (quo *QuestionUpdateOne) RemoveAnswerIDs(ids ...int) *QuestionUpdateOne {
	quo.mutation.RemoveAnswerIDs(ids...)
	return quo
}

// RemoveAnswer removes "Answer" edges to Answer entities.
func (quo *QuestionUpdateOne) RemoveAnswer(a ...*Answer) *QuestionUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return quo.RemoveAnswerIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (quo *QuestionUpdateOne) Select(field string, fields ...string) *QuestionUpdateOne {
	quo.fields = append([]string{field}, fields...)
	return quo
}

// Save executes the query and returns the updated Question entity.
func (quo *QuestionUpdateOne) Save(ctx context.Context) (*Question, error) {
	var (
		err  error
		node *Question
	)
	if len(quo.hooks) == 0 {
		if err = quo.check(); err != nil {
			return nil, err
		}
		node, err = quo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*QuestionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = quo.check(); err != nil {
				return nil, err
			}
			quo.mutation = mutation
			node, err = quo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(quo.hooks) - 1; i >= 0; i-- {
			if quo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = quo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, quo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (quo *QuestionUpdateOne) SaveX(ctx context.Context) *Question {
	node, err := quo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (quo *QuestionUpdateOne) Exec(ctx context.Context) error {
	_, err := quo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (quo *QuestionUpdateOne) ExecX(ctx context.Context) {
	if err := quo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (quo *QuestionUpdateOne) check() error {
	if _, ok := quo.mutation.TheoreticalTestID(); quo.mutation.TheoreticalTestCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"TheoreticalTest\"")
	}
	return nil
}

func (quo *QuestionUpdateOne) sqlSave(ctx context.Context) (_node *Question, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   question.Table,
			Columns: question.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: question.FieldID,
			},
		},
	}
	id, ok := quo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Question.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := quo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, question.FieldID)
		for _, f := range fields {
			if !question.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != question.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := quo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := quo.mutation.Question(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: question.FieldQuestion,
		})
	}
	if quo.mutation.TheoreticalTestCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   question.TheoreticalTestTable,
			Columns: []string{question.TheoreticalTestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: theoreticaltest.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := quo.mutation.TheoreticalTestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   question.TheoreticalTestTable,
			Columns: []string{question.TheoreticalTestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: theoreticaltest.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if quo.mutation.AnswerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   question.AnswerTable,
			Columns: []string{question.AnswerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: answer.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := quo.mutation.RemovedAnswerIDs(); len(nodes) > 0 && !quo.mutation.AnswerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   question.AnswerTable,
			Columns: []string{question.AnswerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: answer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := quo.mutation.AnswerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   question.AnswerTable,
			Columns: []string{question.AnswerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: answer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Question{config: quo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, quo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{question.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
