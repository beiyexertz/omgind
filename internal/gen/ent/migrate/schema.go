// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// SysDictsColumns holds the columns for the "sys_dicts" table.
	SysDictsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 36},
		{Name: "is_del", Type: field.TypeBool, Default: false},
		{Name: "memo", Type: field.TypeString, Size: 1024, Default: ""},
		{Name: "sort", Type: field.TypeInt, Default: 9999},
		{Name: "crtd_at", Type: field.TypeTime},
		{Name: "uptd_at", Type: field.TypeTime},
		{Name: "dltd_at", Type: field.TypeTime, Nullable: true},
		{Name: "status", Type: field.TypeInt, Default: 1},
		{Name: "name_cn", Type: field.TypeString, Size: 128},
		{Name: "name_en", Type: field.TypeString, Size: 128},
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
			{
				Name:    "sysdict_status",
				Unique:  false,
				Columns: []*schema.Column{SysDictsColumns[7]},
			},
		},
	}
	// SysDictItemsColumns holds the columns for the "sys_dict_items" table.
	SysDictItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 36},
		{Name: "is_del", Type: field.TypeBool, Default: false},
		{Name: "memo", Type: field.TypeString, Size: 1024, Default: ""},
		{Name: "sort", Type: field.TypeInt, Default: 9999},
		{Name: "crtd_at", Type: field.TypeTime},
		{Name: "uptd_at", Type: field.TypeTime},
		{Name: "dltd_at", Type: field.TypeTime, Nullable: true},
		{Name: "label", Type: field.TypeString, Size: 128},
		{Name: "val", Type: field.TypeInt},
		{Name: "status", Type: field.TypeInt},
		{Name: "dict_id", Type: field.TypeString, Size: 36},
	}
	// SysDictItemsTable holds the schema information for the "sys_dict_items" table.
	SysDictItemsTable = &schema.Table{
		Name:        "sys_dict_items",
		Columns:     SysDictItemsColumns,
		PrimaryKey:  []*schema.Column{SysDictItemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
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
	// SysJwtBlocksColumns holds the columns for the "sys_jwt_blocks" table.
	SysJwtBlocksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 36},
		{Name: "is_del", Type: field.TypeBool, Default: false},
		{Name: "memo", Type: field.TypeString, Size: 1024, Default: ""},
		{Name: "crtd_at", Type: field.TypeTime},
		{Name: "uptd_at", Type: field.TypeTime},
		{Name: "dltd_at", Type: field.TypeTime, Nullable: true},
		{Name: "status", Type: field.TypeInt, Default: 1},
		{Name: "jwt", Type: field.TypeString, Size: 2147483647},
	}
	// SysJwtBlocksTable holds the schema information for the "sys_jwt_blocks" table.
	SysJwtBlocksTable = &schema.Table{
		Name:        "sys_jwt_blocks",
		Columns:     SysJwtBlocksColumns,
		PrimaryKey:  []*schema.Column{SysJwtBlocksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "sysjwtblock_id",
				Unique:  true,
				Columns: []*schema.Column{SysJwtBlocksColumns[0]},
			},
			{
				Name:    "sysjwtblock_is_del",
				Unique:  false,
				Columns: []*schema.Column{SysJwtBlocksColumns[1]},
			},
			{
				Name:    "sysjwtblock_crtd_at",
				Unique:  false,
				Columns: []*schema.Column{SysJwtBlocksColumns[3]},
			},
			{
				Name:    "sysjwtblock_uptd_at",
				Unique:  false,
				Columns: []*schema.Column{SysJwtBlocksColumns[4]},
			},
			{
				Name:    "sysjwtblock_dltd_at",
				Unique:  false,
				Columns: []*schema.Column{SysJwtBlocksColumns[5]},
			},
			{
				Name:    "sysjwtblock_status",
				Unique:  false,
				Columns: []*schema.Column{SysJwtBlocksColumns[6]},
			},
		},
	}
	// SysLoggingsColumns holds the columns for the "sys_loggings" table.
	SysLoggingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 36},
		{Name: "is_del", Type: field.TypeBool, Default: false},
		{Name: "memo", Type: field.TypeString, Size: 1024, Default: ""},
		{Name: "level", Type: field.TypeString, Size: 32},
		{Name: "trace_id", Type: field.TypeString, Size: 128},
		{Name: "user_id", Type: field.TypeString, Size: 128},
		{Name: "tag", Type: field.TypeString, Size: 128},
		{Name: "version", Type: field.TypeString, Size: 64},
		{Name: "message", Type: field.TypeString},
		{Name: "data", Type: field.TypeString},
		{Name: "error_stack", Type: field.TypeString},
		{Name: "crtd_at", Type: field.TypeTime},
	}
	// SysLoggingsTable holds the schema information for the "sys_loggings" table.
	SysLoggingsTable = &schema.Table{
		Name:        "sys_loggings",
		Columns:     SysLoggingsColumns,
		PrimaryKey:  []*schema.Column{SysLoggingsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "syslogging_id",
				Unique:  true,
				Columns: []*schema.Column{SysLoggingsColumns[0]},
			},
			{
				Name:    "syslogging_is_del",
				Unique:  false,
				Columns: []*schema.Column{SysLoggingsColumns[1]},
			},
			{
				Name:    "syslogging_level",
				Unique:  false,
				Columns: []*schema.Column{SysLoggingsColumns[3]},
			},
			{
				Name:    "syslogging_trace_id",
				Unique:  false,
				Columns: []*schema.Column{SysLoggingsColumns[4]},
			},
			{
				Name:    "syslogging_user_id",
				Unique:  false,
				Columns: []*schema.Column{SysLoggingsColumns[5]},
			},
			{
				Name:    "syslogging_tag",
				Unique:  false,
				Columns: []*schema.Column{SysLoggingsColumns[6]},
			},
			{
				Name:    "syslogging_crtd_at",
				Unique:  false,
				Columns: []*schema.Column{SysLoggingsColumns[11]},
			},
		},
	}
	// SysMenusColumns holds the columns for the "sys_menus" table.
	SysMenusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 36},
		{Name: "is_del", Type: field.TypeBool, Default: false},
		{Name: "memo", Type: field.TypeString, Size: 1024, Default: ""},
		{Name: "sort", Type: field.TypeInt, Default: 9999},
		{Name: "crtd_at", Type: field.TypeTime},
		{Name: "uptd_at", Type: field.TypeTime},
		{Name: "dltd_at", Type: field.TypeTime, Nullable: true},
		{Name: "status", Type: field.TypeInt, Default: 1},
		{Name: "name", Type: field.TypeString, Size: 64},
		{Name: "icon", Type: field.TypeString, Size: 256},
		{Name: "router", Type: field.TypeString, Size: 4096},
		{Name: "is_show", Type: field.TypeBool, Default: true},
		{Name: "pid", Type: field.TypeString, Nullable: true, Size: 36},
		{Name: "ppath", Type: field.TypeString, Nullable: true, Size: 160},
	}
	// SysMenusTable holds the schema information for the "sys_menus" table.
	SysMenusTable = &schema.Table{
		Name:        "sys_menus",
		Columns:     SysMenusColumns,
		PrimaryKey:  []*schema.Column{SysMenusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "sysmenu_id",
				Unique:  true,
				Columns: []*schema.Column{SysMenusColumns[0]},
			},
			{
				Name:    "sysmenu_is_del",
				Unique:  false,
				Columns: []*schema.Column{SysMenusColumns[1]},
			},
			{
				Name:    "sysmenu_sort",
				Unique:  false,
				Columns: []*schema.Column{SysMenusColumns[3]},
			},
			{
				Name:    "sysmenu_crtd_at",
				Unique:  false,
				Columns: []*schema.Column{SysMenusColumns[4]},
			},
			{
				Name:    "sysmenu_uptd_at",
				Unique:  false,
				Columns: []*schema.Column{SysMenusColumns[5]},
			},
			{
				Name:    "sysmenu_dltd_at",
				Unique:  false,
				Columns: []*schema.Column{SysMenusColumns[6]},
			},
			{
				Name:    "sysmenu_status",
				Unique:  false,
				Columns: []*schema.Column{SysMenusColumns[7]},
			},
			{
				Name:    "sysmenu_pid",
				Unique:  false,
				Columns: []*schema.Column{SysMenusColumns[12]},
			},
			{
				Name:    "sysmenu_pid_name",
				Unique:  true,
				Columns: []*schema.Column{SysMenusColumns[12], SysMenusColumns[8]},
			},
		},
	}
	// SysMenuActionsColumns holds the columns for the "sys_menu_actions" table.
	SysMenuActionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 36},
		{Name: "is_del", Type: field.TypeBool, Default: false},
		{Name: "sort", Type: field.TypeInt, Default: 9999},
		{Name: "status", Type: field.TypeInt, Default: 1},
		{Name: "memo", Type: field.TypeString, Size: 1024, Default: ""},
		{Name: "crtd_at", Type: field.TypeTime},
		{Name: "uptd_at", Type: field.TypeTime},
		{Name: "dltd_at", Type: field.TypeTime, Nullable: true},
		{Name: "menu_id", Type: field.TypeString, Size: 36},
		{Name: "code", Type: field.TypeString, Size: 128},
		{Name: "name", Type: field.TypeString, Size: 128},
	}
	// SysMenuActionsTable holds the schema information for the "sys_menu_actions" table.
	SysMenuActionsTable = &schema.Table{
		Name:        "sys_menu_actions",
		Columns:     SysMenuActionsColumns,
		PrimaryKey:  []*schema.Column{SysMenuActionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "sysmenuaction_id",
				Unique:  true,
				Columns: []*schema.Column{SysMenuActionsColumns[0]},
			},
			{
				Name:    "sysmenuaction_is_del",
				Unique:  false,
				Columns: []*schema.Column{SysMenuActionsColumns[1]},
			},
			{
				Name:    "sysmenuaction_sort",
				Unique:  false,
				Columns: []*schema.Column{SysMenuActionsColumns[2]},
			},
			{
				Name:    "sysmenuaction_status",
				Unique:  false,
				Columns: []*schema.Column{SysMenuActionsColumns[3]},
			},
			{
				Name:    "sysmenuaction_crtd_at",
				Unique:  false,
				Columns: []*schema.Column{SysMenuActionsColumns[5]},
			},
			{
				Name:    "sysmenuaction_uptd_at",
				Unique:  false,
				Columns: []*schema.Column{SysMenuActionsColumns[6]},
			},
			{
				Name:    "sysmenuaction_dltd_at",
				Unique:  false,
				Columns: []*schema.Column{SysMenuActionsColumns[7]},
			},
		},
	}
	// SysMenuActionResourcesColumns holds the columns for the "sys_menu_action_resources" table.
	SysMenuActionResourcesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 36},
		{Name: "is_del", Type: field.TypeBool, Default: false},
		{Name: "sort", Type: field.TypeInt, Default: 9999},
		{Name: "memo", Type: field.TypeString, Size: 1024, Default: ""},
		{Name: "crtd_at", Type: field.TypeTime},
		{Name: "uptd_at", Type: field.TypeTime},
		{Name: "dltd_at", Type: field.TypeTime, Nullable: true},
		{Name: "status", Type: field.TypeInt, Default: 1},
		{Name: "method", Type: field.TypeString, Size: 128},
		{Name: "path", Type: field.TypeString, Size: 256},
		{Name: "action_id", Type: field.TypeString, Size: 36},
	}
	// SysMenuActionResourcesTable holds the schema information for the "sys_menu_action_resources" table.
	SysMenuActionResourcesTable = &schema.Table{
		Name:        "sys_menu_action_resources",
		Columns:     SysMenuActionResourcesColumns,
		PrimaryKey:  []*schema.Column{SysMenuActionResourcesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "sysmenuactionresource_id",
				Unique:  true,
				Columns: []*schema.Column{SysMenuActionResourcesColumns[0]},
			},
			{
				Name:    "sysmenuactionresource_is_del",
				Unique:  false,
				Columns: []*schema.Column{SysMenuActionResourcesColumns[1]},
			},
			{
				Name:    "sysmenuactionresource_sort",
				Unique:  false,
				Columns: []*schema.Column{SysMenuActionResourcesColumns[2]},
			},
			{
				Name:    "sysmenuactionresource_crtd_at",
				Unique:  false,
				Columns: []*schema.Column{SysMenuActionResourcesColumns[4]},
			},
			{
				Name:    "sysmenuactionresource_uptd_at",
				Unique:  false,
				Columns: []*schema.Column{SysMenuActionResourcesColumns[5]},
			},
			{
				Name:    "sysmenuactionresource_dltd_at",
				Unique:  false,
				Columns: []*schema.Column{SysMenuActionResourcesColumns[6]},
			},
			{
				Name:    "sysmenuactionresource_status",
				Unique:  false,
				Columns: []*schema.Column{SysMenuActionResourcesColumns[7]},
			},
		},
	}
	// SysRolesColumns holds the columns for the "sys_roles" table.
	SysRolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 36},
		{Name: "is_del", Type: field.TypeBool, Default: false},
		{Name: "status", Type: field.TypeInt, Default: 1},
		{Name: "sort", Type: field.TypeInt, Default: 9999},
		{Name: "memo", Type: field.TypeString, Size: 1024, Default: ""},
		{Name: "crtd_at", Type: field.TypeTime},
		{Name: "uptd_at", Type: field.TypeTime},
		{Name: "dltd_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString, Size: 64},
	}
	// SysRolesTable holds the schema information for the "sys_roles" table.
	SysRolesTable = &schema.Table{
		Name:        "sys_roles",
		Columns:     SysRolesColumns,
		PrimaryKey:  []*schema.Column{SysRolesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "sysrole_id",
				Unique:  true,
				Columns: []*schema.Column{SysRolesColumns[0]},
			},
			{
				Name:    "sysrole_is_del",
				Unique:  false,
				Columns: []*schema.Column{SysRolesColumns[1]},
			},
			{
				Name:    "sysrole_status",
				Unique:  false,
				Columns: []*schema.Column{SysRolesColumns[2]},
			},
			{
				Name:    "sysrole_sort",
				Unique:  false,
				Columns: []*schema.Column{SysRolesColumns[3]},
			},
			{
				Name:    "sysrole_crtd_at",
				Unique:  false,
				Columns: []*schema.Column{SysRolesColumns[5]},
			},
			{
				Name:    "sysrole_uptd_at",
				Unique:  false,
				Columns: []*schema.Column{SysRolesColumns[6]},
			},
			{
				Name:    "sysrole_dltd_at",
				Unique:  false,
				Columns: []*schema.Column{SysRolesColumns[7]},
			},
		},
	}
	// SysRoleMenusColumns holds the columns for the "sys_role_menus" table.
	SysRoleMenusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 36},
		{Name: "is_del", Type: field.TypeBool, Default: false},
		{Name: "crtd_at", Type: field.TypeTime},
		{Name: "uptd_at", Type: field.TypeTime},
		{Name: "dltd_at", Type: field.TypeTime, Nullable: true},
		{Name: "role_id", Type: field.TypeString, Size: 36},
		{Name: "menu_id", Type: field.TypeString, Size: 36},
		{Name: "action_id", Type: field.TypeString, Nullable: true, Size: 36},
	}
	// SysRoleMenusTable holds the schema information for the "sys_role_menus" table.
	SysRoleMenusTable = &schema.Table{
		Name:        "sys_role_menus",
		Columns:     SysRoleMenusColumns,
		PrimaryKey:  []*schema.Column{SysRoleMenusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "sysrolemenu_id",
				Unique:  true,
				Columns: []*schema.Column{SysRoleMenusColumns[0]},
			},
			{
				Name:    "sysrolemenu_is_del",
				Unique:  false,
				Columns: []*schema.Column{SysRoleMenusColumns[1]},
			},
			{
				Name:    "sysrolemenu_crtd_at",
				Unique:  false,
				Columns: []*schema.Column{SysRoleMenusColumns[2]},
			},
			{
				Name:    "sysrolemenu_uptd_at",
				Unique:  false,
				Columns: []*schema.Column{SysRoleMenusColumns[3]},
			},
			{
				Name:    "sysrolemenu_dltd_at",
				Unique:  false,
				Columns: []*schema.Column{SysRoleMenusColumns[4]},
			},
		},
	}
	// SysUsersColumns holds the columns for the "sys_users" table.
	SysUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 36},
		{Name: "is_del", Type: field.TypeBool, Default: false},
		{Name: "sort", Type: field.TypeInt, Default: 9999},
		{Name: "crtd_at", Type: field.TypeTime},
		{Name: "uptd_at", Type: field.TypeTime},
		{Name: "dltd_at", Type: field.TypeTime, Nullable: true},
		{Name: "status", Type: field.TypeInt, Default: 1},
		{Name: "user_name", Type: field.TypeString, Size: 128},
		{Name: "real_name", Type: field.TypeString, Nullable: true, Size: 64},
		{Name: "first_name", Type: field.TypeString, Nullable: true, Size: 31},
		{Name: "last_name", Type: field.TypeString, Nullable: true, Size: 31},
		{Name: "passwd", Type: field.TypeString, Size: 256},
		{Name: "email", Type: field.TypeString, Size: 256},
		{Name: "phone", Type: field.TypeString, Size: 20},
		{Name: "salt", Type: field.TypeString},
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
				Name:    "sysuser_status",
				Unique:  false,
				Columns: []*schema.Column{SysUsersColumns[6]},
			},
			{
				Name:    "sysuser_user_name",
				Unique:  true,
				Columns: []*schema.Column{SysUsersColumns[7]},
			},
		},
	}
	// SysUserRolesColumns holds the columns for the "sys_user_roles" table.
	SysUserRolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 36},
		{Name: "is_del", Type: field.TypeBool, Default: false},
		{Name: "crtd_at", Type: field.TypeTime},
		{Name: "uptd_at", Type: field.TypeTime},
		{Name: "dltd_at", Type: field.TypeTime, Nullable: true},
		{Name: "user_id", Type: field.TypeString, Size: 36},
		{Name: "role_id", Type: field.TypeString, Size: 36},
	}
	// SysUserRolesTable holds the schema information for the "sys_user_roles" table.
	SysUserRolesTable = &schema.Table{
		Name:        "sys_user_roles",
		Columns:     SysUserRolesColumns,
		PrimaryKey:  []*schema.Column{SysUserRolesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "sysuserrole_id",
				Unique:  true,
				Columns: []*schema.Column{SysUserRolesColumns[0]},
			},
			{
				Name:    "sysuserrole_is_del",
				Unique:  false,
				Columns: []*schema.Column{SysUserRolesColumns[1]},
			},
			{
				Name:    "sysuserrole_crtd_at",
				Unique:  false,
				Columns: []*schema.Column{SysUserRolesColumns[2]},
			},
			{
				Name:    "sysuserrole_uptd_at",
				Unique:  false,
				Columns: []*schema.Column{SysUserRolesColumns[3]},
			},
			{
				Name:    "sysuserrole_dltd_at",
				Unique:  false,
				Columns: []*schema.Column{SysUserRolesColumns[4]},
			},
		},
	}
	// XxxDemosColumns holds the columns for the "xxx_demos" table.
	XxxDemosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 36},
		{Name: "is_del", Type: field.TypeBool, Default: false},
		{Name: "memo", Type: field.TypeString, Size: 1024, Default: ""},
		{Name: "sort", Type: field.TypeInt, Default: 9999},
		{Name: "crtd_at", Type: field.TypeTime},
		{Name: "uptd_at", Type: field.TypeTime},
		{Name: "dltd_at", Type: field.TypeTime, Nullable: true},
		{Name: "code", Type: field.TypeString, Size: 128},
		{Name: "name", Type: field.TypeString, Size: 128},
		{Name: "status", Type: field.TypeInt, Default: 1},
	}
	// XxxDemosTable holds the schema information for the "xxx_demos" table.
	XxxDemosTable = &schema.Table{
		Name:        "xxx_demos",
		Columns:     XxxDemosColumns,
		PrimaryKey:  []*schema.Column{XxxDemosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "xxxdemo_id",
				Unique:  true,
				Columns: []*schema.Column{XxxDemosColumns[0]},
			},
			{
				Name:    "xxxdemo_is_del",
				Unique:  false,
				Columns: []*schema.Column{XxxDemosColumns[1]},
			},
			{
				Name:    "xxxdemo_sort",
				Unique:  false,
				Columns: []*schema.Column{XxxDemosColumns[3]},
			},
			{
				Name:    "xxxdemo_crtd_at",
				Unique:  false,
				Columns: []*schema.Column{XxxDemosColumns[4]},
			},
			{
				Name:    "xxxdemo_uptd_at",
				Unique:  false,
				Columns: []*schema.Column{XxxDemosColumns[5]},
			},
			{
				Name:    "xxxdemo_dltd_at",
				Unique:  false,
				Columns: []*schema.Column{XxxDemosColumns[6]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		SysDictsTable,
		SysDictItemsTable,
		SysJwtBlocksTable,
		SysLoggingsTable,
		SysMenusTable,
		SysMenuActionsTable,
		SysMenuActionResourcesTable,
		SysRolesTable,
		SysRoleMenusTable,
		SysUsersTable,
		SysUserRolesTable,
		XxxDemosTable,
	}
)

func init() {
}
