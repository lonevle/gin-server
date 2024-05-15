package main

import (
	"github.com/lonevle/gin-server/core"   // 核心包
	"github.com/lonevle/gin-server/global" // 全局变量
	"github.com/lonevle/gin-server/utils"
	"go.uber.org/zap"
)

func init() {
	global.GS_ROOT = utils.GetProgramPath() // 初始化程序根目录
	global.GS_VP = core.Viper()             // 初始化Viper
	global.GS_LOG = core.Zap()              // 初始化zap日志库
	zap.ReplaceGlobals(global.GS_LOG)
}

func main() {
	zap.L().Info("gin-server启动成功", zap.String("root", global.GS_ROOT)) // 启动日志
	zap.L().Sugar().Infof("gin-server启动成功")
	zap.S().Info("gin-server启动成功")
	zap.S().Infof("gin-server启动成功%s", "ok")

	// global.GS_LOG.Info("gin-server启动成功") // 启动日志
	// global.GS_LOG.Warn("警告信息")
	// global.GS_LOG.Error("错误信息")
	// global.GS_LOG.DPanic("DPanic信息")
	// global.GS_LOG.Panic("Panic信息")
	// global.GS_LOG.Fatal("Fatal信息")

}
