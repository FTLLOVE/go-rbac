package response

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
)

type Json struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Method    string      `json:"method"`
	Path      string      `json:"path"`
	Timestamp int64       `json:"timestamp"`
}

func Result(r *ghttp.Request, code int, message string, data interface{}) {
	method := r.Method
	path := r.Request.URL.Path
	if err := r.Response.WriteJson(Json{
		Code:      code,
		Message:   message,
		Data:      data,
		Method:    method,
		Path:      path,
		Timestamp: gtime.Timestamp(),
	}); err != nil {
		g.Log().Error(err.Error())
	}

	r.Exit()
}

func Ok(r *ghttp.Request) {
	Result(r, 0, "success", nil)
}

func OkWithData(r *ghttp.Request, data interface{}) {
	Result(r, 0, "success", data)
}

func Fail(r *ghttp.Request) {
	Result(r, 1, "fail", nil)
}

func FailWithMsg(r *ghttp.Request, message string) {
	Result(r, 1, message, nil)
}

func FailNoAuth(r *ghttp.Request) {
	Result(r, 2, "无权限", nil)
}
