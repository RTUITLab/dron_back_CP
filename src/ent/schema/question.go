package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Question holds the schema definition for the Question entity.
type Question struct {
	ent.Schema
}

// Fields of the Question.
func (Question) Fields() []ent.Field {
	return []ent.Field{
		field.Int("theorical_test_id"),
		field.Text("question"),
	}
}

func (Question) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Question"},
	}
}


// Edges of the Question.
func (Question) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("TheoreticalTest", TheoreticalTest.Type).
			Ref("Question").
			Unique().
			Field("theorical_test_id").
			Required(),
		edge.To("Answer", Answer.Type).
			Annotations(
				entsql.Annotation{OnDelete: entsql.Cascade},
			),
		edge.To("TryAnswer", TryAnswer.Type),
	}
}
