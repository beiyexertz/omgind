// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// SysCasbinRulesColumns holds the columns for the "sys_casbin_rules" table.
	SysCasbinRulesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 36},
		{Name: "is_del", Type: field.TypeBool, Default: false},
		{Name: "sort", Type: field.TypeInt32, Default: 9999},
		{Name: "crtd_at", Type: field.TypeTime},
		{Name: "uptd_at", Type: field.TypeTime},
		{Name: "dltd_at", Type: field.TypeTime, Nullable: true},
		{Name: "p_type", Type: field.TypeString, Nullable: true, Size: 128},
		{Name: "v0", Type: field.TypeString, Nullable: true, Size: 128},
		{Name: "v1", Type: field.TypeString, Nullable: true, Size: 128},
		{Name: "v2", Type: field.TypeString, Nullable: true, Size: 128},
		{Name: "v3", Type: field.TypeString, Nullable: true, Size: 128},
		{Name: "v4", Type: field.TypeString, Nullable: true, Size: 128},
		{Name: "v5", Type: field.TypeString, Nullable: true, Size: 128},
	}
	// SysCasbinRulesTable holds the schema information for the "sys_casbin_rules" table.
	SysCasbinRulesTable = &schema.Table{
		Name:        "sys_casbin_rules",
		Columns:     SysCasbinRulesColumns,
		PrimaryKey:  []*schema.Column{SysCasbinRulesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "syscasbinrule_id",
				Unique:  true,
				Columns: []*schema.Column{SysCasbinRulesColumns[0]},
			},
			{
				Name:    "syscasbinrule_is_del",
				Unique:  false,
				Columns: []*schema.Column{SysCasbinRulesColumns[1]},
			},
			{
				Name:    "syscasbinrule_sort",
				Unique:  false,
				Columns: []*schema.Column{SysCasbinRulesColumns[2]},
			},
			{
				Name:    "syscasbinrule_crtd_at",
				Unique:  false,
				Columns: []*schema.Column{SysCasbinRulesColumns[3]},
			},
			{
				Name:    "syscasbinrule_uptd_at",
				Unique:  false,
				Columns: []*schema.Column{SysCasbinRulesColumns[4]},
			},
			{
				Name:    "syscasbinrule_dltd_at",
				Unique:  false,
				Columns: []*schema.Column{SysCasbinRulesColumns[5]},
			},
		},
	}
	// SysDictsColumns holds the columns for the "sys_dicts" table.
	SysDictsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 36},
		{Name: "is_del", Type: field.TypeBool, Default: false},
		{Name: "memo", Type: field.TypeString, Size: 1024, Default: ""},
		{Name: "sort", Type: field.TypeInt32, Default: 9999},
		{Name: "crtd_at", Type: field.TypeTime},
		{Name: "uptd_at", Type: field.TypeTime},
		{Name: "dltd_at", Type: field.TypeTime, Nullable: true},
		{Name: "name_cn", Type: field.TypeString, Size: 128},
		{Name: "name_en", Type: field.TypeString, Size: 128},
		{Name: "status", Type: field.TypeBool, Default: true},
	}
	// SysDictsTable holds the schema information for the "sys_dicts" table.
	SysDictsTable = &schema.Table{
		Name:        "sys_dicts",
		Columns:     SysDictsColumns,
		PrimaryKey:  []*schema.Column{SysDictsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "sysdict_id",
				Unique:  true,
				Columns: []*schema.Column{SysDictsColumns[0]},
			},
			{
				Name:    "sysdict_is_del",
				Unique:  false,
				Columns: []*schema.Column{SysDictsColumns[1]},
			},
			{
				Name:    "sysdict_sort",
				Unique:  false,
				Columns: []*schema.Column{SysDictsColumns[3]},
			},
			{
				Name:    "sysdict_crtd_at",
				Unique:  false,
				Columns: []*schema.Column{SysDictsColumns[4]},
			},
			{
				Name:    "sysdict_uptd_at",
				Unique:  false,
				Columns: []*schema.Column{SysDictsColumns[5]},
			},
			{
				Name:    "sysdict_dltd_at",
				Unique:  false,
				Columns: []*schema.Column{SysDictsColumns[6]},
			},
		},
	}
	// SysDictItemsColumns holds the columns for the "sys_dict_items" table.
	SysDictItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 36},
		{Name: "is_del", Type: field.TypeBool, Default: false},
		{Name: "memo", Type: field.TypeString, Size: 1024, Default: ""},
		{Name: "sort", Type: field.TypeInt32, Default: 9999},
		{Name: "crtd_at", Type: field.TypeTime},
		{Name: "uptd_at", Type: field.TypeTime},
		{Name: "dltd_at", Type: field.TypeTime, Nullable: true},
		{Name: "label", Type: field.TypeString, Size: 128},
		{Name: "val", Type: field.TypeInt},
		{Name: "status", Type: field.TypeBool},
		{Name: "sys_dict_sys_dict_items", Type: field.TypeString, Nullable: true, Size: 36},
	}
	// SysDictItemsTable holds the schema information for the "sys_dict_items" table.
	SysDictItemsTable = &schema.Table{
		Name:       "sys_dict_items",
		Columns:    SysDictItemsColumns,
		PrimaryKey: []*schema.Column{SysDictItemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_dict_items_sys_dicts_SysDictItems",
				Columns:    []*schema.Column{SysDictItemsColumns[10]},
				RefColumns: []*schema.Column{SysDictsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "sysdictitem_id",
				Unique:  true,
				Columns: []*schema.Column{SysDictItemsColumns[0]},
			},
			{
				Name:    "sysdictitem_is_del",
				Unique:  false,
				Columns: []*schema.Column{SysDictItemsColumns[1]},
			},
			{
				Name:    "sysdictitem_sort",
				Unique:  false,
				Columns: []*schema.Column{SysDictItemsColumns[3]},
			},
			{
				Name:    "sysdictitem_crtd_at",
				Unique:  false,
				Columns: []*schema.Column{SysDictItemsColumns[4]},
			},
			{
				Name:    "sysdictitem_uptd_at",
				Unique:  false,
				Columns: []*schema.Column{SysDictItemsColumns[5]},
			},
			{
				Name:    "sysdictitem_dltd_at",
				Unique:  false,
				Columns: []*schema.Column{SysDictItemsColumns[6]},
			},
		},
	}
	// SysMenusColumns holds the columns for the "sys_menus" table.
	SysMenusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// SysMenusTable holds the schema information for the "sys_menus" table.
	SysMenusTable = &schema.Table{
		Name:        "sys_menus",
		Columns:     SysMenusColumns,
		PrimaryKey:  []*schema.Column{SysMenusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// SysMenuActionsColumns holds the columns for the "sys_menu_actions" table.
	SysMenuActionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// SysMenuActionsTable holds the schema information for the "sys_menu_actions" table.
	SysMenuActionsTable = &schema.Table{
		Name:        "sys_menu_actions",
		Columns:     SysMenuActionsColumns,
		PrimaryKey:  []*schema.Column{SysMenuActionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// SysMenuActionResourcesColumns holds the columns for the "sys_menu_action_resources" table.
	SysMenuActionResourcesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// SysMenuActionResourcesTable holds the schema information for the "sys_menu_action_resources" table.
	SysMenuActionResourcesTable = &schema.Table{
		Name:        "sys_menu_action_resources",
		Columns:     SysMenuActionResourcesColumns,
		PrimaryKey:  []*schema.Column{SysMenuActionResourcesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// SysRolesColumns holds the columns for the "sys_roles" table.
	SysRolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// SysRolesTable holds the schema information for the "sys_roles" table.
	SysRolesTable = &schema.Table{
		Name:        "sys_roles",
		Columns:     SysRolesColumns,
		PrimaryKey:  []*schema.Column{SysRolesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// SysRoleMenusColumns holds the columns for the "sys_role_menus" table.
	SysRoleMenusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// SysRoleMenusTable holds the schema information for the "sys_role_menus" table.
	SysRoleMenusTable = &schema.Table{
		Name:        "sys_role_menus",
		Columns:     SysRoleMenusColumns,
		PrimaryKey:  []*schema.Column{SysRoleMenusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// SysUsersColumns holds the columns for the "sys_users" table.
	SysUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 36},
		{Name: "is_del", Type: field.TypeBool, Default: false},
		{Name: "sort", Type: field.TypeInt32, Default: 9999},
		{Name: "crtd_at", Type: field.TypeTime},
		{Name: "uptd_at", Type: field.TypeTime},
		{Name: "dltd_at", Type: field.TypeTime, Nullable: true},
		{Name: "user_name", Type: field.TypeString, Size: 128},
		{Name: "real_name", Type: field.TypeString, Size: 64},
		{Name: "first_name", Type: field.TypeString, Nullable: true, Size: 31},
		{Name: "last_name", Type: field.TypeString, Nullable: true, Size: 31},
		{Name: "passwd", Type: field.TypeString, Size: 256},
		{Name: "email", Type: field.TypeString, Size: 256},
		{Name: "phone", Type: field.TypeString, Size: 20},
	}
	// SysUsersTable holds the schema information for the "sys_users" table.
	SysUsersTable = &schema.Table{
		Name:        "sys_users",
		Columns:     SysUsersColumns,
		PrimaryKey:  []*schema.Column{SysUsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "sysuser_id",
				Unique:  true,
				Columns: []*schema.Column{SysUsersColumns[0]},
			},
			{
				Name:    "sysuser_is_del",
				Unique:  false,
				Columns: []*schema.Column{SysUsersColumns[1]},
			},
			{
				Name:    "sysuser_sort",
				Unique:  false,
				Columns: []*schema.Column{SysUsersColumns[2]},
			},
			{
				Name:    "sysuser_crtd_at",
				Unique:  false,
				Columns: []*schema.Column{SysUsersColumns[3]},
			},
			{
				Name:    "sysuser_uptd_at",
				Unique:  false,
				Columns: []*schema.Column{SysUsersColumns[4]},
			},
			{
				Name:    "sysuser_dltd_at",
				Unique:  false,
				Columns: []*schema.Column{SysUsersColumns[5]},
			},
			{
				Name:    "sysuser_user_name",
				Unique:  true,
				Columns: []*schema.Column{SysUsersColumns[6]},
			},
		},
	}
	// SysUserRolesColumns holds the columns for the "sys_user_roles" table.
	SysUserRolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// SysUserRolesTable holds the schema information for the "sys_user_roles" table.
	SysUserRolesTable = &schema.Table{
		Name:        "sys_user_roles",
		Columns:     SysUserRolesColumns,
		PrimaryKey:  []*schema.Column{SysUserRolesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		SysCasbinRulesTable,
		SysDictsTable,
		SysDictItemsTable,
		SysMenusTable,
		SysMenuActionsTable,
		SysMenuActionResourcesTable,
		SysRolesTable,
		SysRoleMenusTable,
		SysUsersTable,
		SysUserRolesTable,
	}
)

func init() {
	SysDictItemsTable.ForeignKeys[0].RefTable = SysDictsTable
}
