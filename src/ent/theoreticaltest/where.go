// Code generated by entc, DO NOT EDIT.

package theoreticaltest

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/0B1t322/CP-Rosseti-Back/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.TheoreticalTest {
	return predicate.TheoreticalTest(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.TheoreticalTest {
	return predicate.TheoreticalTest(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.TheoreticalTest {
	return predicate.TheoreticalTest(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.TheoreticalTest {
	return predicate.TheoreticalTest(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.TheoreticalTest {
	return predicate.TheoreticalTest(func(s *sql.Selector) {
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
func IDGT(id int) predicate.TheoreticalTest {
	return predicate.TheoreticalTest(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.TheoreticalTest {
	return predicate.TheoreticalTest(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.TheoreticalTest {
	return predicate.TheoreticalTest(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.TheoreticalTest {
	return predicate.TheoreticalTest(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// TestID applies equality check predicate on the "test_id" field. It's identical to TestIDEQ.
func TestID(v int) predicate.TheoreticalTest {
	return predicate.TheoreticalTest(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTestID), v))
	})
}

// TestIDEQ applies the EQ predicate on the "test_id" field.
func TestIDEQ(v int) predicate.TheoreticalTest {
	return predicate.TheoreticalTest(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTestID), v))
	})
}

// TestIDNEQ applies the NEQ predicate on the "test_id" field.
func TestIDNEQ(v int) predicate.TheoreticalTest {
	return predicate.TheoreticalTest(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTestID), v))
	})
}

// TestIDIn applies the In predicate on the "test_id" field.
func TestIDIn(vs ...int) predicate.TheoreticalTest {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TheoreticalTest(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTestID), v...))
	})
}

// TestIDNotIn applies the NotIn predicate on the "test_id" field.
func TestIDNotIn(vs ...int) predicate.TheoreticalTest {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TheoreticalTest(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTestID), v...))
	})
}

// HasTest applies the HasEdge predicate on the "Test" edge.
func HasTest() predicate.TheoreticalTest {
	return predicate.TheoreticalTest(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TestTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TestTable, TestColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTestWith applies the HasEdge predicate on the "Test" edge with a given conditions (other predicates).
func HasTestWith(preds ...predicate.Test) predicate.TheoreticalTest {
	return predicate.TheoreticalTest(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TestInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TestTable, TestColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasQuestion applies the HasEdge predicate on the "Question" edge.
func HasQuestion() predicate.TheoreticalTest {
	return predicate.TheoreticalTest(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(QuestionTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, QuestionTable, QuestionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasQuestionWith applies the HasEdge predicate on the "Question" edge with a given conditions (other predicates).
func HasQuestionWith(preds ...predicate.Question) predicate.TheoreticalTest {
	return predicate.TheoreticalTest(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(QuestionInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, QuestionTable, QuestionColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TheoreticalTest) predicate.TheoreticalTest {
	return predicate.TheoreticalTest(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TheoreticalTest) predicate.TheoreticalTest {
	return predicate.TheoreticalTest(func(s *sql.Selector) {
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
func Not(p predicate.TheoreticalTest) predicate.TheoreticalTest {
	return predicate.TheoreticalTest(func(s *sql.Selector) {
		p(s.Not())
	})
}
