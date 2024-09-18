package router

import (
	"calender_reminders/router/system"
)

type RouterGroup struct {
	System system.RouterGroup
	//Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
