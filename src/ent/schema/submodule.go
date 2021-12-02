package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SubModule holds the schema definition for the SubModule entity.
type SubModule struct {
	ent.Schema
}

// Fields of the SubModule.
func (SubModule) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Text("text"),
	}
}

func (SubModule) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "SubModule"},
	}
}

// Edges of the SubModule.
func (SubModule) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Module", Module.Type).
			Ref("SubModules").
			Unique(),
		edge.To("Test", SubModuleTest.Type).
			Unique(),
	}
}
