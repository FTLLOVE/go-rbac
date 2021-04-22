package api

import (
	"github.com/gogf/gf/net/ghttp"
	"rbac/app/service"
	"rbac/library/response"
)

type menuApi struct {
}

var Menu = new(menuApi)

// findMenuTree 获取菜单
func (*menuApi) FindMenuTree(r *ghttp.Request) {
	result, err := service.Menu.FindMenuTree()

	if err != nil {
		response.Fail(r)
	} else {
		response.OkWithData(r, result)
	}
}

// findMenu 根据用户获取指定的菜单列表
func (*menuApi) FindMenu(r *ghttp.Request) {
	userId := r.GetInt("userId")
	result, err := service.Menu.FindMenu(userId)
	if err != nil {
		response.Fail(r)
	} else {
		response.OkWithData(r, result)
	}

}
