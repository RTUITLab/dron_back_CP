// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/0B1t322/CP-Rosseti-Back/ent/predicate"
	"github.com/0B1t322/CP-Rosseti-Back/ent/role"
	"github.com/0B1t322/CP-Rosseti-Back/ent/theoreticaltry"
	"github.com/0B1t322/CP-Rosseti-Back/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetLogin sets the "login" field.
func (uu *UserUpdate) SetLogin(s string) *UserUpdate {
	uu.mutation.SetLogin(s)
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetRoleID sets the "Role" edge to the Role entity by ID.
func (uu *UserUpdate) SetRoleID(id int) *UserUpdate {
	uu.mutation.SetRoleID(id)
	return uu
}

// SetNillableRoleID sets the "Role" edge to the Role entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillableRoleID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetRoleID(*id)
	}
	return uu
}

// SetRole sets the "Role" edge to the Role entity.
func (uu *UserUpdate) SetRole(r *Role) *UserUpdate {
	return uu.SetRoleID(r.ID)
}

// AddTheoTryIDs adds the "TheoTry" edge to the TheoreticalTry entity by IDs.
func (uu *UserUpdate) AddTheoTryIDs(ids ...int) *UserUpdate {
	uu.mutation.AddTheoTryIDs(ids...)
	return uu
}

// AddTheoTry adds the "TheoTry" edges to the TheoreticalTry entity.
func (uu *UserUpdate) AddTheoTry(t ...*TheoreticalTry) *UserUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.AddTheoTryIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearRole clears the "Role" edge to the Role entity.
func (uu *UserUpdate) ClearRole() *UserUpdate {
	uu.mutation.ClearRole()
	return uu
}

// ClearTheoTry clears all "TheoTry" edges to the TheoreticalTry entity.
func (uu *UserUpdate) ClearTheoTry() *UserUpdate {
	uu.mutation.ClearTheoTry()
	return uu
}

// RemoveTheoTryIDs removes the "TheoTry" edge to TheoreticalTry entities by IDs.
func (uu *UserUpdate) RemoveTheoTryIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveTheoTryIDs(ids...)
	return uu
}

// RemoveTheoTry removes "TheoTry" edges to TheoreticalTry entities.
func (uu *UserUpdate) RemoveTheoTry(t ...*TheoreticalTry) *UserUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.RemoveTheoTryIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uu.hooks) == 0 {
		if err = uu.check(); err != nil {
			return 0, err
		}
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uu.check(); err != nil {
				return 0, err
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			if uu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Login(); ok {
		if err := user.LoginValidator(v); err != nil {
			return &ValidationError{Name: "login", err: fmt.Errorf("ent: validator failed for field \"login\": %w", err)}
		}
	}
	if v, ok := uu.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf("ent: validator failed for field \"password\": %w", err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Login(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldLogin,
		})
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPassword,
		})
	}
	if uu.mutation.RoleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.RoleTable,
			Columns: []string{user.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: role.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.RoleTable,
			Columns: []string{user.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: role.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.TheoTryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TheoTryTable,
			Columns: []string{user.TheoTryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: theoreticaltry.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedTheoTryIDs(); len(nodes) > 0 && !uu.mutation.TheoTryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TheoTryTable,
			Columns: []string{user.TheoTryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: theoreticaltry.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.TheoTryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TheoTryTable,
			Columns: []string{user.TheoTryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: theoreticaltry.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetLogin sets the "login" field.
func (uuo *UserUpdateOne) SetLogin(s string) *UserUpdateOne {
	uuo.mutation.SetLogin(s)
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetRoleID sets the "Role" edge to the Role entity by ID.
func (uuo *UserUpdateOne) SetRoleID(id int) *UserUpdateOne {
	uuo.mutation.SetRoleID(id)
	return uuo
}

// SetNillableRoleID sets the "Role" edge to the Role entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableRoleID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetRoleID(*id)
	}
	return uuo
}

// SetRole sets the "Role" edge to the Role entity.
func (uuo *UserUpdateOne) SetRole(r *Role) *UserUpdateOne {
	return uuo.SetRoleID(r.ID)
}

// AddTheoTryIDs adds the "TheoTry" edge to the TheoreticalTry entity by IDs.
func (uuo *UserUpdateOne) AddTheoTryIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddTheoTryIDs(ids...)
	return uuo
}

// AddTheoTry adds the "TheoTry" edges to the TheoreticalTry entity.
func (uuo *UserUpdateOne) AddTheoTry(t ...*TheoreticalTry) *UserUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.AddTheoTryIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearRole clears the "Role" edge to the Role entity.
func (uuo *UserUpdateOne) ClearRole() *UserUpdateOne {
	uuo.mutation.ClearRole()
	return uuo
}

// ClearTheoTry clears all "TheoTry" edges to the TheoreticalTry entity.
func (uuo *UserUpdateOne) ClearTheoTry() *UserUpdateOne {
	uuo.mutation.ClearTheoTry()
	return uuo
}

// RemoveTheoTryIDs removes the "TheoTry" edge to TheoreticalTry entities by IDs.
func (uuo *UserUpdateOne) RemoveTheoTryIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveTheoTryIDs(ids...)
	return uuo
}

// RemoveTheoTry removes "TheoTry" edges to TheoreticalTry entities.
func (uuo *UserUpdateOne) RemoveTheoTry(t ...*TheoreticalTry) *UserUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.RemoveTheoTryIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	if len(uuo.hooks) == 0 {
		if err = uuo.check(); err != nil {
			return nil, err
		}
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uuo.check(); err != nil {
				return nil, err
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			if uuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Login(); ok {
		if err := user.LoginValidator(v); err != nil {
			return &ValidationError{Name: "login", err: fmt.Errorf("ent: validator failed for field \"login\": %w", err)}
		}
	}
	if v, ok := uuo.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf("ent: validator failed for field \"password\": %w", err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing User.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Login(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldLogin,
		})
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPassword,
		})
	}
	if uuo.mutation.RoleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.RoleTable,
			Columns: []string{user.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: role.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.RoleTable,
			Columns: []string{user.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: role.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.TheoTryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TheoTryTable,
			Columns: []string{user.TheoTryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: theoreticaltry.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedTheoTryIDs(); len(nodes) > 0 && !uuo.mutation.TheoTryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TheoTryTable,
			Columns: []string{user.TheoTryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: theoreticaltry.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.TheoTryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TheoTryTable,
			Columns: []string{user.TheoTryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: theoreticaltry.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
