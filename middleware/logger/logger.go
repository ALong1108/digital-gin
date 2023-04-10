package logger

import (
	"digital-gin/library"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

const (
	traceID       = "Trace-Id"
	pathRawFormat = "%s?%s"
	defaultFormat = "%3d | %13v | %15s | %-7s %#v | trace:%s | goroutine:%s |\n%s"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Start timer
		var (
			startTime = time.Now()
			path      = ctx.Request.URL.Path
			raw       = ctx.Request.URL.RawQuery
		)
		if len(raw) > 0 {
			path = fmt.Sprintf(pathRawFormat, path, raw)
		}

		// Process request
		ctx.Next()

		// Stop timer
		var latency = time.Now().Sub(startTime)
		if latency > time.Minute {
			latency = latency.Truncate(time.Second)
		}

		Infof(defaultFormat,
			ctx.Writer.Status(),
			latency,
			ctx.ClientIP(),
			ctx.Request.Method, path,
			ctx.GetHeader(traceID),
			library.GoroutineID(),
			ctx.Errors.ByType(gin.ErrorTypePrivate).String())
	}
}

func Debug(v ...any) {
	if logLevel > debugLogLevel {
		return
	}
	output(debugLogLevel, v)
}

func Debugf(format string, v ...any) {
	if logLevel > debugLogLevel {
		return
	}
	output(debugLogLevel, fmt.Sprintf(format, v...))
}

func Info(v ...any) {
	if logLevel > infoLogLevel {
		return
	}
	output(infoLogLevel, v)
}

func Infof(format string, v ...any) {
	if logLevel > infoLogLevel {
		return
	}
	output(infoLogLevel, fmt.Sprintf(format, v...))
}

func Warn(v ...interface{}) {
	if logLevel > warnLogLevel {
		return
	}
	output(warnLogLevel, v)
}

func Warnf(format string, v ...any) {
	if logLevel > warnLogLevel {
		return
	}
	output(warnLogLevel, fmt.Sprintf(format, v...))
}

func Error(v ...any) {
	if logLevel > errorLogLevel {
		return
	}
	output(errorLogLevel, v)
}

func Errorf(format string, v ...any) {
	if logLevel > errorLogLevel {
		return
	}
	output(errorLogLevel, fmt.Sprintf(format, v...))
}

func Fatal(v ...any) {
	if logLevel > fatalLogLevel {
		return
	}
	output(fatalLogLevel, v)
}

func Fatalf(format string, v ...any) {
	if logLevel > fatalLogLevel {
		return
	}
	output(fatalLogLevel, fmt.Sprintf(format, v...))
}

const callDepth = 2

func output(level int8, v ...any) {
	if level > fatalLogLevel {
		return
	}

	defer gin.Recovery()

	var str string
	if len(v) == 1 {
		var ok bool
		if str, ok = v[0].(string); !ok {
			str = fmt.Sprintln(v...)
		}
	} else {
		str = fmt.Sprintln(v...)
	}

	_ = logInstance.l.Output(callDepth, logLevelToStrMap[level]+str)
}
