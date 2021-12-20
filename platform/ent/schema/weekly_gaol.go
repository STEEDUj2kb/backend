package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Weekly_Gaol holds the schema definition for the Weekly_Gaol entity.
type Weekly_Gaol struct {
	ent.Schema
}

// Fields of the Weekly_Gaol.
func (Weekly_Gaol) Fields() []ent.Field {
	return []ent.Field{
		field.String("goal"),
		field.Int("score"),
		field.Int("nth"),
	}
}

// Edges of the Weekly_Gaol.
func (Weekly_Gaol) Edges() []ent.Edge {
	return []ent.Edge{
		// Weekly goal have multiple Daily Goals
		edge.To("dgaols", Daily_Gaol.Type),

		// Weeky Goal reference a Study
		edge.From("study", Study.Type).
			Ref("wgoals").
			Unique(),
	}
}
