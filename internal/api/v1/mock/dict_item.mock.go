package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// DictItemSet 注入DictItem
var DictItemSet = wire.NewSet(wire.Struct(new(DictItem), "*"))

// DictItem 字典项
type DictItem struct {
}

// Query 查询数据
// @Tags 字典项
// @Summary 查询数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param current query int true "分页索引" default(1)
// @Param pageSize query int true "分页大小" default(10)
// @Param queryValue query string false "查询值"
// @Success 200 {array} schema.DictItem "查询结果：{list:列表数据,pagination:{current:页索引,pageSize:页大小,total:总数量}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v1/dict-items [get]
func (a *DictItem) Query(c *gin.Context) {
}

// Get 查询指定数据
// @Tags 字典项
// @Summary 查询指定数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "唯一标识"
// @Success 200 {object} schema.DictItem
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 404 {object} schema.ErrorResult "{error:{code:0,message:资源不存在}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v1/dict-items/{id} [get]
func (a *DictItem) Get(c *gin.Context) {
}

// Create 创建数据
// @Tags 字典项
// @Summary 创建数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param body body schema.DictItem true "创建数据"
// @Success 200 {object} schema.IDResult
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:无效的请求参数}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v1/dict-items [post]
func (a *DictItem) Create(c *gin.Context) {
}

// Update 更新数据
// @Tags 字典项
// @Summary 更新数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "唯一标识"
// @Param body body schema.DictItem true "更新数据"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:无效的请求参数}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v1/dict-items/{id} [put]
func (a *DictItem) Update(c *gin.Context) {
}

// Delete 删除数据
// @Tags 字典项
// @Summary 删除数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "唯一标识"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v1/dict-items/{id} [delete]
func (a *DictItem) Delete(c *gin.Context) {
}