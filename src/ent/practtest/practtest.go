// Code generated by entc, DO NOT EDIT.

package practtest

const (
	// Label holds the string label denoting the practtest type in the database.
	Label = "pract_test"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTestID holds the string denoting the test_id field in the database.
	FieldTestID = "test_id"
	// FieldConfig holds the string denoting the config field in the database.
	FieldConfig = "config"
	// FieldDuration holds the string denoting the duration field in the database.
	FieldDuration = "duration"
	// EdgeTest holds the string denoting the test edge name in mutations.
	EdgeTest = "Test"
	// EdgeTask holds the string denoting the task edge name in mutations.
	EdgeTask = "Task"
	// Table holds the table name of the practtest in the database.
	Table = "PractTest"
	// TestTable is the table that holds the Test relation/edge.
	TestTable = "PractTest"
	// TestInverseTable is the table name for the Test entity.
	// It exists in this package in order to avoid circular dependency with the "test" package.
	TestInverseTable = "Test"
	// TestColumn is the table column denoting the Test relation/edge.
	TestColumn = "test_id"
	// TaskTable is the table that holds the Task relation/edge.
	TaskTable = "Task"
	// TaskInverseTable is the table name for the Task entity.
	// It exists in this package in order to avoid circular dependency with the "task" package.
	TaskInverseTable = "Task"
	// TaskColumn is the table column denoting the Task relation/edge.
	TaskColumn = "pract_test_id"
)

// Columns holds all SQL columns for practtest fields.
var Columns = []string{
	FieldID,
	FieldTestID,
	FieldConfig,
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
