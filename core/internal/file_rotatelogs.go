package internal

import (
	"os"

	"github.com/lonevle/gin-server/global"
	"github.com/lonevle/gin-server/utils"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var FileRotatelogs = new(fileRotatelogs)

type fileRotatelogs struct{}

func (r *fileRotatelogs) GetWriteSyncer(level string) zapcore.WriteSyncer {
	var filename = utils.JoinPath(global.GS_CONFIG.Zap.Director, level+".log") // 文件名

	fileWriter := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    global.GS_CONFIG.Zap.MaxSize,    // 每个日志文件最大MB
		MaxBackups: global.GS_CONFIG.Zap.MaxBackups, // 保留3个
		MaxAge:     global.GS_CONFIG.Zap.MaxAge,     // 最多保留28天
		Compress:   global.GS_CONFIG.Zap.Compress,   //是否压缩处理
	}
	if global.GS_CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter))
	}
	return zapcore.AddSync(fileWriter)
}
