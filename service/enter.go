package service

import (
	"calender_reminders/service/example"
	"calender_reminders/service/system"
)

type Service struct {
	SystemServiceGroup  system.SysGroup
	ExampleServiceGroup example.ServiceGroup
}

var ServiceApp = new(Service)
