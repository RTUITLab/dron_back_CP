// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/0B1t322/CP-Rosseti-Back/ent/role"
	"github.com/0B1t322/CP-Rosseti-Back/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Login holds the value of the "login" field.
	Login string `json:"login,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges     UserEdges `json:"edges"`
	role_user *int
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Role holds the value of the Role edge.
	Role *Role `json:"Role,omitempty"`
	// TheoTry holds the value of the TheoTry edge.
	TheoTry []*TheoreticalTry `json:"TheoTry,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// RoleOrErr returns the Role value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) RoleOrErr() (*Role, error) {
	if e.loadedTypes[0] {
		if e.Role == nil {
			// The edge Role was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: role.Label}
		}
		return e.Role, nil
	}
	return nil, &NotLoadedError{edge: "Role"}
}

// TheoTryOrErr returns the TheoTry value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) TheoTryOrErr() ([]*TheoreticalTry, error) {
	if e.loadedTypes[1] {
		return e.TheoTry, nil
	}
	return nil, &NotLoadedError{edge: "TheoTry"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldLogin, user.FieldPassword:
			values[i] = new(sql.NullString)
		case user.ForeignKeys[0]: // role_user
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldLogin:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field login", values[i])
			} else if value.Valid {
				u.Login = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field role_user", value)
			} else if value.Valid {
				u.role_user = new(int)
				*u.role_user = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryRole queries the "Role" edge of the User entity.
func (u *User) QueryRole() *RoleQuery {
	return (&UserClient{config: u.config}).QueryRole(u)
}

// QueryTheoTry queries the "TheoTry" edge of the User entity.
func (u *User) QueryTheoTry() *TheoreticalTryQuery {
	return (&UserClient{config: u.config}).QueryTheoTry(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", login=")
	builder.WriteString(u.Login)
	builder.WriteString(", password=")
	builder.WriteString(u.Password)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
