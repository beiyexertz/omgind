// Code generated by entc, DO NOT EDIT.

package sysmenuactionresource

import (
	"time"
)

const (
	// Label holds the string label denoting the sysmenuactionresource type in the database.
	Label = "sys_menu_action_resource"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIsDel holds the string denoting the is_del field in the database.
	FieldIsDel = "is_del"
	// FieldSort holds the string denoting the sort field in the database.
	FieldSort = "sort"
	// FieldMemo holds the string denoting the memo field in the database.
	FieldMemo = "memo"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "crtd_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "uptd_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "dltd_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldMethod holds the string denoting the method field in the database.
	FieldMethod = "method"
	// FieldPath holds the string denoting the path field in the database.
	FieldPath = "path"
	// FieldActionID holds the string denoting the action_id field in the database.
	FieldActionID = "action_id"
	// Table holds the table name of the sysmenuactionresource in the database.
	Table = "sys_menu_action_resources"
)

// Columns holds all SQL columns for sysmenuactionresource fields.
var Columns = []string{
	FieldID,
	FieldIsDel,
	FieldSort,
	FieldMemo,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldStatus,
	FieldMethod,
	FieldPath,
	FieldActionID,
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
	// DefaultSort holds the default value on creation for the "sort" field.
	DefaultSort int32
	// DefaultMemo holds the default value on creation for the "memo" field.
	DefaultMemo string
	// MemoValidator is a validator for the "memo" field. It is called by the builders before save.
	MemoValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int32
	// MethodValidator is a validator for the "method" field. It is called by the builders before save.
	MethodValidator func(string) error
	// PathValidator is a validator for the "path" field. It is called by the builders before save.
	PathValidator func(string) error
	// ActionIDValidator is a validator for the "action_id" field. It is called by the builders before save.
	ActionIDValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)
