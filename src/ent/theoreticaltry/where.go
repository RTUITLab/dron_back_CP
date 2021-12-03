// Code generated by entc, DO NOT EDIT.

package theoreticaltry

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/0B1t322/CP-Rosseti-Back/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.TheoreticalTry {
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.TheoreticalTry {
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.TheoreticalTry {
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.TheoreticalTry {
	return predicate.TheoreticalTry(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.TheoreticalTry {
	return predicate.TheoreticalTry(func(s *sql.Selector) {
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
func IDGT(id int) predicate.TheoreticalTry {
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.TheoreticalTry {
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.TheoreticalTry {
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.TheoreticalTry {
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Start applies equality check predicate on the "start" field. It's identical to StartEQ.
func Start(v time.Time) predicate.TheoreticalTry {
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStart), v))
	})
}

// End applies equality check predicate on the "end" field. It's identical to EndEQ.
func End(v time.Time) predicate.TheoreticalTry {
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnd), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.TheoreticalTry {
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// TheoreticalTestID applies equality check predicate on the "theoretical_test_id" field. It's identical to TheoreticalTestIDEQ.
func TheoreticalTestID(v int) predicate.TheoreticalTry {
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTheoreticalTestID), v))
	})
}

// StartEQ applies the EQ predicate on the "start" field.
func StartEQ(v time.Time) predicate.TheoreticalTry {
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStart), v))
	})
}

// StartNEQ applies the NEQ predicate on the "start" field.
func StartNEQ(v time.Time) predicate.TheoreticalTry {
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStart), v))
	})
}

// StartIn applies the In predicate on the "start" field.
func StartIn(vs ...time.Time) predicate.TheoreticalTry {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStart), v...))
	})
}

// StartNotIn applies the NotIn predicate on the "start" field.
func StartNotIn(vs ...time.Time) predicate.TheoreticalTry {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStart), v...))
	})
}

// StartGT applies the GT predicate on the "start" field.
func StartGT(v time.Time) predicate.TheoreticalTry {
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStart), v))
	})
}

// StartGTE applies the GTE predicate on the "start" field.
func StartGTE(v time.Time) predicate.TheoreticalTry {
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStart), v))
	})
}

// StartLT applies the LT predicate on the "start" field.
func StartLT(v time.Time) predicate.TheoreticalTry {
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStart), v))
	})
}

// StartLTE applies the LTE predicate on the "start" field.
func StartLTE(v time.Time) predicate.TheoreticalTry {
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStart), v))
	})
}

// EndEQ applies the EQ predicate on the "end" field.
func EndEQ(v time.Time) predicate.TheoreticalTry {
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnd), v))
	})
}

// EndNEQ applies the NEQ predicate on the "end" field.
func EndNEQ(v time.Time) predicate.TheoreticalTry {
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEnd), v))
	})
}

// EndIn applies the In predicate on the "end" field.
func EndIn(vs ...time.Time) predicate.TheoreticalTry {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEnd), v...))
	})
}

// EndNotIn applies the NotIn predicate on the "end" field.
func EndNotIn(vs ...time.Time) predicate.TheoreticalTry {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEnd), v...))
	})
}

// EndGT applies the GT predicate on the "end" field.
func EndGT(v time.Time) predicate.TheoreticalTry {
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEnd), v))
	})
}

// EndGTE applies the GTE predicate on the "end" field.
func EndGTE(v time.Time) predicate.TheoreticalTry {
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEnd), v))
	})
}

// EndLT applies the LT predicate on the "end" field.
func EndLT(v time.Time) predicate.TheoreticalTry {
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEnd), v))
	})
}

// EndLTE applies the LTE predicate on the "end" field.
func EndLTE(v time.Time) predicate.TheoreticalTry {
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEnd), v))
	})
}

// EndIsNil applies the IsNil predicate on the "end" field.
func EndIsNil() predicate.TheoreticalTry {
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldEnd)))
	})
}

// EndNotNil applies the NotNil predicate on the "end" field.
func EndNotNil() predicate.TheoreticalTry {
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldEnd)))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.TheoreticalTry {
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.TheoreticalTry {
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.TheoreticalTry {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.TheoreticalTry {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// TheoreticalTestIDEQ applies the EQ predicate on the "theoretical_test_id" field.
func TheoreticalTestIDEQ(v int) predicate.TheoreticalTry {
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTheoreticalTestID), v))
	})
}

// TheoreticalTestIDNEQ applies the NEQ predicate on the "theoretical_test_id" field.
func TheoreticalTestIDNEQ(v int) predicate.TheoreticalTry {
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTheoreticalTestID), v))
	})
}

// TheoreticalTestIDIn applies the In predicate on the "theoretical_test_id" field.
func TheoreticalTestIDIn(vs ...int) predicate.TheoreticalTry {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTheoreticalTestID), v...))
	})
}

// TheoreticalTestIDNotIn applies the NotIn predicate on the "theoretical_test_id" field.
func TheoreticalTestIDNotIn(vs ...int) predicate.TheoreticalTry {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTheoreticalTestID), v...))
	})
}

// HasTheoreticalTest applies the HasEdge predicate on the "TheoreticalTest" edge.
func HasTheoreticalTest() predicate.TheoreticalTry {
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TheoreticalTestTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TheoreticalTestTable, TheoreticalTestColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTheoreticalTestWith applies the HasEdge predicate on the "TheoreticalTest" edge with a given conditions (other predicates).
func HasTheoreticalTestWith(preds ...predicate.TheoreticalTest) predicate.TheoreticalTry {
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TheoreticalTestInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TheoreticalTestTable, TheoreticalTestColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "User" edge.
func HasUser() predicate.TheoreticalTry {
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "User" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.TheoreticalTry {
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTryAnswer applies the HasEdge predicate on the "TryAnswer" edge.
func HasTryAnswer() predicate.TheoreticalTry {
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TryAnswerTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TryAnswerTable, TryAnswerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTryAnswerWith applies the HasEdge predicate on the "TryAnswer" edge with a given conditions (other predicates).
func HasTryAnswerWith(preds ...predicate.TryAnswer) predicate.TheoreticalTry {
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TryAnswerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TryAnswerTable, TryAnswerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TheoreticalTry) predicate.TheoreticalTry {
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TheoreticalTry) predicate.TheoreticalTry {
	return predicate.TheoreticalTry(func(s *sql.Selector) {
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
func Not(p predicate.TheoreticalTry) predicate.TheoreticalTry {
	return predicate.TheoreticalTry(func(s *sql.Selector) {
		p(s.Not())
	})
}
