// Code generated by entc, DO NOT EDIT.

package task

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/0B1t322/CP-Rosseti-Back/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// PractTestID applies equality check predicate on the "pract_test_id" field. It's identical to PractTestIDEQ.
func PractTestID(v int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPractTestID), v))
	})
}

// Task applies equality check predicate on the "task" field. It's identical to TaskEQ.
func Task(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTask), v))
	})
}

// PractTestIDEQ applies the EQ predicate on the "pract_test_id" field.
func PractTestIDEQ(v int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPractTestID), v))
	})
}

// PractTestIDNEQ applies the NEQ predicate on the "pract_test_id" field.
func PractTestIDNEQ(v int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPractTestID), v))
	})
}

// PractTestIDIn applies the In predicate on the "pract_test_id" field.
func PractTestIDIn(vs ...int) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPractTestID), v...))
	})
}

// PractTestIDNotIn applies the NotIn predicate on the "pract_test_id" field.
func PractTestIDNotIn(vs ...int) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPractTestID), v...))
	})
}

// TaskEQ applies the EQ predicate on the "task" field.
func TaskEQ(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTask), v))
	})
}

// TaskNEQ applies the NEQ predicate on the "task" field.
func TaskNEQ(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTask), v))
	})
}

// TaskIn applies the In predicate on the "task" field.
func TaskIn(vs ...string) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTask), v...))
	})
}

// TaskNotIn applies the NotIn predicate on the "task" field.
func TaskNotIn(vs ...string) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTask), v...))
	})
}

// TaskGT applies the GT predicate on the "task" field.
func TaskGT(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTask), v))
	})
}

// TaskGTE applies the GTE predicate on the "task" field.
func TaskGTE(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTask), v))
	})
}

// TaskLT applies the LT predicate on the "task" field.
func TaskLT(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTask), v))
	})
}

// TaskLTE applies the LTE predicate on the "task" field.
func TaskLTE(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTask), v))
	})
}

// TaskContains applies the Contains predicate on the "task" field.
func TaskContains(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTask), v))
	})
}

// TaskHasPrefix applies the HasPrefix predicate on the "task" field.
func TaskHasPrefix(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTask), v))
	})
}

// TaskHasSuffix applies the HasSuffix predicate on the "task" field.
func TaskHasSuffix(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTask), v))
	})
}

// TaskEqualFold applies the EqualFold predicate on the "task" field.
func TaskEqualFold(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTask), v))
	})
}

// TaskContainsFold applies the ContainsFold predicate on the "task" field.
func TaskContainsFold(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTask), v))
	})
}

// HasTest applies the HasEdge predicate on the "Test" edge.
func HasTest() predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TestTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, TestTable, TestColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTestWith applies the HasEdge predicate on the "Test" edge with a given conditions (other predicates).
func HasTestWith(preds ...predicate.PractTest) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TestInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, TestTable, TestColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Task) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Task) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
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
func Not(p predicate.Task) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		p(s.Not())
	})
}
