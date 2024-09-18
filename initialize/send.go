package initialize

import (
	"calender_reminders/model/system"
	"calender_reminders/service"
	"calender_reminders/utils"
	"fmt"
	"sync"
	"time"
)

type TimedTaskStatus int32

const (
	TimedTaskStatus_Pending  TimedTaskStatus = 1 // 待执行
	TimedTaskStatus_Running  TimedTaskStatus = 2 // 执行中
	TimedTaskStatus_Finished TimedTaskStatus = 3 // 已完成
	TimedTaskStatus_Error    TimedTaskStatus = 4 // 发生错误
	TimedTaskStatus_Stop     TimedTaskStatus = 5 // 已停止

)

var timedTaskWaitGroup sync.WaitGroup

func StartOssTimedTask() {
	for {
		// 获取需要发送的任务
		tasks, _ := service.ServiceApp.SystemServiceGroup.GetPendingTasks()
		for _, task := range tasks {
			timedTaskWaitGroup.Add(1)
			go func(task *system.Calender) {
				defer timedTaskWaitGroup.Done()
				err := runTimedTask(task)
				if err != nil {
					_ = service.ServiceApp.SystemServiceGroup.UpdateStatus(task.Id, int(TimedTaskStatus_Error))
					return
				}
			}(task)
		}
		timedTaskWaitGroup.Wait()
		time.Sleep(10 * time.Second)
	}
}

func runTimedTask(task *system.Calender) error {
	//获取用户id
	user := service.ServiceApp.SystemServiceGroup.GetUser(task.UserId)
	//锁定当前任务
	err := service.ServiceApp.SystemServiceGroup.UpdateStatus(task.Id, int(TimedTaskStatus_Running))
	if err != nil {
		return err
	}
	text := fmt.Sprintf("你的提醒：%s,时间：%s", task.Context, task.RemindTime)
	//执行任务
	switch task.Type {
	case 0:
		err = utils.SendEmail(user.Email, text)
	case 1:

	case 2:
	}
	if err != nil {
		return err
	}
	err = service.ServiceApp.SystemServiceGroup.UpdateStatus(task.Id, int(TimedTaskStatus_Finished))
	if err != nil {
		return err
	}
	return nil
}
