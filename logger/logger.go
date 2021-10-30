package logger

import "log"

func Errorf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func Infof(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func Debugf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}
