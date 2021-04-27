// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package app

import (
	"github.com/wanhello/omgind/internal/app/api"
	// "github.com/wanhello/omgind/internal/app/api/mock"
	"github.com/google/wire"
	"github.com/wanhello/omgind/internal/app/module/adapter"
	"github.com/wanhello/omgind/internal/app/router"
	"github.com/wanhello/omgind/internal/app/service"

	"github.com/wanhello/omgind/internal/app/model/gormx/repo"
)

// BuildInjector 生成注入器
func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		// mock.MockSet,
		InitGormDB,
		InitRedisCli,
		InitVcode,
		repo.RepoSet,
		InitAuth,
		InitCasbin,
		InitGinEngine,
		service.ServiceSet,
		api.APISet,
		router.RouterSet,
		adapter.CasbinAdapterSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
