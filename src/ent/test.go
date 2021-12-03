// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/0B1t322/CP-Rosseti-Back/ent/test"
)

// Test is the model entity for the Test schema.
type Test struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// TestType holds the value of the "TestType" field.
	TestType string `json:"TestType,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TestQuery when eager-loading is set.
	Edges TestEdges `json:"edges"`
}

// TestEdges holds the relations/edges for other nodes in the graph.
type TestEdges struct {
	// ModuleTest holds the value of the ModuleTest edge.
	ModuleTest []*ModuleTest `json:"ModuleTest,omitempty"`
	// SubmoduleTest holds the value of the SubmoduleTest edge.
	SubmoduleTest []*SubModuleTest `json:"SubmoduleTest,omitempty"`
	// TherTest holds the value of the TherTest edge.
	TherTest []*TheoreticalTest `json:"TherTest,omitempty"`
	// PractTest holds the value of the PractTest edge.
	PractTest []*PractTest `json:"PractTest,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// ModuleTestOrErr returns the ModuleTest value or an error if the edge
// was not loaded in eager-loading.
func (e TestEdges) ModuleTestOrErr() ([]*ModuleTest, error) {
	if e.loadedTypes[0] {
		return e.ModuleTest, nil
	}
	return nil, &NotLoadedError{edge: "ModuleTest"}
}

// SubmoduleTestOrErr returns the SubmoduleTest value or an error if the edge
// was not loaded in eager-loading.
func (e TestEdges) SubmoduleTestOrErr() ([]*SubModuleTest, error) {
	if e.loadedTypes[1] {
		return e.SubmoduleTest, nil
	}
	return nil, &NotLoadedError{edge: "SubmoduleTest"}
}

// TherTestOrErr returns the TherTest value or an error if the edge
// was not loaded in eager-loading.
func (e TestEdges) TherTestOrErr() ([]*TheoreticalTest, error) {
	if e.loadedTypes[2] {
		return e.TherTest, nil
	}
	return nil, &NotLoadedError{edge: "TherTest"}
}

// PractTestOrErr returns the PractTest value or an error if the edge
// was not loaded in eager-loading.
func (e TestEdges) PractTestOrErr() ([]*PractTest, error) {
	if e.loadedTypes[3] {
		return e.PractTest, nil
	}
	return nil, &NotLoadedError{edge: "PractTest"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Test) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case test.FieldID:
			values[i] = new(sql.NullInt64)
		case test.FieldTestType:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Test", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Test fields.
func (t *Test) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case test.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case test.FieldTestType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field TestType", values[i])
			} else if value.Valid {
				t.TestType = value.String
			}
		}
	}
	return nil
}

// QueryModuleTest queries the "ModuleTest" edge of the Test entity.
func (t *Test) QueryModuleTest() *ModuleTestQuery {
	return (&TestClient{config: t.config}).QueryModuleTest(t)
}

// QuerySubmoduleTest queries the "SubmoduleTest" edge of the Test entity.
func (t *Test) QuerySubmoduleTest() *SubModuleTestQuery {
	return (&TestClient{config: t.config}).QuerySubmoduleTest(t)
}

// QueryTherTest queries the "TherTest" edge of the Test entity.
func (t *Test) QueryTherTest() *TheoreticalTestQuery {
	return (&TestClient{config: t.config}).QueryTherTest(t)
}

// QueryPractTest queries the "PractTest" edge of the Test entity.
func (t *Test) QueryPractTest() *PractTestQuery {
	return (&TestClient{config: t.config}).QueryPractTest(t)
}

// Update returns a builder for updating this Test.
// Note that you need to call Test.Unwrap() before calling this method if this Test
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Test) Update() *TestUpdateOne {
	return (&TestClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Test entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Test) Unwrap() *Test {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Test is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Test) String() string {
	var builder strings.Builder
	builder.WriteString("Test(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", TestType=")
	builder.WriteString(t.TestType)
	builder.WriteByte(')')
	return builder.String()
}

// Tests is a parsable slice of Test.
type Tests []*Test

func (t Tests) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
