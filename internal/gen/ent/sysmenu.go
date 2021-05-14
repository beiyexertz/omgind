// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/wanhello/omgind/internal/gen/ent/sysmenu"
	"github.com/wanhello/omgind/pkg/helper/pulid"
)

// SysMenu is the model entity for the SysMenu schema.
type SysMenu struct {
	config `json:"-"`
	// ID of the ent.
	ID pulid.ID `json:"id,omitempty"`
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
	// Status holds the value of the "status" field.
	// 状态,
	Status int32 `json:"status,omitempty"`
	// Name holds the value of the "name" field.
	// 菜单名称
	Name string `json:"name,omitempty"`
	// Icon holds the value of the "icon" field.
	// 菜单图标
	Icon string `json:"icon,omitempty"`
	// Router holds the value of the "router" field.
	// 前端路由
	Router string `json:"router,omitempty"`
	// IsShow holds the value of the "is_show" field.
	// 是否显示
	IsShow bool `json:"is_show,omitempty"`
	// ParentID holds the value of the "parent_id" field.
	// 父级id
	ParentID *string `json:"parent_id,omitempty"`
	// ParentPath holds the value of the "parent_path" field.
	// 父级路径: 1/2/3
	ParentPath *string `json:"parent_path,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SysMenu) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case sysmenu.FieldID:
			values[i] = new(pulid.ID)
		case sysmenu.FieldIsShow:
			values[i] = new(sql.NullBool)
		case sysmenu.FieldSort, sysmenu.FieldStatus:
			values[i] = new(sql.NullInt64)
		case sysmenu.FieldMemo, sysmenu.FieldName, sysmenu.FieldIcon, sysmenu.FieldRouter, sysmenu.FieldParentID, sysmenu.FieldParentPath:
			values[i] = new(sql.NullString)
		case sysmenu.FieldCreatedAt, sysmenu.FieldUpdatedAt, sysmenu.FieldDeletedAt:
			values[i] = new(sql.NullTime)
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
			if value, ok := values[i].(*pulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				sm.ID = *value
			}
		case sysmenu.FieldMemo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field memo", values[i])
			} else if value.Valid {
				sm.Memo = value.String
			}
		case sysmenu.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				sm.Sort = int32(value.Int64)
			}
		case sysmenu.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sm.CreatedAt = value.Time
			}
		case sysmenu.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sm.UpdatedAt = value.Time
			}
		case sysmenu.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				sm.DeletedAt = new(time.Time)
				*sm.DeletedAt = value.Time
			}
		case sysmenu.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				sm.Status = int32(value.Int64)
			}
		case sysmenu.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				sm.Name = value.String
			}
		case sysmenu.FieldIcon:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field icon", values[i])
			} else if value.Valid {
				sm.Icon = value.String
			}
		case sysmenu.FieldRouter:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field router", values[i])
			} else if value.Valid {
				sm.Router = value.String
			}
		case sysmenu.FieldIsShow:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_show", values[i])
			} else if value.Valid {
				sm.IsShow = value.Bool
			}
		case sysmenu.FieldParentID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field parent_id", values[i])
			} else if value.Valid {
				sm.ParentID = new(string)
				*sm.ParentID = value.String
			}
		case sysmenu.FieldParentPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field parent_path", values[i])
			} else if value.Valid {
				sm.ParentPath = new(string)
				*sm.ParentPath = value.String
			}
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
	builder.WriteString(", memo=")
	builder.WriteString(sm.Memo)
	builder.WriteString(", sort=")
	builder.WriteString(fmt.Sprintf("%v", sm.Sort))
	builder.WriteString(", created_at=")
	builder.WriteString(sm.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(sm.UpdatedAt.Format(time.ANSIC))
	if v := sm.DeletedAt; v != nil {
		builder.WriteString(", deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", sm.Status))
	builder.WriteString(", name=")
	builder.WriteString(sm.Name)
	builder.WriteString(", icon=")
	builder.WriteString(sm.Icon)
	builder.WriteString(", router=")
	builder.WriteString(sm.Router)
	builder.WriteString(", is_show=")
	builder.WriteString(fmt.Sprintf("%v", sm.IsShow))
	if v := sm.ParentID; v != nil {
		builder.WriteString(", parent_id=")
		builder.WriteString(*v)
	}
	if v := sm.ParentPath; v != nil {
		builder.WriteString(", parent_path=")
		builder.WriteString(*v)
	}
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
