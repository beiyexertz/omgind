package repo

import (
	"context"

	"github.com/wanhello/omgind/internal/app/model/gormx/entity"
	"github.com/wanhello/omgind/internal/app/schema"
	"github.com/wanhello/omgind/pkg/errors"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

// UserRoleSet 注入UserRole
var UserRoleSet = wire.NewSet(wire.Struct(new(UserRole), "*"))

// UserRole 用户角色存储
type UserRole struct {
	DB *gorm.DB
}

func (a *UserRole) getQueryOption(opts ...schema.UserRoleQueryOptions) schema.UserRoleQueryOptions {
	var opt schema.UserRoleQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *UserRole) Query(ctx context.Context, params schema.UserRoleQueryParam, opts ...schema.UserRoleQueryOptions) (*schema.UserRoleQueryResult, error) {
	opt := a.getQueryOption(opts...)

	db := entity.GetUserRoleDB(ctx, a.DB)
	if v := params.UserID; v != "" {
		db = db.Where("user_id=?", v)
	}
	if v := params.UserIDs; len(v) > 0 {
		db = db.Where("user_id IN (?)", v)
	}

	opt.OrderFields = append(opt.OrderFields, schema.NewOrderField("id", schema.OrderByDESC))
	db = db.Order(ParseOrder(opt.OrderFields))

	var list entity.UserRoles
	pr, err := WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	qr := &schema.UserRoleQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaUserRoles(),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *UserRole) Get(ctx context.Context, id string, opts ...schema.UserRoleQueryOptions) (*schema.UserRole, error) {
	db := entity.GetUserRoleDB(ctx, a.DB).Where("id=?", id)
	var item entity.UserRole
	ok, err := FindOne(ctx, db, &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaUserRole(), nil
}

// Create 创建数据
func (a *UserRole) Create(ctx context.Context, item schema.UserRole) error {
	eitem := entity.SchemaUserRole(item).ToUserRole()
	result := entity.GetUserRoleDB(ctx, a.DB).Create(eitem)
	return errors.WithStack(result.Error)
}

// Update 更新数据
func (a *UserRole) Update(ctx context.Context, id string, item schema.UserRole) error {
	eitem := entity.SchemaUserRole(item).ToUserRole()
	result := entity.GetUserRoleDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
	return errors.WithStack(result.Error)
}

// Delete 删除数据
func (a *UserRole) Delete(ctx context.Context, id string) error {
	result := entity.GetUserRoleDB(ctx, a.DB).Where("id=?", id).Delete(entity.UserRole{})
	return errors.WithStack(result.Error)
}

// DeleteByUserID 根据用户ID删除数据
func (a *UserRole) DeleteByUserID(ctx context.Context, userID string) error {
	result := entity.GetUserRoleDB(ctx, a.DB).Where("user_id=?", userID).Delete(entity.UserRole{})
	return errors.WithStack(result.Error)
}
