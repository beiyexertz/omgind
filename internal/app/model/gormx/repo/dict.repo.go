package repo

import (
	"context"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/wanhello/omgind/internal/app/model/gormx/entity"
	"github.com/wanhello/omgind/internal/app/schema"
	"github.com/wanhello/omgind/pkg/errors"
)

// var _ model.IDict = (*Dict)(nil)

// DictSet 注入Dict
var DictSet = wire.NewSet(wire.Struct(new(Dict), "*"))

// Dict 字典存储
type Dict struct {
	DB *gorm.DB
}

func (a *Dict) getQueryOption(opts ...schema.DictQueryOptions) schema.DictQueryOptions {
	var opt schema.DictQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *Dict) Query(ctx context.Context, params schema.DictQueryParam, opts ...schema.DictQueryOptions) (*schema.DictQueryResult, error) {
	opt := a.getQueryOption(opts...)

	db := entity.GetDictDB(ctx, a.DB)
	if v := params.IDs; len(v) > 0 {
		db = db.Where("id IN (?)", v)
	}

	if v := params.NameCn; v != "" {
		db = db.Where("name_cn=?", v)
	}

	if v := params.NameEn; v != "" {
		db = db.Where("name_en=?", v)
	}

	if v := params.Status; v > 0 {
		db = db.Where("status=?", v)
	}

	if v := params.QueryValue; v != "" {
		v = "%" + v + "%"
		db = db.Where("name_cn LIKE ? OR name_en LIKE ? OR memo LIKE ?", v, v, v)
	}

	opt.OrderFields = append(opt.OrderFields, schema.NewOrderField("id", schema.OrderByDESC))
	db = db.Order(ParseOrder(opt.OrderFields))

	var list entity.Dicts
	pr, err := WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	qr := &schema.DictQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaDicts(),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *Dict) Get(ctx context.Context, id string, opts ...schema.DictQueryOptions) (*schema.Dict, error) {
	db := entity.GetDictDB(ctx, a.DB).Where("id=?", id)
	var item entity.Dict
	ok, err := FindOne(ctx, db, &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaDict(), nil
}

// Create 创建数据
func (a *Dict) Create(ctx context.Context, item schema.Dict) error {
	eitem := entity.SchemaDict(item).ToDict()
	result := entity.GetDictDB(ctx, a.DB).Create(eitem)
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Update 更新数据
func (a *Dict) Update(ctx context.Context, id string, item schema.Dict) error {
	eitem := entity.SchemaDict(item).ToDict()
	result := entity.GetDictDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Delete 删除数据
func (a *Dict) Delete(ctx context.Context, id string) error {
	result := entity.GetDictDB(ctx, a.DB).Where("id=?", id).Delete(entity.Dict{})
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
