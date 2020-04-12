package logger

import (
	"fmt"
)

func Log(a ...interface{}) {
	fmt.Printf("\033[1;30m%v\n\033[0m", a...)
}

func Info(a ...interface{}) {
	fmt.Printf("\033[1;36m%v\n\033[0m", a...)
}

func Warn(a ...interface{}) {
	fmt.Printf("\033[1;33m%v\n\033[0m", a...)
}

func Error(a ...interface{}) {
	fmt.Printf("\033[1;31m%v\n\033[0m", a...)
}
