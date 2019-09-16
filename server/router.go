package server

import (
	"DuckyGo/api"
	"DuckyGo/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	r.StaticFile("/favicon.ico", "static/favicon.ico")

	// 中间件, 顺序不能改
	// 这条是为了防止因为记录日志而导致性能损失
	if os.Getenv("LOG_LEVEL") == "DEBUG" {
		r.Use(middleware.SaveLog())
	}
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.GET("ping", api.Ping)
		v1.GET("/add", api.Add)
		// 用户注册---------------------------------------------------------------------------------------------
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)
		// ----------------------------------------------------------------------------------------------------



		// 留言板Crud ---------------------------------------------------------------------------------------------
		// 查询留言
		v1.GET("msg/:id", api.GetMsg)
		// 查询所有留言
		v1.GET("msg", api.GetMegs)
		// -------------------------------------------------------------------------------------------------------



		// 获得所有二手房信息
		v1.GET("lianjia", api.GetsHouse)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("user/me", api.UserMe)
			auth.DELETE("user/logout", api.UserLogout)
			auth.PUT("user/changepassword", api.ChangePassword)

			// 创建留言
			auth.POST("msg", api.CreateMsg)

			// 修改留言
			auth.PUT("msg/:id", api.ChangeMsg)
			// 删除留言
			auth.DELETE("msg/:id", api.DeleteMsg)

			// 需要是管理员
			admin := auth.Group("")
			admin.Use(middleware.AuthAdmin())
			{

			}
		}

	}
	return r
}
