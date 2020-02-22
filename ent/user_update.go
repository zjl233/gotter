// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/zjl233/gotter/ent/authtoken"
	"github.com/zjl233/gotter/ent/predicate"
	"github.com/zjl233/gotter/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config

	updated_at *time.Time

	password_hash *string
	nickname      *string
	status        *user.Status
	tokens        map[int]struct{}
	removedTokens map[int]struct{}
	predicates    []predicate.User
}

// Where adds a new predicate for the builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.predicates = append(uu.predicates, ps...)
	return uu
}

// SetUpdatedAt sets the updated_at field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.updated_at = &t
	return uu
}

// SetPasswordHash sets the password_hash field.
func (uu *UserUpdate) SetPasswordHash(s string) *UserUpdate {
	uu.password_hash = &s
	return uu
}

// SetNickname sets the nickname field.
func (uu *UserUpdate) SetNickname(s string) *UserUpdate {
	uu.nickname = &s
	return uu
}

// SetStatus sets the status field.
func (uu *UserUpdate) SetStatus(u user.Status) *UserUpdate {
	uu.status = &u
	return uu
}

// SetNillableStatus sets the status field if the given value is not nil.
func (uu *UserUpdate) SetNillableStatus(u *user.Status) *UserUpdate {
	if u != nil {
		uu.SetStatus(*u)
	}
	return uu
}

// AddTokenIDs adds the tokens edge to AuthToken by ids.
func (uu *UserUpdate) AddTokenIDs(ids ...int) *UserUpdate {
	if uu.tokens == nil {
		uu.tokens = make(map[int]struct{})
	}
	for i := range ids {
		uu.tokens[ids[i]] = struct{}{}
	}
	return uu
}

// AddTokens adds the tokens edges to AuthToken.
func (uu *UserUpdate) AddTokens(a ...*AuthToken) *UserUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uu.AddTokenIDs(ids...)
}

// RemoveTokenIDs removes the tokens edge to AuthToken by ids.
func (uu *UserUpdate) RemoveTokenIDs(ids ...int) *UserUpdate {
	if uu.removedTokens == nil {
		uu.removedTokens = make(map[int]struct{})
	}
	for i := range ids {
		uu.removedTokens[ids[i]] = struct{}{}
	}
	return uu
}

// RemoveTokens removes tokens edges to AuthToken.
func (uu *UserUpdate) RemoveTokens(a ...*AuthToken) *UserUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uu.RemoveTokenIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	if uu.updated_at == nil {
		v := user.UpdateDefaultUpdatedAt()
		uu.updated_at = &v
	}
	if uu.nickname != nil {
		if err := user.NicknameValidator(*uu.nickname); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"nickname\": %v", err)
		}
	}
	if uu.status != nil {
		if err := user.StatusValidator(*uu.status); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"status\": %v", err)
		}
	}
	return uu.sqlSave(ctx)
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
	if ps := uu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value := uu.updated_at; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: user.FieldUpdatedAt,
		})
	}
	if value := uu.password_hash; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldPasswordHash,
		})
	}
	if value := uu.nickname; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldNickname,
		})
	}
	if value := uu.status; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  *value,
			Column: user.FieldStatus,
		})
	}
	if nodes := uu.removedTokens; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TokensTable,
			Columns: []string{user.TokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: authtoken.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.tokens; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TokensTable,
			Columns: []string{user.TokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: authtoken.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	id int

	updated_at *time.Time

	password_hash *string
	nickname      *string
	status        *user.Status
	tokens        map[int]struct{}
	removedTokens map[int]struct{}
}

// SetUpdatedAt sets the updated_at field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.updated_at = &t
	return uuo
}

// SetPasswordHash sets the password_hash field.
func (uuo *UserUpdateOne) SetPasswordHash(s string) *UserUpdateOne {
	uuo.password_hash = &s
	return uuo
}

// SetNickname sets the nickname field.
func (uuo *UserUpdateOne) SetNickname(s string) *UserUpdateOne {
	uuo.nickname = &s
	return uuo
}

// SetStatus sets the status field.
func (uuo *UserUpdateOne) SetStatus(u user.Status) *UserUpdateOne {
	uuo.status = &u
	return uuo
}

// SetNillableStatus sets the status field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableStatus(u *user.Status) *UserUpdateOne {
	if u != nil {
		uuo.SetStatus(*u)
	}
	return uuo
}

// AddTokenIDs adds the tokens edge to AuthToken by ids.
func (uuo *UserUpdateOne) AddTokenIDs(ids ...int) *UserUpdateOne {
	if uuo.tokens == nil {
		uuo.tokens = make(map[int]struct{})
	}
	for i := range ids {
		uuo.tokens[ids[i]] = struct{}{}
	}
	return uuo
}

// AddTokens adds the tokens edges to AuthToken.
func (uuo *UserUpdateOne) AddTokens(a ...*AuthToken) *UserUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uuo.AddTokenIDs(ids...)
}

// RemoveTokenIDs removes the tokens edge to AuthToken by ids.
func (uuo *UserUpdateOne) RemoveTokenIDs(ids ...int) *UserUpdateOne {
	if uuo.removedTokens == nil {
		uuo.removedTokens = make(map[int]struct{})
	}
	for i := range ids {
		uuo.removedTokens[ids[i]] = struct{}{}
	}
	return uuo
}

// RemoveTokens removes tokens edges to AuthToken.
func (uuo *UserUpdateOne) RemoveTokens(a ...*AuthToken) *UserUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uuo.RemoveTokenIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	if uuo.updated_at == nil {
		v := user.UpdateDefaultUpdatedAt()
		uuo.updated_at = &v
	}
	if uuo.nickname != nil {
		if err := user.NicknameValidator(*uuo.nickname); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"nickname\": %v", err)
		}
	}
	if uuo.status != nil {
		if err := user.StatusValidator(*uuo.status); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"status\": %v", err)
		}
	}
	return uuo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	u, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return u
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

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (u *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Value:  uuo.id,
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if value := uuo.updated_at; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: user.FieldUpdatedAt,
		})
	}
	if value := uuo.password_hash; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldPasswordHash,
		})
	}
	if value := uuo.nickname; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldNickname,
		})
	}
	if value := uuo.status; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  *value,
			Column: user.FieldStatus,
		})
	}
	if nodes := uuo.removedTokens; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TokensTable,
			Columns: []string{user.TokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: authtoken.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.tokens; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TokensTable,
			Columns: []string{user.TokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: authtoken.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	u = &User{config: uuo.config}
	_spec.Assign = u.assignValues
	_spec.ScanValues = u.scanValues()
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return u, nil
}
