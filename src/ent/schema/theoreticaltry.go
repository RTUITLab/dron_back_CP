package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TheoreticalTry holds the schema definition for the TheoreticalTry entity.
type TheoreticalTry struct {
	ent.Schema
}

// Fields of the TheoreticalTry.
func (TheoreticalTry) Fields() []ent.Field {
	return []ent.Field{
		field.Time("start"),
		field.Time("end").Optional(),
		field.Int("user_id"),
		field.Int("theoretical_test_id"),
	}
}

func (TheoreticalTry) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "TheoreticalTry"},
	}
}

// Edges of the TheoreticalTry.
func (TheoreticalTry) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("TheoreticalTest", TheoreticalTest.Type).
			Ref("TheoTry").
			Unique().
			Field("theoretical_test_id").
			Required(),
		edge.From("User", User.Type).
			Ref("TheoTry").
			Unique().
			Field("user_id").
			Required(),
		edge.To("TryAnswer", TryAnswer.Type),
	}
}
