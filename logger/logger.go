package logger

import (
	"fmt"
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogLevel zapcore.Level

const (
	// We cannt override the default log level in zaplog, because they are globally
	// so, we need to skip the zap default log level
	skipZapLogLevel LogLevel = LogLevel(zapcore.FatalLevel) + iota
	AlertLevel
	CritLevel
	ErrLevel
	WarnLevel
	NoticeLevel
	InfoLevel
	DebugLevel
)

type Logger interface {
	Alert(format string, args ...any)
	Crit(format string, args ...any)
	Error(format string, args ...any)
	Warn(format string, args ...any)
	Notice(format string, args ...any)
	Info(format string, args ...any)
	Debug(format string, args ...any)
}

func New(writer io.Writer) Logger {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeLevel = capitalLevelEncoder
	cfg.StacktraceKey = zapcore.OmitKey
	core := zapcore.NewCore(zapcore.NewJSONEncoder(cfg), zapcore.AddSync(writer), zapcore.Level(AlertLevel))
	return &lwrapper{zap.New(core)}
}

type lwrapper struct {
	*zap.Logger
}

func (l *lwrapper) Alert(format string, args ...any) {
	l.Log(zapcore.Level(AlertLevel), fmt.Sprintf(format, args...))
}

func (l *lwrapper) Crit(format string, args ...any) {
	l.Log(zapcore.Level(CritLevel), fmt.Sprintf(format, args...))
}
func (l *lwrapper) Warn(format string, args ...any) {
	l.Log(zapcore.Level(WarnLevel), fmt.Sprintf(format, args...))
}
func (l *lwrapper) Notice(format string, args ...any) {
	l.Log(zapcore.Level(NoticeLevel), fmt.Sprintf(format, args...))
}
func (l *lwrapper) Debug(format string, args ...any) {
	l.Log(zapcore.Level(DebugLevel), fmt.Sprintf(format, args...))
}

func (l *lwrapper) Info(format string, args ...any) {
	l.Log(zapcore.Level(InfoLevel), fmt.Sprintf(format, args...))
}

func (l *lwrapper) Error(format string, args ...any) {
	l.Log(zapcore.Level(ErrLevel), fmt.Sprintf(format, args...))
}

func capitalLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	switch level {
	case zapcore.Level(AlertLevel):
		enc.AppendString("Alert")
	case zapcore.Level(CritLevel):
		enc.AppendString("Critical")
	case zapcore.Level(ErrLevel):
		enc.AppendString("Error")
	case zapcore.Level(WarnLevel):
		enc.AppendString("Warning")
	case zapcore.Level(NoticeLevel):
		enc.AppendString("Notice")
	case zapcore.Level(InfoLevel):
		enc.AppendString("Info")
	case zapcore.Level(DebugLevel):
		enc.AppendString("Debug")
	}
}
