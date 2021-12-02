// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/0B1t322/CP-Rosseti-Back/ent/question"
	"github.com/0B1t322/CP-Rosseti-Back/ent/theoreticaltest"
)

// Question is the model entity for the Question schema.
type Question struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// TheoricalTestID holds the value of the "theorical_test_id" field.
	TheoricalTestID int `json:"theorical_test_id,omitempty"`
	// Question holds the value of the "question" field.
	Question string `json:"question,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the QuestionQuery when eager-loading is set.
	Edges QuestionEdges `json:"edges"`
}

// QuestionEdges holds the relations/edges for other nodes in the graph.
type QuestionEdges struct {
	// TheoreticalTest holds the value of the TheoreticalTest edge.
	TheoreticalTest *TheoreticalTest `json:"TheoreticalTest,omitempty"`
	// Answer holds the value of the Answer edge.
	Answer []*Answer `json:"Answer,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TheoreticalTestOrErr returns the TheoreticalTest value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e QuestionEdges) TheoreticalTestOrErr() (*TheoreticalTest, error) {
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

// AnswerOrErr returns the Answer value or an error if the edge
// was not loaded in eager-loading.
func (e QuestionEdges) AnswerOrErr() ([]*Answer, error) {
	if e.loadedTypes[1] {
		return e.Answer, nil
	}
	return nil, &NotLoadedError{edge: "Answer"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Question) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case question.FieldID, question.FieldTheoricalTestID:
			values[i] = new(sql.NullInt64)
		case question.FieldQuestion:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Question", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Question fields.
func (q *Question) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case question.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			q.ID = int(value.Int64)
		case question.FieldTheoricalTestID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field theorical_test_id", values[i])
			} else if value.Valid {
				q.TheoricalTestID = int(value.Int64)
			}
		case question.FieldQuestion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field question", values[i])
			} else if value.Valid {
				q.Question = value.String
			}
		}
	}
	return nil
}

// QueryTheoreticalTest queries the "TheoreticalTest" edge of the Question entity.
func (q *Question) QueryTheoreticalTest() *TheoreticalTestQuery {
	return (&QuestionClient{config: q.config}).QueryTheoreticalTest(q)
}

// QueryAnswer queries the "Answer" edge of the Question entity.
func (q *Question) QueryAnswer() *AnswerQuery {
	return (&QuestionClient{config: q.config}).QueryAnswer(q)
}

// Update returns a builder for updating this Question.
// Note that you need to call Question.Unwrap() before calling this method if this Question
// was returned from a transaction, and the transaction was committed or rolled back.
func (q *Question) Update() *QuestionUpdateOne {
	return (&QuestionClient{config: q.config}).UpdateOne(q)
}

// Unwrap unwraps the Question entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (q *Question) Unwrap() *Question {
	tx, ok := q.config.driver.(*txDriver)
	if !ok {
		panic("ent: Question is not a transactional entity")
	}
	q.config.driver = tx.drv
	return q
}

// String implements the fmt.Stringer.
func (q *Question) String() string {
	var builder strings.Builder
	builder.WriteString("Question(")
	builder.WriteString(fmt.Sprintf("id=%v", q.ID))
	builder.WriteString(", theorical_test_id=")
	builder.WriteString(fmt.Sprintf("%v", q.TheoricalTestID))
	builder.WriteString(", question=")
	builder.WriteString(q.Question)
	builder.WriteByte(')')
	return builder.String()
}

// Questions is a parsable slice of Question.
type Questions []*Question

func (q Questions) config(cfg config) {
	for _i := range q {
		q[_i].config = cfg
	}
}
