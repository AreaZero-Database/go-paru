package zap

import (
	"github.com/AreaZero-Database/go-paru/log"
	"go.uber.org/zap"
)

var _ log.Logger = (*Logger)(nil)

type Logger struct {
	log  *zap.Logger
	slog *zap.SugaredLogger
}

func NewLogger(path string, rotationMaxSize int, compress bool) *Logger {
	zLogger := createZapLogger(path, rotationMaxSize, compress)
	return &Logger{
		log:  zLogger,
		slog: zLogger.Sugar(),
	}
}

// Debug log
func (z *Logger) Debug(v ...any) {
	z.slog.Debug(v...)
}

// DebugF log
func (z *Logger) DebugF(format string, v ...any) {
	z.slog.Debugf(format, v...)
}

// Info log
func (z *Logger) Info(v ...any) {
	z.slog.Info(v...)
}

// InfoF log
func (z *Logger) InfoF(format string, v ...any) {
	z.slog.Infof(format, v...)
}

// Warn log
func (z *Logger) Warn(v ...any) {
	z.slog.Warn(v...)
}

// WarnF log
func (z *Logger) WarnF(format string, v ...any) {
	z.slog.Warnf(format, v...)
}

// Error log
func (z *Logger) Error(v ...any) {
	z.slog.Error(v...)
}

// ErrorF log
func (z *Logger) ErrorF(format string, v ...any) {
	z.slog.Errorf(format, v...)
}

func (z *Logger) GetLogger() *zap.Logger {
	return z.log
}

func (z *Logger) GetSugaredLogger() *zap.SugaredLogger {
	return z.slog
}
