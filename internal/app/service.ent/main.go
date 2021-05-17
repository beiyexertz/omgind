package service_ent

import "github.com/google/wire"

// ServiceSet 注入
var ServiceEntSet = wire.NewSet(
	DemoSet,
	DictSet,
)
