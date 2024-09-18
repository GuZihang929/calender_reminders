package system

import (
	"calender_reminders/controller"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	sysRouter := Router.Group("sys")
	//获取路由函数
	var loginController = controller.ApiGroupApp.SystemApiGroup
	{
		sysRouter.POST("/add", loginController.Add)
		sysRouter.DELETE("/del", loginController.Del)
		sysRouter.GET("/check", loginController.Check)
		sysRouter.PUT("/update", loginController.Update)
	}
	return sysRouter
}
