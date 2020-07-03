package middleware

import (
	"DuckyGo/auth"
	"DuckyGo/cache"
	"DuckyGo/conf"
	"DuckyGo/serializer"
	"DuckyGo/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"strings"
)

// JwtRequired 需要在Header中传递token
func JwtRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获得token
		userToken := c.Request.Header.Get("Authorization")
		// 判断请求头中是否有token
		if userToken == "" {
			c.JSON(http.StatusOK, serializer.Response{
				Code: serializer.UserNotPermissionError,
				Msg:  "令牌不能为空！",
			}.Result())
			c.Abort()
			return
		}

		split := strings.Split(userToken, " ")
		if len(split) != 2 || split[0] != "Bearer" {
			c.JSON(http.StatusOK, serializer.Response{
				Code: serializer.UserNotPermissionError,
				Msg:  "令牌格式不正确",
			}.Result())
			c.Abort()
			return
		}

		// 解码token值
		token, err := jwt.ParseWithClaims(split[1], &auth.Jwt{}, func(token *jwt.Token) (interface{}, error) { return conf.SigningKey, nil })
		if err != nil || token.Valid != true {
			// 过期或者非正确处理
			c.JSON(http.StatusOK, serializer.Response{
				Code: serializer.UserNotPermissionError,
				Msg:  "令牌错误！",
			}.Result())
			c.Abort()
			return
		}

		// 判断令牌是否在黑名单里面
		if result, _ := cache.RedisClient.SIsMember("jwt:baned", token.Raw).Result(); result {
			c.JSON(http.StatusOK, serializer.Response{
				Code: serializer.UserNotPermissionError,
				Msg:  "令牌已注销!",
			}.Result())
			c.Abort()
			return
		}

		// 将Token也放入Context, 用于注销添加黑名单
		c.Set("token", token.Raw)

		// 将结构体地址存入上下文
		if jwtStruct, ok := token.Claims.(*auth.Jwt); ok {
			c.Set("user", &jwtStruct.Data)
		}
	}
}

// WebsocketRequiredToken websocket服务token不能放到请求头, 需要从query获取token来验证
func WebSocketRequiredToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		userToken := c.Query("token")
		var upgrader = websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.JSON(http.StatusOK, serializer.Response{
				Code: serializer.UserBadRequest,
				Msg:  "这是一个ws接口, 请使用websocket访问。",
			}.Result())
			c.Abort()
			return
		}

		// 检查token
		if userToken == "" {
			conn.WriteMessage(websocket.TextMessage, util.ResponseJsonMarshal(serializer.Response{
				Code: serializer.UserInputError,
				Msg:  "Token错误",
			}.Result()))
			c.Abort()
			conn.Close()
			return
		}

		// 解码token值
		token, err := jwt.ParseWithClaims(userToken, &auth.Jwt{}, func(token *jwt.Token) (interface{}, error) { return conf.SigningKey, nil })
		if err != nil || token.Valid != true {
			// 过期或者非正确处理
			conn.WriteMessage(websocket.TextMessage, util.ResponseJsonMarshal(serializer.Response{
				Code: serializer.UserNotPermissionError,
				Msg:  "令牌错误！",
			}.Result()))
			c.Abort()
			conn.Close()
			return
		}

		// 判断令牌是否在黑名单里面
		if result, _ := cache.RedisClient.SIsMember("jwt:baned", token.Raw).Result(); result {
			conn.WriteMessage(websocket.TextMessage, util.ResponseJsonMarshal(serializer.Response{
				Code: serializer.UserNotPermissionError,
				Msg:  "令牌已注销!",
			}.Result()))
			c.Abort()
			conn.Close()
			return
		}

		conn.WriteMessage(websocket.TextMessage, util.ResponseJsonMarshal(serializer.Response{
			Code: 0,
			Data: "连接成功！",
		}.Result()))

		// 将Token也放入Context, 用于注销添加黑名单
		c.Set("token", token.Raw)
		c.Set("conn", conn)

		// 将结构体地址存入上下文
		if jwtStruct, ok := token.Claims.(*auth.Jwt); ok {
			c.Set("user", &jwtStruct.Data)
		}

	}
}
