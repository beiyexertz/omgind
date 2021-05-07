// Code generated by entc, DO NOT EDIT.

package entcd

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/wanhello/omgind/internal/app/model/entcd/sysuser"
)

// SysUser is the model entity for the SysUser schema.
type SysUser struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SysUser) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case sysuser.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SysUser", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SysUser fields.
func (su *SysUser) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sysuser.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			su.ID = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this SysUser.
// Note that you need to call SysUser.Unwrap() before calling this method if this SysUser
// was returned from a transaction, and the transaction was committed or rolled back.
func (su *SysUser) Update() *SysUserUpdateOne {
	return (&SysUserClient{config: su.config}).UpdateOne(su)
}

// Unwrap unwraps the SysUser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (su *SysUser) Unwrap() *SysUser {
	tx, ok := su.config.driver.(*txDriver)
	if !ok {
		panic("entcd: SysUser is not a transactional entity")
	}
	su.config.driver = tx.drv
	return su
}

// String implements the fmt.Stringer.
func (su *SysUser) String() string {
	var builder strings.Builder
	builder.WriteString("SysUser(")
	builder.WriteString(fmt.Sprintf("id=%v", su.ID))
	builder.WriteByte(')')
	return builder.String()
}

// SysUsers is a parsable slice of SysUser.
type SysUsers []*SysUser

func (su SysUsers) config(cfg config) {
	for _i := range su {
		su[_i].config = cfg
	}
}
