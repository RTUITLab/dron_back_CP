// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/0B1t322/CP-Rosseti-Back/ent/predicate"
	"github.com/0B1t322/CP-Rosseti-Back/ent/submoduletest"
)

// SubModuleTestDelete is the builder for deleting a SubModuleTest entity.
type SubModuleTestDelete struct {
	config
	hooks    []Hook
	mutation *SubModuleTestMutation
}

// Where appends a list predicates to the SubModuleTestDelete builder.
func (smtd *SubModuleTestDelete) Where(ps ...predicate.SubModuleTest) *SubModuleTestDelete {
	smtd.mutation.Where(ps...)
	return smtd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (smtd *SubModuleTestDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(smtd.hooks) == 0 {
		affected, err = smtd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SubModuleTestMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			smtd.mutation = mutation
			affected, err = smtd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(smtd.hooks) - 1; i >= 0; i-- {
			if smtd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = smtd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, smtd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (smtd *SubModuleTestDelete) ExecX(ctx context.Context) int {
	n, err := smtd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (smtd *SubModuleTestDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: submoduletest.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: submoduletest.FieldID,
			},
		},
	}
	if ps := smtd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, smtd.driver, _spec)
}

// SubModuleTestDeleteOne is the builder for deleting a single SubModuleTest entity.
type SubModuleTestDeleteOne struct {
	smtd *SubModuleTestDelete
}

// Exec executes the deletion query.
func (smtdo *SubModuleTestDeleteOne) Exec(ctx context.Context) error {
	n, err := smtdo.smtd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{submoduletest.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (smtdo *SubModuleTestDeleteOne) ExecX(ctx context.Context) {
	smtdo.smtd.ExecX(ctx)
}
