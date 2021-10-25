package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogFormat struct {
	Service string
	Message string
	Error   error
}

func NewLogFormat(service string) *LogFormat {
	return &LogFormat{
		Service: service,
	}
}

func (lf *LogFormat) SetError(err error) *LogFormat {
	if lf == nil {
		return nil
	}
	lf.Error = err
	return lf
}

func (lf *LogFormat) SetMsg(msg string) *LogFormat {
	if lf == nil {
		return nil
	}
	lf.Message = msg
	return lf
}

func (lf *LogFormat) FlushErrMsg() {
	if lf == nil {
		return
	}
	lf.Message = ""
	lf.Error = nil
}

func (lf LogFormat) toZapFields() []zapcore.Field {
	return []zapcore.Field{zap.String("service", lf.Service), zap.String("message", lf.Message), zap.Error(lf.Error)}
}

// ------------------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- LOG FORMAT FUNCS ----------------------------------------------------------

var nilStr = ""

func FatalFmt(lf LogFormat) {
	zap.L().Fatal(nilStr, lf.toZapFields()...)
}

func PanicFmt(lf LogFormat) {
	zap.L().Panic(nilStr, lf.toZapFields()...)
}

func ErrorFmt(lf LogFormat) {
	zap.L().Error(nilStr, lf.toZapFields()...)
}

func WarnFmt(lf LogFormat) {
	zap.L().Error(nilStr, lf.toZapFields()...)
}

func InfoFmt(lf LogFormat) {
	zap.L().Error(nilStr, lf.toZapFields()...)
}

func DebugFmt(lf LogFormat) {
	zap.L().Debug(nilStr, lf.toZapFields()...)
}
