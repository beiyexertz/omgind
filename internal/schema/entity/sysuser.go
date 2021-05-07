package entity

import "entgo.io/ent"

// SysUser holds the schema definition for the SysUser entity.
type SysUser struct {
	ent.Schema
}

// Fields of the SysUser.
func (SysUser) Fields() []ent.Field {
	return nil
}

// Edges of the SysUser.
func (SysUser) Edges() []ent.Edge {
	return nil
}
