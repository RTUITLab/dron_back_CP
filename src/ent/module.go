// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/0B1t322/CP-Rosseti-Back/ent/module"
)

// Module is the model entity for the Module schema.
type Module struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ModuleQuery when eager-loading is set.
	Edges ModuleEdges `json:"edges"`
}

// ModuleEdges holds the relations/edges for other nodes in the graph.
type ModuleEdges struct {
	// ModuleDependcies holds the value of the ModuleDependcies edge.
	ModuleDependcies []*ModuleDependcies `json:"ModuleDependcies,omitempty"`
	// ModuleDependOn holds the value of the ModuleDependOn edge.
	ModuleDependOn []*ModuleDependcies `json:"ModuleDependOn,omitempty"`
	// SubModules holds the value of the SubModules edge.
	SubModules []*SubModule `json:"SubModules,omitempty"`
	// Test holds the value of the Test edge.
	Test []*ModuleTest `json:"Test,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// ModuleDependciesOrErr returns the ModuleDependcies value or an error if the edge
// was not loaded in eager-loading.
func (e ModuleEdges) ModuleDependciesOrErr() ([]*ModuleDependcies, error) {
	if e.loadedTypes[0] {
		return e.ModuleDependcies, nil
	}
	return nil, &NotLoadedError{edge: "ModuleDependcies"}
}

// ModuleDependOnOrErr returns the ModuleDependOn value or an error if the edge
// was not loaded in eager-loading.
func (e ModuleEdges) ModuleDependOnOrErr() ([]*ModuleDependcies, error) {
	if e.loadedTypes[1] {
		return e.ModuleDependOn, nil
	}
	return nil, &NotLoadedError{edge: "ModuleDependOn"}
}

// SubModulesOrErr returns the SubModules value or an error if the edge
// was not loaded in eager-loading.
func (e ModuleEdges) SubModulesOrErr() ([]*SubModule, error) {
	if e.loadedTypes[2] {
		return e.SubModules, nil
	}
	return nil, &NotLoadedError{edge: "SubModules"}
}

// TestOrErr returns the Test value or an error if the edge
// was not loaded in eager-loading.
func (e ModuleEdges) TestOrErr() ([]*ModuleTest, error) {
	if e.loadedTypes[3] {
		return e.Test, nil
	}
	return nil, &NotLoadedError{edge: "Test"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Module) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case module.FieldID:
			values[i] = new(sql.NullInt64)
		case module.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Module", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Module fields.
func (m *Module) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case module.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case module.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				m.Name = value.String
			}
		}
	}
	return nil
}

// QueryModuleDependcies queries the "ModuleDependcies" edge of the Module entity.
func (m *Module) QueryModuleDependcies() *ModuleDependciesQuery {
	return (&ModuleClient{config: m.config}).QueryModuleDependcies(m)
}

// QueryModuleDependOn queries the "ModuleDependOn" edge of the Module entity.
func (m *Module) QueryModuleDependOn() *ModuleDependciesQuery {
	return (&ModuleClient{config: m.config}).QueryModuleDependOn(m)
}

// QuerySubModules queries the "SubModules" edge of the Module entity.
func (m *Module) QuerySubModules() *SubModuleQuery {
	return (&ModuleClient{config: m.config}).QuerySubModules(m)
}

// QueryTest queries the "Test" edge of the Module entity.
func (m *Module) QueryTest() *ModuleTestQuery {
	return (&ModuleClient{config: m.config}).QueryTest(m)
}

// Update returns a builder for updating this Module.
// Note that you need to call Module.Unwrap() before calling this method if this Module
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Module) Update() *ModuleUpdateOne {
	return (&ModuleClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the Module entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Module) Unwrap() *Module {
	tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Module is not a transactional entity")
	}
	m.config.driver = tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Module) String() string {
	var builder strings.Builder
	builder.WriteString("Module(")
	builder.WriteString(fmt.Sprintf("id=%v", m.ID))
	builder.WriteString(", name=")
	builder.WriteString(m.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Modules is a parsable slice of Module.
type Modules []*Module

func (m Modules) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}
