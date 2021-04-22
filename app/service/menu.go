package service

import (
	"github.com/gogf/gf/container/gset"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"rbac/app"
	"rbac/app/dao"
	"rbac/app/model"
	"strings"
)

type menuService struct {
}

var Menu = new(menuService)

// addMenu 新增菜单
func (m *menuService) AddMenu(menu *model.TbMenu) error {
	// parentId:0 证明是根节点
	if menu.ParentId == 0 {
		menu.Level = 1
		menu.Path = ""
	} else {
		parentMenu, err := dao.TbMenu.FindOne(menu.ParentId)
		if err != nil {
			return err
		}
		menu.Level = parentMenu.Level + 1

		if parentMenu.Path != "" {
			menu.Path = parentMenu.Path + "," + gconv.String(parentMenu.Id)
		} else {
			menu.Path = gconv.String(parentMenu.Id)
		}
	}

	_, err := dao.TbMenu.Insert(menu)

	if err != nil {
		return err
	}
	return nil
}

// findMenuTree 获取菜单树形结构列表
func (m *menuService) FindMenuTree() ([]*app.MenuEntity, error) {
	menuList, err := dao.TbMenu.Order("level,sort").FindAll()

	if err != nil {
		return nil, err
	}

	menuEntityList, err := m.transferMenuEntity(menuList, 0)

	if err != nil {
		return nil, err
	}

	return menuEntityList, nil

}

// findMenu 获取用户的菜单权限列表
func (m *menuService) FindMenu(userId int) ([]*app.MenuEntity, error) {
	userRoles, err := dao.TbUserRole.FindAll(g.Map{"user_id": userId})
	if err != nil {
		return nil, err
	}
	if len(userRoles) != 0 {
		roleMenus, err := dao.TbRoleMenu.FindAll(g.Map{"role_id": userRoles[0].RoleId})

		if err != nil {
			return nil, err
		}

		if len(roleMenus) != 0 {
			menuIds := gset.New(true)
			for _, el := range roleMenus {
				menuIds.Add(el.MenuId)
			}

			menus, err := dao.TbMenu.FindAll(g.Map{
				"id": menuIds.Slice(),
			})

			if err != nil {
				return nil, err
			}

			if len(menus) != 0 {
				allMenuIds := gset.New(true)

				for _, el := range menus {
					allMenuIds.Add(el.Id)
					if el.Path != "" {
						pathIds := strings.Split(el.Path, ",")
						for _, el := range pathIds {
							allMenuIds.Add(el)
						}
					}
				}

				menus, err := dao.TbMenu.FindAll(g.Map{"id": allMenuIds.Slice()})
				if err != nil {
					return nil, err
				}

				result, err := m.transferMenuEntity(menus, 0)

				if err != nil {
					return nil, err
				}
				return result, nil

			}
		}
	}
	return nil, nil
}

func (*menuService) IsMenuExist(userId int, menuCode string) (int, error) {
	db := g.DB()
	count, err := db.GetAll("SELECT * FROM tb_menu WHERE id IN ( SELECT menu_id FROM tb_role_menu WHERE role_id IN ( SELECT role_id FROM tb_user_role WHERE user_id = ? ) ) AND menu_code = ?", userId, menuCode)

	if err != nil {
		return 0, err
	}
	return count.Len(), nil
}

// transferMenuEntity []*tbMenu转换成[]*menuEntity
func (m *menuService) transferMenuEntity(menuList []*model.TbMenu, parentId int64) ([]*app.MenuEntity, error) {
	var resultList []*app.MenuEntity
	if len(menuList) != 0 {
		for _, el := range menuList {
			if parentId == el.ParentId {
				//menuEntity := new(app.MenuEntity)
				var menuEntity = new(app.MenuEntity)
				_ = gconv.Struct(el, &menuEntity.Menu)
				childMenuList, _ := m.transferMenuEntity(menuList, el.Id)
				if len(childMenuList) != 0 {
					menuEntity.ChildMenu = childMenuList
				}
				resultList = append(resultList, menuEntity)
			}
		}
	}
	return resultList, nil
}
