// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/wanhello/omgind/internal/gen/ent/sysuser"
)

// SysUser is the model entity for the SysUser schema.
type SysUser struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
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
	// UserName holds the value of the "user_name" field.
	// 用户名
	UserName string `json:"user_name,omitempty"`
	// RealName holds the value of the "real_name" field.
	RealName *string `json:"real_name,omitempty"`
	// FirstName holds the value of the "first_name" field.
	// 名
	FirstName *string `json:"first_name,omitempty"`
	// LastName holds the value of the "last_name" field.
	// 姓
	LastName *string `json:"last_name,omitempty"`
	// Password holds the value of the "Password" field.
	// 密码
	Password string `json:"-"`
	// Email holds the value of the "Email" field.
	// 电子邮箱
	Email string `json:"email,omitempty"`
	// Phone holds the value of the "Phone" field.
	// 电话号码
	Phone string `json:"phone,omitempty"`
	// Salt holds the value of the "salt" field.
	// 盐
	Salt string `json:"salt,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SysUser) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case sysuser.FieldSort, sysuser.FieldStatus:
			values[i] = new(sql.NullInt64)
		case sysuser.FieldID, sysuser.FieldUserName, sysuser.FieldRealName, sysuser.FieldFirstName, sysuser.FieldLastName, sysuser.FieldPassword, sysuser.FieldEmail, sysuser.FieldPhone, sysuser.FieldSalt:
			values[i] = new(sql.NullString)
		case sysuser.FieldCreatedAt, sysuser.FieldUpdatedAt, sysuser.FieldDeletedAt:
			values[i] = new(sql.NullTime)
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
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				su.ID = value.String
			}
		case sysuser.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				su.Sort = int32(value.Int64)
			}
		case sysuser.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				su.CreatedAt = value.Time
			}
		case sysuser.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				su.UpdatedAt = value.Time
			}
		case sysuser.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				su.DeletedAt = new(time.Time)
				*su.DeletedAt = value.Time
			}
		case sysuser.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				su.Status = int32(value.Int64)
			}
		case sysuser.FieldUserName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_name", values[i])
			} else if value.Valid {
				su.UserName = value.String
			}
		case sysuser.FieldRealName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field real_name", values[i])
			} else if value.Valid {
				su.RealName = new(string)
				*su.RealName = value.String
			}
		case sysuser.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field first_name", values[i])
			} else if value.Valid {
				su.FirstName = new(string)
				*su.FirstName = value.String
			}
		case sysuser.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				su.LastName = new(string)
				*su.LastName = value.String
			}
		case sysuser.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Password", values[i])
			} else if value.Valid {
				su.Password = value.String
			}
		case sysuser.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Email", values[i])
			} else if value.Valid {
				su.Email = value.String
			}
		case sysuser.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Phone", values[i])
			} else if value.Valid {
				su.Phone = value.String
			}
		case sysuser.FieldSalt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field salt", values[i])
			} else if value.Valid {
				su.Salt = value.String
			}
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
		panic("ent: SysUser is not a transactional entity")
	}
	su.config.driver = tx.drv
	return su
}

// String implements the fmt.Stringer.
func (su *SysUser) String() string {
	var builder strings.Builder
	builder.WriteString("SysUser(")
	builder.WriteString(fmt.Sprintf("id=%v", su.ID))
	builder.WriteString(", sort=")
	builder.WriteString(fmt.Sprintf("%v", su.Sort))
	builder.WriteString(", created_at=")
	builder.WriteString(su.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(su.UpdatedAt.Format(time.ANSIC))
	if v := su.DeletedAt; v != nil {
		builder.WriteString(", deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", su.Status))
	builder.WriteString(", user_name=")
	builder.WriteString(su.UserName)
	if v := su.RealName; v != nil {
		builder.WriteString(", real_name=")
		builder.WriteString(*v)
	}
	if v := su.FirstName; v != nil {
		builder.WriteString(", first_name=")
		builder.WriteString(*v)
	}
	if v := su.LastName; v != nil {
		builder.WriteString(", last_name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", Password=<sensitive>")
	builder.WriteString(", Email=")
	builder.WriteString(su.Email)
	builder.WriteString(", Phone=")
	builder.WriteString(su.Phone)
	builder.WriteString(", salt=")
	builder.WriteString(su.Salt)
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
