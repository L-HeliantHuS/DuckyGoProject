package api

import (
	"DuckyGo/model"
	"DuckyGo/service"
	"github.com/gin-gonic/gin"
)

// CreateComment 创建评论
func CreateComment(c *gin.Context) {
	var service service.CreateCommentService
	if err := c.ShouldBind(&service); err == nil {
		// 获取上下文中存取的User结构体
		get, _ := c.Get("user")
		user := get.(*model.User)
		res := service.Create(user)
		c.JSON(200, res.Result())
	} else {
		c.JSON(200, ErrorResponse(err).Result())
	}
}

// GetComment 获取评论
func GetComment(c *gin.Context) {
	var service service.GetCommentService
	if err := c.ShouldBind(&service); err == nil {
		oid := c.Param("oid")
		pn := c.Query("pn")

		if pn == "" {
			pn = "1"
		}

		res := service.Get(oid, pn)
		c.JSON(200, res.Result())
	} else {
		c.JSON(200, ErrorResponse(err).Result())
	}
}

// DeleteComment 删除评论
func DeleteComment(c *gin.Context) {
	var service service.DeleteCommentService
	if err := c.ShouldBind(&service); err == nil {
		id := c.Param("id")
		get, _ := c.Get("user")
		user := get.(*model.User)
		res := service.Delete(id, user)
		c.JSON(200, res.Result())
	} else {
		c.JSON(200, ErrorResponse(err).Result())
	}
}
