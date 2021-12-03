package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TryAnswer holds the schema definition for the TryAnswer entity.
type TryAnswer struct {
	ent.Schema
}

// Fields of the TryAnswer.
func (TryAnswer) Fields() []ent.Field {
	return []ent.Field{
		field.Int("try_id"),
		field.Int("question_id"),
		field.Text("asnwer"),
		field.Bool("correct").Optional(),
	}
}

func (TryAnswer) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "TryAnswer"},
	}
}

// Edges of the TryAnswer.
func (TryAnswer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("TheoreticalTry", TheoreticalTry.Type).
			Ref("TryAnswer").
			Unique().
			Field("try_id").
			Required(),
		edge.From("Question", Question.Type).
			Ref("TryAnswer").
			Unique().
			Field("question_id").
			Required(),
	}
}
