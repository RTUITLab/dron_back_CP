// Code generated by entc, DO NOT EDIT.

package test

const (
	// Label holds the string label denoting the test type in the database.
	Label = "test"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTestType holds the string denoting the testtype field in the database.
	FieldTestType = "test_type"
	// EdgeModuleTest holds the string denoting the moduletest edge name in mutations.
	EdgeModuleTest = "ModuleTest"
	// EdgeSubmoduleTest holds the string denoting the submoduletest edge name in mutations.
	EdgeSubmoduleTest = "SubmoduleTest"
	// EdgeTherTest holds the string denoting the thertest edge name in mutations.
	EdgeTherTest = "TherTest"
	// EdgePractTest holds the string denoting the practtest edge name in mutations.
	EdgePractTest = "PractTest"
	// Table holds the table name of the test in the database.
	Table = "Test"
	// ModuleTestTable is the table that holds the ModuleTest relation/edge.
	ModuleTestTable = "ModuleTest"
	// ModuleTestInverseTable is the table name for the ModuleTest entity.
	// It exists in this package in order to avoid circular dependency with the "moduletest" package.
	ModuleTestInverseTable = "ModuleTest"
	// ModuleTestColumn is the table column denoting the ModuleTest relation/edge.
	ModuleTestColumn = "test_id"
	// SubmoduleTestTable is the table that holds the SubmoduleTest relation/edge.
	SubmoduleTestTable = "SubModuleTest"
	// SubmoduleTestInverseTable is the table name for the SubModuleTest entity.
	// It exists in this package in order to avoid circular dependency with the "submoduletest" package.
	SubmoduleTestInverseTable = "SubModuleTest"
	// SubmoduleTestColumn is the table column denoting the SubmoduleTest relation/edge.
	SubmoduleTestColumn = "test_id"
	// TherTestTable is the table that holds the TherTest relation/edge.
	TherTestTable = "TheoreticalTest"
	// TherTestInverseTable is the table name for the TheoreticalTest entity.
	// It exists in this package in order to avoid circular dependency with the "theoreticaltest" package.
	TherTestInverseTable = "TheoreticalTest"
	// TherTestColumn is the table column denoting the TherTest relation/edge.
	TherTestColumn = "test_id"
	// PractTestTable is the table that holds the PractTest relation/edge.
	PractTestTable = "PractTest"
	// PractTestInverseTable is the table name for the PractTest entity.
	// It exists in this package in order to avoid circular dependency with the "practtest" package.
	PractTestInverseTable = "PractTest"
	// PractTestColumn is the table column denoting the PractTest relation/edge.
	PractTestColumn = "test_id"
)

// Columns holds all SQL columns for test fields.
var Columns = []string{
	FieldID,
	FieldTestType,
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
