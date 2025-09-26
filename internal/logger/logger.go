package logger

import (
	"fmt"
	"log"
	"os"

	"github.com/to-the-moshpit/sandpit-turtle/internal/environment"
)

type Logger interface {
	Debug(string)
	Debugf(string, ...any)
	Info(string)
	Infof(string, ...any)
	Warn(string)
	Warnf(string, ...any)
	Error(string)
	Errorf(string, ...any)
}

type logger struct {
	level       LogLevel
	debugLogger *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
}

func (l *logger) logDebug(message string) {
	if l.level <= debugLevel {
		ci := getCallerInformation(3)
		l.debugLogger.Printf("%v@%v(%v): %v", ci.file, ci.lineNo, ci.funcName, message)
	}
}

func (l *logger) logInfo(message string) {
	if l.level <= infoLevel {
		l.infoLogger.Println(message)
	}
}

func (l *logger) logWarn(message string) {
	if l.level <= warnLevel {
		ci := getCallerInformation(3)
		l.warnLogger.Printf("%v: %v", ci.funcName, message)
	}
}

func (l *logger) logError(message string) {
	if l.level <= errorLevel {
		ci := getCallerInformation(3)
		l.errorLogger.Printf("%v@%v(%v): %v", ci.file, ci.lineNo, ci.funcName, message)
	}
}

// Debug implements Logger.
func (l *logger) Debug(message string) {
	l.logDebug(message)
}

// Debugf implements Logger.
func (l *logger) Debugf(format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	l.logDebug(message)
}

// Error implements Logger.
func (l *logger) Error(message string) {
	l.logError(message)
}

// Errorf implements Logger.
func (l *logger) Errorf(format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	l.logError(message)
}

// Info implements Logger.
func (l *logger) Info(message string) {
	l.logInfo(message)
}

// Infof implements Logger.
func (l *logger) Infof(format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	l.logInfo(message)
}

// Warn implements Logger.
func (l *logger) Warn(message string) {
	l.logWarn(message)
}

// Warnf implements Logger.
func (l *logger) Warnf(format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	l.logWarn(message)
}

var loggerInstance *logger

func New(env *environment.Env) Logger {
	if loggerInstance != nil {
		return loggerInstance
	}

	loggerInstance = &logger{
		level:       resolveLogLevel(env.LogLevel),
		debugLogger: log.New(os.Stdout, "[DEBUG] ", log.Default().Flags()),
		infoLogger:  log.New(os.Stdout, "[INFO]  ", log.Default().Flags()),
		warnLogger:  log.New(os.Stdout, "[WARN]  ", log.Default().Flags()),
		errorLogger: log.New(os.Stdout, "[ERROR] ", log.Default().Flags()),
	}

	return loggerInstance
}
