package schema

import (
	"time"

	"github.com/STEEDUj2kb/v1/pkg/repository"
	"github.com/google/uuid"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		// field.Int("id").
		// 	StructTag(`json:"-"`),
		field.UUID("uuid", uuid.UUID{}).
			Default(uuid.New),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.String("name"),
		field.String("email").
			Unique(),
		field.Int("user_status").
			Max(1).
			Default(1),
		field.Enum("user_role").
			Values(string(repository.Admin), string(repository.User)).
			Default(string(repository.User)),
		field.String("password_hash").
			Sensitive(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
