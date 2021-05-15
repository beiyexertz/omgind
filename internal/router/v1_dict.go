package router

import (
	"github.com/gin-gonic/gin"
	api_v1 "github.com/wanhello/omgind/internal/api/v1"
)

func (r *Router) initDictRouterV1(urg *gin.RouterGroup, api *api_v1.Dict, pathcomp string) {

	gDict := urg.Group(pathcomp)
	{
		gDict.GET("", api.Query)
		gDict.GET(":id", api.Get)
		gDict.POST("", api.Create)
		gDict.PUT(":id", api.Update)
		gDict.DELETE(":id", api.Delete)
		gDict.PATCH(":id/enable", api.Enable)
		gDict.PATCH(":id/disable", api.Disable)
	}

}
