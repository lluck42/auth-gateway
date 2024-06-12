package main

import (
	"fmt"
	"gq_volunteer_platform/bean"
	"gq_volunteer_platform/middleware"
	"gq_volunteer_platform/route"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func init() {
	// 配置项初始化
	bean.ConfigInit()
	// 日志初始化
	bean.LogInit()
	// mysql 初始化
	bean.MysqlInit()
}

func main() {

	gin.SetMode(bean.Config.App.SetMode)
	gin.DisableConsoleColor()
	// 记录到指定文件
	// if f, err := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err != nil {
	// 	panic("创建 access.log 失败")1
	// } else {
	// 	gin.DefaultWriter = f
	// }

	router := gin.Default()
	router.TrustedPlatform = "X-Real-Ip" // 用于获取 nginx 设置的 X-Real-Ip（real client ip）

	router.MaxMultipartMemory = 200 << 20 // 上传文件大小限制 200 MiB

	// 管理端路由
	route.Admin(router)
	// 反向代理 后端的服务
	router.NoRoute(middleware.Record, middleware.Login, middleware.Permission, func(c *gin.Context) {
		fmt.Println("router.NoRoute")
		proxy := httputil.NewSingleHostReverseProxy(&url.URL{
			Scheme: bean.Config.ReverseProxy.Scheme,
			Host:   bean.Config.ReverseProxy.Host,
		})

		proxy.ServeHTTP(c.Writer, c.Request)
	})

	router.Run(bean.Config.App.Addr) // listen and serve on 0.0.0.0:8080
}
