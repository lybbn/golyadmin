package router

import (
	"gitee.com/lybbn/golyadmin/router/system"
)

type RouterGroup struct {
	System system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
