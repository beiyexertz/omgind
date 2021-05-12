package entity

import (
	"context"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/wanhello/omgind/internal/app/schema"
	"github.com/wanhello/omgind/pkg/helper/structure"
)

// GetDictDB 获取Dict存储
func GetDictDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, defDB, new(Dict))
}

// SchemaDict 字典对象
type SchemaDict schema.Dict

// ToDict 转换为实体
func (a SchemaDict) ToDict() *Dict {
	item := new(Dict)
	structure.Copy(a, item)
	return item
}

// Dict 字典实体
type Dict struct {
	ID      string `gorm:"column:id;primary_key;size:36;"`
	NameCn  string `gorm:"column:name_cn;size:128;"`  // 字典名（中）
	NameEn  string `gorm:"column:name_en;size:128;"`  // 字典名（英）
	Status  int    `gorm:"column:status;default:1;"`  // 状态
	Memo    string `gorm:"column:memo;size:128;"`     // 备注
	Sort    int    `gorm:"column:sort;default:9999;"` // 排序
	Creator string `gorm:"column:creator;"`           // 创建者

	IsDel bool `gorm:"column:is_del;default:false;index;not null;"`

	CreatedAt time.Time  `gorm:"column:created_at;index;"`
	UpdatedAt time.Time  `gorm:"column:updated_at;index;"`
	DeletedAt *time.Time `gorm:"column:deleted_at;index;"`
}

// ToSchemaDict 转换为demo对象
func (a Dict) ToSchemaDict() *schema.Dict {
	item := new(schema.Dict)
	structure.Copy(a, item)
	return item
}

// Dicts 字典实体列表
type Dicts []*Dict

// ToSchemaDicts 转换为对象列表
func (a Dicts) ToSchemaDicts() schema.Dicts {
	list := make(schema.Dicts, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaDict()
	}
	return list
}
