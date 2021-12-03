package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.Int("pract_test_id"),
		field.Text("task"),
	}
}

func (Task) Annotations() []schema.Annotation {
	return []schema.Annotation {
		entsql.Annotation{Table: "Task"},
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("Test", PractTest.Type).
			Ref("Task").
			Unique().
			Field("pract_test_id").
			Required(),
	}
}
