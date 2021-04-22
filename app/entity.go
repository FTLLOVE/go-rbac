package app

import "rbac/app/model"

type MenuEntity struct {
	Menu      *model.TbMenu
	ChildMenu []*MenuEntity
}
