// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/0B1t322/CP-Rosseti-Back/ent/theoreticaltest"
	"github.com/0B1t322/CP-Rosseti-Back/ent/theoreticaltry"
	"github.com/0B1t322/CP-Rosseti-Back/ent/user"
)

// TheoreticalTry is the model entity for the TheoreticalTry schema.
type TheoreticalTry struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Start holds the value of the "start" field.
	Start time.Time `json:"start,omitempty"`
	// End holds the value of the "end" field.
	End time.Time `json:"end,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// TheoreticalTestID holds the value of the "theoretical_test_id" field.
	TheoreticalTestID int `json:"theoretical_test_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TheoreticalTryQuery when eager-loading is set.
	Edges TheoreticalTryEdges `json:"edges"`
}

// TheoreticalTryEdges holds the relations/edges for other nodes in the graph.
type TheoreticalTryEdges struct {
	// TheoreticalTest holds the value of the TheoreticalTest edge.
	TheoreticalTest *TheoreticalTest `json:"TheoreticalTest,omitempty"`
	// User holds the value of the User edge.
	User *User `json:"User,omitempty"`
	// TryAnswer holds the value of the TryAnswer edge.
	TryAnswer []*TryAnswer `json:"TryAnswer,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// TheoreticalTestOrErr returns the TheoreticalTest value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TheoreticalTryEdges) TheoreticalTestOrErr() (*TheoreticalTest, error) {
	if e.loadedTypes[0] {
		if e.TheoreticalTest == nil {
			// The edge TheoreticalTest was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: theoreticaltest.Label}
		}
		return e.TheoreticalTest, nil
	}
	return nil, &NotLoadedError{edge: "TheoreticalTest"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TheoreticalTryEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.User == nil {
			// The edge User was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "User"}
}

// TryAnswerOrErr returns the TryAnswer value or an error if the edge
// was not loaded in eager-loading.
func (e TheoreticalTryEdges) TryAnswerOrErr() ([]*TryAnswer, error) {
	if e.loadedTypes[2] {
		return e.TryAnswer, nil
	}
	return nil, &NotLoadedError{edge: "TryAnswer"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TheoreticalTry) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case theoreticaltry.FieldID, theoreticaltry.FieldUserID, theoreticaltry.FieldTheoreticalTestID:
			values[i] = new(sql.NullInt64)
		case theoreticaltry.FieldStart, theoreticaltry.FieldEnd:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TheoreticalTry", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TheoreticalTry fields.
func (tt *TheoreticalTry) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case theoreticaltry.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tt.ID = int(value.Int64)
		case theoreticaltry.FieldStart:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start", values[i])
			} else if value.Valid {
				tt.Start = value.Time
			}
		case theoreticaltry.FieldEnd:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end", values[i])
			} else if value.Valid {
				tt.End = value.Time
			}
		case theoreticaltry.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				tt.UserID = int(value.Int64)
			}
		case theoreticaltry.FieldTheoreticalTestID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field theoretical_test_id", values[i])
			} else if value.Valid {
				tt.TheoreticalTestID = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryTheoreticalTest queries the "TheoreticalTest" edge of the TheoreticalTry entity.
func (tt *TheoreticalTry) QueryTheoreticalTest() *TheoreticalTestQuery {
	return (&TheoreticalTryClient{config: tt.config}).QueryTheoreticalTest(tt)
}

// QueryUser queries the "User" edge of the TheoreticalTry entity.
func (tt *TheoreticalTry) QueryUser() *UserQuery {
	return (&TheoreticalTryClient{config: tt.config}).QueryUser(tt)
}

// QueryTryAnswer queries the "TryAnswer" edge of the TheoreticalTry entity.
func (tt *TheoreticalTry) QueryTryAnswer() *TryAnswerQuery {
	return (&TheoreticalTryClient{config: tt.config}).QueryTryAnswer(tt)
}

// Update returns a builder for updating this TheoreticalTry.
// Note that you need to call TheoreticalTry.Unwrap() before calling this method if this TheoreticalTry
// was returned from a transaction, and the transaction was committed or rolled back.
func (tt *TheoreticalTry) Update() *TheoreticalTryUpdateOne {
	return (&TheoreticalTryClient{config: tt.config}).UpdateOne(tt)
}

// Unwrap unwraps the TheoreticalTry entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tt *TheoreticalTry) Unwrap() *TheoreticalTry {
	tx, ok := tt.config.driver.(*txDriver)
	if !ok {
		panic("ent: TheoreticalTry is not a transactional entity")
	}
	tt.config.driver = tx.drv
	return tt
}

// String implements the fmt.Stringer.
func (tt *TheoreticalTry) String() string {
	var builder strings.Builder
	builder.WriteString("TheoreticalTry(")
	builder.WriteString(fmt.Sprintf("id=%v", tt.ID))
	builder.WriteString(", start=")
	builder.WriteString(tt.Start.Format(time.ANSIC))
	builder.WriteString(", end=")
	builder.WriteString(tt.End.Format(time.ANSIC))
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", tt.UserID))
	builder.WriteString(", theoretical_test_id=")
	builder.WriteString(fmt.Sprintf("%v", tt.TheoreticalTestID))
	builder.WriteByte(')')
	return builder.String()
}

// TheoreticalTries is a parsable slice of TheoreticalTry.
type TheoreticalTries []*TheoreticalTry

func (tt TheoreticalTries) config(cfg config) {
	for _i := range tt {
		tt[_i].config = cfg
	}
}
