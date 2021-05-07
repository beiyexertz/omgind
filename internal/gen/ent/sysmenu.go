// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/wanhello/omgind/internal/gen/ent/sysmenu"
)

// SysMenu is the model entity for the SysMenu schema.
type SysMenu struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SysMenu) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case sysmenu.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SysMenu", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SysMenu fields.
func (sm *SysMenu) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sysmenu.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sm.ID = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this SysMenu.
// Note that you need to call SysMenu.Unwrap() before calling this method if this SysMenu
// was returned from a transaction, and the transaction was committed or rolled back.
func (sm *SysMenu) Update() *SysMenuUpdateOne {
	return (&SysMenuClient{config: sm.config}).UpdateOne(sm)
}

// Unwrap unwraps the SysMenu entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sm *SysMenu) Unwrap() *SysMenu {
	tx, ok := sm.config.driver.(*txDriver)
	if !ok {
		panic("ent: SysMenu is not a transactional entity")
	}
	sm.config.driver = tx.drv
	return sm
}

// String implements the fmt.Stringer.
func (sm *SysMenu) String() string {
	var builder strings.Builder
	builder.WriteString("SysMenu(")
	builder.WriteString(fmt.Sprintf("id=%v", sm.ID))
	builder.WriteByte(')')
	return builder.String()
}

// SysMenus is a parsable slice of SysMenu.
type SysMenus []*SysMenu

func (sm SysMenus) config(cfg config) {
	for _i := range sm {
		sm[_i].config = cfg
	}
}
