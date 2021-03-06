// Code generated by entc, DO NOT EDIT.

package tryanswer

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/0B1t322/CP-Rosseti-Back/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.TryAnswer {
	return predicate.TryAnswer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.TryAnswer {
	return predicate.TryAnswer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.TryAnswer {
	return predicate.TryAnswer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.TryAnswer {
	return predicate.TryAnswer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.TryAnswer {
	return predicate.TryAnswer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.TryAnswer {
	return predicate.TryAnswer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.TryAnswer {
	return predicate.TryAnswer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.TryAnswer {
	return predicate.TryAnswer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.TryAnswer {
	return predicate.TryAnswer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// TryID applies equality check predicate on the "try_id" field. It's identical to TryIDEQ.
func TryID(v int) predicate.TryAnswer {
	return predicate.TryAnswer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTryID), v))
	})
}

// QuestionID applies equality check predicate on the "question_id" field. It's identical to QuestionIDEQ.
func QuestionID(v int) predicate.TryAnswer {
	return predicate.TryAnswer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldQuestionID), v))
	})
}

// Asnwer applies equality check predicate on the "asnwer" field. It's identical to AsnwerEQ.
func Asnwer(v string) predicate.TryAnswer {
	return predicate.TryAnswer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAsnwer), v))
	})
}

// Correct applies equality check predicate on the "correct" field. It's identical to CorrectEQ.
func Correct(v bool) predicate.TryAnswer {
	return predicate.TryAnswer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCorrect), v))
	})
}

// TryIDEQ applies the EQ predicate on the "try_id" field.
func TryIDEQ(v int) predicate.TryAnswer {
	return predicate.TryAnswer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTryID), v))
	})
}

// TryIDNEQ applies the NEQ predicate on the "try_id" field.
func TryIDNEQ(v int) predicate.TryAnswer {
	return predicate.TryAnswer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTryID), v))
	})
}

// TryIDIn applies the In predicate on the "try_id" field.
func TryIDIn(vs ...int) predicate.TryAnswer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TryAnswer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTryID), v...))
	})
}

// TryIDNotIn applies the NotIn predicate on the "try_id" field.
func TryIDNotIn(vs ...int) predicate.TryAnswer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TryAnswer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTryID), v...))
	})
}

// QuestionIDEQ applies the EQ predicate on the "question_id" field.
func QuestionIDEQ(v int) predicate.TryAnswer {
	return predicate.TryAnswer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldQuestionID), v))
	})
}

// QuestionIDNEQ applies the NEQ predicate on the "question_id" field.
func QuestionIDNEQ(v int) predicate.TryAnswer {
	return predicate.TryAnswer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldQuestionID), v))
	})
}

// QuestionIDIn applies the In predicate on the "question_id" field.
func QuestionIDIn(vs ...int) predicate.TryAnswer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TryAnswer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldQuestionID), v...))
	})
}

// QuestionIDNotIn applies the NotIn predicate on the "question_id" field.
func QuestionIDNotIn(vs ...int) predicate.TryAnswer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TryAnswer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldQuestionID), v...))
	})
}

// AsnwerEQ applies the EQ predicate on the "asnwer" field.
func AsnwerEQ(v string) predicate.TryAnswer {
	return predicate.TryAnswer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAsnwer), v))
	})
}

// AsnwerNEQ applies the NEQ predicate on the "asnwer" field.
func AsnwerNEQ(v string) predicate.TryAnswer {
	return predicate.TryAnswer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAsnwer), v))
	})
}

// AsnwerIn applies the In predicate on the "asnwer" field.
func AsnwerIn(vs ...string) predicate.TryAnswer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TryAnswer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAsnwer), v...))
	})
}

// AsnwerNotIn applies the NotIn predicate on the "asnwer" field.
func AsnwerNotIn(vs ...string) predicate.TryAnswer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TryAnswer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAsnwer), v...))
	})
}

// AsnwerGT applies the GT predicate on the "asnwer" field.
func AsnwerGT(v string) predicate.TryAnswer {
	return predicate.TryAnswer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAsnwer), v))
	})
}

// AsnwerGTE applies the GTE predicate on the "asnwer" field.
func AsnwerGTE(v string) predicate.TryAnswer {
	return predicate.TryAnswer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAsnwer), v))
	})
}

// AsnwerLT applies the LT predicate on the "asnwer" field.
func AsnwerLT(v string) predicate.TryAnswer {
	return predicate.TryAnswer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAsnwer), v))
	})
}

// AsnwerLTE applies the LTE predicate on the "asnwer" field.
func AsnwerLTE(v string) predicate.TryAnswer {
	return predicate.TryAnswer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAsnwer), v))
	})
}

// AsnwerContains applies the Contains predicate on the "asnwer" field.
func AsnwerContains(v string) predicate.TryAnswer {
	return predicate.TryAnswer(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAsnwer), v))
	})
}

// AsnwerHasPrefix applies the HasPrefix predicate on the "asnwer" field.
func AsnwerHasPrefix(v string) predicate.TryAnswer {
	return predicate.TryAnswer(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAsnwer), v))
	})
}

// AsnwerHasSuffix applies the HasSuffix predicate on the "asnwer" field.
func AsnwerHasSuffix(v string) predicate.TryAnswer {
	return predicate.TryAnswer(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAsnwer), v))
	})
}

// AsnwerEqualFold applies the EqualFold predicate on the "asnwer" field.
func AsnwerEqualFold(v string) predicate.TryAnswer {
	return predicate.TryAnswer(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAsnwer), v))
	})
}

// AsnwerContainsFold applies the ContainsFold predicate on the "asnwer" field.
func AsnwerContainsFold(v string) predicate.TryAnswer {
	return predicate.TryAnswer(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAsnwer), v))
	})
}

// CorrectEQ applies the EQ predicate on the "correct" field.
func CorrectEQ(v bool) predicate.TryAnswer {
	return predicate.TryAnswer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCorrect), v))
	})
}

// CorrectNEQ applies the NEQ predicate on the "correct" field.
func CorrectNEQ(v bool) predicate.TryAnswer {
	return predicate.TryAnswer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCorrect), v))
	})
}

// CorrectIsNil applies the IsNil predicate on the "correct" field.
func CorrectIsNil() predicate.TryAnswer {
	return predicate.TryAnswer(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCorrect)))
	})
}

// CorrectNotNil applies the NotNil predicate on the "correct" field.
func CorrectNotNil() predicate.TryAnswer {
	return predicate.TryAnswer(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCorrect)))
	})
}

// HasTheoreticalTry applies the HasEdge predicate on the "TheoreticalTry" edge.
func HasTheoreticalTry() predicate.TryAnswer {
	return predicate.TryAnswer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TheoreticalTryTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TheoreticalTryTable, TheoreticalTryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTheoreticalTryWith applies the HasEdge predicate on the "TheoreticalTry" edge with a given conditions (other predicates).
func HasTheoreticalTryWith(preds ...predicate.TheoreticalTry) predicate.TryAnswer {
	return predicate.TryAnswer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TheoreticalTryInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TheoreticalTryTable, TheoreticalTryColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasQuestion applies the HasEdge predicate on the "Question" edge.
func HasQuestion() predicate.TryAnswer {
	return predicate.TryAnswer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(QuestionTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, QuestionTable, QuestionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasQuestionWith applies the HasEdge predicate on the "Question" edge with a given conditions (other predicates).
func HasQuestionWith(preds ...predicate.Question) predicate.TryAnswer {
	return predicate.TryAnswer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(QuestionInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, QuestionTable, QuestionColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TryAnswer) predicate.TryAnswer {
	return predicate.TryAnswer(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TryAnswer) predicate.TryAnswer {
	return predicate.TryAnswer(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TryAnswer) predicate.TryAnswer {
	return predicate.TryAnswer(func(s *sql.Selector) {
		p(s.Not())
	})
}
