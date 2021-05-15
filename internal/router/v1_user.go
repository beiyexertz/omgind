package router

import (
	"github.com/gin-gonic/gin"
	api_v1 "github.com/wanhello/omgind/internal/api/v1"
)

func (r *Router) initUserRouterV1(urg *gin.RouterGroup, api *api_v1.User, pathcomp string) {

	gUser := urg.Group(pathcomp)
	{
		gUser.GET("", api.Query)
		gUser.GET(":id", api.Get)
		gUser.POST("", api.Create)
		gUser.PUT(":id", api.Update)
		gUser.DELETE(":id", api.Delete)
		gUser.PATCH(":id/enable", api.Enable)
		gUser.PATCH(":id/disable", api.Disable)
	}

}
