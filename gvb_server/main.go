package main

import (
	"fmt"
	"gvb_server/core"
	_ "gvb_server/docs"
	"gvb_server/flag"
	"gvb_server/global"
	"gvb_server/routers"
)

// @title gvb_server API文档
// @version 1.0
// @description gvb_server 文档
// @host 127.0.0.01:8080
// @BasePath /
func main() {
	// 读取配置文件
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()
	// 使用全局日志变量记录日志
	//global.Log.Warnln("嘿嘿嘿")
	//global.Log.Error("嘻嘻嘻")
	//global.Log.Infof("哈哈哈")
	// 直接使用logrus包记录日志（这里假设您希望同时使用logrus的直接调用）
	//logrus.Warnln("嘿嘿嘿")
	//logrus.Error("嘻嘻嘻")
	//logrus.Infof("哈哈哈")
	// 初始化数据库连接
	global.DB = core.InitGorm()
	global.Redis = core.ConnectRedis()
	global.EsClient = core.EsConnect()
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}
	// 初始化路由
	router := routers.InitRouter()
	// 获取服务地址
	addr := global.Config.System.Addr()
	// 使用全局日志变量记录服务启动地址
	global.Log.Infof("The server is running on: %s", addr)
	// 启动服务器
	err := router.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
	// 打印数据库连接信息，这里可能需要调整以适应实际的DB对象结构
	fmt.Println(global.DB)

}
