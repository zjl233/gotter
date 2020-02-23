// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/zjl233/gotter/ent/authtoken"
	"github.com/zjl233/gotter/ent/user"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	created_at    *time.Time
	updated_at    *time.Time
	username      *string
	password_hash *string
	tokens        map[int]struct{}
}

// SetCreatedAt sets the created_at field.
func (uc *UserCreate) SetCreatedAt(t time.Time) *UserCreate {
	uc.created_at = &t
	return uc
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (uc *UserCreate) SetNillableCreatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreatedAt(*t)
	}
	return uc
}

// SetUpdatedAt sets the updated_at field.
func (uc *UserCreate) SetUpdatedAt(t time.Time) *UserCreate {
	uc.updated_at = &t
	return uc
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUpdatedAt(*t)
	}
	return uc
}

// SetUsername sets the username field.
func (uc *UserCreate) SetUsername(s string) *UserCreate {
	uc.username = &s
	return uc
}

// SetPasswordHash sets the password_hash field.
func (uc *UserCreate) SetPasswordHash(s string) *UserCreate {
	uc.password_hash = &s
	return uc
}

// AddTokenIDs adds the tokens edge to AuthToken by ids.
func (uc *UserCreate) AddTokenIDs(ids ...int) *UserCreate {
	if uc.tokens == nil {
		uc.tokens = make(map[int]struct{})
	}
	for i := range ids {
		uc.tokens[ids[i]] = struct{}{}
	}
	return uc
}

// AddTokens adds the tokens edges to AuthToken.
func (uc *UserCreate) AddTokens(a ...*AuthToken) *UserCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uc.AddTokenIDs(ids...)
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	if uc.created_at == nil {
		v := user.DefaultCreatedAt()
		uc.created_at = &v
	}
	if uc.updated_at == nil {
		v := user.DefaultUpdatedAt()
		uc.updated_at = &v
	}
	if uc.username == nil {
		return nil, errors.New("ent: missing required field \"username\"")
	}
	if err := user.UsernameValidator(*uc.username); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"username\": %v", err)
	}
	if uc.password_hash == nil {
		return nil, errors.New("ent: missing required field \"password_hash\"")
	}
	return uc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	var (
		u     = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		}
	)
	if value := uc.created_at; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: user.FieldCreatedAt,
		})
		u.CreatedAt = *value
	}
	if value := uc.updated_at; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: user.FieldUpdatedAt,
		})
		u.UpdatedAt = *value
	}
	if value := uc.username; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldUsername,
		})
		u.Username = *value
	}
	if value := uc.password_hash; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldPasswordHash,
		})
		u.PasswordHash = *value
	}
	if nodes := uc.tokens; len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	u.ID = int(id)
	return u, nil
}
