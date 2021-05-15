package router

import (
	"github.com/gin-gonic/gin"
	api_v1 "github.com/wanhello/omgind/internal/api/v1"
)

func (r *Router) initRoleRouterV1(urg *gin.RouterGroup, api *api_v1.Role, pathcomp string) {
	gRole := urg.Group(pathcomp)
	{
		gRole.GET("", api.Query)
		gRole.GET(":id", api.Get)
		gRole.POST("", api.Create)
		gRole.PUT(":id", api.Update)
		gRole.DELETE(":id", api.Delete)
		gRole.PATCH(":id/enable", api.Enable)
		gRole.PATCH(":id/disable", api.Disable)
	}

}
