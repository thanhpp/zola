package core

import (
	"encoding/json"
	"log"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LogFormat data format
type LogFormat struct {
	ServiceName string      `json:"srv"`
	Package     string      `json:"pkg,omitempty"`
	Action      string      `json:"act,omitempty"`
	Data        interface{} `json:"data,omitempty"`
	Err         error       `json:"err,omitempty"`
	ErrMsg      string      `json:"errmsg,omitempty"`
	Success     interface{} `json:"suc,omitempty"`
	Message     string      `json:"msg,omitempty"`
	Source      string
}

var (
	logger *LogFormat
)

// GetLogger return log format object
func GetLogger() *LogFormat {
	return logger
}

// InitLogger return logger object
func InitLogger(cfg *MainConfig) (lg *LogFormat) {
	var zapConfig zap.Config
	lg = &LogFormat{ServiceName: cfg.Name}
	if strings.ToUpper(cfg.Environment) == "DEVELOPMENT" {
		zapConfig = zap.NewDevelopmentConfig()
	} else {
		zapConfig = zap.NewProductionConfig()
	}

	// level
	switch strings.ToUpper(cfg.Log.Level) {
	case "FATAL":
		zapConfig.Level.SetLevel(zapcore.FatalLevel)
	case "PANIC":
		zapConfig.Level.SetLevel(zapcore.PanicLevel)
	case "DPANIC":
		zapConfig.Level.SetLevel(zapcore.DPanicLevel)
	case "ERROR":
		zapConfig.Level.SetLevel(zapcore.ErrorLevel)
	case "WARN":
		zapConfig.Level.SetLevel(zapcore.WarnLevel)
	case "INFO":
		zapConfig.Level.SetLevel(zapcore.InfoLevel)
	case "DEBUG":
		zapConfig.Level.SetLevel(zapcore.DebugLevel)
	default:
		zapConfig.Level.SetLevel(zapcore.InfoLevel)
	}

	zapConfig.EncoderConfig = zapcore.EncoderConfig{
		MessageKey:  "message",
		FunctionKey: "function",

		NameKey:    "name",
		EncodeName: zapcore.FullNameEncoder,

		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalColorLevelEncoder,

		TimeKey:    "time",
		EncodeTime: zapcore.ISO8601TimeEncoder,

		CallerKey:    "caller",
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	if cfg.Log.Color {
		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}
	zapLogger, err := zapConfig.Build(zap.AddCallerSkip(1), zap.AddStacktrace(zap.FatalLevel))
	if err != nil {
		log.Fatalf("Logger init failed: %v", err)
	}
	namedLogger := zapLogger.Named(cfg.Name)

	zap.ReplaceGlobals(namedLogger)
	return lg
}

// Fatal log fatal error
func (lg *LogFormat) Fatal(message string) {
	zap.L().Fatal(message)
}

// Panic log panic error
func (lg *LogFormat) Panic(message string) {
	zap.L().Panic(message)
}

// DPanic log panic error and in Development
// config it then panics
func (lg *LogFormat) DPanic(message string) {
	zap.L().DPanic(message)
}

// Error log error message
func (lg *LogFormat) Error(message string) {
	zap.L().Error(message)
}

// Warn log warn message
func (lg *LogFormat) Warn(message string) {
	zap.L().Warn(message)
}

// Info log info message
func (lg *LogFormat) Info(message string) {
	zap.L().Info(message)
}

// Debug log debug message
func (lg *LogFormat) Debug(message string) {
	zap.L().Debug(message)
}

// Fatalf log fatal error
func (lg *LogFormat) Fatalf(template string, args ...interface{}) {
	zap.S().Fatalf(template, args)
}

// Panicf log panic error
func (lg *LogFormat) Panicf(template string, args ...interface{}) {
	zap.S().Panicf(template, args)
}

// DPanicf log panic error and in Development
// config it then panics
func (lg *LogFormat) DPanicf(template string, args ...interface{}) {
	zap.S().DPanicf(template, args)
}

// Errorf log error message
func (lg *LogFormat) Errorf(template string, args ...interface{}) {
	zap.S().Error(template, args)
}

// Warnf log warn message
func (lg *LogFormat) Warnf(template string, args ...interface{}) {
	zap.S().Warnf(template, args)
}

// Infof log info message
func (lg *LogFormat) Infof(template string, args ...interface{}) {
	zap.S().Infof(template, args)
}

// Debugf log debug message
func (lg *LogFormat) Debugf(template string, args ...interface{}) {
	zap.S().Debugf(template, args)
}

func (lg *LogFormat) Dataf(data interface{}) {
	js, _ := json.Marshal(data)
	zap.S().Infof("%s", js)
}
