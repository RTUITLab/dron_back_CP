// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/0B1t322/CP-Rosseti-Back/ent/module"
	"github.com/0B1t322/CP-Rosseti-Back/ent/moduledependcies"
)

// ModuleDependciesCreate is the builder for creating a ModuleDependcies entity.
type ModuleDependciesCreate struct {
	config
	mutation *ModuleDependciesMutation
	hooks    []Hook
}

// SetDependentID sets the "dependent_id" field.
func (mdc *ModuleDependciesCreate) SetDependentID(i int) *ModuleDependciesCreate {
	mdc.mutation.SetDependentID(i)
	return mdc
}

// SetDependentOnID sets the "dependent_on_id" field.
func (mdc *ModuleDependciesCreate) SetDependentOnID(i int) *ModuleDependciesCreate {
	mdc.mutation.SetDependentOnID(i)
	return mdc
}

// SetModuleDependciesID sets the "ModuleDependcies" edge to the Module entity by ID.
func (mdc *ModuleDependciesCreate) SetModuleDependciesID(id int) *ModuleDependciesCreate {
	mdc.mutation.SetModuleDependciesID(id)
	return mdc
}

// SetModuleDependcies sets the "ModuleDependcies" edge to the Module entity.
func (mdc *ModuleDependciesCreate) SetModuleDependcies(m *Module) *ModuleDependciesCreate {
	return mdc.SetModuleDependciesID(m.ID)
}

// SetModuleDependOnID sets the "ModuleDependOn" edge to the Module entity by ID.
func (mdc *ModuleDependciesCreate) SetModuleDependOnID(id int) *ModuleDependciesCreate {
	mdc.mutation.SetModuleDependOnID(id)
	return mdc
}

// SetModuleDependOn sets the "ModuleDependOn" edge to the Module entity.
func (mdc *ModuleDependciesCreate) SetModuleDependOn(m *Module) *ModuleDependciesCreate {
	return mdc.SetModuleDependOnID(m.ID)
}

// Mutation returns the ModuleDependciesMutation object of the builder.
func (mdc *ModuleDependciesCreate) Mutation() *ModuleDependciesMutation {
	return mdc.mutation
}

// Save creates the ModuleDependcies in the database.
func (mdc *ModuleDependciesCreate) Save(ctx context.Context) (*ModuleDependcies, error) {
	var (
		err  error
		node *ModuleDependcies
	)
	if len(mdc.hooks) == 0 {
		if err = mdc.check(); err != nil {
			return nil, err
		}
		node, err = mdc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ModuleDependciesMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mdc.check(); err != nil {
				return nil, err
			}
			mdc.mutation = mutation
			if node, err = mdc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mdc.hooks) - 1; i >= 0; i-- {
			if mdc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mdc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mdc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mdc *ModuleDependciesCreate) SaveX(ctx context.Context) *ModuleDependcies {
	v, err := mdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mdc *ModuleDependciesCreate) Exec(ctx context.Context) error {
	_, err := mdc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mdc *ModuleDependciesCreate) ExecX(ctx context.Context) {
	if err := mdc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mdc *ModuleDependciesCreate) check() error {
	if _, ok := mdc.mutation.DependentID(); !ok {
		return &ValidationError{Name: "dependent_id", err: errors.New(`ent: missing required field "dependent_id"`)}
	}
	if _, ok := mdc.mutation.DependentOnID(); !ok {
		return &ValidationError{Name: "dependent_on_id", err: errors.New(`ent: missing required field "dependent_on_id"`)}
	}
	if _, ok := mdc.mutation.ModuleDependciesID(); !ok {
		return &ValidationError{Name: "ModuleDependcies", err: errors.New("ent: missing required edge \"ModuleDependcies\"")}
	}
	if _, ok := mdc.mutation.ModuleDependOnID(); !ok {
		return &ValidationError{Name: "ModuleDependOn", err: errors.New("ent: missing required edge \"ModuleDependOn\"")}
	}
	return nil
}

func (mdc *ModuleDependciesCreate) sqlSave(ctx context.Context) (*ModuleDependcies, error) {
	_node, _spec := mdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mdc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (mdc *ModuleDependciesCreate) createSpec() (*ModuleDependcies, *sqlgraph.CreateSpec) {
	var (
		_node = &ModuleDependcies{config: mdc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: moduledependcies.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: moduledependcies.FieldID,
			},
		}
	)
	if nodes := mdc.mutation.ModuleDependciesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   moduledependcies.ModuleDependciesTable,
			Columns: []string{moduledependcies.ModuleDependciesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: module.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.DependentID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mdc.mutation.ModuleDependOnIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   moduledependcies.ModuleDependOnTable,
			Columns: []string{moduledependcies.ModuleDependOnColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: module.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.DependentOnID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ModuleDependciesCreateBulk is the builder for creating many ModuleDependcies entities in bulk.
type ModuleDependciesCreateBulk struct {
	config
	builders []*ModuleDependciesCreate
}

// Save creates the ModuleDependcies entities in the database.
func (mdcb *ModuleDependciesCreateBulk) Save(ctx context.Context) ([]*ModuleDependcies, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mdcb.builders))
	nodes := make([]*ModuleDependcies, len(mdcb.builders))
	mutators := make([]Mutator, len(mdcb.builders))
	for i := range mdcb.builders {
		func(i int, root context.Context) {
			builder := mdcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ModuleDependciesMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mdcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mdcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mdcb *ModuleDependciesCreateBulk) SaveX(ctx context.Context) []*ModuleDependcies {
	v, err := mdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mdcb *ModuleDependciesCreateBulk) Exec(ctx context.Context) error {
	_, err := mdcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mdcb *ModuleDependciesCreateBulk) ExecX(ctx context.Context) {
	if err := mdcb.Exec(ctx); err != nil {
		panic(err)
	}
}
