package router_v2

import (
	"github.com/gin-gonic/gin"
	api_v2 "github.com/wanhello/omgind/internal/api/v2"
)

func InitDemoRouter(router *gin.RouterGroup, demoapi *api_v2.Demo, path string) {
	demoRouter := router.Group(path)
	{
		demoRouter.GET("", demoapi.Query)
		demoRouter.GET(":id", demoapi.Get)
		demoRouter.POST("", demoapi.Create)
		demoRouter.PUT(":id", demoapi.Update)
		demoRouter.DELETE(":id", demoapi.Delete)
		demoRouter.PATCH(":id/enable", demoapi.Enable)
		demoRouter.PATCH(":id/disable", demoapi.Disable)
	}
}
