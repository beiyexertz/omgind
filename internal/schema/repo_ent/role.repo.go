package repo_ent

import (
	"context"
	"time"

	"github.com/google/wire"
	"github.com/wanhello/omgind/internal/app/schema"
	"github.com/wanhello/omgind/internal/gen/ent"
	"github.com/wanhello/omgind/internal/gen/ent/sysrole"
	"github.com/wanhello/omgind/pkg/errors"
	"github.com/wanhello/omgind/pkg/helper/structure"
)

// RoleSet 注入Role
var RoleSet = wire.NewSet(wire.Struct(new(Role), "*"))

// Role 角色存储
type Role struct {
	EntCli *ent.Client
}


func (a *Role) toSchemaRole(role *ent.SysRole) *schema.Role {
	item := new(schema.Role)
	structure.Copy(role, item)
	return item
}

func (a *Role) toSchemaRoles(roles ent.SysRoles) []*schema.Role {
	list := make([]*schema.Role, len(roles))
	for i, item := range roles {
		list[i] = a.toSchemaRole(item)
	}
	return list
}


func (a *Role) ToEntCreateSysRoleInput(schrole *schema.Role) *ent.CreateSysRoleInput {
	createinput := new(ent.CreateSysRoleInput)
	structure.Copy(schrole, &createinput)

	return createinput
}

func (a *Role) ToEntUpdateSysRoleInput(schrole *schema.Role) *ent.UpdateSysRoleInput {
	updateinput := new(ent.UpdateSysRoleInput)
	structure.Copy(schrole, &updateinput)

	return updateinput
}


func (a *Role) getQueryOption(opts ...schema.RoleQueryOptions) schema.RoleQueryOptions {
	var opt schema.RoleQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *Role) Query(ctx context.Context, params schema.RoleQueryParam, opts ...schema.RoleQueryOptions) (*schema.RoleQueryResult, error) {

	opt := a.getQueryOption(opts...)

	query := a.EntCli.SysRole.Query()

	if v := params.IDs; len(v) > 0 {
		query = query.Where(sysrole.IDIn(v...))
	}
	if v := params.Name; v != "" {
		query = query.Where(sysrole.NameEQ(v))
	}
	if v := params.UserID; v != "" {
		// TODO::  subquery  子查询

		//query = query.Where(func(s *sql.Selector) {
		//	ur_table := sql.Table(sysuserrole.Table)
		//	//var sb strings.Builder
		//	s.Where()
		//})
		
		//subQuery := entity.GetUserRoleDB(ctx, a.DB).
		//	Where("deleted_at is null").
		//	Where("user_id=?", v).
		//	Select("role_id").SubQuery()
		//db = db.Where("id IN ?", subQuery)
	}
	if v := params.QueryValue; v != "" {
		query = query.Where(sysrole.Or(sysrole.NameContains(v), sysrole.MemoContains(v)))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// get total
	pr := &schema.PaginationResult{Total: count}

	if params.PaginationParam.OnlyCount {
		return &schema.RoleQueryResult{PageResult: pr}, nil
	}

	opt.OrderFields = append(opt.OrderFields, schema.NewOrderField("id", schema.OrderByDESC))
	query = query.Order(ParseOrder(opt.OrderFields)...)

	pr.Current = params.PaginationParam.GetCurrent()
	pr.PageSize = params.PaginationParam.GetPageSize()

	if params.Offset() > count {
		return &schema.RoleQueryResult{PageResult: pr}, nil
	}

	query = query.Limit(params.Limit()).Offset(params.Offset())


	list, err1 := query.All(ctx)
	if err1 != nil {
		return nil, errors.WithStack(err)
	}
	rlist := ent.SysRoles(list)

	qr := &schema.RoleQueryResult{
		PageResult: pr,
		Data:       a.toSchemaRoles(rlist),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *Role) Get(ctx context.Context, id string, opts ...schema.RoleQueryOptions) (*schema.Role, error) {

	role, err := a.EntCli.SysRole.Query().Where(sysrole.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return a.toSchemaRole(role), nil
}

// Create 创建数据
func (a *Role) Create(ctx context.Context, item schema.Role) (*schema.Role, error) {
	//item.CreatedAt = time.Now()
	//item.UpdatedAt = time.Now()

	iteminput := a.ToEntCreateSysRoleInput(&item)
	iteminput.CreatedAt = nil
	iteminput.UpdatedAt = nil

	role, err := a.EntCli.SysRole.Create().SetInput(*iteminput).Save(ctx)

	if err != nil {
		return nil, err
	}

	sch_role := a.toSchemaRole(role)
	return sch_role, nil

}

// Update 更新数据
func (a *Role) Update(ctx context.Context, id string, item schema.Role) (*schema.Role,error) {

	oitem, err := a.EntCli.SysRole.Query().Where(sysrole.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, errors.ErrNotFound
	}

	//item.UpdatedAt = time.Now()
	itemInput := a.ToEntUpdateSysRoleInput(&item)
	itemInput.UpdatedAt = nil
	role, err := oitem.Update().SetInput(*itemInput).Save(ctx)
	if err != nil {
		return nil, err
	}

	sch_role := a.toSchemaRole(role)
	return sch_role, nil
}

// Delete 删除数据
func (a *Role) Delete(ctx context.Context, id string) error {
	role, err := a.EntCli.SysRole.Query().Where(sysrole.IDEQ(id)).Only(ctx)
	if err != nil {
		return errors.ErrNotFound
	}
	_, err1 := role.Update().SetDeletedAt(time.Now()).Save(ctx)
	return errors.WithStack(err1)

}

// UpdateStatus 更新状态
func (a *Role) UpdateStatus(ctx context.Context, id string, status int) error {
	_, err1 := a.EntCli.SysRole.UpdateOneID(id).SetStatus(status).Save(ctx)
	return errors.WithStack(err1)
}
