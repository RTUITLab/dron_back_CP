package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// PractTryAnswer holds the schema definition for the PractTryAnswer entity.
type PractTryAnswer struct {
	ent.Schema
}

// Fields of the PractTryAnswer.
func (PractTryAnswer) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("correct"),
		field.Int("task_id"),
		field.Int("try_id"),
	}
}

// Edges of the PractTryAnswer.
func (PractTryAnswer) Edges() []ent.Edge {
	return nil
}
