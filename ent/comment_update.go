// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/zjl233/gotter/ent/comment"
	"github.com/zjl233/gotter/ent/post"
	"github.com/zjl233/gotter/ent/predicate"
	"github.com/zjl233/gotter/ent/user"
)

// CommentUpdate is the builder for updating Comment entities.
type CommentUpdate struct {
	config
	content       *string
	author        map[int]struct{}
	post          map[int]struct{}
	clearedAuthor bool
	clearedPost   bool
	predicates    []predicate.Comment
}

// Where adds a new predicate for the builder.
func (cu *CommentUpdate) Where(ps ...predicate.Comment) *CommentUpdate {
	cu.predicates = append(cu.predicates, ps...)
	return cu
}

// SetContent sets the content field.
func (cu *CommentUpdate) SetContent(s string) *CommentUpdate {
	cu.content = &s
	return cu
}

// SetAuthorID sets the author edge to User by id.
func (cu *CommentUpdate) SetAuthorID(id int) *CommentUpdate {
	if cu.author == nil {
		cu.author = make(map[int]struct{})
	}
	cu.author[id] = struct{}{}
	return cu
}

// SetAuthor sets the author edge to User.
func (cu *CommentUpdate) SetAuthor(u *User) *CommentUpdate {
	return cu.SetAuthorID(u.ID)
}

// SetPostID sets the post edge to Post by id.
func (cu *CommentUpdate) SetPostID(id int) *CommentUpdate {
	if cu.post == nil {
		cu.post = make(map[int]struct{})
	}
	cu.post[id] = struct{}{}
	return cu
}

// SetPost sets the post edge to Post.
func (cu *CommentUpdate) SetPost(p *Post) *CommentUpdate {
	return cu.SetPostID(p.ID)
}

// ClearAuthor clears the author edge to User.
func (cu *CommentUpdate) ClearAuthor() *CommentUpdate {
	cu.clearedAuthor = true
	return cu
}

// ClearPost clears the post edge to Post.
func (cu *CommentUpdate) ClearPost() *CommentUpdate {
	cu.clearedPost = true
	return cu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (cu *CommentUpdate) Save(ctx context.Context) (int, error) {
	if cu.content != nil {
		if err := comment.ContentValidator(*cu.content); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"content\": %v", err)
		}
	}
	if len(cu.author) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"author\"")
	}
	if cu.clearedAuthor && cu.author == nil {
		return 0, errors.New("ent: clearing a unique edge \"author\"")
	}
	if len(cu.post) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"post\"")
	}
	if cu.clearedPost && cu.post == nil {
		return 0, errors.New("ent: clearing a unique edge \"post\"")
	}
	return cu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CommentUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CommentUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CommentUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CommentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   comment.Table,
			Columns: comment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: comment.FieldID,
			},
		},
	}
	if ps := cu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value := cu.content; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: comment.FieldContent,
		})
	}
	if cu.clearedAuthor {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.author; len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.clearedPost {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.post; len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// CommentUpdateOne is the builder for updating a single Comment entity.
type CommentUpdateOne struct {
	config
	id            int
	content       *string
	author        map[int]struct{}
	post          map[int]struct{}
	clearedAuthor bool
	clearedPost   bool
}

// SetContent sets the content field.
func (cuo *CommentUpdateOne) SetContent(s string) *CommentUpdateOne {
	cuo.content = &s
	return cuo
}

// SetAuthorID sets the author edge to User by id.
func (cuo *CommentUpdateOne) SetAuthorID(id int) *CommentUpdateOne {
	if cuo.author == nil {
		cuo.author = make(map[int]struct{})
	}
	cuo.author[id] = struct{}{}
	return cuo
}

// SetAuthor sets the author edge to User.
func (cuo *CommentUpdateOne) SetAuthor(u *User) *CommentUpdateOne {
	return cuo.SetAuthorID(u.ID)
}

// SetPostID sets the post edge to Post by id.
func (cuo *CommentUpdateOne) SetPostID(id int) *CommentUpdateOne {
	if cuo.post == nil {
		cuo.post = make(map[int]struct{})
	}
	cuo.post[id] = struct{}{}
	return cuo
}

// SetPost sets the post edge to Post.
func (cuo *CommentUpdateOne) SetPost(p *Post) *CommentUpdateOne {
	return cuo.SetPostID(p.ID)
}

// ClearAuthor clears the author edge to User.
func (cuo *CommentUpdateOne) ClearAuthor() *CommentUpdateOne {
	cuo.clearedAuthor = true
	return cuo
}

// ClearPost clears the post edge to Post.
func (cuo *CommentUpdateOne) ClearPost() *CommentUpdateOne {
	cuo.clearedPost = true
	return cuo
}

// Save executes the query and returns the updated entity.
func (cuo *CommentUpdateOne) Save(ctx context.Context) (*Comment, error) {
	if cuo.content != nil {
		if err := comment.ContentValidator(*cuo.content); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"content\": %v", err)
		}
	}
	if len(cuo.author) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"author\"")
	}
	if cuo.clearedAuthor && cuo.author == nil {
		return nil, errors.New("ent: clearing a unique edge \"author\"")
	}
	if len(cuo.post) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"post\"")
	}
	if cuo.clearedPost && cuo.post == nil {
		return nil, errors.New("ent: clearing a unique edge \"post\"")
	}
	return cuo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CommentUpdateOne) SaveX(ctx context.Context) *Comment {
	c, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return c
}

// Exec executes the query on the entity.
func (cuo *CommentUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CommentUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CommentUpdateOne) sqlSave(ctx context.Context) (c *Comment, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   comment.Table,
			Columns: comment.Columns,
			ID: &sqlgraph.FieldSpec{
				Value:  cuo.id,
				Type:   field.TypeInt,
				Column: comment.FieldID,
			},
		},
	}
	if value := cuo.content; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: comment.FieldContent,
		})
	}
	if cuo.clearedAuthor {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.author; len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.clearedPost {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.post; len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	c = &Comment{config: cuo.config}
	_spec.Assign = c.assignValues
	_spec.ScanValues = c.scanValues()
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return c, nil
}
