package api

import (
	"DuckyGo/service"
	"github.com/gin-gonic/gin"
)

// CreatePay 创建订单
func CreatePay(c *gin.Context) {
	var service service.CreatePayService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res.Result())
	} else {
		c.JSON(200, ErrorResponse(err).Result())
	}
}

// NotifyPay 回调通知
func NotifyPay(c *gin.Context) {
	var service service.NotifyPayService
	if err := c.ShouldBind(&service); err == nil {
		service.Notify(c.Request)
		c.String(200, "success")
	} else {
		c.JSON(200, ErrorResponse(err).Result())
	}
}

// ReturnPay 同步通知
func ReturnPay(c *gin.Context) {
	var service service.ReturnPayService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Return(c.Request.Form)
		c.JSON(200, res.Result())
	} else {
		c.JSON(200, ErrorResponse(err).Result())
	}
}

