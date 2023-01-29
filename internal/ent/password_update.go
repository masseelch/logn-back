// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/logn-soft/logn-back/internal/ent/password"
	"github.com/logn-soft/logn-back/internal/ent/predicate"
	"github.com/logn-soft/logn-back/internal/ent/user"
)

// PasswordUpdate is the builder for updating Password entities.
type PasswordUpdate struct {
	config
	hooks    []Hook
	mutation *PasswordMutation
}

// Where appends a list predicates to the PasswordUpdate builder.
func (pu *PasswordUpdate) Where(ps ...predicate.Password) *PasswordUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetHash sets the "hash" field.
func (pu *PasswordUpdate) SetHash(s string) *PasswordUpdate {
	pu.mutation.SetHash(s)
	return pu
}

// SetSalt sets the "salt" field.
func (pu *PasswordUpdate) SetSalt(s string) *PasswordUpdate {
	pu.mutation.SetSalt(s)
	return pu
}

// SetCreatedAt sets the "created_at" field.
func (pu *PasswordUpdate) SetCreatedAt(t time.Time) *PasswordUpdate {
	pu.mutation.SetCreatedAt(t)
	return pu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pu *PasswordUpdate) SetNillableCreatedAt(t *time.Time) *PasswordUpdate {
	if t != nil {
		pu.SetCreatedAt(*t)
	}
	return pu
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (pu *PasswordUpdate) AddUserIDs(ids ...int) *PasswordUpdate {
	pu.mutation.AddUserIDs(ids...)
	return pu
}

// AddUsers adds the "users" edges to the User entity.
func (pu *PasswordUpdate) AddUsers(u ...*User) *PasswordUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pu.AddUserIDs(ids...)
}

// Mutation returns the PasswordMutation object of the builder.
func (pu *PasswordUpdate) Mutation() *PasswordMutation {
	return pu.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (pu *PasswordUpdate) ClearUsers() *PasswordUpdate {
	pu.mutation.ClearUsers()
	return pu
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (pu *PasswordUpdate) RemoveUserIDs(ids ...int) *PasswordUpdate {
	pu.mutation.RemoveUserIDs(ids...)
	return pu
}

// RemoveUsers removes "users" edges to User entities.
func (pu *PasswordUpdate) RemoveUsers(u ...*User) *PasswordUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pu.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PasswordUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, PasswordMutation](ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PasswordUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PasswordUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PasswordUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PasswordUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   password.Table,
			Columns: password.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: password.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Hash(); ok {
		_spec.SetField(password.FieldHash, field.TypeString, value)
	}
	if value, ok := pu.mutation.Salt(); ok {
		_spec.SetField(password.FieldSalt, field.TypeString, value)
	}
	if value, ok := pu.mutation.CreatedAt(); ok {
		_spec.SetField(password.FieldCreatedAt, field.TypeTime, value)
	}
	if pu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   password.UsersTable,
			Columns: password.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedUsersIDs(); len(nodes) > 0 && !pu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   password.UsersTable,
			Columns: password.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   password.UsersTable,
			Columns: password.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{password.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PasswordUpdateOne is the builder for updating a single Password entity.
type PasswordUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PasswordMutation
}

// SetHash sets the "hash" field.
func (puo *PasswordUpdateOne) SetHash(s string) *PasswordUpdateOne {
	puo.mutation.SetHash(s)
	return puo
}

// SetSalt sets the "salt" field.
func (puo *PasswordUpdateOne) SetSalt(s string) *PasswordUpdateOne {
	puo.mutation.SetSalt(s)
	return puo
}

// SetCreatedAt sets the "created_at" field.
func (puo *PasswordUpdateOne) SetCreatedAt(t time.Time) *PasswordUpdateOne {
	puo.mutation.SetCreatedAt(t)
	return puo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (puo *PasswordUpdateOne) SetNillableCreatedAt(t *time.Time) *PasswordUpdateOne {
	if t != nil {
		puo.SetCreatedAt(*t)
	}
	return puo
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (puo *PasswordUpdateOne) AddUserIDs(ids ...int) *PasswordUpdateOne {
	puo.mutation.AddUserIDs(ids...)
	return puo
}

// AddUsers adds the "users" edges to the User entity.
func (puo *PasswordUpdateOne) AddUsers(u ...*User) *PasswordUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return puo.AddUserIDs(ids...)
}

// Mutation returns the PasswordMutation object of the builder.
func (puo *PasswordUpdateOne) Mutation() *PasswordMutation {
	return puo.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (puo *PasswordUpdateOne) ClearUsers() *PasswordUpdateOne {
	puo.mutation.ClearUsers()
	return puo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (puo *PasswordUpdateOne) RemoveUserIDs(ids ...int) *PasswordUpdateOne {
	puo.mutation.RemoveUserIDs(ids...)
	return puo
}

// RemoveUsers removes "users" edges to User entities.
func (puo *PasswordUpdateOne) RemoveUsers(u ...*User) *PasswordUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return puo.RemoveUserIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PasswordUpdateOne) Select(field string, fields ...string) *PasswordUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Password entity.
func (puo *PasswordUpdateOne) Save(ctx context.Context) (*Password, error) {
	return withHooks[*Password, PasswordMutation](ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PasswordUpdateOne) SaveX(ctx context.Context) *Password {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PasswordUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PasswordUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PasswordUpdateOne) sqlSave(ctx context.Context) (_node *Password, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   password.Table,
			Columns: password.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: password.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Password.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, password.FieldID)
		for _, f := range fields {
			if !password.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != password.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Hash(); ok {
		_spec.SetField(password.FieldHash, field.TypeString, value)
	}
	if value, ok := puo.mutation.Salt(); ok {
		_spec.SetField(password.FieldSalt, field.TypeString, value)
	}
	if value, ok := puo.mutation.CreatedAt(); ok {
		_spec.SetField(password.FieldCreatedAt, field.TypeTime, value)
	}
	if puo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   password.UsersTable,
			Columns: password.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !puo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   password.UsersTable,
			Columns: password.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   password.UsersTable,
			Columns: password.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Password{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{password.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}