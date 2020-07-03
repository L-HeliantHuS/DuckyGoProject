package v2

import (
	"DuckyGo/model"
	"DuckyGo/serializer"
	"DuckyGo/util"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

type WebSocketService struct {
}

var (
	GlobalConn = make(map[uint]*websocket.Conn)
)

// UpgradeWS ws提升
func (WebSocketService) UpgradeWS(user *model.User, conn *websocket.Conn) {
	defer func() {
		delete(GlobalConn, user.ID)
		conn.Close()
	}()


	// 登录检测
	if _, ok := GlobalConn[user.ID]; ok {
		GlobalConn[user.ID].WriteMessage(websocket.TextMessage, util.ResponseJsonMarshal(serializer.Response{
			Code: serializer.UserNotPermissionError,
			Msg:  "在其他地方已经登录.",
		}.Result()))
		GlobalConn[user.ID].Close()
	}

	GlobalConn[user.ID] = conn

	fmt.Printf("当前在线数: %d\n", len(GlobalConn))

	for {
		// 初始化接收message的结构体
		temp := serializer.WebSocketMessage{}

		_, message, err := conn.ReadMessage()
		if err != nil {
			break
		}

		// 心跳检测
		if string(message) == "ping" {
			conn.WriteMessage(websocket.TextMessage, []byte("pong"))
			continue
		}


		// serializer2struct
		err = json.Unmarshal(message, &temp)
		if err != nil {
			conn.WriteMessage(websocket.TextMessage, util.ResponseJsonMarshal(serializer.Response{
				Code: serializer.UserInputError,
				Msg:  "Json格式有误",
			}))
			continue
		}

		// 创建回复用的响应结构体
		responseStruct := serializer.Response{
			Data: user.Nickname + " msg is:" + temp.Message,
		}.Result()

		for connItem, _ := range GlobalConn {
			GlobalConn[connItem].WriteMessage(websocket.TextMessage, util.ResponseJsonMarshal(responseStruct))
		}
	}

}
