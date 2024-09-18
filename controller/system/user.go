package system

import (
	"calender_reminders/global"
	"calender_reminders/service"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type SysUserController struct {
}

type AddReq struct {
	Context    string    `json:"context"`
	RemindTime time.Time `json:"remind_time"`
}

func (sc *SysUserController) Add(c *gin.Context) {
	//获取用户Id
	user_id := c.GetInt("UserId")
	req := AddReq{}
	err := c.ShouldBind(&req)
	if err != nil {
		global.FAST_LOG.Error("ShouldBind is err: " + err.Error())
	}
	// 将信息添加到数据库中
	err = service.ServiceApp.SystemServiceGroup.Add(user_id, req.Context, req.RemindTime)
	if err != nil {
		global.FAST_LOG.Error("Add is err: " + err.Error())
		c.JSON(200, gin.H{
			"code": -1,
			"data": "",
			"msg":  err,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"data": "success",
		"msg":  "",
	})
}

func (sc *SysUserController) Del(c *gin.Context) {
	userid := c.Param("user_id")
	id, _ := strconv.Atoi(userid)
	err := service.ServiceApp.SystemServiceGroup.Del(id)
	if err != nil {
		global.FAST_LOG.Error("Del is err: " + err.Error())
		c.JSON(200, gin.H{
			"code": -1,
			"data": "",
			"msg":  err,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"date": "",
		"mgs":  "删除成功",
	})
}

func (sc *SysUserController) Check(c *gin.Context) {
	user_id := c.GetInt("UserId")

	list, err := service.ServiceApp.SystemServiceGroup.Check(user_id)
	if err != nil {
		c.JSON(200, gin.H{
			"code": -1,
			"date": "",
			"msg":  err,
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 0,
		"date": list,
		"msg":  "",
	})

}

type UpdateReq struct {
	Id         int       `json:"id"`
	Context    string    `json:"context"`
	RemindTime time.Time `json:"remind_time"`
}

func (sc *SysUserController) Update(c *gin.Context) {
	req := &UpdateReq{}
	err := c.ShouldBind(&req)
	if err != nil {
		global.FAST_LOG.Error("ShouldBind is err：" + err.Error())
		c.JSON(200, gin.H{
			"code": -1,
			"date": "",
			"mgs":  err,
		})
		return
	}

	err = service.ServiceApp.SystemServiceGroup.Update(req.Id, req.Context, req.RemindTime)
	if err != nil {
		c.JSON(200, gin.H{
			"code": -1,
			"date": "",
			"msg":  err,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"date": "",
		"msg":  "Update success",
	})

}
