package logger

type logLevelType = int8

const (
	debugLogLevel logLevelType = iota
	infoLogLevel
	warnLogLevel
	errorLogLevel
	fatalLogLevel
)

const (
	debugStr = "DEBUG"
	infoStr  = "INFO"
	warnStr  = "WARN"
	errorStr = "ERROR"
	fatalStr = "FATAL"
)

var logLevelToStrMap = map[logLevelType]string{
	debugLogLevel: debugStr,
	infoLogLevel:  infoStr,
	warnLogLevel:  warnStr,
	errorLogLevel: errorStr,
	fatalLogLevel: fatalStr,
}

var logLevel int8
