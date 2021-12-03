package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// PractTry holds the schema definition for the PractTry entity.
type PractTry struct {
	ent.Schema
}

// Fields of the PractTry.
func (PractTry) Fields() []ent.Field {
	return []ent.Field{
		field.Time("start"),
		field.Time("end").Optional(),
		field.Int("user_id"),
		field.Int("pract_test_id"),
	}
}

// Edges of the PractTry.
func (PractTry) Edges() []ent.Edge {
	return nil
}
