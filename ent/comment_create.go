// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/zjl233/gotter/ent/comment"
	"github.com/zjl233/gotter/ent/post"
	"github.com/zjl233/gotter/ent/user"
)

// CommentCreate is the builder for creating a Comment entity.
type CommentCreate struct {
	config
	created_at *time.Time
	updated_at *time.Time
	content    *string
	author     map[int]struct{}
	post       map[int]struct{}
}

// SetCreatedAt sets the created_at field.
func (cc *CommentCreate) SetCreatedAt(t time.Time) *CommentCreate {
	cc.created_at = &t
	return cc
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (cc *CommentCreate) SetNillableCreatedAt(t *time.Time) *CommentCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the updated_at field.
func (cc *CommentCreate) SetUpdatedAt(t time.Time) *CommentCreate {
	cc.updated_at = &t
	return cc
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (cc *CommentCreate) SetNillableUpdatedAt(t *time.Time) *CommentCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetContent sets the content field.
func (cc *CommentCreate) SetContent(s string) *CommentCreate {
	cc.content = &s
	return cc
}

// SetAuthorID sets the author edge to User by id.
func (cc *CommentCreate) SetAuthorID(id int) *CommentCreate {
	if cc.author == nil {
		cc.author = make(map[int]struct{})
	}
	cc.author[id] = struct{}{}
	return cc
}

// SetAuthor sets the author edge to User.
func (cc *CommentCreate) SetAuthor(u *User) *CommentCreate {
	return cc.SetAuthorID(u.ID)
}

// SetPostID sets the post edge to Post by id.
func (cc *CommentCreate) SetPostID(id int) *CommentCreate {
	if cc.post == nil {
		cc.post = make(map[int]struct{})
	}
	cc.post[id] = struct{}{}
	return cc
}

// SetPost sets the post edge to Post.
func (cc *CommentCreate) SetPost(p *Post) *CommentCreate {
	return cc.SetPostID(p.ID)
}

// Save creates the Comment in the database.
func (cc *CommentCreate) Save(ctx context.Context) (*Comment, error) {
	if cc.created_at == nil {
		v := comment.DefaultCreatedAt()
		cc.created_at = &v
	}
	if cc.updated_at == nil {
		v := comment.DefaultUpdatedAt()
		cc.updated_at = &v
	}
	if cc.content == nil {
		return nil, errors.New("ent: missing required field \"content\"")
	}
	if err := comment.ContentValidator(*cc.content); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"content\": %v", err)
	}
	if len(cc.author) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"author\"")
	}
	if cc.author == nil {
		return nil, errors.New("ent: missing required edge \"author\"")
	}
	if len(cc.post) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"post\"")
	}
	if cc.post == nil {
		return nil, errors.New("ent: missing required edge \"post\"")
	}
	return cc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CommentCreate) SaveX(ctx context.Context) *Comment {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cc *CommentCreate) sqlSave(ctx context.Context) (*Comment, error) {
	var (
		c     = &Comment{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: comment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: comment.FieldID,
			},
		}
	)
	if value := cc.created_at; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: comment.FieldCreatedAt,
		})
		c.CreatedAt = *value
	}
	if value := cc.updated_at; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: comment.FieldUpdatedAt,
		})
		c.UpdatedAt = *value
	}
	if value := cc.content; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: comment.FieldContent,
		})
		c.Content = *value
	}
	if nodes := cc.author; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.AuthorTable,
			Columns: []string{comment.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.post; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.PostTable,
			Columns: []string{comment.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	c.ID = int(id)
	return c, nil
}
