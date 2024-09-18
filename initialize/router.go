package initialize

import (
	"calender_reminders/router"

	//_ "Library_Project/docs"
	"calender_reminders/global"
	"calender_reminders/middleware"
	"github.com/gin-gonic/gin"
)

// 初始化总路由
// todo 初始化总路由
func Routers() *gin.Engine {
	Router := gin.New()

	Router.Use(middleware.GinLogger())

	//Router.GET("api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//加载文件
	// 全局注册路由组实例位置
	// todo:  注册路有组实例
	systemRouter := router.RouterGroupApp.System
	//    私有路有组有拦截
	PrivateGroup := Router.Group("api")
	PrivateGroup.Use(middleware.JWTAuth())
	{
		systemRouter.InitUserRouter(PrivateGroup)
	}
	global.FAST_LOG.Info("router register success")

	return Router
}
