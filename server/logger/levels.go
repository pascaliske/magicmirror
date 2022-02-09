package logger

import (
	"strings"

	"github.com/fatih/color"
)

type LogLevel int

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

var logLevel LogLevel = InfoLevel

func SetLevel(level string) {
	switch strings.ToUpper(level) {
	case DebugLevel.String():
		logLevel = DebugLevel
	case InfoLevel.String():
		logLevel = InfoLevel
	case WarnLevel.String():
		logLevel = WarnLevel
	case ErrorLevel.String():
		logLevel = ErrorLevel
	case FatalLevel.String():
		logLevel = FatalLevel
	}
}

func (level LogLevel) String() string {
	switch level {
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	case FatalLevel:
		return "FATAL"
	}
	return ""
}

func (level LogLevel) StringWithColor() string {
	switch level {
	case DebugLevel:
		return color.WhiteString(level.String())
	case InfoLevel:
		return color.BlueString(level.String())
	case WarnLevel:
		return color.YellowString(level.String())
	case ErrorLevel:
		return color.RedString(level.String())
	case FatalLevel:
		return color.RedString(level.String())
	}
	return ""
}
