package repo

import (
	"context"
	"time"

	"github.com/google/wire"
	"github.com/wanhello/omgind/pkg/helper/structure"

	"github.com/wanhello/omgind/internal/app/schema"
	"github.com/wanhello/omgind/internal/gen/ent"
	"github.com/wanhello/omgind/internal/gen/ent/sysmenuaction"
	"github.com/wanhello/omgind/pkg/errors"
)

// MenuActionSet 注入MenuAction
var MenuActionSet = wire.NewSet(wire.Struct(new(MenuAction), "*"))

// MenuAction 菜单动作存储
type MenuAction struct {
	EntCli *ent.Client
	//TxCli *ent.Tx
}

func (a *MenuAction) toSchemaSysMenuAction(ma *ent.SysMenuAction) *schema.MenuAction {
	item := new(schema.MenuAction)
	structure.Copy(ma, item)
	return item
}

func (a *MenuAction) toSchemaSysMenuActions(mas ent.SysMenuActions) []*schema.MenuAction {
	list := make([]*schema.MenuAction, len(mas))
	for i, item := range mas {
		list[i] = a.toSchemaSysMenuAction(item)
	}
	return list
}


func (a *MenuAction) ToEntCreateSysMenuActionInput(ma *schema.MenuAction) *ent.CreateSysMenuActionInput {
	createinput := new(ent.CreateSysMenuActionInput)
	structure.Copy(ma, &createinput)

	return createinput
}

func (a *MenuAction) ToEntUpdateSysMenuActionInput(ma *schema.MenuAction) *ent.UpdateSysMenuActionInput {
	updateinput := new(ent.UpdateSysMenuActionInput)
	structure.Copy(ma, &updateinput)

	return updateinput
}


func (a *MenuAction) getQueryOption(opts ...schema.MenuActionQueryOptions) schema.MenuActionQueryOptions {
	var opt schema.MenuActionQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *MenuAction) Query(ctx context.Context, params schema.MenuActionQueryParam, opts ...schema.MenuActionQueryOptions) (*schema.MenuActionQueryResult, error) {
	opt := a.getQueryOption(opts...)

	query := a.EntCli.SysMenuAction.Query().Where(sysmenuaction.DeletedAtIsNil())

	if v := params.MenuID; v != "" {
		query = query.Where(sysmenuaction.MenuIDEQ(v))
	}
	if v := params.IDs; len(v) > 0 {
		query = query.Where(sysmenuaction.IDIn(v...))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// get total
	pr := &schema.PaginationResult{Total: count}
	if params.PaginationParam.OnlyCount {
		return &schema.MenuActionQueryResult{PageResult: pr}, nil
	}

	opt.OrderFields = append(opt.OrderFields, schema.NewOrderField("id", schema.OrderByASC))
	query = query.Order(ParseOrder(opt.OrderFields)...)

	pr.Current = params.PaginationParam.GetCurrent()
	pr.PageSize = params.PaginationParam.GetPageSize()
	if params.Offset() > count {
		return &schema.MenuActionQueryResult{PageResult: pr}, nil
	}

	query = query.Limit(params.Limit()).Offset(params.Offset())

	list, err := query.All(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	rlist := ent.SysMenuActions(list)

	qr := &schema.MenuActionQueryResult{
		PageResult: pr,
		Data:       a.toSchemaSysMenuActions(rlist),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *MenuAction) Get(ctx context.Context, id string, opts ...schema.MenuActionQueryOptions) (*schema.MenuAction, error) {

	menu_action, err := a.EntCli.SysMenuAction.Query().Where(sysmenuaction.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}

	return a.toSchemaSysMenuAction(menu_action), nil
}

// Create 创建数据
func (a *MenuAction) Create(ctx context.Context, item schema.MenuAction) (*schema.MenuAction, error) {

	iteminput := a.ToEntCreateSysMenuActionInput(&item)

	sys_ma, err := a.EntCli.SysMenuAction.Create().SetInput(*iteminput).Save(ctx)

	if err != nil {
		return nil, err
	}
	sch_ma := a.toSchemaSysMenuAction(sys_ma)
	return sch_ma, nil
}

// Update 更新数据
func (a *MenuAction) Update(ctx context.Context, id string, item schema.MenuAction) (*schema.MenuAction, error) {
	oitem, err := a.EntCli.SysMenuAction.Query().Where(sysmenuaction.IDEQ(id)).Only(ctx)

	if err != nil {
		return nil, err
	}

	iteminput := a.ToEntUpdateSysMenuActionInput(&item)
	sys_ma, err := oitem.Update().SetInput(*iteminput).Save(ctx)
	if err != nil {
		return nil, err
	}
	sch_ma := a.toSchemaSysMenuAction(sys_ma)

	return sch_ma, nil
}

// Delete 删除数据
func (a *MenuAction) Delete(ctx context.Context, id string) error {

	sys_ma, err := a.EntCli.SysMenuAction.Query().Where(sysmenuaction.IDEQ(id)).Only(ctx)

	if err != nil {
		return err
	}
	_, err1 := sys_ma.Update().SetDeletedAt(time.Now()).Save(ctx)
	return errors.WithStack(err1)
}

// DeleteByMenuID 根据菜单ID删除数据
func (a *MenuAction) DeleteByMenuID(ctx context.Context, menuID string) error {
	_, err := a.EntCli.SysMenuAction.Update().Where(sysmenuaction.MenuIDEQ(menuID)).SetDeletedAt(time.Now()).Save(ctx)

	return errors.WithStack(err)
}
