package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// AuthToken holds the schema definition for the AuthToken entity.
type AuthToken struct {
	ent.Schema
}

// Mixin of the AuthToken.
func (AuthToken) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the AuthToken.
func (AuthToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("token").Sensitive(),
		field.Time("expires_at").Immutable().Nillable().Optional(),
	}
}

// Edges of the AuthToken.
func (AuthToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("tokens").Required().Unique(),
	}
}
