package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("content").NotEmpty().MaxLen(255),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return nil
}
