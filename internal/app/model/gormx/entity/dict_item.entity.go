package entity

import (
	"context"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/wanhello/omgind/internal/app/schema"
	"github.com/wanhello/omgind/pkg/helper/structure"
)

// GetDictItemDB 获取DictItem存储
func GetDictItemDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, defDB, new(DictItem))
}

// SchemaDictItem 字典项对象
type SchemaDictItem schema.DictItem

// ToDictItem 转换为实体
func (a SchemaDictItem) ToDictItem() *DictItem {
	item := new(DictItem)
	structure.Copy(a, item)
	return item
}

// DictItem 字典项实体
type DictItem struct {
	ID string `gorm:"column:id;primary_key;size:36;"`

	Label  string `gorm:"column:label;size:128;"`      // 显示值
	Value  int    `gorm:"column:value;"`               // 字典值
	Status int    `gorm:"column:status;default:true;"` // 状态(1:启用 2:禁用)
	DictId string `gorm:"column:dict_id;size:36;"`     // dict.id
	Memo   string `gorm:"column:memo;size:128;"`       // 备注
	Sort   int    `gorm:"column:sort;default:9999;"`   // 排序

	IsDel     bool       `json:"is_del"`
	Creator   string     `gorm:"column:creator;"` // 创建者
	CreatedAt time.Time  `gorm:"column:created_at;index;"`
	UpdatedAt time.Time  `gorm:"column:updated_at;index;"`
	DeletedAt *time.Time `gorm:"column:deleted_at;index;"`
}

// ToSchemaDictItem 转换为demo对象
func (a DictItem) ToSchemaDictItem() *schema.DictItem {
	item := new(schema.DictItem)
	structure.Copy(a, item)
	return item
}

// DictItems 字典项实体列表
type DictItems []*DictItem

// ToSchemaDictItems 转换为对象列表
func (a DictItems) ToSchemaDictItems() schema.DictItems {
	list := make(schema.DictItems, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaDictItem()
	}
	return list
}
