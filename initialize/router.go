package initialize

import "github.com/gin-gonic/gin"

// 初始化路由
func Routers() *gin.Engine {
	Router := gin.New()
	Router.Use(gin.Recovery()) // 捕获异常
	if gin.Mode() == gin.DebugMode {
		Router.Use(gin.Logger())
	}
	return Router
}
