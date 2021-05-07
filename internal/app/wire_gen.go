// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package app

import (
	"github.com/wanhello/omgind/internal/app/api"
	"github.com/wanhello/omgind/internal/app/model/gormx/repo"
	"github.com/wanhello/omgind/internal/app/module/adapter"
	"github.com/wanhello/omgind/internal/app/router"
	"github.com/wanhello/omgind/internal/app/service"
)

import (
	_ "github.com/wanhello/omgind/internal/app/swagger"
)

// Injectors from wire.go:

// BuildInjector 生成注入器
func BuildInjector() (*Injector, func(), error) {
	auther, cleanup, err := InitAuth()
	if err != nil {
		return nil, nil, err
	}
	db, cleanup2, err := InitGormDB()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	role := &repo.Role{
		DB: db,
	}
	roleMenu := &repo.RoleMenu{
		DB: db,
	}
	menuActionResource := &repo.MenuActionResource{
		DB: db,
	}
	user := &repo.User{
		DB: db,
	}
	userRole := &repo.UserRole{
		DB: db,
	}
	casbinAdapter := &adapter.CasbinAdapter{
		RoleModel:         role,
		RoleMenuModel:     roleMenu,
		MenuResourceModel: menuActionResource,
		UserModel:         user,
		UserRoleModel:     userRole,
	}
	syncedEnforcer, cleanup3, err := InitCasbin(casbinAdapter)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	demo := &repo.Demo{
		DB: db,
	}
	serviceDemo := &service.Demo{
		DemoModel: demo,
	}
	apiDemo := &api.Demo{
		DemoSrv: serviceDemo,
	}
	menu := &repo.Menu{
		DB: db,
	}
	menuAction := &repo.MenuAction{
		DB: db,
	}
	cmdable, cleanup4, err := InitRedisCli()
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	vcode := InitVcode(cmdable)
	signIn := &service.SignIn{
		Auth:            auther,
		UserModel:       user,
		UserRoleModel:   userRole,
		RoleModel:       role,
		RoleMenuModel:   roleMenu,
		MenuModel:       menu,
		MenuActionModel: menuAction,
		Vcode:           vcode,
	}
	apiSignIn := &api.SignIn{
		SigninSrv: signIn,
		Vcode:     vcode,
	}
	trans := &repo.Trans{
		DB: db,
	}
	serviceMenu := &service.Menu{
		TransModel:              trans,
		MenuModel:               menu,
		MenuActionModel:         menuAction,
		MenuActionResourceModel: menuActionResource,
	}
	apiMenu := &api.Menu{
		MenuSrv: serviceMenu,
	}
	serviceRole := &service.Role{
		Enforcer:      syncedEnforcer,
		TransModel:    trans,
		RoleModel:     role,
		RoleMenuModel: roleMenu,
		UserModel:     user,
	}
	apiRole := &api.Role{
		RoleSrv: serviceRole,
	}
	serviceUser := &service.User{
		Enforcer:      syncedEnforcer,
		TransModel:    trans,
		UserModel:     user,
		UserRoleModel: userRole,
		RoleModel:     role,
	}
	apiUser := &api.User{
		UserSrv: serviceUser,
	}
	routerRouter := &router.Router{
		Auth:           auther,
		CasbinEnforcer: syncedEnforcer,
		DemoAPI:        apiDemo,
		SignInAPI:      apiSignIn,
		MenuAPI:        apiMenu,
		RoleAPI:        apiRole,
		UserAPI:        apiUser,
	}
	engine := InitGinEngine(routerRouter)
	client, cleanup5, err := InitEntClient()
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	injector := &Injector{
		Engine:         engine,
		Auth:           auther,
		CasbinEnforcer: syncedEnforcer,
		MenuSrv:        serviceMenu,
		RedisCli:       cmdable,
		EntCli:         client,
	}
	return injector, func() {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
