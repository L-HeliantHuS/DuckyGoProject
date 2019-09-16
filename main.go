package main

import (
	"DuckyGo/conf"
	"DuckyGo/server"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	r := server.NewRouter()

	// 运行 起在8000端口
	r.Run(":8000")
}


/*
40001  用户输入导致数据库查询错误
40003  用户没权限

50001  数据库创建数据异常
50002  数据库更新数据异常
50003  数据库查询数据异常

*/