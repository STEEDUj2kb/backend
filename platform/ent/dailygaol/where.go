// Code generated by entc, DO NOT EDIT.

package dailygaol

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/STEEDUj2kb/platform/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
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
func IDGT(id int) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// Todo applies equality check predicate on the "todo" field. It's identical to TodoEQ.
func Todo(v string) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTodo), v))
	})
}

// Done applies equality check predicate on the "done" field. It's identical to DoneEQ.
func Done(v bool) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDone), v))
	})
}

// IsRemoved applies equality check predicate on the "is_removed" field. It's identical to IsRemovedEQ.
func IsRemoved(v bool) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsRemoved), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.DailyGaol {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DailyGaol(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.DailyGaol {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DailyGaol(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.DailyGaol {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DailyGaol(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.DailyGaol {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DailyGaol(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// TodoEQ applies the EQ predicate on the "todo" field.
func TodoEQ(v string) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTodo), v))
	})
}

// TodoNEQ applies the NEQ predicate on the "todo" field.
func TodoNEQ(v string) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTodo), v))
	})
}

// TodoIn applies the In predicate on the "todo" field.
func TodoIn(vs ...string) predicate.DailyGaol {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DailyGaol(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTodo), v...))
	})
}

// TodoNotIn applies the NotIn predicate on the "todo" field.
func TodoNotIn(vs ...string) predicate.DailyGaol {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DailyGaol(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTodo), v...))
	})
}

// TodoGT applies the GT predicate on the "todo" field.
func TodoGT(v string) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTodo), v))
	})
}

// TodoGTE applies the GTE predicate on the "todo" field.
func TodoGTE(v string) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTodo), v))
	})
}

// TodoLT applies the LT predicate on the "todo" field.
func TodoLT(v string) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTodo), v))
	})
}

// TodoLTE applies the LTE predicate on the "todo" field.
func TodoLTE(v string) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTodo), v))
	})
}

// TodoContains applies the Contains predicate on the "todo" field.
func TodoContains(v string) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTodo), v))
	})
}

// TodoHasPrefix applies the HasPrefix predicate on the "todo" field.
func TodoHasPrefix(v string) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTodo), v))
	})
}

// TodoHasSuffix applies the HasSuffix predicate on the "todo" field.
func TodoHasSuffix(v string) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTodo), v))
	})
}

// TodoEqualFold applies the EqualFold predicate on the "todo" field.
func TodoEqualFold(v string) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTodo), v))
	})
}

// TodoContainsFold applies the ContainsFold predicate on the "todo" field.
func TodoContainsFold(v string) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTodo), v))
	})
}

// DoneEQ applies the EQ predicate on the "done" field.
func DoneEQ(v bool) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDone), v))
	})
}

// DoneNEQ applies the NEQ predicate on the "done" field.
func DoneNEQ(v bool) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDone), v))
	})
}

// IsRemovedEQ applies the EQ predicate on the "is_removed" field.
func IsRemovedEQ(v bool) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsRemoved), v))
	})
}

// IsRemovedNEQ applies the NEQ predicate on the "is_removed" field.
func IsRemovedNEQ(v bool) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsRemoved), v))
	})
}

// HasStudy applies the HasEdge predicate on the "study" edge.
func HasStudy() predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StudyTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, StudyTable, StudyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStudyWith applies the HasEdge predicate on the "study" edge with a given conditions (other predicates).
func HasStudyWith(preds ...predicate.Study) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StudyInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, StudyTable, StudyColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasWgoal applies the HasEdge predicate on the "wgoal" edge.
func HasWgoal() predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WgoalTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WgoalTable, WgoalColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWgoalWith applies the HasEdge predicate on the "wgoal" edge with a given conditions (other predicates).
func HasWgoalWith(preds ...predicate.WeeklyGaol) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WgoalInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WgoalTable, WgoalColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DailyGaol) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DailyGaol) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
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
func Not(p predicate.DailyGaol) predicate.DailyGaol {
	return predicate.DailyGaol(func(s *sql.Selector) {
		p(s.Not())
	})
}
