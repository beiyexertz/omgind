package service

import (
	"context"
	"fmt"

	"github.com/google/wire"
	"github.com/wanhello/omgind/internal/app/model/gormx/repo"
	"github.com/wanhello/omgind/internal/app/schema"
	"github.com/wanhello/omgind/pkg/errors"

	"github.com/wanhello/omgind/pkg/helper/deepcopier"

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
		DictID: id,
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
		item.DictID = dictID
		err := a.DictItemModel.Create(ctx, *item)
		if err != nil {
			return nil
		}
	}
	return nil
}

// Update 更新数据
func (a *Dict) Update(ctx context.Context, id string, item schema.Dict) error {

	oldItem, err := a.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	} else if oldItem.NameEn != item.NameEn || oldItem.NameCn != item.NameCn {
		if err := a.checkName(ctx, item); err != nil {
			return err
		}
	}

	item.ID = oldItem.ID
	item.Creator = oldItem.Creator
	item.CreatedAt = oldItem.CreatedAt
	return a.TransModel.Exec(ctx, func(ctx context.Context) error {

		err := a.updateDictItems(ctx, id, oldItem.Items, item.Items)
		if err != nil {
			return err
		}
		return a.DictModel.Update(ctx, id, item)
	})
	//return a.DictModel.Update(ctx, id, item)
}

func (a *Dict) updateDictItems(ctx context.Context, dictID string, oldItems, newItems schema.DictItems) error {

	addItems, delItems, updateItems := a.compareDictItems(ctx, oldItems, newItems)

	err := a.createDictItems(ctx, dictID, addItems)
	if err != nil {
		return err
	}

	for _, item := range delItems {
		err := a.DictItemModel.Delete(ctx, item.ID)
		if err != nil {
			return err
		}
	}

	mOldItems := oldItems.ToMap()
	for _, item := range updateItems {
		oitem := mOldItems[item.Label]

		hasChange := oitem.Compare(item)
		deepcopier.Copy(item).Include([]string{"Label", "Value", "Memo", "Status"}).Exclude([]string{"CreatedAt",
			"UpdatedAt", "ID"}).To(oitem)

		if !hasChange {
			err := a.DictItemModel.Update(ctx, oitem.ID, *oitem)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (a *Dict) compareDictItems(ctx context.Context, oldItems, newItems schema.DictItems) (addList, delList, updateList schema.DictItems) {

	mOldItems := oldItems.ToMap()
	mNewItems := newItems.ToMap()
	for k, item := range mNewItems {
		if _, ok := mOldItems[k]; ok {
			updateList = append(updateList, item)
			delete(mOldItems, k)
			continue
		}
		addList = append(addList, item)
	}
	for _, item := range mOldItems {
		delList = append(delList, item)
	}
	return
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

// Delete 删除数据
func (a *Dict) DeleteS(ctx context.Context, id string) error {
	oldItem, err := a.DictModel.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}
	return a.DictModel.Update(ctx, id, *oldItem)
}

func (a *Dict) UpdateStatus(ctx context.Context, id string, status int) error {

	fmt.Printf(" ---- ==== status = %d ", status)

	oldItem, err := a.DictModel.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.DictModel.UpdateStatus(ctx, id, status)
}
