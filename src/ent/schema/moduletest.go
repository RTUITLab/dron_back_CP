package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ModuleTest holds the schema definition for the ModuleTest entity.
type ModuleTest struct {
	ent.Schema
}

// Fields of the ModuleTest.
func (ModuleTest) Fields() []ent.Field {
	return []ent.Field{
		field.Int("module_id"),
		field.Int("test_id"),
	}
}

func (ModuleTest) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "ModuleTest"},
	}
}

// Edges of the ModuleTest.
func (ModuleTest) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Module", Module.Type).
			Ref("Test").
			Unique().
			Field("module_id").
			Required(),
		edge.From("Test", Test.Type).
			Ref("ModuleTest").
			Unique().
			Field("test_id").
			Required(),
	}
}
