package middleware

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"rbac/app/service"
	"rbac/library/response"
)

func AuthMiddleware(r *ghttp.Request) {
	methodCode := r.Request.URL.Path
	token := r.Request.Header.Get("x-token")

	if token == "" {
		response.FailNoAuth(r)
	}

	count, err := service.Menu.IsMenuExist(gconv.Int(token), methodCode)

	if err != nil || count <= 0 {
		response.FailNoAuth(r)
	}

	r.Middleware.Next()
}
