package entity

import "entgo.io/ent"

// SysRole holds the schema definition for the SysRole entity.
type SysRole struct {
	ent.Schema
}

// Fields of the SysRole.
func (SysRole) Fields() []ent.Field {
	return nil
}

// Edges of the SysRole.
func (SysRole) Edges() []ent.Edge {
	return nil
}
