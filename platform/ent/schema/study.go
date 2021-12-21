package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Study holds the schema definition for the Study entity.
type Study struct {
	ent.Schema
}

// Fields of the Study.
func (Study) Fields() []ent.Field {
	return []ent.Field{
		field.Time("started_at").
			Default(time.Now).
			Immutable(),
		field.Time("ended_at").
			Default(nil).
			Optional(),
		field.String("content").
			NotEmpty(),
	}
}

// Edges of the Study.
func (Study) Edges() []ent.Edge {
	return []ent.Edge{
		// Study reference a User
		edge.From("planner", User.Type).
			Ref("studies").
			Unique(),

		// User have many daily goals
		edge.To("dgoals", DailyGaol.Type),

		// User have many weekly goals
		edge.To("wgoals", WeeklyGaol.Type),
	}
}
