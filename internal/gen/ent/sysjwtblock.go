// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/wanhello/omgind/internal/gen/ent/sysjwtblock"
)

// SysJwtBlock is the model entity for the SysJwtBlock schema.
type SysJwtBlock struct {
	config `json:"-"`
	// ID of the ent.
	// 主键
	ID string `json:"id,omitempty"`
	// IsDel holds the value of the "is_del" field.
	// 是否删除
	IsDel bool `json:"is_del,omitempty"`
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
	// Status holds the value of the "status" field.
	// 状态,
	Status int32 `json:"status,omitempty"`
	// Jwt holds the value of the "jwt" field.
	// jwt
	Jwt string `json:"jwt,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SysJwtBlock) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case sysjwtblock.FieldIsDel:
			values[i] = new(sql.NullBool)
		case sysjwtblock.FieldStatus:
			values[i] = new(sql.NullInt64)
		case sysjwtblock.FieldID, sysjwtblock.FieldMemo, sysjwtblock.FieldJwt:
			values[i] = new(sql.NullString)
		case sysjwtblock.FieldCreatedAt, sysjwtblock.FieldUpdatedAt, sysjwtblock.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SysJwtBlock", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SysJwtBlock fields.
func (sjb *SysJwtBlock) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sysjwtblock.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				sjb.ID = value.String
			}
		case sysjwtblock.FieldIsDel:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_del", values[i])
			} else if value.Valid {
				sjb.IsDel = value.Bool
			}
		case sysjwtblock.FieldMemo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field memo", values[i])
			} else if value.Valid {
				sjb.Memo = value.String
			}
		case sysjwtblock.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sjb.CreatedAt = value.Time
			}
		case sysjwtblock.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sjb.UpdatedAt = value.Time
			}
		case sysjwtblock.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				sjb.DeletedAt = new(time.Time)
				*sjb.DeletedAt = value.Time
			}
		case sysjwtblock.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				sjb.Status = int32(value.Int64)
			}
		case sysjwtblock.FieldJwt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field jwt", values[i])
			} else if value.Valid {
				sjb.Jwt = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this SysJwtBlock.
// Note that you need to call SysJwtBlock.Unwrap() before calling this method if this SysJwtBlock
// was returned from a transaction, and the transaction was committed or rolled back.
func (sjb *SysJwtBlock) Update() *SysJwtBlockUpdateOne {
	return (&SysJwtBlockClient{config: sjb.config}).UpdateOne(sjb)
}

// Unwrap unwraps the SysJwtBlock entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sjb *SysJwtBlock) Unwrap() *SysJwtBlock {
	tx, ok := sjb.config.driver.(*txDriver)
	if !ok {
		panic("ent: SysJwtBlock is not a transactional entity")
	}
	sjb.config.driver = tx.drv
	return sjb
}

// String implements the fmt.Stringer.
func (sjb *SysJwtBlock) String() string {
	var builder strings.Builder
	builder.WriteString("SysJwtBlock(")
	builder.WriteString(fmt.Sprintf("id=%v", sjb.ID))
	builder.WriteString(", is_del=")
	builder.WriteString(fmt.Sprintf("%v", sjb.IsDel))
	builder.WriteString(", memo=")
	builder.WriteString(sjb.Memo)
	builder.WriteString(", created_at=")
	builder.WriteString(sjb.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(sjb.UpdatedAt.Format(time.ANSIC))
	if v := sjb.DeletedAt; v != nil {
		builder.WriteString(", deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", sjb.Status))
	builder.WriteString(", jwt=")
	builder.WriteString(sjb.Jwt)
	builder.WriteByte(')')
	return builder.String()
}

// SysJwtBlocks is a parsable slice of SysJwtBlock.
type SysJwtBlocks []*SysJwtBlock

func (sjb SysJwtBlocks) config(cfg config) {
	for _i := range sjb {
		sjb[_i].config = cfg
	}
}