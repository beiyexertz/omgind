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
			gSignIn := pub.Group("signin")
			{
				gSignIn.GET("captchaid", a.SignInAPI.GetCaptcha)
				gSignIn.GET("captcha", a.SignInAPI.ResCaptcha)
				gSignIn.POST("", a.SignInAPI.SignIn)
				gSignIn.POST("exit", a.SignInAPI.Logout)
			}

			gCurrent := pub.Group("current")
			{
				gCurrent.PUT("password", a.SignInAPI.UpdatePassword)
				gCurrent.GET("user", a.SignInAPI.GetUserInfo)
				gCurrent.GET("menutree", a.SignInAPI.QueryUserMenuTree)
			}
			pub.POST("/refresh-token", a.SignInAPI.RefreshToken)
		}

		gDemo := v1.Group("demos")
		{
			gDemo.GET("", a.DemoAPI.Query)
			gDemo.GET(":id", a.DemoAPI.Get)
			gDemo.POST("", a.DemoAPI.Create)
			gDemo.PUT(":id", a.DemoAPI.Update)
			gDemo.DELETE(":id", a.DemoAPI.Delete)
			gDemo.PATCH(":id/enable", a.DemoAPI.Enable)
			gDemo.PATCH(":id/disable", a.DemoAPI.Disable)
		}

		gMenu := v1.Group("menus")
		{
			gMenu.GET("", a.MenuAPI.Query)
			gMenu.GET(":id", a.MenuAPI.Get)
			gMenu.POST("", a.MenuAPI.Create)
			gMenu.PUT(":id", a.MenuAPI.Update)
			gMenu.DELETE(":id", a.MenuAPI.Delete)
			gMenu.PATCH(":id/enable", a.MenuAPI.Enable)
			gMenu.PATCH(":id/disable", a.MenuAPI.Disable)
		}
		v1.GET("/menus.tree", a.MenuAPI.QueryTree)

		gRole := v1.Group("roles")
		{
			gRole.GET("", a.RoleAPI.Query)
			gRole.GET(":id", a.RoleAPI.Get)
			gRole.POST("", a.RoleAPI.Create)
			gRole.PUT(":id", a.RoleAPI.Update)
			gRole.DELETE(":id", a.RoleAPI.Delete)
			gRole.PATCH(":id/enable", a.RoleAPI.Enable)
			gRole.PATCH(":id/disable", a.RoleAPI.Disable)
		}
		v1.GET("/roles.select", a.RoleAPI.QuerySelect)

		gUser := v1.Group("users")
		{
			gUser.GET("", a.UserAPI.Query)
			gUser.GET(":id", a.UserAPI.Get)
			gUser.POST("", a.UserAPI.Create)
			gUser.PUT(":id", a.UserAPI.Update)
			gUser.DELETE(":id", a.UserAPI.Delete)
			gUser.PATCH(":id/enable", a.UserAPI.Enable)
			gUser.PATCH(":id/disable", a.UserAPI.Disable)
		}
	}
}
