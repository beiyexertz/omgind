// Code generated by entc, DO NOT EDIT.

package sysmenu

import (
	"time"
)

const (
	// Label holds the string label denoting the sysmenu type in the database.
	Label = "sys_menu"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIsDel holds the string denoting the is_del field in the database.
	FieldIsDel = "is_del"
	// FieldMemo holds the string denoting the memo field in the database.
	FieldMemo = "memo"
	// FieldSort holds the string denoting the sort field in the database.
	FieldSort = "sort"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "crtd_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "uptd_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "dltd_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldIcon holds the string denoting the icon field in the database.
	FieldIcon = "icon"
	// FieldRouter holds the string denoting the router field in the database.
	FieldRouter = "router"
	// FieldIsShow holds the string denoting the is_show field in the database.
	FieldIsShow = "is_show"
	// FieldParentID holds the string denoting the parent_id field in the database.
	FieldParentID = "pid"
	// FieldParentPath holds the string denoting the parent_path field in the database.
	FieldParentPath = "ppath"
	// Table holds the table name of the sysmenu in the database.
	Table = "sys_menus"
)

// Columns holds all SQL columns for sysmenu fields.
var Columns = []string{
	FieldID,
	FieldIsDel,
	FieldMemo,
	FieldSort,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldStatus,
	FieldName,
	FieldIcon,
	FieldRouter,
	FieldIsShow,
	FieldParentID,
	FieldParentPath,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultIsDel holds the default value on creation for the "is_del" field.
	DefaultIsDel bool
	// DefaultMemo holds the default value on creation for the "memo" field.
	DefaultMemo string
	// MemoValidator is a validator for the "memo" field. It is called by the builders before save.
	MemoValidator func(string) error
	// DefaultSort holds the default value on creation for the "sort" field.
	DefaultSort int32
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int16
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// IconValidator is a validator for the "icon" field. It is called by the builders before save.
	IconValidator func(string) error
	// RouterValidator is a validator for the "router" field. It is called by the builders before save.
	RouterValidator func(string) error
	// DefaultIsShow holds the default value on creation for the "is_show" field.
	DefaultIsShow bool
	// ParentIDValidator is a validator for the "parent_id" field. It is called by the builders before save.
	ParentIDValidator func(string) error
	// ParentPathValidator is a validator for the "parent_path" field. It is called by the builders before save.
	ParentPathValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)
