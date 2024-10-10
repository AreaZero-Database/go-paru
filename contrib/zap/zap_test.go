package zap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatesZapInstance(t *testing.T) {
	logger := NewLogger("/tmp", 10, false)
	assert.NotNil(t, logger)
	assert.NotNil(t, logger.log)
	assert.NotNil(t, logger.slog)
}

func TestZapDebug(t *testing.T) {
	logger := NewLogger("/tmp", 10, false)
	assert.NotPanics(t, func() {
		logger.Debug("debug message")
	})
}

func TestZapDebugF(t *testing.T) {
	logger := NewLogger("/tmp", 10, false)
	assert.NotPanics(t, func() {
		logger.DebugF("debug %s", "message")
	})
}

func TestZapInfo(t *testing.T) {
	logger := NewLogger("/tmp", 10, false)
	assert.NotPanics(t, func() {
		logger.Info("info message")
	})
}

func TestZapInfoF(t *testing.T) {
	logger := NewLogger("/tmp", 10, false)
	assert.NotPanics(t, func() {
		logger.InfoF("info %s", "message")
	})
}

func TestZapWarn(t *testing.T) {
	logger := NewLogger("/tmp", 10, false)
	assert.NotPanics(t, func() {
		logger.Warn("warn message")
	})
}

func TestZapWarnF(t *testing.T) {
	logger := NewLogger("/tmp", 10, false)
	assert.NotPanics(t, func() {
		logger.WarnF("warn %s", "message")
	})
}

func TestZapError(t *testing.T) {
	logger := NewLogger("/tmp", 10, false)
	assert.NotPanics(t, func() {
		logger.Error("error message")
	})
}

func TestZapErrorF(t *testing.T) {
	logger := NewLogger("/tmp", 10, false)
	assert.NotPanics(t, func() {
		logger.ErrorF("error %s", "message")
	})
}

func TestGetLoggerAndSugaredLogger(t *testing.T) {
	logger := NewLogger("/tmp", 10, false)
	assert.NotNil(t, logger.GetLogger())
	assert.NotNil(t, logger.GetSugaredLogger())
}
