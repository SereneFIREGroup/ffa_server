package log

import (
	"fmt"
	"io"
	"os"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/serenefiregroup/ffa_server/pkg/config"
	"github.com/serenefiregroup/ffa_server/pkg/errors"
	"github.com/serenefiregroup/ffa_server/pkg/file"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func LoadLog() error {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Format("2006-01-02T15:04:05.000Z0700"))
	}
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(cfg.EncoderConfig)
	writeSyncer := zapcore.AddSync(NewLoggerWriter())
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	return nil
}

func LoggerHandleFunc() gin.HandlerFunc {
	return ginzap.Ginzap(logger, time.RFC3339, true)
}

func RecoveryHandleFunc() gin.HandlerFunc {
	return ginzap.RecoveryWithZap(logger, true)
}

func Error(format string, args ...interface{}) {
	if logger == nil {
		return
	}
	logger.Error(fmt.Sprintf(format, args...))
}

func Warn(format string, args ...interface{}) {
	if logger == nil {
		return
	}
	logger.Warn(fmt.Sprintf(format, args...))
}

func Info(format string, args ...interface{}) {
	if logger == nil {
		return
	}
	logger.Info(fmt.Sprintf(format, args...))
}

func Debug(format string, args ...interface{}) {
	if logger == nil {
		return
	}
	logger.Debug(fmt.Sprintf(format, args...))
}

func NewLoggerWriter() io.Writer {
	logFile, err := GetLogFile()
	if err != nil {
		panic(err)
	}
	writer := io.MultiWriter(os.Stdout, logFile)
	return writer
}

func GetLogFile() (*os.File, error) {
	logFilePath := config.String("log_file", "")
	if logFilePath == "" {
		err := file.CheckDirAndMkDir("./log")
		if err != nil {
			return nil, errors.Trace(err)
		}
		now := time.Now()
		logFilePath = "./log/%d-%d-%d.log"
		logFilePath = fmt.Sprintf(logFilePath, now.Year(), now.Month(), now.Day())
	}

	logFile, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}
