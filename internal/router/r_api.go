package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wanhello/omgind/internal/app/middleware"
)

// RegisterAPI register api group router
func (a *Router) RegisterAPI(app *gin.Engine) {

	g := app.Group("/api")

	g.Use(middleware.UserAuthMiddleware(a.Auth,
		middleware.AllowPathPrefixSkipper("/api/v1/pub/signin"),
	))

	g.Use(middleware.CasbinMiddleware(a.CasbinEnforcer,
		middleware.AllowPathPrefixSkipper("/api/v1/pub"),
	))

	g.Use(middleware.RateLimiterMiddleware())

	v1 := g.Group("/v1")
	{
		pub := v1.Group("/pub")
		{
			gSignIn := pub.Group("/signin")
			{
				gSignIn.GET("captchaid", a.SignInAPI.GetCaptcha)
				gSignIn.GET("captcha", a.SignInAPI.ResCaptcha)
				gSignIn.POST("", a.SignInAPI.SignIn)
				gSignIn.POST("exit", a.SignInAPI.SignOut)
			}

			gCurrent := pub.Group("current")
			{
				gCurrent.PUT("password", a.SignInAPI.UpdatePassword)
				gCurrent.GET("user", a.SignInAPI.GetUserInfo)
				gCurrent.GET("menutree", a.SignInAPI.QueryUserMenuTree)
			}
			pub.POST("/refresh-token", a.SignInAPI.RefreshToken)
		}

		a.initMenuRouterV1(v1, a.MenuAPI, "menus")
		v1.GET("/menus.tree", a.MenuAPI.QueryTree)

		a.initRoleRouterV1(v1, a.RoleAPI, "roles")
		v1.GET("/roles.select", a.RoleAPI.QuerySelect)

		a.initUserRouterV1(v1, a.UserAPI, "users")

		//gDictItem := v1.Group("dict-items")
		//{
		//	gDictItem.GET("", a.DictItemAPI.Query)
		//	gDictItem.GET(":id", a.DictItemAPI.Get)
		//	gDictItem.POST("", a.DictItemAPI.Create)
		//	gDictItem.PUT(":id", a.DictItemAPI.Update)
		//	gDictItem.DELETE(":id", a.DictItemAPI.Delete)
		//}

	}

	v2 := g.Group("/v2")
	{
		a.initDictRouterV2(v2, a.DictApiV2, "dicts")
		a.initDemoRouterV2(v2, a.DemoAPIV2, "demos")

	}

}
