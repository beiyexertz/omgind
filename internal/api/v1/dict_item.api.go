package api_v1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/wanhello/omgind/internal/app/ginx"
	"github.com/wanhello/omgind/internal/app/schema"
	"github.com/wanhello/omgind/internal/app/service"
)

// DictItemSet 注入DictItem
var DictItemSet = wire.NewSet(wire.Struct(new(DictItem), "*"))

// DictItem 字典项
type DictItem struct {
	DictItemSrv *service.DictItem
}

// Query 查询数据
func (a *DictItem) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.DictItemQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}

	params.Pagination = true
	result, err := a.DictItemSrv.Query(ctx, params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	ginx.ResPage(c, result.Data, result.PageResult)
}

// Get 查询指定数据
func (a *DictItem) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.DictItemSrv.Get(ctx, c.Param("id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

// Create 创建数据
func (a *DictItem) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.DictItem
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	item.Creator = ginx.GetUserID(c)
	result, err := a.DictItemSrv.Create(ctx, item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, result)
}

// Update 更新数据
func (a *DictItem) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.DictItem
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	err := a.DictItemSrv.Update(ctx, c.Param("id"), item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}

// Delete 删除数据
func (a *DictItem) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.DictItemSrv.Delete(ctx, c.Param("id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}
