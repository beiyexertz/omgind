// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package app

import (
	"github.com/wanhello/omgind/internal/api/v2"
	"github.com/wanhello/omgind/internal/app/service"
	"github.com/wanhello/omgind/internal/schema/repo"

	// "github.com/wanhello/omgind/internal/app/api_v2/mock"
	"github.com/google/wire"
	"github.com/wanhello/omgind/internal/app/module/adapter"
	"github.com/wanhello/omgind/internal/router"

)

// BuildInjector 生成注入器
func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		// mock.MockSet,
		InitEntClient,
		InitRedisCli,
		InitVcode,
		//InitInfluxDB,
		//InitRabbitMQ,
		repo.RepoSet,
		InitAuth,
		InitCasbin,
		InitGinEngine,
		service.ServiceSet,
		api_v2.APIV2Set,
		router.RouterSet,
		adapter.CasbinAdapterSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
