// Code generated by entc, DO NOT EDIT.

package task

const (
	// Label holds the string label denoting the task type in the database.
	Label = "task"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPractTestID holds the string denoting the pract_test_id field in the database.
	FieldPractTestID = "pract_test_id"
	// FieldTask holds the string denoting the task field in the database.
	FieldTask = "task"
	// EdgeTest holds the string denoting the test edge name in mutations.
	EdgeTest = "Test"
	// Table holds the table name of the task in the database.
	Table = "Task"
	// TestTable is the table that holds the Test relation/edge.
	TestTable = "Task"
	// TestInverseTable is the table name for the PractTest entity.
	// It exists in this package in order to avoid circular dependency with the "practtest" package.
	TestInverseTable = "PractTest"
	// TestColumn is the table column denoting the Test relation/edge.
	TestColumn = "pract_test_id"
)

// Columns holds all SQL columns for task fields.
var Columns = []string{
	FieldID,
	FieldPractTestID,
	FieldTask,
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
