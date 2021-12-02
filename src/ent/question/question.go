// Code generated by entc, DO NOT EDIT.

package question

const (
	// Label holds the string label denoting the question type in the database.
	Label = "question"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTheoricalTestID holds the string denoting the theorical_test_id field in the database.
	FieldTheoricalTestID = "theorical_test_id"
	// FieldQuestion holds the string denoting the question field in the database.
	FieldQuestion = "question"
	// EdgeTheoreticalTest holds the string denoting the theoreticaltest edge name in mutations.
	EdgeTheoreticalTest = "TheoreticalTest"
	// EdgeAnswer holds the string denoting the answer edge name in mutations.
	EdgeAnswer = "Answer"
	// Table holds the table name of the question in the database.
	Table = "Question"
	// TheoreticalTestTable is the table that holds the TheoreticalTest relation/edge.
	TheoreticalTestTable = "Question"
	// TheoreticalTestInverseTable is the table name for the TheoreticalTest entity.
	// It exists in this package in order to avoid circular dependency with the "theoreticaltest" package.
	TheoreticalTestInverseTable = "TheoreticalTest"
	// TheoreticalTestColumn is the table column denoting the TheoreticalTest relation/edge.
	TheoreticalTestColumn = "theorical_test_id"
	// AnswerTable is the table that holds the Answer relation/edge.
	AnswerTable = "answers"
	// AnswerInverseTable is the table name for the Answer entity.
	// It exists in this package in order to avoid circular dependency with the "answer" package.
	AnswerInverseTable = "answers"
	// AnswerColumn is the table column denoting the Answer relation/edge.
	AnswerColumn = "question_id"
)

// Columns holds all SQL columns for question fields.
var Columns = []string{
	FieldID,
	FieldTheoricalTestID,
	FieldQuestion,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
