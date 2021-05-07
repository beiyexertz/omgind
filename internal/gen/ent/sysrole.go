// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/wanhello/omgind/internal/gen/ent/sysrole"
)

// SysRole is the model entity for the SysRole schema.
type SysRole struct {
	config `json:"-"`
	// ID of the ent.
	// 主键
	ID string `json:"id,omitempty"`
	// IsDel holds the value of the "is_del" field.
	// 是否删除
	IsDel bool `json:"is_del,omitempty"`
	// Status holds the value of the "status" field.
	// 状态,
	Status int32 `json:"sort,omitempty"`
	// Sort holds the value of the "sort" field.
	// 排序, 在数据库里的排序
	Sort int32 `json:"sort,omitempty"`
	// Memo holds the value of the "memo" field.
	// 备注
	Memo string `json:"memo,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	// 创建时间,由程序自动生成
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	// 更新时间,由程序自动生成
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	// 删除时间,
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Name holds the value of the "name" field.
	// 角色名称
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SysRoleQuery when eager-loading is set.
	Edges SysRoleEdges `json:"edges"`
}

// SysRoleEdges holds the relations/edges for other nodes in the graph.
type SysRoleEdges struct {
	// UserRoles holds the value of the userRoles edge.
	UserRoles []*SysUserRole `json:"userRoles,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserRolesOrErr returns the UserRoles value or an error if the edge
// was not loaded in eager-loading.
func (e SysRoleEdges) UserRolesOrErr() ([]*SysUserRole, error) {
	if e.loadedTypes[0] {
		return e.UserRoles, nil
	}
	return nil, &NotLoadedError{edge: "userRoles"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SysRole) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case sysrole.FieldIsDel:
			values[i] = new(sql.NullBool)
		case sysrole.FieldStatus, sysrole.FieldSort:
			values[i] = new(sql.NullInt64)
		case sysrole.FieldID, sysrole.FieldMemo, sysrole.FieldName:
			values[i] = new(sql.NullString)
		case sysrole.FieldCreatedAt, sysrole.FieldUpdatedAt, sysrole.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SysRole", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SysRole fields.
func (sr *SysRole) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sysrole.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				sr.ID = value.String
			}
		case sysrole.FieldIsDel:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_del", values[i])
			} else if value.Valid {
				sr.IsDel = value.Bool
			}
		case sysrole.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				sr.Status = int32(value.Int64)
			}
		case sysrole.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				sr.Sort = int32(value.Int64)
			}
		case sysrole.FieldMemo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field memo", values[i])
			} else if value.Valid {
				sr.Memo = value.String
			}
		case sysrole.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sr.CreatedAt = value.Time
			}
		case sysrole.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sr.UpdatedAt = value.Time
			}
		case sysrole.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				sr.DeletedAt = new(time.Time)
				*sr.DeletedAt = value.Time
			}
		case sysrole.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				sr.Name = value.String
			}
		}
	}
	return nil
}

// QueryUserRoles queries the "userRoles" edge of the SysRole entity.
func (sr *SysRole) QueryUserRoles() *SysUserRoleQuery {
	return (&SysRoleClient{config: sr.config}).QueryUserRoles(sr)
}

// Update returns a builder for updating this SysRole.
// Note that you need to call SysRole.Unwrap() before calling this method if this SysRole
// was returned from a transaction, and the transaction was committed or rolled back.
func (sr *SysRole) Update() *SysRoleUpdateOne {
	return (&SysRoleClient{config: sr.config}).UpdateOne(sr)
}

// Unwrap unwraps the SysRole entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sr *SysRole) Unwrap() *SysRole {
	tx, ok := sr.config.driver.(*txDriver)
	if !ok {
		panic("ent: SysRole is not a transactional entity")
	}
	sr.config.driver = tx.drv
	return sr
}

// String implements the fmt.Stringer.
func (sr *SysRole) String() string {
	var builder strings.Builder
	builder.WriteString("SysRole(")
	builder.WriteString(fmt.Sprintf("id=%v", sr.ID))
	builder.WriteString(", is_del=")
	builder.WriteString(fmt.Sprintf("%v", sr.IsDel))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", sr.Status))
	builder.WriteString(", sort=")
	builder.WriteString(fmt.Sprintf("%v", sr.Sort))
	builder.WriteString(", memo=")
	builder.WriteString(sr.Memo)
	builder.WriteString(", created_at=")
	builder.WriteString(sr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(sr.UpdatedAt.Format(time.ANSIC))
	if v := sr.DeletedAt; v != nil {
		builder.WriteString(", deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", name=")
	builder.WriteString(sr.Name)
	builder.WriteByte(')')
	return builder.String()
}

// SysRoles is a parsable slice of SysRole.
type SysRoles []*SysRole

func (sr SysRoles) config(cfg config) {
	for _i := range sr {
		sr[_i].config = cfg
	}
}
