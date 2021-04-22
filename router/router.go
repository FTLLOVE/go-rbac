package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rbac/app/api"
	"rbac/middleware"
)

func init() {
	s := g.Server()

	s.SetNameToUriType(ghttp.URI_TYPE_CAMEL)

	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/menu", api.Menu)

		group.Group("/role", func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.AuthMiddleware)
			group.ALL("/", api.Role)
		})
	})
}
