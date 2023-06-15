package service

import (
	"gitee.com/lybbn/go-vue-lyadmin/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
