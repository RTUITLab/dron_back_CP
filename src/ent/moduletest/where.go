// Code generated by entc, DO NOT EDIT.

package moduletest

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/0B1t322/CP-Rosseti-Back/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ModuleTest {
	return predicate.ModuleTest(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ModuleTest {
	return predicate.ModuleTest(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ModuleTest {
	return predicate.ModuleTest(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ModuleTest {
	return predicate.ModuleTest(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.ModuleTest {
	return predicate.ModuleTest(func(s *sql.Selector) {
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
func IDGT(id int) predicate.ModuleTest {
	return predicate.ModuleTest(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ModuleTest {
	return predicate.ModuleTest(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ModuleTest {
	return predicate.ModuleTest(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ModuleTest {
	return predicate.ModuleTest(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// ModuleID applies equality check predicate on the "module_id" field. It's identical to ModuleIDEQ.
func ModuleID(v int) predicate.ModuleTest {
	return predicate.ModuleTest(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldModuleID), v))
	})
}

// TestID applies equality check predicate on the "test_id" field. It's identical to TestIDEQ.
func TestID(v int) predicate.ModuleTest {
	return predicate.ModuleTest(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTestID), v))
	})
}

// ModuleIDEQ applies the EQ predicate on the "module_id" field.
func ModuleIDEQ(v int) predicate.ModuleTest {
	return predicate.ModuleTest(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldModuleID), v))
	})
}

// ModuleIDNEQ applies the NEQ predicate on the "module_id" field.
func ModuleIDNEQ(v int) predicate.ModuleTest {
	return predicate.ModuleTest(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldModuleID), v))
	})
}

// ModuleIDIn applies the In predicate on the "module_id" field.
func ModuleIDIn(vs ...int) predicate.ModuleTest {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ModuleTest(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldModuleID), v...))
	})
}

// ModuleIDNotIn applies the NotIn predicate on the "module_id" field.
func ModuleIDNotIn(vs ...int) predicate.ModuleTest {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ModuleTest(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldModuleID), v...))
	})
}

// TestIDEQ applies the EQ predicate on the "test_id" field.
func TestIDEQ(v int) predicate.ModuleTest {
	return predicate.ModuleTest(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTestID), v))
	})
}

// TestIDNEQ applies the NEQ predicate on the "test_id" field.
func TestIDNEQ(v int) predicate.ModuleTest {
	return predicate.ModuleTest(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTestID), v))
	})
}

// TestIDIn applies the In predicate on the "test_id" field.
func TestIDIn(vs ...int) predicate.ModuleTest {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ModuleTest(func(s *sql.Selector) {
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
func TestIDNotIn(vs ...int) predicate.ModuleTest {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ModuleTest(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTestID), v...))
	})
}

// HasModule applies the HasEdge predicate on the "Module" edge.
func HasModule() predicate.ModuleTest {
	return predicate.ModuleTest(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ModuleTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ModuleTable, ModuleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasModuleWith applies the HasEdge predicate on the "Module" edge with a given conditions (other predicates).
func HasModuleWith(preds ...predicate.Module) predicate.ModuleTest {
	return predicate.ModuleTest(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ModuleInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ModuleTable, ModuleColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTest applies the HasEdge predicate on the "Test" edge.
func HasTest() predicate.ModuleTest {
	return predicate.ModuleTest(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TestTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TestTable, TestColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTestWith applies the HasEdge predicate on the "Test" edge with a given conditions (other predicates).
func HasTestWith(preds ...predicate.Test) predicate.ModuleTest {
	return predicate.ModuleTest(func(s *sql.Selector) {
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

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ModuleTest) predicate.ModuleTest {
	return predicate.ModuleTest(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ModuleTest) predicate.ModuleTest {
	return predicate.ModuleTest(func(s *sql.Selector) {
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
func Not(p predicate.ModuleTest) predicate.ModuleTest {
	return predicate.ModuleTest(func(s *sql.Selector) {
		p(s.Not())
	})
}