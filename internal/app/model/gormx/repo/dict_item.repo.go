package repo

import (
	"context"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/wanhello/omgind/internal/app/model/gormx/entity"
	"github.com/wanhello/omgind/internal/app/schema"
	"github.com/wanhello/omgind/pkg/errors"
)

// var _ model.IDictItem = (*DictItem)(nil)

// DictItemSet 注入DictItem
var DictItemSet = wire.NewSet(wire.Struct(new(DictItem), "*"))

// DictItem 字典项存储
type DictItem struct {
	DB *gorm.DB
}

func (a *DictItem) getQueryOption(opts ...schema.DictItemQueryOptions) schema.DictItemQueryOptions {
	var opt schema.DictItemQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *DictItem) Query(ctx context.Context, params schema.DictItemQueryParam, opts ...schema.DictItemQueryOptions) (*schema.DictItemQueryResult, error) {
	opt := a.getQueryOption(opts...)

	db := entity.GetDictItemDB(ctx, a.DB)
	// TODO: 查询条件

	opt.OrderFields = append(opt.OrderFields, schema.NewOrderField("id", schema.OrderByDESC))
	db = db.Order(ParseOrder(opt.OrderFields))

	var list entity.DictItems
	pr, err := WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	qr := &schema.DictItemQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaDictItems(),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *DictItem) Get(ctx context.Context, id string, opts ...schema.DictItemQueryOptions) (*schema.DictItem, error) {
	db := entity.GetDictItemDB(ctx, a.DB).Where("id=?", id)
	var item entity.DictItem
	ok, err := FindOne(ctx, db, &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaDictItem(), nil
}

// Create 创建数据
func (a *DictItem) Create(ctx context.Context, item schema.DictItem) error {
	eitem := entity.SchemaDictItem(item).ToDictItem()
	result := entity.GetDictItemDB(ctx, a.DB).Create(eitem)
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Update 更新数据
func (a *DictItem) Update(ctx context.Context, id string, item schema.DictItem) error {
	eitem := entity.SchemaDictItem(item).ToDictItem()
	result := entity.GetDictItemDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Delete 删除数据
func (a *DictItem) Delete(ctx context.Context, id string) error {
	result := entity.GetDictItemDB(ctx, a.DB).Where("id=?", id).Delete(entity.DictItem{})
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
