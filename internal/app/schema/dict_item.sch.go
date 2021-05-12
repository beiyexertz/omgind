package schema

import "time"

// DictItem 字典项对象
type DictItem struct {
	ID string `json:"id"` // 唯一标识

	Label  string `json:"label" binding:"required"`   // 显示值
	Value  int    `json:"value" binding:"required"`   // 字典值
	Status int    `json:"status" binding:"required"`  // 状态(1:启用 2:禁用)
	DictId string `json:"dict_id" binding:"required"` // dict.id
	Memo   string `json:"memo"`                       // 备注
	Sort   int    `json:"sort"`                       // 排序
	IsDel  bool   `json:"is_del"`

	Creator   string    `json:"creator"`    // 创建者
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间

}

// DictItemQueryParam 查询条件
type DictItemQueryParam struct {
	PaginationParam
	DictId string //字典id
	IDs    []string
}

// DictItemQueryOptions 查询可选参数项
type DictItemQueryOptions struct {
	OrderFields []*OrderField // 排序字段
}

// DictItemQueryResult 查询结果
type DictItemQueryResult struct {
	Data       DictItems
	PageResult *PaginationResult
}

// DictItems 字典项列表
type DictItems []*DictItem

// ToMap
func (a DictItems) ToMap() map[string]*DictItem {
	m := make(map[string]*DictItem)
	for _, item := range a {
		m[item.Label] = item
	}
	return m
}
