package schema

import (
	"entgo.io/ent"

	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PractTest holds the schema definition for the PractTest entity.
type PractTest struct {
	ent.Schema
}

type JSONObject map[string]interface{}

// Fields of the PractTest.
func (PractTest) Fields() []ent.Field {
	return []ent.Field{
		field.Int("test_id"),
		field.JSON("config", JSONObject{}),
	}
}

func (PractTest) Annotations() []schema.Annotation {
	return []schema.Annotation {
		entsql.Annotation{Table: "PractTest"},
	}
}

// Edges of the PractTest.
func (PractTest) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Test", Test.Type).
			Ref("PractTest").
			Unique().
			Field("test_id").
			Required(),
	}
}
