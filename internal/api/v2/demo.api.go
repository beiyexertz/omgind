package api_v2

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/wanhello/omgind/internal/app/ginx"
	"github.com/wanhello/omgind/internal/app/schema"
	service_ent "github.com/wanhello/omgind/internal/app/service.ent"
)

// DemoSet 注入Demo
var DemoSet = wire.NewSet(wire.Struct(new(Demo), "*"))

// Demo 示例程序
type Demo struct {
	DemoSrv *service_ent.Demo
}

// Query 查询数据
func (a *Demo) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.DemoQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}

	params.Pagination = true
	result, err := a.DemoSrv.Query(ctx, params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	fmt.Printf(" ------- ===== %+v \n", result.PageResult)
	fmt.Printf(" ------- ===== %+v \n", result.Data)

	ginx.ResPage(c, result.Data, result.PageResult)
}

// Get 查询指定数据
func (a *Demo) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.DemoSrv.Get(ctx, c.Param("id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

// Create 创建数据
func (a *Demo) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.Demo
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	item.Creator = ginx.GetUserID(c)

	result, err := a.DemoSrv.Create(ctx, item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, result)
}

// Update 更新数据
func (a *Demo) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.Demo
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	_, err := a.DemoSrv.Update(ctx, c.Param("id"), item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
	//ginx.ResSuccess(c, result)
}

// Delete 删除数据
func (a *Demo) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.DemoSrv.Delete(ctx, c.Param("id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}

// Enable 启用数据
func (a *Demo) Enable(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.DemoSrv.UpdateStatus(ctx, c.Param("id"), 1)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}

// Disable 禁用数据
func (a *Demo) Disable(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.DemoSrv.UpdateStatus(ctx, c.Param("id"), 2)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}
