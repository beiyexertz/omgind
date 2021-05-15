package router

import (
	"github.com/gin-gonic/gin"
	api_v1 "github.com/wanhello/omgind/internal/api/v1"
)

func (r *Router) initDemoRouterV1(urg *gin.RouterGroup, api *api_v1.Demo, pathcomp string) {
	gDemo := urg.Group(pathcomp)
	{
		gDemo.GET("", api.Query)
		gDemo.GET(":id", api.Get)
		gDemo.POST("", api.Create)
		gDemo.PUT(":id", api.Update)
		gDemo.DELETE(":id", api.Delete)
		gDemo.PATCH(":id/enable", api.Enable)
		gDemo.PATCH(":id/disable", api.Disable)
	}
}
