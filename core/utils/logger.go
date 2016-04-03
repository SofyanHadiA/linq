package utils

import (
	"fmt"
	"net/http"
	"time"

	logrus "github.com/Sirupsen/logrus"
)


var log = logrus.New()

type Logger struct{
	logLevel int
}

var Log = Logger{logLevel : 2}

func NewLogger(logLevel int) Logger{
	log.Level = logrus.DebugLevel

	Log = Logger{logLevel: logLevel}
	
	return Log
}


func(logger Logger) LogHttp(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)
		logger.Info(fmt.Sprintf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		))
	})
}

func(logger Logger) Debug(message string, obj ...interface{}) {
	if logger.logLevel == 0 {
		if len(obj) > 0 {
			log.Debug(message, fmt.Sprintf("%s", obj))
		} else {
			log.Debug(message)
		}
	}
}

func(logger Logger) Info(message string, obj ...interface{}) {
	if logger.logLevel <= 1 {
		if len(obj) > 0 {
			log.Info(message, fmt.Sprintf("%s", obj))
		} else {
			log.Info(message)
		}
	}
}

func(logger Logger) Warn(message string, err ...interface{}) {
	if logger.logLevel <= 2 {
		if len(err) > 0 {
			log.Warn(message, fmt.Sprintf("%s", err))
		} else {
			log.Warn(message)
		}
	}
}

func(logger Logger) Fatal(message string, err ...interface{}) {
	if logger.logLevel <= 3 {
		if len(err) > 0 {
			log.Fatal(message, fmt.Sprintf("%s", err))
		} else {
			log.Fatal(message)
		}
	}
}
