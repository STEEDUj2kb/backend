package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// DailyGaol holds the schema definition for the DailyGaol entity.
type DailyGaol struct {
	ent.Schema
}

// Fields of the DailyGaol.
func (DailyGaol) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.String("todo").
			NotEmpty(),
		field.Bool("done").
			Default(false),
		field.Bool("is_removed").
			Default(false),
	}
}

// Edges of the DailyGaol.
func (DailyGaol) Edges() []ent.Edge {
	return []ent.Edge{
		// Daily Goal reference a Study
		edge.From("study", Study.Type).
			Ref("dgoals").
			Unique(),

		// Daily Goal reference a Weeky Goal
		edge.From("wgoal", WeeklyGaol.Type).
			Ref("dgaols").
			Unique(),
	}
}
