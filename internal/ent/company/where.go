// Code generated by ent, DO NOT EDIT.

package company

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/logn-soft/logn-back/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Company {
	return predicate.Company(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Company {
	return predicate.Company(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Company {
	return predicate.Company(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Company {
	return predicate.Company(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Company {
	return predicate.Company(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Company {
	return predicate.Company(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Company {
	return predicate.Company(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Company {
	return predicate.Company(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Company {
	return predicate.Company(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Company {
	return predicate.Company(sql.FieldEQ(FieldName, v))
}

// Employ applies equality check predicate on the "employ" field. It's identical to EmployEQ.
func Employ(v int) predicate.Company {
	return predicate.Company(sql.FieldEQ(FieldEmploy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Company {
	return predicate.Company(sql.FieldEQ(FieldCreatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Company {
	return predicate.Company(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Company {
	return predicate.Company(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Company {
	return predicate.Company(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Company {
	return predicate.Company(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Company {
	return predicate.Company(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Company {
	return predicate.Company(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Company {
	return predicate.Company(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Company {
	return predicate.Company(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Company {
	return predicate.Company(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Company {
	return predicate.Company(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Company {
	return predicate.Company(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Company {
	return predicate.Company(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Company {
	return predicate.Company(sql.FieldContainsFold(FieldName, v))
}

// EmployEQ applies the EQ predicate on the "employ" field.
func EmployEQ(v int) predicate.Company {
	return predicate.Company(sql.FieldEQ(FieldEmploy, v))
}

// EmployNEQ applies the NEQ predicate on the "employ" field.
func EmployNEQ(v int) predicate.Company {
	return predicate.Company(sql.FieldNEQ(FieldEmploy, v))
}

// EmployIn applies the In predicate on the "employ" field.
func EmployIn(vs ...int) predicate.Company {
	return predicate.Company(sql.FieldIn(FieldEmploy, vs...))
}

// EmployNotIn applies the NotIn predicate on the "employ" field.
func EmployNotIn(vs ...int) predicate.Company {
	return predicate.Company(sql.FieldNotIn(FieldEmploy, vs...))
}

// EmployGT applies the GT predicate on the "employ" field.
func EmployGT(v int) predicate.Company {
	return predicate.Company(sql.FieldGT(FieldEmploy, v))
}

// EmployGTE applies the GTE predicate on the "employ" field.
func EmployGTE(v int) predicate.Company {
	return predicate.Company(sql.FieldGTE(FieldEmploy, v))
}

// EmployLT applies the LT predicate on the "employ" field.
func EmployLT(v int) predicate.Company {
	return predicate.Company(sql.FieldLT(FieldEmploy, v))
}

// EmployLTE applies the LTE predicate on the "employ" field.
func EmployLTE(v int) predicate.Company {
	return predicate.Company(sql.FieldLTE(FieldEmploy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Company {
	return predicate.Company(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Company {
	return predicate.Company(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Company {
	return predicate.Company(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Company {
	return predicate.Company(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Company {
	return predicate.Company(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Company {
	return predicate.Company(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Company {
	return predicate.Company(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Company {
	return predicate.Company(sql.FieldLTE(FieldCreatedAt, v))
}

// HasSocials applies the HasEdge predicate on the "socials" edge.
func HasSocials() predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, SocialsTable, SocialsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSocialsWith applies the HasEdge predicate on the "socials" edge with a given conditions (other predicates).
func HasSocialsWith(preds ...predicate.Social) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SocialsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, SocialsTable, SocialsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRoles applies the HasEdge predicate on the "roles" edge.
func HasRoles() predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, RolesTable, RolesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRolesWith applies the HasEdge predicate on the "roles" edge with a given conditions (other predicates).
func HasRolesWith(preds ...predicate.Role) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RolesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, RolesTable, RolesPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasVacancies applies the HasEdge predicate on the "vacancies" edge.
func HasVacancies() predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, VacanciesTable, VacanciesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVacanciesWith applies the HasEdge predicate on the "vacancies" edge with a given conditions (other predicates).
func HasVacanciesWith(preds ...predicate.Vacancy) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(VacanciesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, VacanciesTable, VacanciesPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAreas applies the HasEdge predicate on the "areas" edge.
func HasAreas() predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, AreasTable, AreasPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAreasWith applies the HasEdge predicate on the "areas" edge with a given conditions (other predicates).
func HasAreasWith(preds ...predicate.Area) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AreasInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, AreasTable, AreasPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, UsersTable, UsersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UsersInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, UsersTable, UsersPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCommunities applies the HasEdge predicate on the "communities" edge.
func HasCommunities() predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, CommunitiesTable, CommunitiesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCommunitiesWith applies the HasEdge predicate on the "communities" edge with a given conditions (other predicates).
func HasCommunitiesWith(preds ...predicate.Community) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CommunitiesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, CommunitiesTable, CommunitiesPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Company) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Company) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
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
func Not(p predicate.Company) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		p(s.Not())
	})
}
