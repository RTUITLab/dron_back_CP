// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/0B1t322/CP-Rosseti-Back/ent/submodule"
	"github.com/0B1t322/CP-Rosseti-Back/ent/submoduletest"
	"github.com/0B1t322/CP-Rosseti-Back/ent/test"
)

// SubModuleTest is the model entity for the SubModuleTest schema.
type SubModuleTest struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// SubmoduleID holds the value of the "submodule_id" field.
	SubmoduleID int `json:"submodule_id,omitempty"`
	// TestID holds the value of the "test_id" field.
	TestID int `json:"test_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SubModuleTestQuery when eager-loading is set.
	Edges SubModuleTestEdges `json:"edges"`
}

// SubModuleTestEdges holds the relations/edges for other nodes in the graph.
type SubModuleTestEdges struct {
	// SubModule holds the value of the SubModule edge.
	SubModule *SubModule `json:"SubModule,omitempty"`
	// Test holds the value of the Test edge.
	Test *Test `json:"Test,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// SubModuleOrErr returns the SubModule value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SubModuleTestEdges) SubModuleOrErr() (*SubModule, error) {
	if e.loadedTypes[0] {
		if e.SubModule == nil {
			// The edge SubModule was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: submodule.Label}
		}
		return e.SubModule, nil
	}
	return nil, &NotLoadedError{edge: "SubModule"}
}

// TestOrErr returns the Test value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SubModuleTestEdges) TestOrErr() (*Test, error) {
	if e.loadedTypes[1] {
		if e.Test == nil {
			// The edge Test was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: test.Label}
		}
		return e.Test, nil
	}
	return nil, &NotLoadedError{edge: "Test"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SubModuleTest) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case submoduletest.FieldID, submoduletest.FieldSubmoduleID, submoduletest.FieldTestID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SubModuleTest", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SubModuleTest fields.
func (smt *SubModuleTest) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case submoduletest.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			smt.ID = int(value.Int64)
		case submoduletest.FieldSubmoduleID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field submodule_id", values[i])
			} else if value.Valid {
				smt.SubmoduleID = int(value.Int64)
			}
		case submoduletest.FieldTestID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field test_id", values[i])
			} else if value.Valid {
				smt.TestID = int(value.Int64)
			}
		}
	}
	return nil
}

// QuerySubModule queries the "SubModule" edge of the SubModuleTest entity.
func (smt *SubModuleTest) QuerySubModule() *SubModuleQuery {
	return (&SubModuleTestClient{config: smt.config}).QuerySubModule(smt)
}

// QueryTest queries the "Test" edge of the SubModuleTest entity.
func (smt *SubModuleTest) QueryTest() *TestQuery {
	return (&SubModuleTestClient{config: smt.config}).QueryTest(smt)
}

// Update returns a builder for updating this SubModuleTest.
// Note that you need to call SubModuleTest.Unwrap() before calling this method if this SubModuleTest
// was returned from a transaction, and the transaction was committed or rolled back.
func (smt *SubModuleTest) Update() *SubModuleTestUpdateOne {
	return (&SubModuleTestClient{config: smt.config}).UpdateOne(smt)
}

// Unwrap unwraps the SubModuleTest entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (smt *SubModuleTest) Unwrap() *SubModuleTest {
	tx, ok := smt.config.driver.(*txDriver)
	if !ok {
		panic("ent: SubModuleTest is not a transactional entity")
	}
	smt.config.driver = tx.drv
	return smt
}

// String implements the fmt.Stringer.
func (smt *SubModuleTest) String() string {
	var builder strings.Builder
	builder.WriteString("SubModuleTest(")
	builder.WriteString(fmt.Sprintf("id=%v", smt.ID))
	builder.WriteString(", submodule_id=")
	builder.WriteString(fmt.Sprintf("%v", smt.SubmoduleID))
	builder.WriteString(", test_id=")
	builder.WriteString(fmt.Sprintf("%v", smt.TestID))
	builder.WriteByte(')')
	return builder.String()
}

// SubModuleTests is a parsable slice of SubModuleTest.
type SubModuleTests []*SubModuleTest

func (smt SubModuleTests) config(cfg config) {
	for _i := range smt {
		smt[_i].config = cfg
	}
}
