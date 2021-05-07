// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/wanhello/omgind/internal/gen/ent/sysuserrole"
)

// SysUserRole is the model entity for the SysUserRole schema.
type SysUserRole struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SysUserRole) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case sysuserrole.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SysUserRole", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SysUserRole fields.
func (sur *SysUserRole) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sysuserrole.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sur.ID = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this SysUserRole.
// Note that you need to call SysUserRole.Unwrap() before calling this method if this SysUserRole
// was returned from a transaction, and the transaction was committed or rolled back.
func (sur *SysUserRole) Update() *SysUserRoleUpdateOne {
	return (&SysUserRoleClient{config: sur.config}).UpdateOne(sur)
}

// Unwrap unwraps the SysUserRole entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sur *SysUserRole) Unwrap() *SysUserRole {
	tx, ok := sur.config.driver.(*txDriver)
	if !ok {
		panic("ent: SysUserRole is not a transactional entity")
	}
	sur.config.driver = tx.drv
	return sur
}

// String implements the fmt.Stringer.
func (sur *SysUserRole) String() string {
	var builder strings.Builder
	builder.WriteString("SysUserRole(")
	builder.WriteString(fmt.Sprintf("id=%v", sur.ID))
	builder.WriteByte(')')
	return builder.String()
}

// SysUserRoles is a parsable slice of SysUserRole.
type SysUserRoles []*SysUserRole

func (sur SysUserRoles) config(cfg config) {
	for _i := range sur {
		sur[_i].config = cfg
	}
}
