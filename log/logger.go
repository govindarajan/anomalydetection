package log

import (
	"fmt"
	"time"
)

func Debug(msg ...interface{}) {
	print(msg)
}

func Info(msg ...interface{}) {
	print(msg)
}

func Warn(msg ...interface{}) {
	print(msg)
}

func Error(msg ...interface{}) {
	print(msg)
}

func Critical(msg ...interface{}) {
	print(msg)
}

func print(msg ...interface{}) {
	fmt.Println(time.Now(), "Level", msg)
}

var logLevel, processName string

func InitLogger(process, level string) error {
	processName = process
	logLevel = level
	return nil
}
