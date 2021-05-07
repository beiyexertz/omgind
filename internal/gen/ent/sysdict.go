// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/wanhello/omgind/internal/gen/ent/sysdict"
)

// SysDict is the model entity for the SysDict schema.
type SysDict struct {
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
	// Sort holds the value of the "sort" field.
	// 排序, 在数据库里的排序
	Sort int32 `json:"sort,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	// 创建时间,由程序自动生成
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	// 更新时间,由程序自动生成
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	// 删除时间,
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// NameCn holds the value of the "name_cn" field.
	// 字典名（中）
	NameCn string `json:"name_cn,omitempty"`
	// NameEn holds the value of the "name_en" field.
	// 字典名（英）
	NameEn string `json:"name_en,omitempty"`
	// Status holds the value of the "Status" field.
	// 状态
	Status bool `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SysDictQuery when eager-loading is set.
	Edges SysDictEdges `json:"edges"`
}

// SysDictEdges holds the relations/edges for other nodes in the graph.
type SysDictEdges struct {
	// SysDictItems holds the value of the SysDictItems edge.
	SysDictItems []*SysDictItem `json:"SysDictItems,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// SysDictItemsOrErr returns the SysDictItems value or an error if the edge
// was not loaded in eager-loading.
func (e SysDictEdges) SysDictItemsOrErr() ([]*SysDictItem, error) {
	if e.loadedTypes[0] {
		return e.SysDictItems, nil
	}
	return nil, &NotLoadedError{edge: "SysDictItems"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SysDict) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case sysdict.FieldIsDel, sysdict.FieldStatus:
			values[i] = new(sql.NullBool)
		case sysdict.FieldSort:
			values[i] = new(sql.NullInt64)
		case sysdict.FieldID, sysdict.FieldMemo, sysdict.FieldNameCn, sysdict.FieldNameEn:
			values[i] = new(sql.NullString)
		case sysdict.FieldCreatedAt, sysdict.FieldUpdatedAt, sysdict.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SysDict", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SysDict fields.
func (sd *SysDict) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sysdict.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				sd.ID = value.String
			}
		case sysdict.FieldIsDel:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_del", values[i])
			} else if value.Valid {
				sd.IsDel = value.Bool
			}
		case sysdict.FieldMemo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field memo", values[i])
			} else if value.Valid {
				sd.Memo = value.String
			}
		case sysdict.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				sd.Sort = int32(value.Int64)
			}
		case sysdict.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sd.CreatedAt = value.Time
			}
		case sysdict.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sd.UpdatedAt = value.Time
			}
		case sysdict.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				sd.DeletedAt = new(time.Time)
				*sd.DeletedAt = value.Time
			}
		case sysdict.FieldNameCn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_cn", values[i])
			} else if value.Valid {
				sd.NameCn = value.String
			}
		case sysdict.FieldNameEn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_en", values[i])
			} else if value.Valid {
				sd.NameEn = value.String
			}
		case sysdict.FieldStatus:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field Status", values[i])
			} else if value.Valid {
				sd.Status = value.Bool
			}
		}
	}
	return nil
}

// QuerySysDictItems queries the "SysDictItems" edge of the SysDict entity.
func (sd *SysDict) QuerySysDictItems() *SysDictItemQuery {
	return (&SysDictClient{config: sd.config}).QuerySysDictItems(sd)
}

// Update returns a builder for updating this SysDict.
// Note that you need to call SysDict.Unwrap() before calling this method if this SysDict
// was returned from a transaction, and the transaction was committed or rolled back.
func (sd *SysDict) Update() *SysDictUpdateOne {
	return (&SysDictClient{config: sd.config}).UpdateOne(sd)
}

// Unwrap unwraps the SysDict entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sd *SysDict) Unwrap() *SysDict {
	tx, ok := sd.config.driver.(*txDriver)
	if !ok {
		panic("ent: SysDict is not a transactional entity")
	}
	sd.config.driver = tx.drv
	return sd
}

// String implements the fmt.Stringer.
func (sd *SysDict) String() string {
	var builder strings.Builder
	builder.WriteString("SysDict(")
	builder.WriteString(fmt.Sprintf("id=%v", sd.ID))
	builder.WriteString(", is_del=")
	builder.WriteString(fmt.Sprintf("%v", sd.IsDel))
	builder.WriteString(", memo=")
	builder.WriteString(sd.Memo)
	builder.WriteString(", sort=")
	builder.WriteString(fmt.Sprintf("%v", sd.Sort))
	builder.WriteString(", created_at=")
	builder.WriteString(sd.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(sd.UpdatedAt.Format(time.ANSIC))
	if v := sd.DeletedAt; v != nil {
		builder.WriteString(", deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", name_cn=")
	builder.WriteString(sd.NameCn)
	builder.WriteString(", name_en=")
	builder.WriteString(sd.NameEn)
	builder.WriteString(", Status=")
	builder.WriteString(fmt.Sprintf("%v", sd.Status))
	builder.WriteByte(')')
	return builder.String()
}

// SysDicts is a parsable slice of SysDict.
type SysDicts []*SysDict

func (sd SysDicts) config(cfg config) {
	for _i := range sd {
		sd[_i].config = cfg
	}
}
