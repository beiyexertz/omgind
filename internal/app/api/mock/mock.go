package mock

import "github.com/google/wire"

// MockSet 注入mock
var MockSet = wire.NewSet(
	DemoSet,
	SignInSet,
	MenuSet,
	RoleSet,
	UserSet,
	DictSet,
	DictItemSet,
)
