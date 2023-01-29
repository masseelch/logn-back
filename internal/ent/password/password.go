// Code generated by ent, DO NOT EDIT.

package password

import (
	"time"
)

const (
	// Label holds the string label denoting the password type in the database.
	Label = "password"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldHash holds the string denoting the hash field in the database.
	FieldHash = "hash"
	// FieldSalt holds the string denoting the salt field in the database.
	FieldSalt = "salt"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// Table holds the table name of the password in the database.
	Table = "passwords"
	// UsersTable is the table that holds the users relation/edge. The primary key declared below.
	UsersTable = "password_users"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
)

// Columns holds all SQL columns for password fields.
var Columns = []string{
	FieldID,
	FieldHash,
	FieldSalt,
	FieldCreatedAt,
}

var (
	// UsersPrimaryKey and UsersColumn2 are the table columns denoting the
	// primary key for the users relation (M2M).
	UsersPrimaryKey = []string{"password_id", "user_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)