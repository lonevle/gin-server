package core

import (
	"fmt"
	"os"

	"github.com/lonevle/gin-server/core/internal"

	"github.com/lonevle/gin-server/global"
	"github.com/lonevle/gin-server/utils"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.GS_CONFIG.Zap.Director); !ok { // 判断是日志文件夹是否存在
		fmt.Printf("create %v directory\n", global.GS_CONFIG.Zap.Director)
		_ = os.Mkdir(global.GS_CONFIG.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()

	// 接受任意数量的 Core 对象，并返回一个新的 Core 对象，这个新对象将日志信息分发到所有提供的 Core 对象
	logger = zap.New(zapcore.NewTee(cores...))

	if global.GS_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}

	return logger
}
