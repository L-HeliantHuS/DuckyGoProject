package api

import (
	"DuckyGo/service"
	"github.com/gin-gonic/gin"
)

// GetCloudList 获取网盘文件列表
func GetCloudList(c *gin.Context) {
	var service service.GetCloudListService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Gets()
		c.JSON(200, res.Result())
	} else {
		c.JSON(200, ErrorResponse(err).Result())
	}
}

// UploadCloud 上传文件到网盘
func UploadCloud(c *gin.Context) {
	var service service.UploadCloudItem
	if err := c.ShouldBind(&service); err == nil {
		//+ 获取上传的文件指针
		file, header, err := c.Request.FormFile("file")
		//! 错误直接返回
		if err != nil {
			c.JSON(200, ErrorResponse(err).Result())
		}
		//+ 没错误就继续上传
		res := service.Upload(&file, header)
		c.JSON(200, res.Result())
	} else {
		c.JSON(200, ErrorResponse(err).Result())
	}
}
