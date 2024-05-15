package internal

import (
	"os"
	"path/filepath"

	"github.com/lonevle/gin-server/global"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var FileRotatelogs = new(fileRotatelogs)

type fileRotatelogs struct{}

func (r *fileRotatelogs) GetWriteSyncer(level string) zapcore.WriteSyncer {
	var filename string

	// 判断配置文件中的日志路径是否是绝对路径
	if filepath.IsAbs(global.GS_CONFIG.Zap.Director) {
		filename = filepath.Join(global.GS_CONFIG.Zap.Director, level+".log")
	} else {
		// 不是绝对路径，拼接日志文件的绝对路径, 避免注册服务导致日志文件路径错误
		filename = filepath.Join(global.GS_ROOT, global.GS_CONFIG.Zap.Director, level+".log")
	}

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
