package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"regexp"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Immutable().Unique().
			MinLen(1).MaxLen(15).
			Match(regexp.MustCompile("[a-zA-Z_]+$")),
		field.String("password_hash").Sensitive(),
		field.String("nickname").MinLen(1).MaxLen(31),
		field.Enum("status").Values("active", "inactive", "suspend").Default("active"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tokens", AuthToken.Type),
	}
}
