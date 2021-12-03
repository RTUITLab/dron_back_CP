package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Test holds the schema definition for the Test entity.
type Test struct {
	ent.Schema
}

// Fields of the Test.
func (Test) Fields() []ent.Field {
	return []ent.Field{
		field.String("TestType"),
	}
}

func (Test) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Test"},
	}
}

// Edges of the Test.
func (Test) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("ModuleTest", ModuleTest.Type).
			Annotations(
				entsql.Annotation{OnDelete: entsql.Cascade},
			),
		edge.To("SubmoduleTest", SubModuleTest.Type).
			Annotations(
				entsql.Annotation{OnDelete: entsql.Cascade},
			),
		edge.To("TherTest", TheoreticalTest.Type).
			Annotations(
				entsql.Annotation{OnDelete: entsql.Cascade},
			),
		edge.To("PractTest", PractTest.Type).
			Annotations(
				entsql.Annotation{OnDelete: entsql.Cascade},
			),
	}
}
