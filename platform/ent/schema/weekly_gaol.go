package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// WeeklyGaol holds the schema definition for the WeeklyGaol entity.
type WeeklyGaol struct {
	ent.Schema
}

// Fields of the WeeklyGaol.
func (WeeklyGaol) Fields() []ent.Field {
	return []ent.Field{
		field.String("goal"),
		field.Int("score").
			Default(0).
			Optional(),
		field.Int("nth"),
	}
}

// Edges of the WeeklyGaol.
func (WeeklyGaol) Edges() []ent.Edge {
	return []ent.Edge{
		// Weekly goal have multiple Daily Goals
		edge.To("dgaols", DailyGaol.Type),

		// Weeky Goal reference a Study
		edge.From("study", Study.Type).
			Ref("wgoals").
			Unique(),
	}
}
