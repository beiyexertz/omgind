package api_v1

import "github.com/google/wire"

// APISet 注入api
var APIV1Set = wire.NewSet(
	DemoSet,
	SignInSet,
	MenuSet,
	RoleSet,
	UserSet,
)
