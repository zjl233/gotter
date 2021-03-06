// Code generated by entc, DO NOT EDIT.

package authtoken

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/zjl233/gotter/ent/schema"
)

const (
	// Label holds the string label denoting the authtoken type in the database.
	Label = "auth_token"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at vertex property in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at vertex property in the database.
	FieldUpdatedAt = "updated_at"
	// FieldToken holds the string denoting the token vertex property in the database.
	FieldToken = "token"
	// FieldExpiresAt holds the string denoting the expires_at vertex property in the database.
	FieldExpiresAt = "expires_at"

	// Table holds the table name of the authtoken in the database.
	Table = "auth_tokens"
	// UserTable is the table the holds the user relation/edge.
	UserTable = "auth_tokens"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_tokens"
)

// Columns holds all SQL columns for authtoken fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldToken,
	FieldExpiresAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the AuthToken type.
var ForeignKeys = []string{
	"user_tokens",
}

var (
	mixin       = schema.AuthToken{}.Mixin()
	mixinFields = [...][]ent.Field{
		mixin[0].Fields(),
	}
	fields = schema.AuthToken{}.Fields()

	// descCreatedAt is the schema descriptor for created_at field.
	descCreatedAt = mixinFields[0][0].Descriptor()
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt = descCreatedAt.Default.(func() time.Time)

	// descUpdatedAt is the schema descriptor for updated_at field.
	descUpdatedAt = mixinFields[0][1].Descriptor()
	// DefaultUpdatedAt holds the default value on creation for the updated_at field.
	DefaultUpdatedAt = descUpdatedAt.Default.(func() time.Time)
	// UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	UpdateDefaultUpdatedAt = descUpdatedAt.UpdateDefault.(func() time.Time)
)
