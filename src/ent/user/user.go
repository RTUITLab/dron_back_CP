// Code generated by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldLogin holds the string denoting the login field in the database.
	FieldLogin = "login"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// EdgeRole holds the string denoting the role edge name in mutations.
	EdgeRole = "Role"
	// Table holds the table name of the user in the database.
	Table = "User"
	// RoleTable is the table that holds the Role relation/edge.
	RoleTable = "User"
	// RoleInverseTable is the table name for the Role entity.
	// It exists in this package in order to avoid circular dependency with the "role" package.
	RoleInverseTable = "Role"
	// RoleColumn is the table column denoting the Role relation/edge.
	RoleColumn = "role_user"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldLogin,
	FieldPassword,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "User"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"role_user",
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
	// LoginValidator is a validator for the "login" field. It is called by the builders before save.
	LoginValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
)