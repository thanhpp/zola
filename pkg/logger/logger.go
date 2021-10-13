package logger

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	globalLogger, _ = zap.NewProduction()
)

func SetLog(cfg *LogConfig) error {
	if cfg == nil {
		return ErrNilConfig
	}

	var (
		zapCfg zap.Config
	)
	// log level
	switch strings.ToUpper(cfg.Level) {
	case "FATAL":
		zapCfg.Level.SetLevel(zapcore.FatalLevel)
	case "PANIC":
		zapCfg.Level.SetLevel(zapcore.PanicLevel)
	case "DPANIC":
		zapCfg.Level.SetLevel(zapcore.DPanicLevel)
	case "ERROR":
		zapCfg.Level.SetLevel(zapcore.ErrorLevel)
	case "WARN":
		zapCfg.Level.SetLevel(zapcore.WarnLevel)
	case "INFO":
		zapCfg.Level.SetLevel(zapcore.InfoLevel)
	case "DEBUG":
		zapCfg.Level.SetLevel(zapcore.DebugLevel)
	default:
		zapCfg.Level.SetLevel(zapcore.InfoLevel)
	}
	// log color
	if cfg.Color {
		zapCfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}
	// encoder
	zapCfg.EncoderConfig = zapcore.EncoderConfig{
		MessageKey:  "message",
		FunctionKey: "function",

		NameKey:    "name",
		EncodeName: zapcore.FullNameEncoder,

		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,

		TimeKey:    "time",
		EncodeTime: zapcore.ISO8601TimeEncoder,

		CallerKey:    "caller",
		EncodeCaller: zapcore.ShortCallerEncoder,
	}
	// build
	zlg, err := zapCfg.Build(zap.AddCallerSkip(1), zap.AddStacktrace(zap.DPanicLevel))
	if err != nil {
		return err
	}

	globalLogger = zlg

	return nil
}

func Fatal(message string) {
	globalLogger.Sugar().Fatal(message)
	return
}

func Fatalf(template string, args ...interface{}) {
	globalLogger.Sugar().Fatalf(template, args...)
	return
}

func Panic(message string) {
	globalLogger.Sugar().Panic(message)
	return
}

func Panicf(template string, args ...interface{}) {
	globalLogger.Sugar().Panicf(template, args...)
	return
}

func DPanic(message string) {
	globalLogger.Sugar().DPanic(message)
	return
}

func DPanicf(template string, args ...interface{}) {
	globalLogger.Sugar().DPanicf(template, args...)
	return
}

func Error(message string) {
	globalLogger.Sugar().Error(message)
	return
}

func Errorf(template string, args ...interface{}) {
	globalLogger.Sugar().Errorf(template, args...)
	return
}

func Warn(message string) {
	globalLogger.Sugar().Warn(message)
	return
}

func Warnf(template string, args ...interface{}) {
	globalLogger.Sugar().Warnf(template, args...)
	return
}

func Info(message string) {
	globalLogger.Sugar().Info(message)
	return
}

func Infof(template string, args ...interface{}) {
	globalLogger.Sugar().Infof(template, args...)
	return
}

func Debug(message string) {
	globalLogger.Sugar().Debug(message)
	return
}

func Debugf(template string, args ...interface{}) {
	globalLogger.Sugar().Debugf(template, args...)
	return
}
