package repo_ent

import (
	"github.com/google/wire"
)

var RepoSet = wire.NewSet(
	UserSet,
	RoleSet,
	UserRoleSet,
	DemoSet,
	DictSet,
	DictItemSet,
)
