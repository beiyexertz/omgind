package repo_ent

import (
	"context"
	"time"

	"github.com/google/wire"
	"github.com/wanhello/omgind/internal/app/schema"
	"github.com/wanhello/omgind/internal/gen/ent"
	"github.com/wanhello/omgind/internal/gen/ent/xxxdemo"
	"github.com/wanhello/omgind/pkg/errors"
	"github.com/wanhello/omgind/pkg/helper/structure"
)

var DemoSet = wire.NewSet(wire.Struct(new(Demo), "*"))

type Demo struct {
	//DB     *gorm.DB
	EntCli *ent.Client
}

// ToSchemaDemo 转换为
func toSchemaDemo(a *ent.XxxDemo) *schema.Demo {
	item := new(schema.Demo)
	structure.Copy(a, item)
	return item
}

func toSchemaDemos(us ent.XxxDemos) []*schema.Demo {
	list := make([]*schema.Demo, len(us))
	for i, item := range us {
		list[i] = toSchemaDemo(item)
	}
	return list
}

func toEntCreateDemoInput(a *schema.Demo) *ent.CreateXxxDemoInput {
	createinput := new(ent.CreateXxxDemoInput)
	structure.Copy(a, &createinput)

	return createinput
}

func toEntUpdateDemoInput(a *schema.Demo) *ent.UpdateXxxDemoInput {
	updateinput := new(ent.UpdateXxxDemoInput)
	structure.Copy(a, &updateinput)

	return updateinput
}

func (a *Demo) getQueryOption(opts ...schema.DemoQueryOptions) schema.DemoQueryOptions {
	var opt schema.DemoQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *Demo) Query(ctx context.Context, params schema.DemoQueryParam, opts ...schema.DemoQueryOptions) (*schema.DemoQueryResult, error) {

	opt := a.getQueryOption(opts...)

	query := a.EntCli.XxxDemo.Query()

	if v := params.Code; v != "" {
		query = query.Where(xxxdemo.CodeEQ(v))
	}

	if v := params.QueryValue; v != "" {
		query = query.Where(xxxdemo.Or(xxxdemo.CodeContains(v), xxxdemo.NameContains(v)))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// order field
	opt.OrderFields = append(opt.OrderFields, schema.NewOrderField("id", schema.OrderByDESC))

	query = query.Order(ParseOrder(opt.OrderFields)...)

	// get total
	pr := &schema.PaginationResult{Total: count}
	if params.PaginationParam.OnlyCount {
		return &schema.DemoQueryResult{PageResult: pr}, nil
	}

	pr.Current = params.PaginationParam.GetCurrent()
	pr.PageSize = params.PaginationParam.GetPageSize()
	if params.Offset() > count {
		return &schema.DemoQueryResult{PageResult: pr}, nil
	}

	query = query.Limit(params.Limit()).Offset(params.Offset())

	list, err1 := query.All(ctx)
	if err1 != nil {
		return nil, errors.WithStack(err)
	}
	rlist := ent.XxxDemos(list)

	qr := &schema.DemoQueryResult{
		PageResult: pr,
		Data:       toSchemaDemos(rlist),
	}
	return qr, nil
}

// Get 查询指定数据
func (a *Demo) Get(ctx context.Context, id string, opts ...schema.DemoQueryOptions) (*schema.Demo, error) {

	xxxdemo, err := a.EntCli.XxxDemo.Query().Where(xxxdemo.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}

	return toSchemaDemo(xxxdemo), nil
}

// Create 创建数据
func (a *Demo) Create(ctx context.Context, item schema.Demo) (*schema.Demo, error) {

	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()

	iteminput := toEntCreateDemoInput(&item)
	xxxdemo, err := a.EntCli.XxxDemo.Create().SetInput(*iteminput).Save(ctx)

	if err != nil {
		return nil, err
	}
	sch_demo := toSchemaDemo(xxxdemo)
	return sch_demo, nil
}

// Update 更新数据
func (a *Demo) Update(ctx context.Context, id string, item schema.Demo) (*schema.Demo, error) {

	oitem, err := a.EntCli.XxxDemo.Query().Where(xxxdemo.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}

	item.UpdatedAt = time.Now()
	iteminput := toEntUpdateDemoInput(&item)
	xxxdemo, err := oitem.Update().SetInput(*iteminput).Save(ctx)
	sch_demo := toSchemaDemo(xxxdemo)

	return sch_demo, nil
}

// Delete 删除数据
func (a *Demo) Delete(ctx context.Context, id string) error {

	xxxdemo, err := a.EntCli.XxxDemo.Query().Where(xxxdemo.IDEQ(id)).Only(ctx)

	if err != nil {
		return err
	}

	_, err1 := xxxdemo.Update().SetDeletedAt(time.Now()).Save(ctx)

	return errors.WithStack(err1)
}

// UpdateStatus 更新状态
func (a *Demo) UpdateStatus(ctx context.Context, id string, status int) error {

	_, err1 := a.EntCli.XxxDemo.Update().Where(xxxdemo.IDEQ(id)).SetStatus(status).Save(ctx)

	return errors.WithStack(err1)
}
