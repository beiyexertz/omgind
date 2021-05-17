package repo_ent

import (
	"github.com/google/wire"
)

var RepoSet = wire.NewSet(
	DemoSet,
	DictSet,
	DictItemSet,
)
