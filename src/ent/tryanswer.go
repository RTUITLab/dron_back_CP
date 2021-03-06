// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/0B1t322/CP-Rosseti-Back/ent/question"
	"github.com/0B1t322/CP-Rosseti-Back/ent/theoreticaltry"
	"github.com/0B1t322/CP-Rosseti-Back/ent/tryanswer"
)

// TryAnswer is the model entity for the TryAnswer schema.
type TryAnswer struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// TryID holds the value of the "try_id" field.
	TryID int `json:"try_id,omitempty"`
	// QuestionID holds the value of the "question_id" field.
	QuestionID int `json:"question_id,omitempty"`
	// Asnwer holds the value of the "asnwer" field.
	Asnwer string `json:"asnwer,omitempty"`
	// Correct holds the value of the "correct" field.
	Correct bool `json:"correct,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TryAnswerQuery when eager-loading is set.
	Edges TryAnswerEdges `json:"edges"`
}

// TryAnswerEdges holds the relations/edges for other nodes in the graph.
type TryAnswerEdges struct {
	// TheoreticalTry holds the value of the TheoreticalTry edge.
	TheoreticalTry *TheoreticalTry `json:"TheoreticalTry,omitempty"`
	// Question holds the value of the Question edge.
	Question *Question `json:"Question,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TheoreticalTryOrErr returns the TheoreticalTry value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TryAnswerEdges) TheoreticalTryOrErr() (*TheoreticalTry, error) {
	if e.loadedTypes[0] {
		if e.TheoreticalTry == nil {
			// The edge TheoreticalTry was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: theoreticaltry.Label}
		}
		return e.TheoreticalTry, nil
	}
	return nil, &NotLoadedError{edge: "TheoreticalTry"}
}

// QuestionOrErr returns the Question value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TryAnswerEdges) QuestionOrErr() (*Question, error) {
	if e.loadedTypes[1] {
		if e.Question == nil {
			// The edge Question was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: question.Label}
		}
		return e.Question, nil
	}
	return nil, &NotLoadedError{edge: "Question"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TryAnswer) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case tryanswer.FieldCorrect:
			values[i] = new(sql.NullBool)
		case tryanswer.FieldID, tryanswer.FieldTryID, tryanswer.FieldQuestionID:
			values[i] = new(sql.NullInt64)
		case tryanswer.FieldAsnwer:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TryAnswer", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TryAnswer fields.
func (ta *TryAnswer) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tryanswer.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ta.ID = int(value.Int64)
		case tryanswer.FieldTryID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field try_id", values[i])
			} else if value.Valid {
				ta.TryID = int(value.Int64)
			}
		case tryanswer.FieldQuestionID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field question_id", values[i])
			} else if value.Valid {
				ta.QuestionID = int(value.Int64)
			}
		case tryanswer.FieldAsnwer:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field asnwer", values[i])
			} else if value.Valid {
				ta.Asnwer = value.String
			}
		case tryanswer.FieldCorrect:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field correct", values[i])
			} else if value.Valid {
				ta.Correct = value.Bool
			}
		}
	}
	return nil
}

// QueryTheoreticalTry queries the "TheoreticalTry" edge of the TryAnswer entity.
func (ta *TryAnswer) QueryTheoreticalTry() *TheoreticalTryQuery {
	return (&TryAnswerClient{config: ta.config}).QueryTheoreticalTry(ta)
}

// QueryQuestion queries the "Question" edge of the TryAnswer entity.
func (ta *TryAnswer) QueryQuestion() *QuestionQuery {
	return (&TryAnswerClient{config: ta.config}).QueryQuestion(ta)
}

// Update returns a builder for updating this TryAnswer.
// Note that you need to call TryAnswer.Unwrap() before calling this method if this TryAnswer
// was returned from a transaction, and the transaction was committed or rolled back.
func (ta *TryAnswer) Update() *TryAnswerUpdateOne {
	return (&TryAnswerClient{config: ta.config}).UpdateOne(ta)
}

// Unwrap unwraps the TryAnswer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ta *TryAnswer) Unwrap() *TryAnswer {
	tx, ok := ta.config.driver.(*txDriver)
	if !ok {
		panic("ent: TryAnswer is not a transactional entity")
	}
	ta.config.driver = tx.drv
	return ta
}

// String implements the fmt.Stringer.
func (ta *TryAnswer) String() string {
	var builder strings.Builder
	builder.WriteString("TryAnswer(")
	builder.WriteString(fmt.Sprintf("id=%v", ta.ID))
	builder.WriteString(", try_id=")
	builder.WriteString(fmt.Sprintf("%v", ta.TryID))
	builder.WriteString(", question_id=")
	builder.WriteString(fmt.Sprintf("%v", ta.QuestionID))
	builder.WriteString(", asnwer=")
	builder.WriteString(ta.Asnwer)
	builder.WriteString(", correct=")
	builder.WriteString(fmt.Sprintf("%v", ta.Correct))
	builder.WriteByte(')')
	return builder.String()
}

// TryAnswers is a parsable slice of TryAnswer.
type TryAnswers []*TryAnswer

func (ta TryAnswers) config(cfg config) {
	for _i := range ta {
		ta[_i].config = cfg
	}
}
