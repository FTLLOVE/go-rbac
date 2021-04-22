package main

import (
	_ "rbac/boot"
	_ "rbac/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
