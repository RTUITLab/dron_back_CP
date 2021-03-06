// Code generated by entc, DO NOT EDIT.

package moduledependcies

const (
	// Label holds the string label denoting the moduledependcies type in the database.
	Label = "module_dependcies"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDependentID holds the string denoting the dependent_id field in the database.
	FieldDependentID = "dependent_id"
	// FieldDependentOnID holds the string denoting the dependent_on_id field in the database.
	FieldDependentOnID = "dependent_on_id"
	// EdgeModuleDependcies holds the string denoting the moduledependcies edge name in mutations.
	EdgeModuleDependcies = "ModuleDependcies"
	// EdgeModuleDependOn holds the string denoting the moduledependon edge name in mutations.
	EdgeModuleDependOn = "ModuleDependOn"
	// Table holds the table name of the moduledependcies in the database.
	Table = "ModuleDependcies"
	// ModuleDependciesTable is the table that holds the ModuleDependcies relation/edge.
	ModuleDependciesTable = "ModuleDependcies"
	// ModuleDependciesInverseTable is the table name for the Module entity.
	// It exists in this package in order to avoid circular dependency with the "module" package.
	ModuleDependciesInverseTable = "Module"
	// ModuleDependciesColumn is the table column denoting the ModuleDependcies relation/edge.
	ModuleDependciesColumn = "dependent_id"
	// ModuleDependOnTable is the table that holds the ModuleDependOn relation/edge.
	ModuleDependOnTable = "ModuleDependcies"
	// ModuleDependOnInverseTable is the table name for the Module entity.
	// It exists in this package in order to avoid circular dependency with the "module" package.
	ModuleDependOnInverseTable = "Module"
	// ModuleDependOnColumn is the table column denoting the ModuleDependOn relation/edge.
	ModuleDependOnColumn = "dependent_on_id"
)

// Columns holds all SQL columns for moduledependcies fields.
var Columns = []string{
	FieldID,
	FieldDependentID,
	FieldDependentOnID,
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
