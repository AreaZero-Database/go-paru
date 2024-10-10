package zap

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
)

// createZapLogger creates a new zap logger with different log levels and outputs.
func createZapLogger(path string, rotationMaxSize int, compress bool) *zap.Logger {
	encoder := getJSONEncoder()
	// Define level-handling logic.
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})
	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	infoWriterSyncer := getLogWriterSyncer("info", path, rotationMaxSize, compress)
	errorWriterSyncer := getLogWriterSyncer("error", path, rotationMaxSize, compress)

	infoMultiWriteSyncer := zapcore.NewMultiWriteSyncer(infoWriterSyncer)
	errorMultiWriteSyncer := zapcore.NewMultiWriteSyncer(errorWriterSyncer, os.Stdout)

	core := zapcore.NewCore(encoder, infoMultiWriteSyncer, infoLevel)
	errorCore := zapcore.NewCore(encoder, errorMultiWriteSyncer, errorLevel)

	coreArr := []zapcore.Core{core, errorCore}
	return zap.New(zapcore.NewTee(coreArr...), zap.AddCaller(), zap.AddCallerSkip(2))
}

// getJSONEncoder returns a JSON encoder with custom configurations.
func getJSONEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// getLogWriterSyncer returns a WriteSyncer for the given log level.
func getLogWriterSyncer(level string, path string, rotationMaxSize int, compress bool) zapcore.WriteSyncer {
	lumberWriteSyncer := &lumberjack.Logger{
		Filename: filepath.Join(path, fmt.Sprintf("%s.log", level)),
		MaxSize:  rotationMaxSize, // Max size in MB
		Compress: compress,
	}
	return zapcore.Lock(zapcore.AddSync(lumberWriteSyncer))
}
