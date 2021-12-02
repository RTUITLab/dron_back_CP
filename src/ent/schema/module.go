package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Module holds the schema definition for the Module entity.
type Module struct {
	ent.Schema
}

// Fields of the Module.
func (Module) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

func (Module) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Module"},
	}
}

// Edges of the Module.
func (Module) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("ModuleDependcies", ModuleDependcies.Type),
		edge.To("ModuleDependOn", ModuleDependcies.Type),
		edge.To("SubModules", SubModule.Type),
	}
}
