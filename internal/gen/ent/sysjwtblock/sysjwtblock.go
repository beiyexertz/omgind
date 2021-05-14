// Code generated by entc, DO NOT EDIT.

package sysjwtblock

import (
	"time"

	"github.com/wanhello/omgind/pkg/helper/pulid"
)

const (
	// Label holds the string label denoting the sysjwtblock type in the database.
	Label = "sys_jwt_block"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
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
	// FieldJwt holds the string denoting the jwt field in the database.
	FieldJwt = "jwt"
	// Table holds the table name of the sysjwtblock in the database.
	Table = "sys_jwt_blocks"
)

// Columns holds all SQL columns for sysjwtblock fields.
var Columns = []string{
	FieldID,
	FieldMemo,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldStatus,
	FieldJwt,
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int32
	// JwtValidator is a validator for the "jwt" field. It is called by the builders before save.
	JwtValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() pulid.ID
)
