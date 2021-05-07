// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/wanhello/omgind/internal/gen/ent/syscasbinrule"
	"github.com/wanhello/omgind/internal/gen/ent/sysdict"
	"github.com/wanhello/omgind/internal/gen/ent/sysdictitem"
	"github.com/wanhello/omgind/internal/gen/ent/sysmenu"
	"github.com/wanhello/omgind/internal/gen/ent/sysmenuaction"
	"github.com/wanhello/omgind/internal/gen/ent/sysuser"
	"github.com/wanhello/omgind/internal/schema/entity"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	syscasbinruleMixin := entity.SysCasbinRule{}.Mixin()
	syscasbinruleMixinFields0 := syscasbinruleMixin[0].Fields()
	_ = syscasbinruleMixinFields0
	syscasbinruleMixinFields1 := syscasbinruleMixin[1].Fields()
	_ = syscasbinruleMixinFields1
	syscasbinruleMixinFields2 := syscasbinruleMixin[2].Fields()
	_ = syscasbinruleMixinFields2
	syscasbinruleFields := entity.SysCasbinRule{}.Fields()
	_ = syscasbinruleFields
	// syscasbinruleDescIsDel is the schema descriptor for is_del field.
	syscasbinruleDescIsDel := syscasbinruleMixinFields0[1].Descriptor()
	// syscasbinrule.DefaultIsDel holds the default value on creation for the is_del field.
	syscasbinrule.DefaultIsDel = syscasbinruleDescIsDel.Default.(bool)
	// syscasbinruleDescSort is the schema descriptor for sort field.
	syscasbinruleDescSort := syscasbinruleMixinFields1[0].Descriptor()
	// syscasbinrule.DefaultSort holds the default value on creation for the sort field.
	syscasbinrule.DefaultSort = syscasbinruleDescSort.Default.(int32)
	// syscasbinruleDescCreatedAt is the schema descriptor for created_at field.
	syscasbinruleDescCreatedAt := syscasbinruleMixinFields2[0].Descriptor()
	// syscasbinrule.DefaultCreatedAt holds the default value on creation for the created_at field.
	syscasbinrule.DefaultCreatedAt = syscasbinruleDescCreatedAt.Default.(func() time.Time)
	// syscasbinruleDescUpdatedAt is the schema descriptor for updated_at field.
	syscasbinruleDescUpdatedAt := syscasbinruleMixinFields2[1].Descriptor()
	// syscasbinrule.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	syscasbinrule.DefaultUpdatedAt = syscasbinruleDescUpdatedAt.Default.(func() time.Time)
	// syscasbinrule.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	syscasbinrule.UpdateDefaultUpdatedAt = syscasbinruleDescUpdatedAt.UpdateDefault.(func() time.Time)
	// syscasbinruleDescPType is the schema descriptor for PType field.
	syscasbinruleDescPType := syscasbinruleFields[0].Descriptor()
	// syscasbinrule.PTypeValidator is a validator for the "PType" field. It is called by the builders before save.
	syscasbinrule.PTypeValidator = syscasbinruleDescPType.Validators[0].(func(string) error)
	// syscasbinruleDescRoleID is the schema descriptor for RoleID field.
	syscasbinruleDescRoleID := syscasbinruleFields[1].Descriptor()
	// syscasbinrule.RoleIDValidator is a validator for the "RoleID" field. It is called by the builders before save.
	syscasbinrule.RoleIDValidator = syscasbinruleDescRoleID.Validators[0].(func(string) error)
	// syscasbinruleDescPath is the schema descriptor for Path field.
	syscasbinruleDescPath := syscasbinruleFields[2].Descriptor()
	// syscasbinrule.PathValidator is a validator for the "Path" field. It is called by the builders before save.
	syscasbinrule.PathValidator = syscasbinruleDescPath.Validators[0].(func(string) error)
	// syscasbinruleDescMethod is the schema descriptor for Method field.
	syscasbinruleDescMethod := syscasbinruleFields[3].Descriptor()
	// syscasbinrule.MethodValidator is a validator for the "Method" field. It is called by the builders before save.
	syscasbinrule.MethodValidator = syscasbinruleDescMethod.Validators[0].(func(string) error)
	// syscasbinruleDescV3 is the schema descriptor for v3 field.
	syscasbinruleDescV3 := syscasbinruleFields[4].Descriptor()
	// syscasbinrule.V3Validator is a validator for the "v3" field. It is called by the builders before save.
	syscasbinrule.V3Validator = syscasbinruleDescV3.Validators[0].(func(string) error)
	// syscasbinruleDescV4 is the schema descriptor for v4 field.
	syscasbinruleDescV4 := syscasbinruleFields[5].Descriptor()
	// syscasbinrule.V4Validator is a validator for the "v4" field. It is called by the builders before save.
	syscasbinrule.V4Validator = syscasbinruleDescV4.Validators[0].(func(string) error)
	// syscasbinruleDescV5 is the schema descriptor for v5 field.
	syscasbinruleDescV5 := syscasbinruleFields[6].Descriptor()
	// syscasbinrule.V5Validator is a validator for the "v5" field. It is called by the builders before save.
	syscasbinrule.V5Validator = syscasbinruleDescV5.Validators[0].(func(string) error)
	// syscasbinruleDescID is the schema descriptor for id field.
	syscasbinruleDescID := syscasbinruleMixinFields0[0].Descriptor()
	// syscasbinrule.DefaultID holds the default value on creation for the id field.
	syscasbinrule.DefaultID = syscasbinruleDescID.Default.(string)
	// syscasbinrule.IDValidator is a validator for the "id" field. It is called by the builders before save.
	syscasbinrule.IDValidator = func() func(string) error {
		validators := syscasbinruleDescID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(id string) error {
			for _, fn := range fns {
				if err := fn(id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	sysdictMixin := entity.SysDict{}.Mixin()
	sysdictMixinFields0 := sysdictMixin[0].Fields()
	_ = sysdictMixinFields0
	sysdictMixinFields1 := sysdictMixin[1].Fields()
	_ = sysdictMixinFields1
	sysdictMixinFields2 := sysdictMixin[2].Fields()
	_ = sysdictMixinFields2
	sysdictMixinFields3 := sysdictMixin[3].Fields()
	_ = sysdictMixinFields3
	sysdictFields := entity.SysDict{}.Fields()
	_ = sysdictFields
	// sysdictDescIsDel is the schema descriptor for is_del field.
	sysdictDescIsDel := sysdictMixinFields0[1].Descriptor()
	// sysdict.DefaultIsDel holds the default value on creation for the is_del field.
	sysdict.DefaultIsDel = sysdictDescIsDel.Default.(bool)
	// sysdictDescMemo is the schema descriptor for memo field.
	sysdictDescMemo := sysdictMixinFields1[0].Descriptor()
	// sysdict.DefaultMemo holds the default value on creation for the memo field.
	sysdict.DefaultMemo = sysdictDescMemo.Default.(string)
	// sysdict.MemoValidator is a validator for the "memo" field. It is called by the builders before save.
	sysdict.MemoValidator = sysdictDescMemo.Validators[0].(func(string) error)
	// sysdictDescSort is the schema descriptor for sort field.
	sysdictDescSort := sysdictMixinFields2[0].Descriptor()
	// sysdict.DefaultSort holds the default value on creation for the sort field.
	sysdict.DefaultSort = sysdictDescSort.Default.(int32)
	// sysdictDescCreatedAt is the schema descriptor for created_at field.
	sysdictDescCreatedAt := sysdictMixinFields3[0].Descriptor()
	// sysdict.DefaultCreatedAt holds the default value on creation for the created_at field.
	sysdict.DefaultCreatedAt = sysdictDescCreatedAt.Default.(func() time.Time)
	// sysdictDescUpdatedAt is the schema descriptor for updated_at field.
	sysdictDescUpdatedAt := sysdictMixinFields3[1].Descriptor()
	// sysdict.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sysdict.DefaultUpdatedAt = sysdictDescUpdatedAt.Default.(func() time.Time)
	// sysdict.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sysdict.UpdateDefaultUpdatedAt = sysdictDescUpdatedAt.UpdateDefault.(func() time.Time)
	// sysdictDescNameCn is the schema descriptor for name_cn field.
	sysdictDescNameCn := sysdictFields[0].Descriptor()
	// sysdict.NameCnValidator is a validator for the "name_cn" field. It is called by the builders before save.
	sysdict.NameCnValidator = sysdictDescNameCn.Validators[0].(func(string) error)
	// sysdictDescNameEn is the schema descriptor for name_en field.
	sysdictDescNameEn := sysdictFields[1].Descriptor()
	// sysdict.NameEnValidator is a validator for the "name_en" field. It is called by the builders before save.
	sysdict.NameEnValidator = sysdictDescNameEn.Validators[0].(func(string) error)
	// sysdictDescStatus is the schema descriptor for Status field.
	sysdictDescStatus := sysdictFields[2].Descriptor()
	// sysdict.DefaultStatus holds the default value on creation for the Status field.
	sysdict.DefaultStatus = sysdictDescStatus.Default.(bool)
	// sysdictDescID is the schema descriptor for id field.
	sysdictDescID := sysdictMixinFields0[0].Descriptor()
	// sysdict.DefaultID holds the default value on creation for the id field.
	sysdict.DefaultID = sysdictDescID.Default.(string)
	// sysdict.IDValidator is a validator for the "id" field. It is called by the builders before save.
	sysdict.IDValidator = func() func(string) error {
		validators := sysdictDescID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(id string) error {
			for _, fn := range fns {
				if err := fn(id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	sysdictitemMixin := entity.SysDictItem{}.Mixin()
	sysdictitemMixinFields0 := sysdictitemMixin[0].Fields()
	_ = sysdictitemMixinFields0
	sysdictitemMixinFields1 := sysdictitemMixin[1].Fields()
	_ = sysdictitemMixinFields1
	sysdictitemMixinFields2 := sysdictitemMixin[2].Fields()
	_ = sysdictitemMixinFields2
	sysdictitemMixinFields3 := sysdictitemMixin[3].Fields()
	_ = sysdictitemMixinFields3
	sysdictitemFields := entity.SysDictItem{}.Fields()
	_ = sysdictitemFields
	// sysdictitemDescIsDel is the schema descriptor for is_del field.
	sysdictitemDescIsDel := sysdictitemMixinFields0[1].Descriptor()
	// sysdictitem.DefaultIsDel holds the default value on creation for the is_del field.
	sysdictitem.DefaultIsDel = sysdictitemDescIsDel.Default.(bool)
	// sysdictitemDescMemo is the schema descriptor for memo field.
	sysdictitemDescMemo := sysdictitemMixinFields1[0].Descriptor()
	// sysdictitem.DefaultMemo holds the default value on creation for the memo field.
	sysdictitem.DefaultMemo = sysdictitemDescMemo.Default.(string)
	// sysdictitem.MemoValidator is a validator for the "memo" field. It is called by the builders before save.
	sysdictitem.MemoValidator = sysdictitemDescMemo.Validators[0].(func(string) error)
	// sysdictitemDescSort is the schema descriptor for sort field.
	sysdictitemDescSort := sysdictitemMixinFields2[0].Descriptor()
	// sysdictitem.DefaultSort holds the default value on creation for the sort field.
	sysdictitem.DefaultSort = sysdictitemDescSort.Default.(int32)
	// sysdictitemDescCreatedAt is the schema descriptor for created_at field.
	sysdictitemDescCreatedAt := sysdictitemMixinFields3[0].Descriptor()
	// sysdictitem.DefaultCreatedAt holds the default value on creation for the created_at field.
	sysdictitem.DefaultCreatedAt = sysdictitemDescCreatedAt.Default.(func() time.Time)
	// sysdictitemDescUpdatedAt is the schema descriptor for updated_at field.
	sysdictitemDescUpdatedAt := sysdictitemMixinFields3[1].Descriptor()
	// sysdictitem.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sysdictitem.DefaultUpdatedAt = sysdictitemDescUpdatedAt.Default.(func() time.Time)
	// sysdictitem.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sysdictitem.UpdateDefaultUpdatedAt = sysdictitemDescUpdatedAt.UpdateDefault.(func() time.Time)
	// sysdictitemDescLabel is the schema descriptor for label field.
	sysdictitemDescLabel := sysdictitemFields[0].Descriptor()
	// sysdictitem.LabelValidator is a validator for the "label" field. It is called by the builders before save.
	sysdictitem.LabelValidator = sysdictitemDescLabel.Validators[0].(func(string) error)
	// sysdictitemDescID is the schema descriptor for id field.
	sysdictitemDescID := sysdictitemMixinFields0[0].Descriptor()
	// sysdictitem.DefaultID holds the default value on creation for the id field.
	sysdictitem.DefaultID = sysdictitemDescID.Default.(string)
	// sysdictitem.IDValidator is a validator for the "id" field. It is called by the builders before save.
	sysdictitem.IDValidator = func() func(string) error {
		validators := sysdictitemDescID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(id string) error {
			for _, fn := range fns {
				if err := fn(id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	sysmenuMixin := entity.SysMenu{}.Mixin()
	sysmenuMixinFields0 := sysmenuMixin[0].Fields()
	_ = sysmenuMixinFields0
	sysmenuMixinFields1 := sysmenuMixin[1].Fields()
	_ = sysmenuMixinFields1
	sysmenuMixinFields2 := sysmenuMixin[2].Fields()
	_ = sysmenuMixinFields2
	sysmenuMixinFields3 := sysmenuMixin[3].Fields()
	_ = sysmenuMixinFields3
	sysmenuMixinFields4 := sysmenuMixin[4].Fields()
	_ = sysmenuMixinFields4
	sysmenuFields := entity.SysMenu{}.Fields()
	_ = sysmenuFields
	// sysmenuDescIsDel is the schema descriptor for is_del field.
	sysmenuDescIsDel := sysmenuMixinFields0[1].Descriptor()
	// sysmenu.DefaultIsDel holds the default value on creation for the is_del field.
	sysmenu.DefaultIsDel = sysmenuDescIsDel.Default.(bool)
	// sysmenuDescMemo is the schema descriptor for memo field.
	sysmenuDescMemo := sysmenuMixinFields1[0].Descriptor()
	// sysmenu.DefaultMemo holds the default value on creation for the memo field.
	sysmenu.DefaultMemo = sysmenuDescMemo.Default.(string)
	// sysmenu.MemoValidator is a validator for the "memo" field. It is called by the builders before save.
	sysmenu.MemoValidator = sysmenuDescMemo.Validators[0].(func(string) error)
	// sysmenuDescSort is the schema descriptor for sort field.
	sysmenuDescSort := sysmenuMixinFields2[0].Descriptor()
	// sysmenu.DefaultSort holds the default value on creation for the sort field.
	sysmenu.DefaultSort = sysmenuDescSort.Default.(int32)
	// sysmenuDescCreatedAt is the schema descriptor for created_at field.
	sysmenuDescCreatedAt := sysmenuMixinFields3[0].Descriptor()
	// sysmenu.DefaultCreatedAt holds the default value on creation for the created_at field.
	sysmenu.DefaultCreatedAt = sysmenuDescCreatedAt.Default.(func() time.Time)
	// sysmenuDescUpdatedAt is the schema descriptor for updated_at field.
	sysmenuDescUpdatedAt := sysmenuMixinFields3[1].Descriptor()
	// sysmenu.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sysmenu.DefaultUpdatedAt = sysmenuDescUpdatedAt.Default.(func() time.Time)
	// sysmenu.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sysmenu.UpdateDefaultUpdatedAt = sysmenuDescUpdatedAt.UpdateDefault.(func() time.Time)
	// sysmenuDescStatus is the schema descriptor for status field.
	sysmenuDescStatus := sysmenuMixinFields4[0].Descriptor()
	// sysmenu.DefaultStatus holds the default value on creation for the status field.
	sysmenu.DefaultStatus = sysmenuDescStatus.Default.(int32)
	// sysmenuDescName is the schema descriptor for name field.
	sysmenuDescName := sysmenuFields[0].Descriptor()
	// sysmenu.NameValidator is a validator for the "name" field. It is called by the builders before save.
	sysmenu.NameValidator = func() func(string) error {
		validators := sysmenuDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// sysmenuDescIcon is the schema descriptor for icon field.
	sysmenuDescIcon := sysmenuFields[1].Descriptor()
	// sysmenu.IconValidator is a validator for the "icon" field. It is called by the builders before save.
	sysmenu.IconValidator = sysmenuDescIcon.Validators[0].(func(string) error)
	// sysmenuDescRouter is the schema descriptor for router field.
	sysmenuDescRouter := sysmenuFields[2].Descriptor()
	// sysmenu.RouterValidator is a validator for the "router" field. It is called by the builders before save.
	sysmenu.RouterValidator = func() func(string) error {
		validators := sysmenuDescRouter.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(router string) error {
			for _, fn := range fns {
				if err := fn(router); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// sysmenuDescIsShow is the schema descriptor for is_show field.
	sysmenuDescIsShow := sysmenuFields[3].Descriptor()
	// sysmenu.DefaultIsShow holds the default value on creation for the is_show field.
	sysmenu.DefaultIsShow = sysmenuDescIsShow.Default.(bool)
	// sysmenuDescParentID is the schema descriptor for parent_id field.
	sysmenuDescParentID := sysmenuFields[4].Descriptor()
	// sysmenu.ParentIDValidator is a validator for the "parent_id" field. It is called by the builders before save.
	sysmenu.ParentIDValidator = sysmenuDescParentID.Validators[0].(func(string) error)
	// sysmenuDescParentPath is the schema descriptor for parent_path field.
	sysmenuDescParentPath := sysmenuFields[5].Descriptor()
	// sysmenu.ParentPathValidator is a validator for the "parent_path" field. It is called by the builders before save.
	sysmenu.ParentPathValidator = sysmenuDescParentPath.Validators[0].(func(string) error)
	// sysmenuDescID is the schema descriptor for id field.
	sysmenuDescID := sysmenuMixinFields0[0].Descriptor()
	// sysmenu.DefaultID holds the default value on creation for the id field.
	sysmenu.DefaultID = sysmenuDescID.Default.(string)
	// sysmenu.IDValidator is a validator for the "id" field. It is called by the builders before save.
	sysmenu.IDValidator = func() func(string) error {
		validators := sysmenuDescID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(id string) error {
			for _, fn := range fns {
				if err := fn(id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	sysmenuactionMixin := entity.SysMenuAction{}.Mixin()
	sysmenuactionMixinFields0 := sysmenuactionMixin[0].Fields()
	_ = sysmenuactionMixinFields0
	sysmenuactionMixinFields1 := sysmenuactionMixin[1].Fields()
	_ = sysmenuactionMixinFields1
	sysmenuactionMixinFields2 := sysmenuactionMixin[2].Fields()
	_ = sysmenuactionMixinFields2
	sysmenuactionMixinFields3 := sysmenuactionMixin[3].Fields()
	_ = sysmenuactionMixinFields3
	sysmenuactionMixinFields4 := sysmenuactionMixin[4].Fields()
	_ = sysmenuactionMixinFields4
	sysmenuactionFields := entity.SysMenuAction{}.Fields()
	_ = sysmenuactionFields
	// sysmenuactionDescIsDel is the schema descriptor for is_del field.
	sysmenuactionDescIsDel := sysmenuactionMixinFields0[1].Descriptor()
	// sysmenuaction.DefaultIsDel holds the default value on creation for the is_del field.
	sysmenuaction.DefaultIsDel = sysmenuactionDescIsDel.Default.(bool)
	// sysmenuactionDescSort is the schema descriptor for sort field.
	sysmenuactionDescSort := sysmenuactionMixinFields1[0].Descriptor()
	// sysmenuaction.DefaultSort holds the default value on creation for the sort field.
	sysmenuaction.DefaultSort = sysmenuactionDescSort.Default.(int32)
	// sysmenuactionDescStatus is the schema descriptor for status field.
	sysmenuactionDescStatus := sysmenuactionMixinFields2[0].Descriptor()
	// sysmenuaction.DefaultStatus holds the default value on creation for the status field.
	sysmenuaction.DefaultStatus = sysmenuactionDescStatus.Default.(int32)
	// sysmenuactionDescMemo is the schema descriptor for memo field.
	sysmenuactionDescMemo := sysmenuactionMixinFields3[0].Descriptor()
	// sysmenuaction.DefaultMemo holds the default value on creation for the memo field.
	sysmenuaction.DefaultMemo = sysmenuactionDescMemo.Default.(string)
	// sysmenuaction.MemoValidator is a validator for the "memo" field. It is called by the builders before save.
	sysmenuaction.MemoValidator = sysmenuactionDescMemo.Validators[0].(func(string) error)
	// sysmenuactionDescCreatedAt is the schema descriptor for created_at field.
	sysmenuactionDescCreatedAt := sysmenuactionMixinFields4[0].Descriptor()
	// sysmenuaction.DefaultCreatedAt holds the default value on creation for the created_at field.
	sysmenuaction.DefaultCreatedAt = sysmenuactionDescCreatedAt.Default.(func() time.Time)
	// sysmenuactionDescUpdatedAt is the schema descriptor for updated_at field.
	sysmenuactionDescUpdatedAt := sysmenuactionMixinFields4[1].Descriptor()
	// sysmenuaction.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sysmenuaction.DefaultUpdatedAt = sysmenuactionDescUpdatedAt.Default.(func() time.Time)
	// sysmenuaction.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sysmenuaction.UpdateDefaultUpdatedAt = sysmenuactionDescUpdatedAt.UpdateDefault.(func() time.Time)
	// sysmenuactionDescMenuID is the schema descriptor for menu_id field.
	sysmenuactionDescMenuID := sysmenuactionFields[0].Descriptor()
	// sysmenuaction.MenuIDValidator is a validator for the "menu_id" field. It is called by the builders before save.
	sysmenuaction.MenuIDValidator = func() func(string) error {
		validators := sysmenuactionDescMenuID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(menu_id string) error {
			for _, fn := range fns {
				if err := fn(menu_id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// sysmenuactionDescCode is the schema descriptor for code field.
	sysmenuactionDescCode := sysmenuactionFields[1].Descriptor()
	// sysmenuaction.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	sysmenuaction.CodeValidator = func() func(string) error {
		validators := sysmenuactionDescCode.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(code string) error {
			for _, fn := range fns {
				if err := fn(code); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// sysmenuactionDescName is the schema descriptor for name field.
	sysmenuactionDescName := sysmenuactionFields[2].Descriptor()
	// sysmenuaction.NameValidator is a validator for the "name" field. It is called by the builders before save.
	sysmenuaction.NameValidator = func() func(string) error {
		validators := sysmenuactionDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// sysmenuactionDescID is the schema descriptor for id field.
	sysmenuactionDescID := sysmenuactionMixinFields0[0].Descriptor()
	// sysmenuaction.DefaultID holds the default value on creation for the id field.
	sysmenuaction.DefaultID = sysmenuactionDescID.Default.(string)
	// sysmenuaction.IDValidator is a validator for the "id" field. It is called by the builders before save.
	sysmenuaction.IDValidator = func() func(string) error {
		validators := sysmenuactionDescID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(id string) error {
			for _, fn := range fns {
				if err := fn(id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	sysuserMixin := entity.SysUser{}.Mixin()
	sysuserMixinFields0 := sysuserMixin[0].Fields()
	_ = sysuserMixinFields0
	sysuserMixinFields1 := sysuserMixin[1].Fields()
	_ = sysuserMixinFields1
	sysuserMixinFields2 := sysuserMixin[2].Fields()
	_ = sysuserMixinFields2
	sysuserMixinFields3 := sysuserMixin[3].Fields()
	_ = sysuserMixinFields3
	sysuserFields := entity.SysUser{}.Fields()
	_ = sysuserFields
	// sysuserDescIsDel is the schema descriptor for is_del field.
	sysuserDescIsDel := sysuserMixinFields0[1].Descriptor()
	// sysuser.DefaultIsDel holds the default value on creation for the is_del field.
	sysuser.DefaultIsDel = sysuserDescIsDel.Default.(bool)
	// sysuserDescSort is the schema descriptor for sort field.
	sysuserDescSort := sysuserMixinFields1[0].Descriptor()
	// sysuser.DefaultSort holds the default value on creation for the sort field.
	sysuser.DefaultSort = sysuserDescSort.Default.(int32)
	// sysuserDescCreatedAt is the schema descriptor for created_at field.
	sysuserDescCreatedAt := sysuserMixinFields2[0].Descriptor()
	// sysuser.DefaultCreatedAt holds the default value on creation for the created_at field.
	sysuser.DefaultCreatedAt = sysuserDescCreatedAt.Default.(func() time.Time)
	// sysuserDescUpdatedAt is the schema descriptor for updated_at field.
	sysuserDescUpdatedAt := sysuserMixinFields2[1].Descriptor()
	// sysuser.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sysuser.DefaultUpdatedAt = sysuserDescUpdatedAt.Default.(func() time.Time)
	// sysuser.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sysuser.UpdateDefaultUpdatedAt = sysuserDescUpdatedAt.UpdateDefault.(func() time.Time)
	// sysuserDescStatus is the schema descriptor for status field.
	sysuserDescStatus := sysuserMixinFields3[0].Descriptor()
	// sysuser.DefaultStatus holds the default value on creation for the status field.
	sysuser.DefaultStatus = sysuserDescStatus.Default.(int32)
	// sysuserDescUserName is the schema descriptor for user_name field.
	sysuserDescUserName := sysuserFields[0].Descriptor()
	// sysuser.UserNameValidator is a validator for the "user_name" field. It is called by the builders before save.
	sysuser.UserNameValidator = func() func(string) error {
		validators := sysuserDescUserName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(user_name string) error {
			for _, fn := range fns {
				if err := fn(user_name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// sysuserDescRealName is the schema descriptor for real_name field.
	sysuserDescRealName := sysuserFields[1].Descriptor()
	// sysuser.RealNameValidator is a validator for the "real_name" field. It is called by the builders before save.
	sysuser.RealNameValidator = sysuserDescRealName.Validators[0].(func(string) error)
	// sysuserDescFirstName is the schema descriptor for first_name field.
	sysuserDescFirstName := sysuserFields[2].Descriptor()
	// sysuser.FirstNameValidator is a validator for the "first_name" field. It is called by the builders before save.
	sysuser.FirstNameValidator = sysuserDescFirstName.Validators[0].(func(string) error)
	// sysuserDescLastName is the schema descriptor for last_name field.
	sysuserDescLastName := sysuserFields[3].Descriptor()
	// sysuser.LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	sysuser.LastNameValidator = sysuserDescLastName.Validators[0].(func(string) error)
	// sysuserDescPassword is the schema descriptor for Password field.
	sysuserDescPassword := sysuserFields[4].Descriptor()
	// sysuser.PasswordValidator is a validator for the "Password" field. It is called by the builders before save.
	sysuser.PasswordValidator = sysuserDescPassword.Validators[0].(func(string) error)
	// sysuserDescEmail is the schema descriptor for Email field.
	sysuserDescEmail := sysuserFields[5].Descriptor()
	// sysuser.EmailValidator is a validator for the "Email" field. It is called by the builders before save.
	sysuser.EmailValidator = sysuserDescEmail.Validators[0].(func(string) error)
	// sysuserDescPhone is the schema descriptor for Phone field.
	sysuserDescPhone := sysuserFields[6].Descriptor()
	// sysuser.PhoneValidator is a validator for the "Phone" field. It is called by the builders before save.
	sysuser.PhoneValidator = sysuserDescPhone.Validators[0].(func(string) error)
	// sysuserDescSalt is the schema descriptor for salt field.
	sysuserDescSalt := sysuserFields[7].Descriptor()
	// sysuser.DefaultSalt holds the default value on creation for the salt field.
	sysuser.DefaultSalt = sysuserDescSalt.Default.(func() string)
	// sysuserDescID is the schema descriptor for id field.
	sysuserDescID := sysuserMixinFields0[0].Descriptor()
	// sysuser.DefaultID holds the default value on creation for the id field.
	sysuser.DefaultID = sysuserDescID.Default.(string)
	// sysuser.IDValidator is a validator for the "id" field. It is called by the builders before save.
	sysuser.IDValidator = func() func(string) error {
		validators := sysuserDescID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(id string) error {
			for _, fn := range fns {
				if err := fn(id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
}
