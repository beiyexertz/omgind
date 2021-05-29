package schema

import (
	"strings"
	"time"

	"github.com/wanhello/omgind/pkg/helper/json"
)

// Menu 菜单对象
type Menu struct {
	ID         string      `json:"id"`                                         // 唯一标识
	Name       string      `json:"name" binding:"required"`                    // 菜单名称
	Sort       int         `json:"sort"`                                       // 排序值
	Icon       string      `json:"icon"`                                       // 菜单图标
	Router     string      `json:"router"`                                     // 访问路由
	ParentID   string      `json:"parent_id"`                                  // 父级ID
	ParentPath string      `json:"parent_path"`                                // 父级路径
	ShowStatus int         `json:"show_status" binding:"required,max=2,min=1"` // 显示状态(1:显示 2:隐藏)
	Status     int         `json:"status" binding:"required,max=2,min=1"`      // 状态(1:启用 2:禁用)
	Memo       string      `json:"memo"`                                       // 备注
	IsShow 	   *bool       `json:"is_show"`
	Creator    string      `json:"creator" `                                   // 创建者
	CreatedAt  *time.Time   `json:"created_at" `                                // 创建时间
	UpdatedAt  *time.Time   `json:"updated_at" `                                // 更新时间
	Actions    MenuActions `json:"actions" `                                   // 动作列表
}

func (a *Menu) String() string {
	return json.MarshalToString(a)
}

// MenuQueryParam 查询条件
type MenuQueryParam struct {
	PaginationParam
	IDs              []string `form:"-"`          // 唯一标识列表
	Name             string   `form:"-"`          // 菜单名称
	PrefixParentPath string   `form:"-"`          // 父级路径(前缀模糊查询)
	QueryValue       string   `form:"queryValue"` // 模糊查询
	ParentID         *string  `form:"parentID"`   // 父级内码
	ShowStatus       int      `form:"showStatus"` // 显示状态(1:显示 2:隐藏)
	IsShow			 *bool 	  `form:"is_show"`
	Status           int      `form:"status"`     // 状态(1:启用 2:禁用)
}

// MenuQueryOptions 查询可选参数项
type MenuQueryOptions struct {
	OrderFields []*OrderField // 排序字段
}

// MenuQueryResult 查询结果
type MenuQueryResult struct {
	Data       Menus
	PageResult *PaginationResult
}

// Menus 菜单列表
type Menus []*Menu

func (a Menus) Len() int {
	return len(a)
}

func (a Menus) Less(i, j int) bool {
	return a[i].Sort > a[j].Sort
}

func (a Menus) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// ToMap 转换为键值映射
func (a Menus) ToMap() map[string]*Menu {
	m := make(map[string]*Menu)
	for _, item := range a {
		m[item.ID] = item
	}
	return m
}

// SplitParentIDs 拆分父级路径的唯一标识列表
func (a Menus) SplitParentIDs() []string {
	idList := make([]string, 0, len(a))
	mIDList := make(map[string]struct{})

	for _, item := range a {
		if _, ok := mIDList[item.ID]; ok || item.ParentPath == "" {
			continue
		}

		for _, pp := range strings.Split(item.ParentPath, "/") {
			if _, ok := mIDList[pp]; ok {
				continue
			}
			idList = append(idList, pp)
			mIDList[pp] = struct{}{}
		}
	}

	return idList
}

// ToTree 转换为菜单树
func (a Menus) ToTree() MenuTrees {
	list := make(MenuTrees, len(a))
	for i, item := range a {
		list[i] = &MenuTree{
			ID:         item.ID,
			Name:       item.Name,
			Icon:       item.Icon,
			Router:     item.Router,
			ParentID:   item.ParentID,
			ParentPath: item.ParentPath,
			Sort:       item.Sort,
			ShowStatus: item.ShowStatus,
			Status:     item.Status,
			Actions:    item.Actions,
		}
	}
	return list.ToTree()
}

// FillMenuAction 填充菜单动作列表
func (a Menus) FillMenuAction(mActions map[string]MenuActions) Menus {
	for _, item := range a {
		if v, ok := mActions[item.ID]; ok {
			item.Actions = v
		}
	}
	return a
}

// ----------------------------------------MenuTree--------------------------------------

// MenuTree 菜单树
type MenuTree struct {
	ID         string      `yaml:"-" json:"id"`                                  // 唯一标识
	Name       string      `yaml:"name" json:"name"`                             // 菜单名称
	Icon       string      `yaml:"icon" json:"icon"`                             // 菜单图标
	Router     string      `yaml:"router,omitempty" json:"router"`               // 访问路由
	ParentID   string      `yaml:"-" json:"parent_id"`                           // 父级ID
	ParentPath string      `yaml:"-" json:"parent_path"`                         // 父级路径
	Sort       int         `yaml:"sort" json:"sort"`                             // 排序值
	ShowStatus int         `yaml:"-" json:"show_status"`                         // 显示状态(1:显示 2:隐藏)
	IsShow 	   *bool       `json:"is_show"`
	Status     int         `yaml:"-" json:"status"`                              // 状态(1:启用 2:禁用)
	Actions    MenuActions `yaml:"actions,omitempty" json:"actions"`             // 动作列表
	Children   *MenuTrees  `yaml:"children,omitempty" json:"children,omitempty"` // 子级树
}

// MenuTrees 菜单树列表
type MenuTrees []*MenuTree

// ToTree 转换为树形结构
func (a MenuTrees) ToTree() MenuTrees {
	mi := make(map[string]*MenuTree)
	for _, item := range a {
		mi[item.ID] = item
	}

	var list MenuTrees
	for _, item := range a {
		if item.ParentID == "" {
			list = append(list, item)
			continue
		}
		if pitem, ok := mi[item.ParentID]; ok {
			if pitem.Children == nil {
				children := MenuTrees{item}
				pitem.Children = &children
				continue
			}
			*pitem.Children = append(*pitem.Children, item)
		}
	}
	return list
}
