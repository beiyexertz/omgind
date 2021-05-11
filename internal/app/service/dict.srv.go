package service

import (
	"context"

	"github.com/google/wire"
	"github.com/wanhello/omgind/internal/app/model/gormx/repo"
	"github.com/wanhello/omgind/internal/app/schema"
	"github.com/wanhello/omgind/pkg/errors"
	uid "github.com/wanhello/omgind/pkg/helper/uid/ulid"
)

// var _ service.IDict = (*Dict)(nil)

// DictSet 注入Dict
var DictSet = wire.NewSet(wire.Struct(new(Dict), "*"))

// Dict 字典
type Dict struct {
	TransModel    *repo.Trans
	DictModel     *repo.Dict
	DictItemModel *repo.DictItem
}

// Query 查询数据
func (a *Dict) Query(ctx context.Context, params schema.DictQueryParam, opts ...schema.DictQueryOptions) (*schema.DictQueryResult, error) {
	return a.DictModel.Query(ctx, params, opts...)
}

// Get 查询指定数据
func (a *Dict) Get(ctx context.Context, id string, opts ...schema.DictQueryOptions) (*schema.Dict, error) {
	item, err := a.DictModel.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}
	ditems, err := a.QueryItems(ctx, id)
	if err != nil {
		return nil, err
	}
	item.Items = ditems

	return item, nil
}

func (d Dict) QueryItems(ctx context.Context, id string) (schema.DictItems, error) {
	result, err := d.DictItemModel.Query(ctx, schema.DictItemQueryParam{
		DictId: id,
	})
	if err != nil {
		return nil, err
	} else if len(result.Data) == 0 {
		return nil, nil
	}

	return result.Data, nil
}

func (d *Dict) checkName(ctx context.Context, item schema.Dict) error {
	// TODO:: need optimization
	result1, err1 := d.DictModel.Query(ctx, schema.DictQueryParam{
		PaginationParam: schema.PaginationParam{
			OnlyCount: true,
		},
		NameEn: item.NameEn,
	})
	result2, err2 := d.DictModel.Query(ctx, schema.DictQueryParam{
		PaginationParam: schema.PaginationParam{
			OnlyCount: true,
		},
		NameCn: item.NameCn,
	})

	if err1 != nil && err2 != nil {
		return nil
	} else if result1.PageResult.Total > 0 {
		return errors.New400Response("字典名称" + item.NameEn + "已经存在")
	} else if result2.PageResult.Total > 0 {
		return errors.New400Response("字典名称" + item.NameCn + "已经存在")
	}

	return nil
}

// Create 创建数据
func (a *Dict) Create(ctx context.Context, item schema.Dict) (*schema.IDResult, error) {
	if err := a.checkName(ctx, item); err != nil {
		return nil, err
	}
	item.ID = uid.MustString()
	err := a.TransModel.Exec(ctx, func(ctx context.Context) error {
		err := a.createDictItems(ctx, item.ID, item.Items)
		if err != nil {
			return err
		}
		return a.DictModel.Create(ctx, item)
	})
	if err != nil {
		return nil, err
	}

	return schema.NewIDResult(item.ID), nil
}

func (a *Dict) createDictItems(ctx context.Context, dictID string, items schema.DictItems) error {

	for _, item := range items {
		item.ID = uid.MustString()
		item.DictId = dictID
		err := a.DictItemModel.Create(ctx, *item)
		if err != nil {
			return nil
		}
	}
	return nil
}

// Update 更新数据
func (a *Dict) Update(ctx context.Context, id string, item schema.Dict) error {
	oldItem, err := a.DictModel.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}
	// TODO: check?
	item.ID = oldItem.ID
	item.Creator = oldItem.Creator
	item.CreatedAt = oldItem.CreatedAt

	return a.DictModel.Update(ctx, id, item)
}

// Delete 删除数据
func (a *Dict) Delete(ctx context.Context, id string) error {
	oldItem, err := a.DictModel.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.DictModel.Delete(ctx, id)
}
