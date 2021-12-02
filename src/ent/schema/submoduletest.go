package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SubModuleTest holds the schema definition for the SubModuleTest entity.
type SubModuleTest struct {
	ent.Schema
}

// Fields of the SubModuleTest.
func (SubModuleTest) Fields() []ent.Field {
	return []ent.Field{
		field.Int("submodule_id"),
	}
}

func (SubModuleTest) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "SubModuleTest"},
	}
}

// Edges of the SubModuleTest.
func (SubModuleTest) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("SubModule", SubModule.Type).
			Ref("Test").
			Unique().
			Field("submodule_id").
			Required(),
		edge.To("TherTest", TheoreticalTest.Type).
			Unique(),
	}
}
