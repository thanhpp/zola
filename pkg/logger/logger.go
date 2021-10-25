package logger

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SetLog(cfg *LogConfig) error {
	if cfg == nil {
		return ErrNilConfig
	}

	var (
		zapCfg = zap.NewDevelopmentConfig()
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

	zap.ReplaceGlobals(zlg.Named(cfg.LoggerName))

	return nil
}

func Fatal(message string) {
	zap.S().Fatal(message)
}

func Fatalf(template string, args ...interface{}) {
	zap.S().Fatalf(template, args...)
}

func Panic(message string) {
	zap.S().Panic(message)
}

func Panicf(template string, args ...interface{}) {
	zap.S().Panicf(template, args...)
}

func DPanic(message string) {
	zap.S().DPanic(message)
}

func DPanicf(template string, args ...interface{}) {
	zap.S().DPanicf(template, args...)
}

func Error(message string) {
	zap.S().Error(message)
}

func Errorf(template string, args ...interface{}) {
	zap.S().Errorf(template, args...)
}

func Warn(message string) {
	zap.S().Warn(message)
}

func Warnf(template string, args ...interface{}) {
	zap.S().Warnf(template, args...)
}

func Info(message string) {
	zap.S().Info(message)
}

func Infof(template string, args ...interface{}) {
	zap.S().Infof(template, args...)
}

func Debug(message string) {
	zap.S().Debug(message)
}

func Debugf(template string, args ...interface{}) {
	zap.S().Debugf(template, args...)
}
