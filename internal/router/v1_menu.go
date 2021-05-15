package router

import (
	"github.com/gin-gonic/gin"
	api_v1 "github.com/wanhello/omgind/internal/api/v1"
)

func (r *Router) initMenuRouterV1(urg *gin.RouterGroup, api *api_v1.Menu, pathcomp string) {
	gMenu := urg.Group(pathcomp)
	{
		gMenu.GET("", api.Query)
		gMenu.GET(":id", api.Get)
		gMenu.POST("", api.Create)
		gMenu.PUT(":id", api.Update)
		gMenu.DELETE(":id", api.Delete)
		gMenu.PATCH(":id/enable", api.Enable)
		gMenu.PATCH(":id/disable", api.Disable)
	}

}
