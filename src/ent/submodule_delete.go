// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/0B1t322/CP-Rosseti-Back/ent/predicate"
	"github.com/0B1t322/CP-Rosseti-Back/ent/submodule"
)

// SubModuleDelete is the builder for deleting a SubModule entity.
type SubModuleDelete struct {
	config
	hooks    []Hook
	mutation *SubModuleMutation
}

// Where appends a list predicates to the SubModuleDelete builder.
func (smd *SubModuleDelete) Where(ps ...predicate.SubModule) *SubModuleDelete {
	smd.mutation.Where(ps...)
	return smd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (smd *SubModuleDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(smd.hooks) == 0 {
		affected, err = smd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SubModuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			smd.mutation = mutation
			affected, err = smd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(smd.hooks) - 1; i >= 0; i-- {
			if smd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = smd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, smd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (smd *SubModuleDelete) ExecX(ctx context.Context) int {
	n, err := smd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (smd *SubModuleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: submodule.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: submodule.FieldID,
			},
		},
	}
	if ps := smd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, smd.driver, _spec)
}

// SubModuleDeleteOne is the builder for deleting a single SubModule entity.
type SubModuleDeleteOne struct {
	smd *SubModuleDelete
}

// Exec executes the deletion query.
func (smdo *SubModuleDeleteOne) Exec(ctx context.Context) error {
	n, err := smdo.smd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{submodule.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (smdo *SubModuleDeleteOne) ExecX(ctx context.Context) {
	smdo.smd.ExecX(ctx)
}
