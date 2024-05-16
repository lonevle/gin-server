package core

import (
	"github.com/lonevle/gin-server/core/internal"

	"github.com/lonevle/gin-server/global"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Zap() (logger *zap.Logger) {
	cores := internal.Zap.GetZapCores()

	// 接受任意数量的 Core 对象，并返回一个新的 Core 对象，这个新对象将日志信息分发到所有提供的 Core 对象
	logger = zap.New(zapcore.NewTee(cores...))

	if global.GS_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}

	return logger
}
