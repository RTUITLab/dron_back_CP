// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/0B1t322/CP-Rosseti-Back/ent/moduledependcies"
	"github.com/0B1t322/CP-Rosseti-Back/ent/predicate"
)

// ModuleDependciesDelete is the builder for deleting a ModuleDependcies entity.
type ModuleDependciesDelete struct {
	config
	hooks    []Hook
	mutation *ModuleDependciesMutation
}

// Where appends a list predicates to the ModuleDependciesDelete builder.
func (mdd *ModuleDependciesDelete) Where(ps ...predicate.ModuleDependcies) *ModuleDependciesDelete {
	mdd.mutation.Where(ps...)
	return mdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (mdd *ModuleDependciesDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mdd.hooks) == 0 {
		affected, err = mdd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ModuleDependciesMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mdd.mutation = mutation
			affected, err = mdd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mdd.hooks) - 1; i >= 0; i-- {
			if mdd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mdd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mdd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (mdd *ModuleDependciesDelete) ExecX(ctx context.Context) int {
	n, err := mdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (mdd *ModuleDependciesDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: moduledependcies.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: moduledependcies.FieldID,
			},
		},
	}
	if ps := mdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, mdd.driver, _spec)
}

// ModuleDependciesDeleteOne is the builder for deleting a single ModuleDependcies entity.
type ModuleDependciesDeleteOne struct {
	mdd *ModuleDependciesDelete
}

// Exec executes the deletion query.
func (mddo *ModuleDependciesDeleteOne) Exec(ctx context.Context) error {
	n, err := mddo.mdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{moduledependcies.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mddo *ModuleDependciesDeleteOne) ExecX(ctx context.Context) {
	mddo.mdd.ExecX(ctx)
}
