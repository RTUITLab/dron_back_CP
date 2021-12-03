// Code generated by entc, DO NOT EDIT.

package practtest

const (
	// Label holds the string label denoting the practtest type in the database.
	Label = "pract_test"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSubmoduletestID holds the string denoting the submoduletest_id field in the database.
	FieldSubmoduletestID = "submoduletest_id"
	// FieldConfig holds the string denoting the config field in the database.
	FieldConfig = "config"
	// EdgeSubModuleTest holds the string denoting the submoduletest edge name in mutations.
	EdgeSubModuleTest = "SubModuleTest"
	// Table holds the table name of the practtest in the database.
	Table = "PractTest"
	// SubModuleTestTable is the table that holds the SubModuleTest relation/edge.
	SubModuleTestTable = "PractTest"
	// SubModuleTestInverseTable is the table name for the SubModuleTest entity.
	// It exists in this package in order to avoid circular dependency with the "submoduletest" package.
	SubModuleTestInverseTable = "SubModuleTest"
	// SubModuleTestColumn is the table column denoting the SubModuleTest relation/edge.
	SubModuleTestColumn = "submoduletest_id"
)

// Columns holds all SQL columns for practtest fields.
var Columns = []string{
	FieldID,
	FieldSubmoduletestID,
	FieldConfig,
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