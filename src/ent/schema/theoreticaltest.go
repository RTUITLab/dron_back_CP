package schema

import (
	"entgo.io/ent"

	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TheoreticalTest holds the schema definition for the TheoreticalTest entity.
type TheoreticalTest struct {
	ent.Schema
}

// Fields of the TheoreticalTest.
func (TheoreticalTest) Fields() []ent.Field {
	return []ent.Field{
		field.Int("submoduletest_id"),
	}
}

func (TheoreticalTest) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "TheoreticalTest"},
	}
}

// Edges of the TheoreticalTest.
func (TheoreticalTest) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("SubModuleTest", SubModuleTest.Type).
			Ref("TherTest").
			Unique().
			Field("submoduletest_id").
			Required(),
		edge.To("Question", Question.Type),
	}
}
