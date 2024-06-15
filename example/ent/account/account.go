// Code generated by ent, DO NOT EDIT.

package account

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/lesomnus/entpb/example"
)

const (
	// Label holds the string label denoting the account type in the database.
	Label = "account"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDateCreated holds the string denoting the date_created field in the database.
	FieldDateCreated = "date_created"
	// FieldAlias holds the string denoting the alias field in the database.
	FieldAlias = "alias"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeMemberships holds the string denoting the memberships edge name in mutations.
	EdgeMemberships = "memberships"
	// Table holds the table name of the account in the database.
	Table = "accounts"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "accounts"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_accounts"
	// MembershipsTable is the table that holds the memberships relation/edge.
	MembershipsTable = "memberships"
	// MembershipsInverseTable is the table name for the Membership entity.
	// It exists in this package in order to avoid circular dependency with the "membership" package.
	MembershipsInverseTable = "memberships"
	// MembershipsColumn is the table column denoting the memberships relation/edge.
	MembershipsColumn = "account_memberships"
)

// Columns holds all SQL columns for account fields.
var Columns = []string{
	FieldID,
	FieldDateCreated,
	FieldAlias,
	FieldRole,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "accounts"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_accounts",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultDateCreated holds the default value on creation for the "date_created" field.
	DefaultDateCreated func() time.Time
	// DefaultAlias holds the default value on creation for the "alias" field.
	DefaultAlias func() string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// RoleValidator is a validator for the "role" field enum values. It is called by the builders before save.
func RoleValidator(r example.Role) error {
	switch r {
	case "OWNER", "MEMBER":
		return nil
	default:
		return fmt.Errorf("account: invalid enum value for role field: %q", r)
	}
}

// OrderOption defines the ordering options for the Account queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDateCreated orders the results by the date_created field.
func ByDateCreated(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDateCreated, opts...).ToFunc()
}

// ByAlias orders the results by the alias field.
func ByAlias(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAlias, opts...).ToFunc()
}

// ByRole orders the results by the role field.
func ByRole(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRole, opts...).ToFunc()
}

// ByOwnerField orders the results by owner field.
func ByOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerStep(), sql.OrderByField(field, opts...))
	}
}

// ByMembershipsCount orders the results by memberships count.
func ByMembershipsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMembershipsStep(), opts...)
	}
}

// ByMemberships orders the results by memberships terms.
func ByMemberships(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMembershipsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
	)
}
func newMembershipsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MembershipsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MembershipsTable, MembershipsColumn),
	)
}
