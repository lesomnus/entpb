// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/lesomnus/entpb/internal/example/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// DateCreated applies equality check predicate on the "date_created" field. It's identical to DateCreatedEQ.
func DateCreated(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDateCreated, v))
}

// DateCreatedEQ applies the EQ predicate on the "date_created" field.
func DateCreatedEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDateCreated, v))
}

// DateCreatedNEQ applies the NEQ predicate on the "date_created" field.
func DateCreatedNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldDateCreated, v))
}

// DateCreatedIn applies the In predicate on the "date_created" field.
func DateCreatedIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldDateCreated, vs...))
}

// DateCreatedNotIn applies the NotIn predicate on the "date_created" field.
func DateCreatedNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldDateCreated, vs...))
}

// DateCreatedGT applies the GT predicate on the "date_created" field.
func DateCreatedGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldDateCreated, v))
}

// DateCreatedGTE applies the GTE predicate on the "date_created" field.
func DateCreatedGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldDateCreated, v))
}

// DateCreatedLT applies the LT predicate on the "date_created" field.
func DateCreatedLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldDateCreated, v))
}

// DateCreatedLTE applies the LTE predicate on the "date_created" field.
func DateCreatedLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldDateCreated, v))
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newParentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChildren applies the HasEdge predicate on the "children" edge.
func HasChildren() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChildrenWith applies the HasEdge predicate on the "children" edge with a given conditions (other predicates).
func HasChildrenWith(preds ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newChildrenStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasIdentities applies the HasEdge predicate on the "identities" edge.
func HasIdentities() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, IdentitiesTable, IdentitiesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasIdentitiesWith applies the HasEdge predicate on the "identities" edge with a given conditions (other predicates).
func HasIdentitiesWith(preds ...predicate.Identity) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newIdentitiesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAccounts applies the HasEdge predicate on the "accounts" edge.
func HasAccounts() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AccountsTable, AccountsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAccountsWith applies the HasEdge predicate on the "accounts" edge with a given conditions (other predicates).
func HasAccountsWith(preds ...predicate.Account) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newAccountsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMemberships applies the HasEdge predicate on the "memberships" edge.
func HasMemberships() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MembershipsTable, MembershipsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMembershipsWith applies the HasEdge predicate on the "memberships" edge with a given conditions (other predicates).
func HasMembershipsWith(preds ...predicate.Membership) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newMembershipsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}
