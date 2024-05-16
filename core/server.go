package core

import (
	"fmt"

	"github.com/lonevle/gin-server/global"
	"github.com/lonevle/gin-server/initialize"
	"github.com/lonevle/gin-server/utils"
)

func RunServer() {
	Router := initialize.Routers()

	Router.Static("/", utils.JoinPath("./resource/web")) // 静态文件

	address := fmt.Sprintf(":%d", global.GS_CONFIG.System.Addr)

	// s := initServer(address, Router)
	Router.Run(address)
}
