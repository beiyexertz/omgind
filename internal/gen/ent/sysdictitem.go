// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/wanhello/omgind/internal/gen/ent/sysdict"
	"github.com/wanhello/omgind/internal/gen/ent/sysdictitem"
)

// SysDictItem is the model entity for the SysDictItem schema.
type SysDictItem struct {
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
	// Label holds the value of the "label" field.
	// 显示值
	Label string `json:"label,omitempty"`
	// Value holds the value of the "value" field.
	// 字典值
	Value int `json:"value,omitempty"`
	// Status holds the value of the "status" field.
	// 启用状态
	Status bool `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SysDictItemQuery when eager-loading is set.
	Edges                   SysDictItemEdges `json:"edges"`
	sys_dict_sys_dict_items *string
}

// SysDictItemEdges holds the relations/edges for other nodes in the graph.
type SysDictItemEdges struct {
	// SysDict holds the value of the SysDict edge.
	SysDict *SysDict `json:"SysDict,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// SysDictOrErr returns the SysDict value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SysDictItemEdges) SysDictOrErr() (*SysDict, error) {
	if e.loadedTypes[0] {
		if e.SysDict == nil {
			// The edge SysDict was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: sysdict.Label}
		}
		return e.SysDict, nil
	}
	return nil, &NotLoadedError{edge: "SysDict"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SysDictItem) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case sysdictitem.FieldIsDel, sysdictitem.FieldStatus:
			values[i] = new(sql.NullBool)
		case sysdictitem.FieldSort, sysdictitem.FieldValue:
			values[i] = new(sql.NullInt64)
		case sysdictitem.FieldID, sysdictitem.FieldMemo, sysdictitem.FieldLabel:
			values[i] = new(sql.NullString)
		case sysdictitem.FieldCreatedAt, sysdictitem.FieldUpdatedAt, sysdictitem.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case sysdictitem.ForeignKeys[0]: // sys_dict_sys_dict_items
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SysDictItem", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SysDictItem fields.
func (sdi *SysDictItem) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sysdictitem.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				sdi.ID = value.String
			}
		case sysdictitem.FieldIsDel:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_del", values[i])
			} else if value.Valid {
				sdi.IsDel = value.Bool
			}
		case sysdictitem.FieldMemo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field memo", values[i])
			} else if value.Valid {
				sdi.Memo = value.String
			}
		case sysdictitem.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				sdi.Sort = int32(value.Int64)
			}
		case sysdictitem.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sdi.CreatedAt = value.Time
			}
		case sysdictitem.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sdi.UpdatedAt = value.Time
			}
		case sysdictitem.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				sdi.DeletedAt = new(time.Time)
				*sdi.DeletedAt = value.Time
			}
		case sysdictitem.FieldLabel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field label", values[i])
			} else if value.Valid {
				sdi.Label = value.String
			}
		case sysdictitem.FieldValue:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				sdi.Value = int(value.Int64)
			}
		case sysdictitem.FieldStatus:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				sdi.Status = value.Bool
			}
		case sysdictitem.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sys_dict_sys_dict_items", values[i])
			} else if value.Valid {
				sdi.sys_dict_sys_dict_items = new(string)
				*sdi.sys_dict_sys_dict_items = value.String
			}
		}
	}
	return nil
}

// QuerySysDict queries the "SysDict" edge of the SysDictItem entity.
func (sdi *SysDictItem) QuerySysDict() *SysDictQuery {
	return (&SysDictItemClient{config: sdi.config}).QuerySysDict(sdi)
}

// Update returns a builder for updating this SysDictItem.
// Note that you need to call SysDictItem.Unwrap() before calling this method if this SysDictItem
// was returned from a transaction, and the transaction was committed or rolled back.
func (sdi *SysDictItem) Update() *SysDictItemUpdateOne {
	return (&SysDictItemClient{config: sdi.config}).UpdateOne(sdi)
}

// Unwrap unwraps the SysDictItem entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sdi *SysDictItem) Unwrap() *SysDictItem {
	tx, ok := sdi.config.driver.(*txDriver)
	if !ok {
		panic("ent: SysDictItem is not a transactional entity")
	}
	sdi.config.driver = tx.drv
	return sdi
}

// String implements the fmt.Stringer.
func (sdi *SysDictItem) String() string {
	var builder strings.Builder
	builder.WriteString("SysDictItem(")
	builder.WriteString(fmt.Sprintf("id=%v", sdi.ID))
	builder.WriteString(", is_del=")
	builder.WriteString(fmt.Sprintf("%v", sdi.IsDel))
	builder.WriteString(", memo=")
	builder.WriteString(sdi.Memo)
	builder.WriteString(", sort=")
	builder.WriteString(fmt.Sprintf("%v", sdi.Sort))
	builder.WriteString(", created_at=")
	builder.WriteString(sdi.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(sdi.UpdatedAt.Format(time.ANSIC))
	if v := sdi.DeletedAt; v != nil {
		builder.WriteString(", deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", label=")
	builder.WriteString(sdi.Label)
	builder.WriteString(", value=")
	builder.WriteString(fmt.Sprintf("%v", sdi.Value))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", sdi.Status))
	builder.WriteByte(')')
	return builder.String()
}

// SysDictItems is a parsable slice of SysDictItem.
type SysDictItems []*SysDictItem

func (sdi SysDictItems) config(cfg config) {
	for _i := range sdi {
		sdi[_i].config = cfg
	}
}
