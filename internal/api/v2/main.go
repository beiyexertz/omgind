package api_v2

import "github.com/google/wire"

// APISet 注入api
var APIV2Set = wire.NewSet(
	UserSet,
	DemoSet,
	DictSet,
)
