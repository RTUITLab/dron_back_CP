// Code generated by entc, DO NOT EDIT.

package module

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/0B1t322/CP-Rosseti-Back/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Module {
	return predicate.Module(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Module {
	return predicate.Module(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Module {
	return predicate.Module(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Module {
	return predicate.Module(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Module {
	return predicate.Module(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Module {
	return predicate.Module(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Module {
	return predicate.Module(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Module {
	return predicate.Module(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Module {
	return predicate.Module(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Module {
	return predicate.Module(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Module {
	return predicate.Module(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Module {
	return predicate.Module(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Module {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Module(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Module {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Module(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Module {
	return predicate.Module(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Module {
	return predicate.Module(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Module {
	return predicate.Module(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Module {
	return predicate.Module(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Module {
	return predicate.Module(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Module {
	return predicate.Module(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Module {
	return predicate.Module(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Module {
	return predicate.Module(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Module {
	return predicate.Module(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// HasModuleDependcies applies the HasEdge predicate on the "ModuleDependcies" edge.
func HasModuleDependcies() predicate.Module {
	return predicate.Module(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ModuleDependciesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ModuleDependciesTable, ModuleDependciesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasModuleDependciesWith applies the HasEdge predicate on the "ModuleDependcies" edge with a given conditions (other predicates).
func HasModuleDependciesWith(preds ...predicate.ModuleDependcies) predicate.Module {
	return predicate.Module(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ModuleDependciesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ModuleDependciesTable, ModuleDependciesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasModuleDependOn applies the HasEdge predicate on the "ModuleDependOn" edge.
func HasModuleDependOn() predicate.Module {
	return predicate.Module(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ModuleDependOnTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ModuleDependOnTable, ModuleDependOnColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasModuleDependOnWith applies the HasEdge predicate on the "ModuleDependOn" edge with a given conditions (other predicates).
func HasModuleDependOnWith(preds ...predicate.ModuleDependcies) predicate.Module {
	return predicate.Module(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ModuleDependOnInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ModuleDependOnTable, ModuleDependOnColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSubModules applies the HasEdge predicate on the "SubModules" edge.
func HasSubModules() predicate.Module {
	return predicate.Module(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SubModulesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SubModulesTable, SubModulesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubModulesWith applies the HasEdge predicate on the "SubModules" edge with a given conditions (other predicates).
func HasSubModulesWith(preds ...predicate.SubModule) predicate.Module {
	return predicate.Module(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SubModulesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SubModulesTable, SubModulesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTest applies the HasEdge predicate on the "Test" edge.
func HasTest() predicate.Module {
	return predicate.Module(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TestTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TestTable, TestColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTestWith applies the HasEdge predicate on the "Test" edge with a given conditions (other predicates).
func HasTestWith(preds ...predicate.ModuleTest) predicate.Module {
	return predicate.Module(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TestInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TestTable, TestColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Module) predicate.Module {
	return predicate.Module(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Module) predicate.Module {
	return predicate.Module(func(s *sql.Selector) {
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
func Not(p predicate.Module) predicate.Module {
	return predicate.Module(func(s *sql.Selector) {
		p(s.Not())
	})
}
