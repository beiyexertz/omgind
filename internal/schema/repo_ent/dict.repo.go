package repo_ent

import (
	"context"
	"time"

	"github.com/google/wire"
	"github.com/wanhello/omgind/internal/app/schema"
	"github.com/wanhello/omgind/internal/gen/ent"
	"github.com/wanhello/omgind/internal/gen/ent/sysdict"
	"github.com/wanhello/omgind/internal/gen/ent/xxxdemo"
	"github.com/wanhello/omgind/pkg/errors"
	"github.com/wanhello/omgind/pkg/helper/structure"
)

// DictSet 注入Dict
var DictSet = wire.NewSet(wire.Struct(new(Dict), "*"))

// Dict 字典存储
type Dict struct {
	EntCli *ent.Client
}

func (a *Dict) toSchemaDict(dit *ent.SysDict) *schema.Dict {
	item := new(schema.Dict)
	structure.Copy(dit, item)
	return item
}

func (a *Dict) toSchemaDicts(dits ent.SysDicts) []*schema.Dict {
	list := make([]*schema.Dict, len(dits))
	for i, item := range dits {
		list[i] = a.toSchemaDict(item)
	}
	return list
}

func (a *Dict) ToEntCreateSysDictInput(sdi *schema.Dict) *ent.CreateSysDictInput {
	createinput := new(ent.CreateSysDictInput)
	structure.Copy(sdi, &createinput)

	return createinput
}

func (a *Dict) ToEntUpdateSysDictInput(sdi *schema.Dict) *ent.UpdateSysDictInput {
	updateinput := new(ent.UpdateSysDictInput)
	structure.Copy(sdi, &updateinput)

	return updateinput
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

	query := a.EntCli.SysDict.Query()

	if v := params.IDs; len(v) > 0 {
		query = query.Where(sysdict.IDIn(v...))
	}

	if v := params.NameCn; v != "" {
		query = query.Where(sysdict.NameCnNEQ(v))
	}

	if v := params.NameEn; v != "" {
		query = query.Where(sysdict.NameEnEQ(v))
	}

	if v := params.Status; v > 0 {
		query = query.Where(sysdict.StatusEQ(v))
	}

	if v := params.QueryValue; v != "" {
		query = query.Where(sysdict.Or(sysdict.NameCnContains(v), sysdict.NameEnEQ(v), sysdict.MemoContains(v)))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// get total
	pr := &schema.PaginationResult{Total: count}
	if params.PaginationParam.OnlyCount {
		return &schema.DictQueryResult{PageResult: pr}, nil
	}

	opt.OrderFields = append(opt.OrderFields, schema.NewOrderField("id", schema.OrderByDESC))
	query = query.Order(ParseOrder(opt.OrderFields)...)

	pr.Current = params.PaginationParam.GetCurrent()
	pr.PageSize = params.PaginationParam.GetPageSize()
	if params.Offset() > count {
		return &schema.DictQueryResult{PageResult: pr}, nil
	}

	query = query.Limit(params.Limit()).Offset(params.Offset())

	list, err1 := query.All(ctx)
	if err1 != nil {
		return nil, errors.WithStack(err)
	}
	rlist := ent.SysDicts(list)

	qr := &schema.DictQueryResult{
		PageResult: pr,
		Data:       a.toSchemaDicts(rlist),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *Dict) Get(ctx context.Context, id string, opts ...schema.DictQueryOptions) (*schema.Dict, error) {

	dict, err := a.EntCli.SysDict.Query().Where(sysdict.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return a.toSchemaDict(dict), nil
}

// Create 创建数据
func (a *Dict) Create(ctx context.Context, item schema.Dict) (*schema.Dict, error) {
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()

	iteminput := a.toEntCreateSysDictInput(&item)
	sysdict, err := a.EntCli.SysDict.Create().SetInput(*iteminput).Save(ctx)

	if err != nil {
		return nil, err
	}
	sch_dict := a.toSchemaDict(sysdict)
	return sch_dict, nil
}

// Update 更新数据
func (a *Dict) Update(ctx context.Context, id string, item schema.Dict) (*schema.Dict, error) {
	oitem, err := a.EntCli.SysDict.Query().Where(sysdict.IDEQ(id)).Only(ctx)

	if err != nil {
		return nil, err
	}

	item.UpdatedAt = time.Now()
	itemInput := a.toEntUpdateSysDictInput(&item)
	dict, err := oitem.Update().SetInput(*itemInput).Save(ctx)
	sch_dict := a.toSchemaDict(dict)

	return sch_dict, nil
}

// Delete 删除数据
func (a *Dict) Delete(ctx context.Context, id string) error {
	dict, err := a.EntCli.SysDict.Query().Where(sysdict.IDEQ(id)).Only(ctx)

	if err != nil {
		return err
	}

	_, err1 := dict.Update().SetDeletedAt(time.Now()).Save(ctx)

	return errors.WithStack(err1)
}

// UpdateStatus 更新状态
func (a *Dict) UpdateStatus(ctx context.Context, id string, status int) error {
	_, err1 := a.EntCli.XxxDemo.Update().Where(xxxdemo.IDEQ(id)).SetStatus(status).Save(ctx)

	return errors.WithStack(err1)
}
