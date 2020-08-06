package log

import (
	"net/http"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	logFileDir = "./"
)

var (
	logger *zap.Logger
	//sugaredLogger *zap.SugaredLogger
)

func GetLogger() *zap.Logger {
	return logger
}

func InitLogger() {
	encoder := zapcore.NewConsoleEncoder(newEncoderConfig())
	levelEnableFunc := newLevelEnableFunc(zap.DebugLevel, zap.ErrorLevel)

	//控制台输出
	consoleCore := zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), levelEnableFunc)
	//文件输出
	fileSyncCore := zapcore.NewCore(encoder, getLogWriter(logFileDir+"log.log"), levelEnableFunc)

	//多输出目标
	tee := zapcore.NewTee(consoleCore, fileSyncCore)
	logger = zap.New(tee, zap.AddCaller())
	//sugaredLogger = logger.Sugar()
}

//日志格式配置
func newEncoderConfig() zapcore.EncoderConfig {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return encoderConfig
}

//日志输出等级限制
func newLevelEnableFunc(lower, upper zapcore.Level) zap.LevelEnablerFunc {
	return func(lvl zapcore.Level) bool {
		return lvl >= lower && lvl <= upper
	}
}

//日志文件切割
func getLogWriter(fileName string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func logging() {
	InitLogger()
	simpleHttpGet("www.google.com")
	simpleHttpGet("http://www.baidu.com")
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(
			"Error fetching url..",
			zap.String("url", url),
			zap.Error(err))
	} else {
		logger.Info("Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}
