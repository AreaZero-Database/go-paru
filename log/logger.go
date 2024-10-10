package log

// Logger zap interface
type Logger interface {
	Debug(v ...any)
	DebugF(format string, v ...any)
	Info(v ...any)
	InfoF(format string, v ...any)
	Warn(v ...any)
	WarnF(format string, v ...any)
	Error(v ...any)
	ErrorF(format string, v ...any)
}
