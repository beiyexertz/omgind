package repo_ent

import (
	"context"
	"time"

	"github.com/google/wire"
	"github.com/wanhello/omgind/internal/gen/ent/sysuser"
	"github.com/wanhello/omgind/pkg/helper/structure"

	"github.com/wanhello/omgind/internal/app/schema"
	"github.com/wanhello/omgind/internal/gen/ent"
	"github.com/wanhello/omgind/pkg/errors"
)

// UserSet 注入User
var UserSet = wire.NewSet(wire.Struct(new(User), "*"))

// User 用户存储
type User struct {
	EntCli *ent.Client
}

func (a *User) toSchemaSysUser(u *ent.SysUser) *schema.User {
	item := new(schema.User)
	structure.Copy(u, item)
	return item
}

func (a *User) toSchemaSysUsers(us ent.SysUsers) []*schema.User {
	list := make([]*schema.User, len(us))
	for i, item := range us {
		list[i] = a.toSchemaSysUser(item)
	}
	return list
}

func (a *User) ToEntCreateSysUserInput(user *schema.User) *ent.CreateSysUserInput {
	createinput := new(ent.CreateSysUserInput)
	structure.Copy(user, &createinput)

	return createinput
}

func (a *User) ToEntUpdateSysUserInput(user *schema.User) *ent.UpdateSysUserInput {
	updateinput := new(ent.UpdateSysUserInput)
	structure.Copy(user, &updateinput)

	return updateinput
}


func (a *User) getQueryOption(opts ...schema.UserQueryOptions) schema.UserQueryOptions {
	var opt schema.UserQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *User) Query(ctx context.Context, params schema.UserQueryParam, opts ...schema.UserQueryOptions) (*schema.UserQueryResult, error) {
	opt := a.getQueryOption(opts...)

	query := a.EntCli.SysUser.Query().Where(sysuser.DeletedAtIsNil() )

	if v := params.UserName; v != "" {
		query = query.Where(sysuser.UserNameEQ(v))
	}
	if v := params.Status; v > 0 {
		query = query.Where(sysuser.StatusEQ(v))
	}
	if v := params.RoleIDs; len(v) > 0 {
		// TODO::  subquery  子查询

		//subQuery := entity.GetUserRoleDB(ctx, a.DB).
		//	Select("user_id").
		//	Where("role_id IN (?)", v).
		//	SubQuery()
		//db = db.Where("id IN ?", subQuery)
	}
	if v := params.QueryValue; v != "" {
		query = query.Where(sysuser.Or(
			sysuser.UserNameContains(v), sysuser.RealNameContains(v),
			sysuser.PhoneContains(v), sysuser.EmailContains(v),
			))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// get total
	pr := &schema.PaginationResult{Total: count}
	if params.PaginationParam.OnlyCount {
		return &schema.UserQueryResult{PageResult: pr}, nil
	}

	opt.OrderFields = append(opt.OrderFields, schema.NewOrderField("id", schema.OrderByDESC))
	query = query.Order(ParseOrder(opt.OrderFields)...)

	pr.Current = params.PaginationParam.GetCurrent()
	pr.PageSize = params.PaginationParam.GetPageSize()
	if params.Offset() > count {
		return &schema.UserQueryResult{PageResult: pr}, nil
	}

	query = query.Limit(params.Limit()).Offset(params.Offset())

	list, err1 := query.All(ctx)
	if err1 != nil {
		return nil, errors.WithStack(err)
	}
	rlist := ent.SysUsers(list)

	qr := &schema.UserQueryResult{
		PageResult: pr,
		Data:       a.toSchemaSysUsers(rlist),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *User) Get(ctx context.Context, id string, opts ...schema.UserQueryOptions) (*schema.User, error) {
	user, err := a.EntCli.SysUser.Query().Where(sysuser.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return a.toSchemaSysUser(user), nil
}

// Create 创建数据
func (a *User) Create(ctx context.Context, item schema.User) (*schema.User, error) {

	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()

	iteminput := a.ToEntCreateSysUserInput(&item)
	iteminput.CreatedAt = nil
	iteminput.UpdatedAt = nil

	sysuser, err := a.EntCli.SysUser.Create().SetInput(*iteminput).Save(ctx)

	if err != nil {
		return nil, err
	}
	sch_user := a.toSchemaSysUser(sysuser)
	return sch_user, nil
}

// Update 更新数据
func (a *User) Update(ctx context.Context, id string, item schema.User) (*schema.User, error) {

	oitem, err := a.EntCli.SysUser.Query().Where(sysuser.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}

	item.UpdatedAt = time.Now()
	iteminput := a.ToEntUpdateSysUserInput(&item)
	user, err := oitem.Update().SetInput(*iteminput).Save(ctx)
	sch_user := a.toSchemaSysUser(user)
	return sch_user, nil
}

// Delete 删除数据
func (a *User) Delete(ctx context.Context, id string) error {
	role, err := a.EntCli.SysUser.Query().Where(sysuser.IDEQ(id)).Only(ctx)

	if err != nil {
		return err
	}
	_, err1 := role.Update().SetDeletedAt(time.Now()).Save(ctx)
	return errors.WithStack(err1)
}

// UpdateStatus 更新状态
func (a *User) UpdateStatus(ctx context.Context, id string, status int) error {
	_, err := a.EntCli.SysUser.UpdateOneID(id).SetStatus(status).Save(ctx)
	return errors.WithStack(err)
}

// UpdatePassword 更新密码
func (a *User) UpdatePassword(ctx context.Context, id, password string) error {
	_, err := a.EntCli.SysUser.UpdateOneID(id).SetPassword(password).Save(ctx)
	return errors.WithStack(err)
}
