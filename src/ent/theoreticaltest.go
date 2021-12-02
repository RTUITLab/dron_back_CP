// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/0B1t322/CP-Rosseti-Back/ent/submoduletest"
	"github.com/0B1t322/CP-Rosseti-Back/ent/theoreticaltest"
)

// TheoreticalTest is the model entity for the TheoreticalTest schema.
type TheoreticalTest struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// SubmoduletestID holds the value of the "submoduletest_id" field.
	SubmoduletestID int `json:"submoduletest_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TheoreticalTestQuery when eager-loading is set.
	Edges TheoreticalTestEdges `json:"edges"`
}

// TheoreticalTestEdges holds the relations/edges for other nodes in the graph.
type TheoreticalTestEdges struct {
	// SubModuleTest holds the value of the SubModuleTest edge.
	SubModuleTest *SubModuleTest `json:"SubModuleTest,omitempty"`
	// Question holds the value of the Question edge.
	Question []*Question `json:"Question,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// SubModuleTestOrErr returns the SubModuleTest value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TheoreticalTestEdges) SubModuleTestOrErr() (*SubModuleTest, error) {
	if e.loadedTypes[0] {
		if e.SubModuleTest == nil {
			// The edge SubModuleTest was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: submoduletest.Label}
		}
		return e.SubModuleTest, nil
	}
	return nil, &NotLoadedError{edge: "SubModuleTest"}
}

// QuestionOrErr returns the Question value or an error if the edge
// was not loaded in eager-loading.
func (e TheoreticalTestEdges) QuestionOrErr() ([]*Question, error) {
	if e.loadedTypes[1] {
		return e.Question, nil
	}
	return nil, &NotLoadedError{edge: "Question"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TheoreticalTest) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case theoreticaltest.FieldID, theoreticaltest.FieldSubmoduletestID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TheoreticalTest", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TheoreticalTest fields.
func (tt *TheoreticalTest) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case theoreticaltest.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tt.ID = int(value.Int64)
		case theoreticaltest.FieldSubmoduletestID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field submoduletest_id", values[i])
			} else if value.Valid {
				tt.SubmoduletestID = int(value.Int64)
			}
		}
	}
	return nil
}

// QuerySubModuleTest queries the "SubModuleTest" edge of the TheoreticalTest entity.
func (tt *TheoreticalTest) QuerySubModuleTest() *SubModuleTestQuery {
	return (&TheoreticalTestClient{config: tt.config}).QuerySubModuleTest(tt)
}

// QueryQuestion queries the "Question" edge of the TheoreticalTest entity.
func (tt *TheoreticalTest) QueryQuestion() *QuestionQuery {
	return (&TheoreticalTestClient{config: tt.config}).QueryQuestion(tt)
}

// Update returns a builder for updating this TheoreticalTest.
// Note that you need to call TheoreticalTest.Unwrap() before calling this method if this TheoreticalTest
// was returned from a transaction, and the transaction was committed or rolled back.
func (tt *TheoreticalTest) Update() *TheoreticalTestUpdateOne {
	return (&TheoreticalTestClient{config: tt.config}).UpdateOne(tt)
}

// Unwrap unwraps the TheoreticalTest entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tt *TheoreticalTest) Unwrap() *TheoreticalTest {
	tx, ok := tt.config.driver.(*txDriver)
	if !ok {
		panic("ent: TheoreticalTest is not a transactional entity")
	}
	tt.config.driver = tx.drv
	return tt
}

// String implements the fmt.Stringer.
func (tt *TheoreticalTest) String() string {
	var builder strings.Builder
	builder.WriteString("TheoreticalTest(")
	builder.WriteString(fmt.Sprintf("id=%v", tt.ID))
	builder.WriteString(", submoduletest_id=")
	builder.WriteString(fmt.Sprintf("%v", tt.SubmoduletestID))
	builder.WriteByte(')')
	return builder.String()
}

// TheoreticalTests is a parsable slice of TheoreticalTest.
type TheoreticalTests []*TheoreticalTest

func (tt TheoreticalTests) config(cfg config) {
	for _i := range tt {
		tt[_i].config = cfg
	}
}