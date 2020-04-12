package logger

import (
	"fmt"
)

type Logger struct {
}

func (l *Logger) Log(a ...interface{}) {
	fmt.Printf("\033[1;37m%v\033[0m", a)
}

func (l *Logger) Info(a ...interface{}) {
	fmt.Printf("\033[1;36m%v\033[0m", a)
}

func (l *Logger) Warn(a ...interface{}) {
	fmt.Printf("\033[1;33m%v\033[0m", a)
}

func (l *Logger) Error(a ...interface{}) {
	fmt.Printf("\033[1;31m%v\033[0m", a)
}
