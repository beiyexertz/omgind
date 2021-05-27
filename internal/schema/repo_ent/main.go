package repo_ent

import (
	"github.com/google/wire"
)

var RepoSet = wire.NewSet(
	UserSet,
	RoleSet,
	UserRoleSet,
	MenuSet,

	DemoSet,
	DictSet,
	DictItemSet,
)
