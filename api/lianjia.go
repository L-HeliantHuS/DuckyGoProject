package api

import (
	"DuckyGo/service"
	"github.com/gin-gonic/gin"
)

// GetsHouse 获取所有二手房信息
func GetsHouse(c *gin.Context) {
	var service service.GetsHouseService
	if err := c.ShouldBind(&service); err == nil {
		    res := service.Gets()
			c.JSON(200, res.Result())
	} else {
		c.JSON(200, ErrorResponse(err).Result())
	}
}