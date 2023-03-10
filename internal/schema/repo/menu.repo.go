package repo

import (
	"context"
	"time"

	"github.com/google/wire"
	"github.com/wanhello/omgind/internal/gen/ent"
	"github.com/wanhello/omgind/internal/gen/ent/sysmenu"
	"github.com/wanhello/omgind/pkg/errors"
	"github.com/wanhello/omgind/pkg/helper/structure"

	"github.com/wanhello/omgind/internal/app/schema"
)

// MenuSet 注入Menu
var MenuSet = wire.NewSet(wire.Struct(new(Menu), "*"))

// Menu 菜单存储
type Menu struct {
	EntCli *ent.Client
}

func (a *Menu) toSchemaSysMenu(dit *ent.SysMenu) *schema.Menu {
	item := new(schema.Menu)
	structure.Copy(dit, item)
	return item
}

func (a *Menu) toSchemaSysMenus(dits ent.SysMenus) []*schema.Menu {
	list := make([]*schema.Menu, len(dits))
	for i, item := range dits {
		list[i] = a.toSchemaSysMenu(item)
	}
	return list
}

func (a *Menu) ToEntCreateSysMenuInput(m *schema.Menu) *ent.CreateSysMenuInput {
	createinput := new(ent.CreateSysMenuInput)
	structure.Copy(m, &createinput)

	return createinput
}

func (a *Menu) ToEntUpdateSysMenuInput(m *schema.Menu) *ent.UpdateSysMenuInput {
	updateinput := new(ent.UpdateSysMenuInput)
	structure.Copy(m, &updateinput)

	return updateinput
}


func (a *Menu) getQueryOption(opts ...schema.MenuQueryOptions) schema.MenuQueryOptions {
	var opt schema.MenuQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *Menu) Query(ctx context.Context, params schema.MenuQueryParam, opts ...schema.MenuQueryOptions) (*schema.MenuQueryResult, error) {
	opt := a.getQueryOption(opts...)

	query := a.EntCli.SysMenu.Query().Where(sysmenu.DeletedAtIsNil())

	if v := params.IDs; len(v) > 0 {
		query = query.Where(sysmenu.IDIn(v...))
	}
	if v := params.Name; v != "" {
		query = query.Where(sysmenu.NameEQ(v))
	}
	if v := params.ParentID; v != nil {
		query = query.Where(sysmenu.ParentIDEQ(*v))
	}

	if v := params.PrefixParentPath; v != "" {
		query = query.Where(sysmenu.ParentPathContains(v))
	}
	if v := params.Status; v != 0 {
		query = query.Where(sysmenu.Status(v))
	}
	if v := params.IsShow; v != nil {
		query = query.Where(sysmenu.IsShowEQ(*v))
	}

	if v := params.QueryValue; v != "" {
		query = query.Where(sysmenu.Or(
			sysmenu.NameContains(v),
			sysmenu.MemoContains(v),
			))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// get total
	pr := &schema.PaginationResult{Total: count}
	if params.PaginationParam.OnlyCount {
		return &schema.MenuQueryResult{PageResult: pr}, nil
	}

	opt.OrderFields = append(opt.OrderFields, schema.NewOrderField("id", schema.OrderByDESC))
	query = query.Order(ParseOrder(opt.OrderFields)...)

	pr.Current = params.PaginationParam.GetCurrent()
	pr.PageSize = params.PaginationParam.GetPageSize()
	if params.Offset() > count {
		return &schema.MenuQueryResult{PageResult: pr}, nil
	}

	query = query.Limit(params.Limit()).Offset(params.Offset())

	list, err := query.All(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	rlist := ent.SysMenus(list)

	qr := &schema.MenuQueryResult{
		PageResult: pr,
		Data:       a.toSchemaSysMenus(rlist),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *Menu) Get(ctx context.Context, id string, opts ...schema.MenuQueryOptions) (*schema.Menu, error) {

	menu, err := a.EntCli.SysMenu.Query().Where(sysmenu.IDEQ(id)).Only(ctx)

	if err != nil {
		return nil, err
	}
	return a.toSchemaSysMenu(menu), nil
}

// Create 创建数据
func (a *Menu) Create(ctx context.Context, item schema.Menu) (*schema.Menu, error) {

	iteminput := a.ToEntCreateSysMenuInput(&item)
	iteminput.CreatedAt = nil
	iteminput.UpdatedAt = nil
	sys_menu, err := a.EntCli.SysMenu.Create().SetInput(*iteminput).Save(ctx)
	if err != nil {
		return nil, err
	}
	sch_menu := a.toSchemaSysMenu(sys_menu)
	return sch_menu, nil
}

// Update 更新数据
func (a *Menu) Update(ctx context.Context, id string, item schema.Menu) (*schema.Menu,error) {

	oitem, err := a.EntCli.SysMenu.Query().Where(sysmenu.IDEQ(id)).Only(ctx)

	if err != nil {
		return nil, err
	}

	iteminput := a.ToEntUpdateSysMenuInput(&item)
	sys_menu, err := oitem.Update().SetInput(*iteminput).Save(ctx)
	if err != nil {
		return nil, err
	}
	sch_menu := a.toSchemaSysMenu(sys_menu)
	return sch_menu, nil
}

// UpdateParentPath 更新父级路径
func (a *Menu) UpdateParentPath(ctx context.Context, id, parentPath string) error {
	_, err := a.EntCli.SysMenu.UpdateOneID(id).SetParentPath(parentPath).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Delete 删除数据
func (a *Menu) Delete(ctx context.Context, id string) error {
	sys_menu, err := a.EntCli.SysMenu.Query().Where(sysmenu.IDEQ(id)).Only(ctx)
	if err != nil {
		return err
	}
	_, err1 := sys_menu.Update().SetDeletedAt(time.Now()).Save(ctx)
	return errors.WithStack(err1)
}

// UpdateStatus 更新状态
func (a *Menu) UpdateStatus(ctx context.Context, id string, status int16) error {
	_, err := a.EntCli.SysMenu.UpdateOneID(id).SetStatus(status).Save(ctx)
	return errors.WithStack(err)
}

