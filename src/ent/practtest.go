// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/0B1t322/CP-Rosseti-Back/ent/practtest"
	"github.com/0B1t322/CP-Rosseti-Back/ent/schema"
	"github.com/0B1t322/CP-Rosseti-Back/ent/task"
	"github.com/0B1t322/CP-Rosseti-Back/ent/test"
)

// PractTest is the model entity for the PractTest schema.
type PractTest struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// TestID holds the value of the "test_id" field.
	TestID int `json:"test_id,omitempty"`
	// Config holds the value of the "config" field.
	Config schema.JSONObject `json:"config,omitempty"`
	// Duration holds the value of the "duration" field.
	Duration int `json:"duration,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PractTestQuery when eager-loading is set.
	Edges PractTestEdges `json:"edges"`
}

// PractTestEdges holds the relations/edges for other nodes in the graph.
type PractTestEdges struct {
	// Test holds the value of the Test edge.
	Test *Test `json:"Test,omitempty"`
	// Task holds the value of the Task edge.
	Task *Task `json:"Task,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TestOrErr returns the Test value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PractTestEdges) TestOrErr() (*Test, error) {
	if e.loadedTypes[0] {
		if e.Test == nil {
			// The edge Test was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: test.Label}
		}
		return e.Test, nil
	}
	return nil, &NotLoadedError{edge: "Test"}
}

// TaskOrErr returns the Task value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PractTestEdges) TaskOrErr() (*Task, error) {
	if e.loadedTypes[1] {
		if e.Task == nil {
			// The edge Task was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: task.Label}
		}
		return e.Task, nil
	}
	return nil, &NotLoadedError{edge: "Task"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PractTest) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case practtest.FieldConfig:
			values[i] = new([]byte)
		case practtest.FieldID, practtest.FieldTestID, practtest.FieldDuration:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type PractTest", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PractTest fields.
func (pt *PractTest) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case practtest.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pt.ID = int(value.Int64)
		case practtest.FieldTestID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field test_id", values[i])
			} else if value.Valid {
				pt.TestID = int(value.Int64)
			}
		case practtest.FieldConfig:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field config", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pt.Config); err != nil {
					return fmt.Errorf("unmarshal field config: %w", err)
				}
			}
		case practtest.FieldDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration", values[i])
			} else if value.Valid {
				pt.Duration = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryTest queries the "Test" edge of the PractTest entity.
func (pt *PractTest) QueryTest() *TestQuery {
	return (&PractTestClient{config: pt.config}).QueryTest(pt)
}

// QueryTask queries the "Task" edge of the PractTest entity.
func (pt *PractTest) QueryTask() *TaskQuery {
	return (&PractTestClient{config: pt.config}).QueryTask(pt)
}

// Update returns a builder for updating this PractTest.
// Note that you need to call PractTest.Unwrap() before calling this method if this PractTest
// was returned from a transaction, and the transaction was committed or rolled back.
func (pt *PractTest) Update() *PractTestUpdateOne {
	return (&PractTestClient{config: pt.config}).UpdateOne(pt)
}

// Unwrap unwraps the PractTest entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pt *PractTest) Unwrap() *PractTest {
	tx, ok := pt.config.driver.(*txDriver)
	if !ok {
		panic("ent: PractTest is not a transactional entity")
	}
	pt.config.driver = tx.drv
	return pt
}

// String implements the fmt.Stringer.
func (pt *PractTest) String() string {
	var builder strings.Builder
	builder.WriteString("PractTest(")
	builder.WriteString(fmt.Sprintf("id=%v", pt.ID))
	builder.WriteString(", test_id=")
	builder.WriteString(fmt.Sprintf("%v", pt.TestID))
	builder.WriteString(", config=")
	builder.WriteString(fmt.Sprintf("%v", pt.Config))
	builder.WriteString(", duration=")
	builder.WriteString(fmt.Sprintf("%v", pt.Duration))
	builder.WriteByte(')')
	return builder.String()
}

// PractTests is a parsable slice of PractTest.
type PractTests []*PractTest

func (pt PractTests) config(cfg config) {
	for _i := range pt {
		pt[_i].config = cfg
	}
}
