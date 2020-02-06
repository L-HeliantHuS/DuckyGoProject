package api

import (
	"DuckyGo/service"
	"github.com/gin-gonic/gin"
)

// OSSGetKey 获取OSS直传KEY
func OSSGetKey(c *gin.Context) {
	var service service.GetOSSKeyService
	if err := c.ShouldBind(&service); err == nil {
        res := service.POST()
        c.JSON(200, res.Result())
	} else {
		c.JSON(200, ErrorResponse(err).Result())
	}
}

// GetOSSFileList 获取OSS里面的所有文件
func GetOSSFileList(c *gin.Context) {
	var service service.GetOSSFileListService
	if err := c.ShouldBind(&service); err == nil {
        res := service.Get()
        c.JSON(200, res.Result())
	} else {
		c.JSON(200, ErrorResponse(err).Result())
	}
}