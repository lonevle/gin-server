package internal

import (
	"time"

	"github.com/lonevle/gin-server/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Zap = new(_zap)

type _zap struct{}

// 获取zap Core对象
func (z *_zap) GetEncoderCore(l zapcore.Level, level zap.LevelEnablerFunc) zapcore.Core {
	writer := FileRotatelogs.GetWriteSyncer(l.String()) // 日志分割

	return zapcore.NewCore(z.GetEncoder(), writer, level) // 这里的第三个参数给一个函数
	// 程序发出一个日志, 所有的core都会收到
	// core根据不同函数, 只有一个函数日志等级正确会返回true, 并写入
	// 达成了不同日志等级可以写入到不同文件的目的
}

func (z *_zap) GetEncoder() zapcore.Encoder {
	if global.GS_CONFIG.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(z.GetEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(z.GetEncoderConfig())
}

// 编码器配置
func (z *_zap) GetEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.GS_CONFIG.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    global.GS_CONFIG.Zap.ZapEncodeLevel(), // 日志编码器
		EncodeTime:     z.CustomTimeEncoder,                   // 时间格式 // zapcore.ISO8601TimeEncoder
		EncodeDuration: zapcore.SecondsDurationEncoder,        // 持续时间格式
		EncodeCaller:   zapcore.FullCallerEncoder,             // 堆栈编码器
	}
}

// CustomTimeEncoder 自定义日志输出时间格式
func (z *_zap) CustomTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(global.GS_CONFIG.Zap.Prefix + t.Format("2006-01-02 15:04:05.000"))
}

// GetLevelPriority 根据 zapcore.Level 获取 zap.LevelEnablerFunc
func (z *_zap) GetLevelPriority(level zapcore.Level) zap.LevelEnablerFunc {
	switch level {
	case zapcore.DebugLevel:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	case zapcore.InfoLevel:
		return func(level zapcore.Level) bool { // 日志级别
			return level == zap.InfoLevel
		}
	case zapcore.WarnLevel:
		return func(level zapcore.Level) bool { // 警告级别
			return level == zap.WarnLevel
		}
	case zapcore.ErrorLevel:
		return func(level zapcore.Level) bool { // 错误级别
			return level == zap.ErrorLevel
		}
	case zapcore.DPanicLevel:
		return func(level zapcore.Level) bool { // dpanic级别
			return level == zap.DPanicLevel
		}
	case zapcore.PanicLevel:
		return func(level zapcore.Level) bool { // panic级别
			return level == zap.PanicLevel
		}
	case zapcore.FatalLevel:
		return func(level zapcore.Level) bool { // 终止级别
			return level == zap.FatalLevel
		}
	default:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	}
}

// GetZapCores 根据配置文件的Level获取 []zapcore.Core
func (z *_zap) GetZapCores() []zapcore.Core {
	cores := make([]zapcore.Core, 0, 7)

	// 从配置的最低等级遍历到Fatal级日志, 创建多个core
	for level := global.GS_CONFIG.Zap.TransportLevel(); level <= zapcore.FatalLevel; level++ {

		// 对每个core进行初始化, 并添加到数组
		cores = append(cores, z.GetEncoderCore(level, z.GetLevelPriority(level))) // 这里根据不同日志等级返回不同的判断函数
	}
	return cores
}
