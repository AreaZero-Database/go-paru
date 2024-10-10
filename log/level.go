package log

import "strings"

// Level is a zap level.
type Level int8

// LevelKey is zap level key.
const LevelKey = "level"

const (
	// LevelDebug is zap debug level.
	LevelDebug Level = iota
	// LevelInfo is zap info level.
	LevelInfo
	// LevelWarn is zap warn level.
	LevelWarn
	// LevelError is zap error level.
	LevelError
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	default:
		return ""
	}
}

// ParseLevel parses a level string into a zap Level value.
func ParseLevel(s string) Level {
	switch strings.ToUpper(s) {
	case "DEBUG":
		return LevelDebug
	case "INFO":
		return LevelInfo
	case "WARN":
		return LevelWarn
	case "ERROR":
		return LevelError
	}
	return LevelInfo
}
