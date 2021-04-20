package service

import "github.com/google/wire"

// ServiceSet 注入
var ServiceSet = wire.NewSet(
	DemoSet,
	SignInSet,
	MenuSet,
	RoleSet,
	UserSet,
)
