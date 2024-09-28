package utils

import "log"

var logger = log.Default()

type Logger struct{}

func (l Logger) Debug(msg string) {
	logger.Print(msg)
}

func (l Logger) Info(msg string) {
	logger.Print(msg)
}
