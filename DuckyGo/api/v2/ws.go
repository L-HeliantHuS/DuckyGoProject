package v2

import (
	"DuckyGo/model"
	v2 "DuckyGo/service/v2"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// WebSocketService WebSocket的服务
func WebSocket(c *gin.Context) {
	var service v2.WebSocketService
	var conn *websocket.Conn
	// 获取websocket的socket连接
	wsConn, exists := c.Get("conn")
	if exists {
		if realConn, ok := wsConn.(*websocket.Conn); ok {
			conn = realConn
		}
	}

	// 获取登录后的用户结构体
	u, _ := c.Get("user")

	service.UpgradeWS(u.(*model.User), conn)
}
