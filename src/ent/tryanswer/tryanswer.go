// Code generated by entc, DO NOT EDIT.

package tryanswer

const (
	// Label holds the string label denoting the tryanswer type in the database.
	Label = "try_answer"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTryID holds the string denoting the try_id field in the database.
	FieldTryID = "try_id"
	// FieldQuestionID holds the string denoting the question_id field in the database.
	FieldQuestionID = "question_id"
	// FieldAsnwer holds the string denoting the asnwer field in the database.
	FieldAsnwer = "asnwer"
	// FieldCorrect holds the string denoting the correct field in the database.
	FieldCorrect = "correct"
	// EdgeTheoreticalTry holds the string denoting the theoreticaltry edge name in mutations.
	EdgeTheoreticalTry = "TheoreticalTry"
	// EdgeQuestion holds the string denoting the question edge name in mutations.
	EdgeQuestion = "Question"
	// Table holds the table name of the tryanswer in the database.
	Table = "TryAnswer"
	// TheoreticalTryTable is the table that holds the TheoreticalTry relation/edge.
	TheoreticalTryTable = "TryAnswer"
	// TheoreticalTryInverseTable is the table name for the TheoreticalTry entity.
	// It exists in this package in order to avoid circular dependency with the "theoreticaltry" package.
	TheoreticalTryInverseTable = "TheoreticalTry"
	// TheoreticalTryColumn is the table column denoting the TheoreticalTry relation/edge.
	TheoreticalTryColumn = "try_id"
	// QuestionTable is the table that holds the Question relation/edge.
	QuestionTable = "TryAnswer"
	// QuestionInverseTable is the table name for the Question entity.
	// It exists in this package in order to avoid circular dependency with the "question" package.
	QuestionInverseTable = "Question"
	// QuestionColumn is the table column denoting the Question relation/edge.
	QuestionColumn = "question_id"
)

// Columns holds all SQL columns for tryanswer fields.
var Columns = []string{
	FieldID,
	FieldTryID,
	FieldQuestionID,
	FieldAsnwer,
	FieldCorrect,
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
