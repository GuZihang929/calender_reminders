package main

import (
	"calender_reminders/core"
	"calender_reminders/global"
	"calender_reminders/initialize"
	"fmt"
	"go.uber.org/zap"
	"os/exec"
)

// @title                       图书管理
// @version                     0.0.1
// @description                测试
// host							127.0.0.1:8181
func main() {
	cmd := exec.Command("swag", "init")
	fmt.Println("重新加载swagger成功", cmd.Args)
	global.FAST_VP = initialize.Viper()
	global.FAST_LOG = initialize.Zap()
	zap.ReplaceGlobals(global.FAST_LOG)
	global.FAST_DB = initialize.Gorm()

	core.RunWindowsServer()
}
