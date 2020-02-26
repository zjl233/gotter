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
		field.String("account").
			Unique().
			Immutable().
			NotEmpty().
			MaxLen(15).
			Match(regexp.MustCompile("[0-9a-zA-Z_]+$")),
		field.String("password_hash").
			Sensitive(),
		field.String("name").
			NotEmpty().
			MaxLen(15),
		field.String("profile_img").
			Default("/static/img/default-user-profile-img.png"),
		field.String("bkg_wall_img").
			Default("/static/img/default-user-bkg-img.jpg"),

		//field.Enum("status").Values("active", "inactive", "suspend").Default("active"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tokens", AuthToken.Type),
		edge.To("posts", Post.Type),
		edge.To("comments", Comment.Type),
		edge.To("following", User.Type).
			From("followers"),
	}
}
