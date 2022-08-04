package service

import (
	"fhgo/service/danmu"
	"fhgo/service/system"
)

type ServiceGroup struct {
	SystemService system.ServiceGroup
	DanMuService  danmu.ServiceGroup
}

var ServiceGroupApi = new(ServiceGroup)

var Group = new(ServiceGroup)
