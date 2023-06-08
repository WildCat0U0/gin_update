package bootstrap

import (
	"Gin_Start/global"
	"Gin_Start/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var level zapcore.Level  // 日志级别
var options []zap.Option // 日志配置选项

func InitializeLog() *zap.Logger {
	//创建根目录
	createRootDir()
	setLogLevel()
	if global.App.Config.Log.ShowLine {
		options = append(options, zap.AddCaller()) // 显示代码行号
	}
	return zap.New(getZapCore(), options...)
}

func createRootDir() {
	if ok, _ := utils.PathExists(global.App.Config.Log.RootDir); !ok {
		_ = os.Mkdir(global.App.Config.Log.RootDir, os.ModePerm)
	}
}

func setLogLevel() {
	switch global.App.Config.Log.Level {
	case "debug":
		level = zap.DebugLevel
		options = append(options, zap.AddStacktrace(level)) // 添加堆栈信息
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
		options = append(options, zap.AddStacktrace(level))
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

}

//编码器是什么 编码器是将日志信息转换成字节流的组件

// 扩展 Zap
func getZapCore() zapcore.Core {
	var encoder zapcore.Encoder

	// 调整编码器默认配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("[" + "2006-01-02 15:04:05.000" + "]"))
	}
	encoderConfig.EncodeLevel = func(l zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(global.App.Config.App.Env + "." + l.String())
	}

	// 设置编码器
	if global.App.Config.Log.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig) // JSON 编码器
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig) // 控制台编码器
	}

	return zapcore.NewCore(encoder, getLogWriter(), level) // 创建日志核心
}

// 使用 lumberjack 作为日志写入器
func getLogWriter() zapcore.WriteSyncer {
	file := &lumberjack.Logger{
		Filename:   global.App.Config.Log.RootDir + "/" + global.App.Config.Log.Filename,
		MaxSize:    global.App.Config.Log.MaxSize,    // megabytes
		MaxBackups: global.App.Config.Log.MaxBackups, // 最多保留多少个备份
		MaxAge:     global.App.Config.Log.MaxAge,     // days
		Compress:   global.App.Config.Log.Compress,   // 是否压缩 disabled by default
	}

	return zapcore.AddSync(file) // 添加文件写入器
}
