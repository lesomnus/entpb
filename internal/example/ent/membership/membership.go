// Code generated by ent, DO NOT EDIT.

package membership

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the membership type in the database.
	Label = "membership"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDateCreated holds the string denoting the date_created field in the database.
	FieldDateCreated = "date_created"
	// FieldAccountID holds the string denoting the account_id field in the database.
	FieldAccountID = "account_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeAccount holds the string denoting the account edge name in mutations.
	EdgeAccount = "account"
	// Table holds the table name of the membership in the database.
	Table = "memberships"
	// AccountTable is the table that holds the account relation/edge.
	AccountTable = "memberships"
	// AccountInverseTable is the table name for the Account entity.
	// It exists in this package in order to avoid circular dependency with the "account" package.
	AccountInverseTable = "accounts"
	// AccountColumn is the table column denoting the account relation/edge.
	AccountColumn = "account_id"
)

// Columns holds all SQL columns for membership fields.
var Columns = []string{
	FieldID,
	FieldDateCreated,
	FieldAccountID,
	FieldName,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "memberships"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_memberships",
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
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName func() string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Membership queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDateCreated orders the results by the date_created field.
func ByDateCreated(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDateCreated, opts...).ToFunc()
}

// ByAccountID orders the results by the account_id field.
func ByAccountID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAccountID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByAccountField orders the results by account field.
func ByAccountField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAccountStep(), sql.OrderByField(field, opts...))
	}
}
func newAccountStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AccountInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, AccountTable, AccountColumn),
	)
}
