package api

import (
	"github.com/gogf/gf/net/ghttp"
	"rbac/library/response"
)

type roleApi struct {
}

var Role = new(roleApi)

// add 新增角色
func (*roleApi) Add(r *ghttp.Request) {
	response.Ok(r)
}

// update 更新角色
func (*roleApi) Update(r *ghttp.Request) {
	response.Ok(r)
}

// findAll 查询角色列表
func (*roleApi) FindAll(r *ghttp.Request) {
	response.Ok(r)
}
