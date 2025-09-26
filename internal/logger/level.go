package logger

import "log"

const (
	debug string = "debug"
	info  string = "info"
	warn  string = "warn"
	err   string = "error"
	none  string = "none"
)

type LogLevel = int

const (
	debugLevel LogLevel = iota
	infoLevel
	warnLevel
	errorLevel
	noneLevel
)

func resolveLogLevel(level string) LogLevel {
	var logLevel LogLevel
	switch level {
	case string(debug):
		logLevel = debugLevel
	case string(info):
		logLevel = infoLevel
	case string(warn):
		logLevel = warnLevel
	case string(err):
		logLevel = errorLevel
	case string(none):
		logLevel = noneLevel
	default:
		logLevel = infoLevel
	}

	log.Printf("LogLevel: %v=%v", level, logLevel)

	return logLevel
}
