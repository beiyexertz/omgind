package router

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/wanhello/omgind/internal/api/v1"
	api_v2 "github.com/wanhello/omgind/internal/api/v2"
	"github.com/wanhello/omgind/pkg/auth"
)

var _ IRouter = (*Router)(nil)

// RouterSet 注入router
var RouterSet = wire.NewSet(wire.Struct(new(Router), "*"), wire.Bind(new(IRouter), new(*Router)))

// IRouter 注册路由
type IRouter interface {
	Register(app *gin.Engine) error
	Prefixes() []string
}

// Router 路由管理器
type Router struct {
	Auth           auth.Auther
	CasbinEnforcer *casbin.SyncedEnforcer
	DemoAPI        *api_v1.Demo
	SignInAPI      *api_v1.SignIn
	MenuAPI        *api_v1.Menu
	RoleAPI        *api_v1.Role
	UserAPI        *api_v1.User
	DictAPI        *api_v1.Dict
	//DictItemAPI    *api_v1.DictItem

	DemoAPIV2      *api_v2.Demo

}

// Register 注册路由
func (a *Router) Register(app *gin.Engine) error {
	a.RegisterAPI(app)
	return nil
}

// Prefixes 路由前缀列表
func (a *Router) Prefixes() []string {
	return []string{
		"/api/",
	}
}
