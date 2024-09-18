package core

import (
	"calender_reminders/global"
	"calender_reminders/initialize"
	"fmt"
	"go.uber.org/zap"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {

	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.FAST_CONFIG.System.Port)

	// 初始化swagger
	initialize.Swagger(Router)
	s := initServer(address, Router)
	// 保证文本顺序输出
	time.Sleep(10 * time.Microsecond)
	global.FAST_LOG.Info("server run success on", zap.String("address", address))
	// 创建定时发送任务
	go initialize.StartOssTimedTask()
	global.FAST_LOG.Error(s.ListenAndServe().Error())

}
