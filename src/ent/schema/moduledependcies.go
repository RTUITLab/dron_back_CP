package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ModuleDependcies holds the schema definition for the ModuleDependcies entity.
type ModuleDependcies struct {
	ent.Schema
}

// Fields of the ModuleDependcies.
func (ModuleDependcies) Fields() []ent.Field {
	return []ent.Field{
		field.Int("dependent_id"),
		field.Int("dependent_on_id"),
	}
}

func (ModuleDependcies) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "ModuleDependcies"},
	}
}

// Edges of the ModuleDependcies.
func (ModuleDependcies) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("ModuleDependcies", Module.Type).
			Ref("ModuleDependcies").
			Unique().
			Field("dependent_id").Required(),
		edge.From("ModuleDependOn", Module.Type).
			Ref("ModuleDependOn").
			Unique().
			Field("dependent_on_id").Required(),
	}
}
