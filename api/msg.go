package api

import (
	"DuckyGo/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// CreateMsg 创建留言
func CreateMsg(c *gin.Context) {
	var service service.CreateMsgService
	s := sessions.Default(c)
	if err := c.ShouldBind(&service); err == nil {
		if uid, ok := s.Get("user_id").(uint); ok {
			res := service.Create(uid)
			c.JSON(200, res.Result())
		}
	} else {
		c.JSON(200, ErrorResponse(err).Result())
	}
}

// GetMsg 获取留言
func GetMsg(c *gin.Context) {
	var service service.GetMsgService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Get(c.Param("id"))
		c.JSON(200, res.Result())
	} else {
		c.JSON(200, ErrorResponse(err).Result())
	}
}

// GetMegs 获取所有留言
func GetMegs(c *gin.Context) {
	var service service.GetMegsService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Gets()
		c.JSON(200, res.Result())
	} else {
		c.JSON(200, ErrorResponse(err).Result())
	}
}

// ChangeMsg 修改留言
func ChangeMsg(c *gin.Context) {
	var service service.ChangeMsgService
	s := sessions.Default(c)
	if err := c.ShouldBind(&service); err == nil {
		if uid, ok := s.Get("user_id").(uint); ok {
			res := service.Change(c.Param("id"), uid)
			c.JSON(200, res)
		}
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}


// DeleteMsg 删除留言
func DeleteMsg(c *gin.Context) {
	var service service.DeleteMsgService
	s := sessions.Default(c)
	if err := c.ShouldBind(&service); err == nil {
		if uid, ok := s.Get("user_id").(uint); ok {
			res := service.Delete(c.Param("id"), uid)
			c.JSON(200, res.Result())
		}
	} else {
		c.JSON(200, ErrorResponse(err).Result())
	}
}