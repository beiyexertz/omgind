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
	DictModel *repo.Dict
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

	return item, nil
}

// Create 创建数据
func (a *Dict) Create(ctx context.Context, item schema.Dict) (*schema.IDResult, error) {
	// TODO: check?
	item.ID = uid.MustString()
	err := a.DictModel.Create(ctx, item)
	if err != nil {
		return nil, err
	}

	return schema.NewIDResult(item.ID), nil
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
