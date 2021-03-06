// Code generated by entc, DO NOT EDIT.

package theoreticaltest

const (
	// Label holds the string label denoting the theoreticaltest type in the database.
	Label = "theoretical_test"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTestID holds the string denoting the test_id field in the database.
	FieldTestID = "test_id"
	// FieldDuration holds the string denoting the duration field in the database.
	FieldDuration = "duration"
	// EdgeTest holds the string denoting the test edge name in mutations.
	EdgeTest = "Test"
	// EdgeQuestion holds the string denoting the question edge name in mutations.
	EdgeQuestion = "Question"
	// EdgeTheoTry holds the string denoting the theotry edge name in mutations.
	EdgeTheoTry = "TheoTry"
	// Table holds the table name of the theoreticaltest in the database.
	Table = "TheoreticalTest"
	// TestTable is the table that holds the Test relation/edge.
	TestTable = "TheoreticalTest"
	// TestInverseTable is the table name for the Test entity.
	// It exists in this package in order to avoid circular dependency with the "test" package.
	TestInverseTable = "Test"
	// TestColumn is the table column denoting the Test relation/edge.
	TestColumn = "test_id"
	// QuestionTable is the table that holds the Question relation/edge.
	QuestionTable = "Question"
	// QuestionInverseTable is the table name for the Question entity.
	// It exists in this package in order to avoid circular dependency with the "question" package.
	QuestionInverseTable = "Question"
	// QuestionColumn is the table column denoting the Question relation/edge.
	QuestionColumn = "theorical_test_id"
	// TheoTryTable is the table that holds the TheoTry relation/edge.
	TheoTryTable = "TheoreticalTry"
	// TheoTryInverseTable is the table name for the TheoreticalTry entity.
	// It exists in this package in order to avoid circular dependency with the "theoreticaltry" package.
	TheoTryInverseTable = "TheoreticalTry"
	// TheoTryColumn is the table column denoting the TheoTry relation/edge.
	TheoTryColumn = "theoretical_test_id"
)

// Columns holds all SQL columns for theoreticaltest fields.
var Columns = []string{
	FieldID,
	FieldTestID,
	FieldDuration,
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
