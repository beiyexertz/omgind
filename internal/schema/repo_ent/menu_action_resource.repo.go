package repo_ent

import (
	"context"

	"github.com/google/wire"
	"github.com/wanhello/omgind/internal/gen/ent/sysmenuactionresource"
	"github.com/wanhello/omgind/pkg/helper/structure"

	"github.com/wanhello/omgind/pkg/errors"
	"github.com/wanhello/omgind/internal/app/model/gormx/entity"
	"github.com/wanhello/omgind/internal/app/schema"
	"github.com/wanhello/omgind/internal/gen/ent"

)

// MenuActionResourceSet 注入MenuActionResource
var MenuActionResourceSet = wire.NewSet(wire.Struct(new(MenuActionResource), "*"))

// MenuActionResource 菜单动作关联资源存储
type MenuActionResource struct {
	EntCli *ent.Client
	TxCli *ent.Tx
}


func (a *MenuActionResource) toSchemaSysMenuActionResource(ma *ent.SysMenuActionResource) *schema.MenuActionResource {
	item := new(schema.MenuActionResource)
	structure.Copy(ma, item)
	return item
}

func (a *MenuActionResource) toSchemaSysMenuActionResources(mas ent.SysMenuActionResources) []*schema.
	MenuActionResource {
	list := make([]*schema.MenuActionResource, len(mas))
	for i, item := range mas {
		list[i] = a.toSchemaSysMenuActionResource(item)
	}
	return list
}


func (a *MenuActionResource) ToEntCreateSysMenuActionInput(ma *schema.MenuActionResource) *ent.CreateSysMenuActionResourceInput {
	createinput := new(ent.CreateSysMenuActionResourceInput)
	structure.Copy(ma, &createinput)

	return createinput
}

func (a *MenuActionResource) ToEntUpdateSysMenuActionInput(ma *schema.MenuActionResource) *ent.UpdateSysMenuActionResourceInput {
	updateinput := new(ent.UpdateSysMenuActionResourceInput)
	structure.Copy(ma, &updateinput)

	return updateinput
}


func (a *MenuActionResource) getQueryOption(opts ...schema.MenuActionResourceQueryOptions) schema.MenuActionResourceQueryOptions {
	var opt schema.MenuActionResourceQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *MenuActionResource) Query(ctx context.Context, params schema.MenuActionResourceQueryParam, opts ...schema.MenuActionResourceQueryOptions) (*schema.MenuActionResourceQueryResult, error) {
	opt := a.getQueryOption(opts...)

	query := a.EntCli.SysMenuActionResource.Query()

	if v := params.MenuID; v != "" {
		// TODO::

		//subQuery := entity.GetMenuActionDB(ctx, a.DB).
		//	Where("menu_id=?", v).
		//	Select("id").SubQuery()
		//db = db.Where("action_id IN ?", subQuery)

	}
	if v := params.MenuIDs; len(v) > 0 {
		// TODO::
		//subQuery := entity.GetMenuActionDB(ctx, a.DB).Where("menu_id IN (?)", v).Select("id").SubQuery()
		//db = db.Where("action_id IN ?", subQuery)
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// get total
	pr := &schema.PaginationResult{Total: count}
	if params.PaginationParam.OnlyCount {
		return &schema.MenuActionResourceQueryResult{PageResult: pr}, nil
	}

	opt.OrderFields = append(opt.OrderFields, schema.NewOrderField("id", schema.OrderByASC))
	query = query.Order(ParseOrder(opt.OrderFields)...)

	pr.Current = params.PaginationParam.GetCurrent()
	pr.PageSize = params.PaginationParam.GetPageSize()
	if params.Offset() > count {
		return &schema.MenuActionResourceQueryResult{PageResult: pr}, nil
	}
	query = query.Limit(params.Limit()).Offset(params.Offset())

	list, err := query.All(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	rlist := ent.SysMenuActionResources(list)

	qr := &schema.MenuActionResourceQueryResult{
		PageResult: pr,
		Data:       a.toSchemaSysMenuActionResources(rlist),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *MenuActionResource) Get(ctx context.Context, id string, opts ...schema.MenuActionResourceQueryOptions) (*schema.MenuActionResource, error) {
	
	sys_mar, err := a.EntCli.SysMenuActionResource.Query().Where(sysmenuactionresource.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return a.toSchemaSysMenuActionResource(sys_mar), nil
}

// Create 创建数据
func (a *MenuActionResource) Create(ctx context.Context, item schema.MenuActionResource) error {
	eitem := entity.SchemaMenuActionResource(item).ToMenuActionResource()
	result := entity.GetMenuActionResourceDB(ctx, a.DB).Create(eitem)
	return errors.WithStack(result.Error)
}

// Update 更新数据
func (a *MenuActionResource) Update(ctx context.Context, id string, item schema.MenuActionResource) error {
	eitem := entity.SchemaMenuActionResource(item).ToMenuActionResource()
	result := entity.GetMenuActionResourceDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
	return errors.WithStack(result.Error)
}

// Delete 删除数据
func (a *MenuActionResource) Delete(ctx context.Context, id string) error {
	result := entity.GetMenuActionResourceDB(ctx, a.DB).Where("id=?", id).Delete(entity.MenuActionResource{})
	return errors.WithStack(result.Error)
}

// DeleteByActionID 根据动作ID删除数据
func (a *MenuActionResource) DeleteByActionID(ctx context.Context, actionID string) error {
	result := entity.GetMenuActionResourceDB(ctx, a.DB).Where("action_id =?", actionID).Delete(entity.MenuActionResource{})
	return errors.WithStack(result.Error)
}

// DeleteByMenuID 根据菜单ID删除数据
func (a *MenuActionResource) DeleteByMenuID(ctx context.Context, menuID string) error {
	subQuery := entity.GetMenuActionDB(ctx, a.DB).Where("menu_id=?", menuID).Select("id").SubQuery()
	result := entity.GetMenuActionResourceDB(ctx, a.DB).Where("action_id IN ?", subQuery).Delete(entity.MenuActionResource{})
	return errors.WithStack(result.Error)
}
