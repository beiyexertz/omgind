// Code generated by entc, DO NOT EDIT.

package xxxdemo

import (
	"time"
)

const (
	// Label holds the string label denoting the xxxdemo type in the database.
	Label = "xxx_demo"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
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
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// Table holds the table name of the xxxdemo in the database.
	Table = "xxx_demos"
)

// Columns holds all SQL columns for xxxdemo fields.
var Columns = []string{
	FieldID,
	FieldMemo,
	FieldSort,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldCode,
	FieldName,
	FieldStatus,
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
	// CodeValidator is a validator for the "code" field. It is called by the builders before save.
	CodeValidator func(string) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)
