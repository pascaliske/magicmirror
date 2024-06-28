package logger

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func Debug(message string, args ...interface{}) {
	logMessage(DebugLevel, message, args)
}

func Info(message string, args ...interface{}) {
	logMessage(InfoLevel, message, args)
}

func Warn(message string, args ...interface{}) {
	logMessage(WarnLevel, message, args)
}

func Error(message string, args ...interface{}) {
	logMessage(ErrorLevel, message, args)
}

func Fatal(message string, args ...interface{}) {
	logMessage(FatalLevel, message, args)
}

func Raw(message string, args ...interface{}) {
	fmt.Printf("%s\n", fmt.Sprintf(message, args...))
}

func logMessage(level LogLevel, message string, args []interface{}) {
	// skip log if below level
	if level.Disabled() {
		return
	}

	// colorize args
	for i, arg := range args {
		args[i] = color.CyanString("%v", arg)
	}

	// prepare formatted time
	timeValue := time.Now().Local().Format(time.RFC3339)
	levelValue := level.StringWithColor()
	messageValue := fmt.Sprintf(message, args...)

	// output message
	fmt.Printf("%s [%s] %s\n", timeValue, levelValue, messageValue)
}
