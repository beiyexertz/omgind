// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package app

import (
	"github.com/wanhello/omgind/internal/api/v2"
	service_ent "github.com/wanhello/omgind/internal/app/service.ent"
	"github.com/wanhello/omgind/internal/schema/repo_ent"

	// "github.com/wanhello/omgind/internal/app/api_v1/mock"
	"github.com/google/wire"
	"github.com/wanhello/omgind/internal/app/module/adapter"
	"github.com/wanhello/omgind/internal/router"

	"github.com/wanhello/omgind/internal/app/model/gormx/repo"
)

// BuildInjector 生成注入器
func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		// mock.MockSet,
		InitGormDB,
		InitEntClient,
		InitRedisCli,
		InitVcode,
		//InitInfluxDB,
		//InitRabbitMQ,
		repo.RepoSet,
		repo_ent.RepoSet,
		InitAuth,
		InitCasbin,
		InitGinEngine,
		//service.ServiceSet,
		service_ent.ServiceEntSet,
		//api_v1.APIV1Set,
		api_v2.APIV2Set,
		router.RouterSet,
		adapter.CasbinAdapterSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
