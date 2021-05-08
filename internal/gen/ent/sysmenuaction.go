// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/wanhello/omgind/internal/gen/ent/sysmenu"
	"github.com/wanhello/omgind/internal/gen/ent/sysmenuaction"
)

// SysMenuAction is the model entity for the SysMenuAction schema.
type SysMenuAction struct {
	config `json:"-"`
	// ID of the ent.
	// 主键
	ID string `json:"id,omitempty"`
	// IsDel holds the value of the "is_del" field.
	// 是否删除
	IsDel bool `json:"is_del,omitempty"`
	// Sort holds the value of the "sort" field.
	// 排序, 在数据库里的排序
	Sort int32 `json:"sort,omitempty"`
	// Status holds the value of the "status" field.
	// 状态,
	Status int32 `json:"status,omitempty"`
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
	// MenuID holds the value of the "menu_id" field.
	// 菜单ID, sys_menu.id
	MenuID string `json:"menu_id,omitempty"`
	// Code holds the value of the "code" field.
	// 动作编号
	Code string `json:"code,omitempty"`
	// Name holds the value of the "name" field.
	// 动作名称
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SysMenuActionQuery when eager-loading is set.
	Edges SysMenuActionEdges `json:"edges"`
}

// SysMenuActionEdges holds the relations/edges for other nodes in the graph.
type SysMenuActionEdges struct {
	// Resources holds the value of the resources edge.
	Resources []*SysMenuActionResource `json:"resources,omitempty"`
	// Menu holds the value of the menu edge.
	Menu *SysMenu `json:"menu,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ResourcesOrErr returns the Resources value or an error if the edge
// was not loaded in eager-loading.
func (e SysMenuActionEdges) ResourcesOrErr() ([]*SysMenuActionResource, error) {
	if e.loadedTypes[0] {
		return e.Resources, nil
	}
	return nil, &NotLoadedError{edge: "resources"}
}

// MenuOrErr returns the Menu value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SysMenuActionEdges) MenuOrErr() (*SysMenu, error) {
	if e.loadedTypes[1] {
		if e.Menu == nil {
			// The edge menu was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: sysmenu.Label}
		}
		return e.Menu, nil
	}
	return nil, &NotLoadedError{edge: "menu"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SysMenuAction) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case sysmenuaction.FieldIsDel:
			values[i] = new(sql.NullBool)
		case sysmenuaction.FieldSort, sysmenuaction.FieldStatus:
			values[i] = new(sql.NullInt64)
		case sysmenuaction.FieldID, sysmenuaction.FieldMemo, sysmenuaction.FieldMenuID, sysmenuaction.FieldCode, sysmenuaction.FieldName:
			values[i] = new(sql.NullString)
		case sysmenuaction.FieldCreatedAt, sysmenuaction.FieldUpdatedAt, sysmenuaction.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SysMenuAction", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SysMenuAction fields.
func (sma *SysMenuAction) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sysmenuaction.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				sma.ID = value.String
			}
		case sysmenuaction.FieldIsDel:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_del", values[i])
			} else if value.Valid {
				sma.IsDel = value.Bool
			}
		case sysmenuaction.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				sma.Sort = int32(value.Int64)
			}
		case sysmenuaction.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				sma.Status = int32(value.Int64)
			}
		case sysmenuaction.FieldMemo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field memo", values[i])
			} else if value.Valid {
				sma.Memo = value.String
			}
		case sysmenuaction.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sma.CreatedAt = value.Time
			}
		case sysmenuaction.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sma.UpdatedAt = value.Time
			}
		case sysmenuaction.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				sma.DeletedAt = new(time.Time)
				*sma.DeletedAt = value.Time
			}
		case sysmenuaction.FieldMenuID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field menu_id", values[i])
			} else if value.Valid {
				sma.MenuID = value.String
			}
		case sysmenuaction.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				sma.Code = value.String
			}
		case sysmenuaction.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				sma.Name = value.String
			}
		}
	}
	return nil
}

// QueryResources queries the "resources" edge of the SysMenuAction entity.
func (sma *SysMenuAction) QueryResources() *SysMenuActionResourceQuery {
	return (&SysMenuActionClient{config: sma.config}).QueryResources(sma)
}

// QueryMenu queries the "menu" edge of the SysMenuAction entity.
func (sma *SysMenuAction) QueryMenu() *SysMenuQuery {
	return (&SysMenuActionClient{config: sma.config}).QueryMenu(sma)
}

// Update returns a builder for updating this SysMenuAction.
// Note that you need to call SysMenuAction.Unwrap() before calling this method if this SysMenuAction
// was returned from a transaction, and the transaction was committed or rolled back.
func (sma *SysMenuAction) Update() *SysMenuActionUpdateOne {
	return (&SysMenuActionClient{config: sma.config}).UpdateOne(sma)
}

// Unwrap unwraps the SysMenuAction entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sma *SysMenuAction) Unwrap() *SysMenuAction {
	tx, ok := sma.config.driver.(*txDriver)
	if !ok {
		panic("ent: SysMenuAction is not a transactional entity")
	}
	sma.config.driver = tx.drv
	return sma
}

// String implements the fmt.Stringer.
func (sma *SysMenuAction) String() string {
	var builder strings.Builder
	builder.WriteString("SysMenuAction(")
	builder.WriteString(fmt.Sprintf("id=%v", sma.ID))
	builder.WriteString(", is_del=")
	builder.WriteString(fmt.Sprintf("%v", sma.IsDel))
	builder.WriteString(", sort=")
	builder.WriteString(fmt.Sprintf("%v", sma.Sort))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", sma.Status))
	builder.WriteString(", memo=")
	builder.WriteString(sma.Memo)
	builder.WriteString(", created_at=")
	builder.WriteString(sma.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(sma.UpdatedAt.Format(time.ANSIC))
	if v := sma.DeletedAt; v != nil {
		builder.WriteString(", deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", menu_id=")
	builder.WriteString(sma.MenuID)
	builder.WriteString(", code=")
	builder.WriteString(sma.Code)
	builder.WriteString(", name=")
	builder.WriteString(sma.Name)
	builder.WriteByte(')')
	return builder.String()
}

// SysMenuActions is a parsable slice of SysMenuAction.
type SysMenuActions []*SysMenuAction

func (sma SysMenuActions) config(cfg config) {
	for _i := range sma {
		sma[_i].config = cfg
	}
}
