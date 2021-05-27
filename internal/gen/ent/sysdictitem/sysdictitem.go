// Code generated by entc, DO NOT EDIT.

package sysdictitem

import (
	"time"
)

const (
	// Label holds the string label denoting the sysdictitem type in the database.
	Label = "sys_dict_item"
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
	// FieldLabel holds the string denoting the label field in the database.
	FieldLabel = "label"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "val"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldDictID holds the string denoting the dict_id field in the database.
	FieldDictID = "dict_id"
	// Table holds the table name of the sysdictitem in the database.
	Table = "sys_dict_items"
)

// Columns holds all SQL columns for sysdictitem fields.
var Columns = []string{
	FieldID,
	FieldIsDel,
	FieldMemo,
	FieldSort,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldLabel,
	FieldValue,
	FieldStatus,
	FieldDictID,
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
	DefaultSort int
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// LabelValidator is a validator for the "label" field. It is called by the builders before save.
	LabelValidator func(string) error
	// DictIDValidator is a validator for the "dict_id" field. It is called by the builders before save.
	DictIDValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)
