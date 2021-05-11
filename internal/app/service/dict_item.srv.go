package service

import (
	"context"

	"github.com/google/wire"
	"github.com/wanhello/omgind/internal/app/model/gormx/repo"
	"github.com/wanhello/omgind/internal/app/schema"
	"github.com/wanhello/omgind/pkg/errors"
	uid "github.com/wanhello/omgind/pkg/helper/uid/ulid"
)

// var _ service.IDictItem = (*DictItem)(nil)

// DictItemSet 注入DictItem
var DictItemSet = wire.NewSet(wire.Struct(new(DictItem), "*"))

// DictItem 字典项
type DictItem struct {
	DictItemModel *repo.DictItem
}

// Query 查询数据
func (a *DictItem) Query(ctx context.Context, params schema.DictItemQueryParam, opts ...schema.DictItemQueryOptions) (*schema.DictItemQueryResult, error) {
	return a.DictItemModel.Query(ctx, params, opts...)
}

// Get 查询指定数据
func (a *DictItem) Get(ctx context.Context, id string, opts ...schema.DictItemQueryOptions) (*schema.DictItem, error) {
	item, err := a.DictItemModel.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

// Create 创建数据
func (a *DictItem) Create(ctx context.Context, item schema.DictItem) (*schema.IDResult, error) {
	// TODO: check?
	item.ID = uid.MustString()
	err := a.DictItemModel.Create(ctx, item)
	if err != nil {
		return nil, err
	}

	return schema.NewIDResult(item.ID), nil
}

// Update 更新数据
func (a *DictItem) Update(ctx context.Context, id string, item schema.DictItem) error {
	oldItem, err := a.DictItemModel.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}
	// TODO: check?
	item.ID = oldItem.ID
	item.Creator = oldItem.Creator
	item.CreatedAt = oldItem.CreatedAt

	return a.DictItemModel.Update(ctx, id, item)
}

// Delete 删除数据
func (a *DictItem) Delete(ctx context.Context, id string) error {
	oldItem, err := a.DictItemModel.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.DictItemModel.Delete(ctx, id)
}

func (a *DictItem) DeleteS(ctx context.Context, id string) error {
	oldItem, err := a.DictItemModel.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}
	oldItem.IsDel = true
	return a.DictItemModel.Update(ctx, id, *oldItem)
}
