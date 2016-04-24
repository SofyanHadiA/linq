package utils

import (
	"fmt"
	"net/http"
	"time"

	Logrus "github.com/Sirupsen/logrus"
)

var logrus = Logrus.New()

type Logger struct {
	logLevel int
}

var Log = Logger{logLevel: 0}

func SetLogLevel(logLevel int) Logger {
	logrus.Level = Logrus.DebugLevel
	Log = Logger{logLevel: logLevel}

	return Log
}

func (log Logger) LogHttp(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)
		log.Info(fmt.Sprintf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		))
	})
}

func (log Logger) Debug(message string, obj ...interface{}) {
	if log.logLevel == 0 {
		if len(obj) > 0 {
			logrus.Debug(message, fmt.Sprintf("%s", obj))
		} else {
			logrus.Debug(message)
		}
	}
}

func (log Logger) Info(message string, obj ...interface{}) {
	if log.logLevel <= 1 {
		if len(obj) > 0 {
			logrus.Info(message, fmt.Sprintf("%s", obj))
		} else {
			logrus.Info(message)
		}
	}
}

func (log Logger) Warn(message string, err ...interface{}) {
	if log.logLevel <= 2 {
		if len(err) > 0 {
			logrus.Warn(message, fmt.Sprintf("%s", err))
		} else {
			logrus.Warn(message)
		}
	}
}

func (log Logger) Fatal(message string, err ...interface{}) {
	if log.logLevel <= 3 {
		if len(err) > 0 {
			logrus.Fatal(message, fmt.Sprintf("%s", err))
		} else {
			logrus.Fatal(message)
		}
	}
}
