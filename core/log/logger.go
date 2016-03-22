package log

import (
	"fmt"
	"net/http"
	"time"

	logrus "github.com/Sirupsen/logrus"
)

const (
	DEBUG LogLevel = 0
	INFO  LogLevel = 1
	WARN  LogLevel = 2
	FATAL LogLevel = 3
)

type LogLevel int

var _logLevel LogLevel = 1

var log = logrus.New()

func SetLogLevel(logLevel LogLevel) {
	_logLevel = logLevel

	log.Level = logrus.DebugLevel
}

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)
		Info(fmt.Sprintf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		))
	})
}

func Debug(message string, obj ...interface{}) {
	if _logLevel == 0 {
		if len(obj) > 0 {
			log.Debug(message, fmt.Sprintf("%s", obj))
		} else {
			log.Debug(message)
		}
	}
}

func Info(message string, obj ...interface{}) {
	if _logLevel <= 1 {
		if len(obj) > 0 {
			log.Info(message, fmt.Sprintf("%s", obj))
		} else {
			log.Info(message)
		}
	}
}

func Warn(message string, err ...interface{}) {
	if _logLevel <= 2 {
		if len(err) > 0 {
			log.Warn(message, fmt.Sprintf("%s", err))
		} else {
			log.Warn(message)
		}
	}
}

func Fatal(message string, err ...interface{}) {
	if _logLevel <= 3 {
		if len(err) > 0 {
			log.Fatal(message, fmt.Sprintf("%s", err))
		} else {
			log.Fatal(message)
		}
	}
}
