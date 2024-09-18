package controller

import (
	"calender_reminders/controller/system"
)

type ApiGroup struct {
	SystemApiGroup system.SystemControllerGroup
}

var ApiGroupApp = new(ApiGroup)
