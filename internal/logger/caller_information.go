package logger

import (
	"log"
	"runtime"
)

type callerInfo struct {
	file     string
	funcName string
	lineNo   int
}

func getCallerInformation(skip int) callerInfo {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		log.Println(("LOGGER: runtime.Caller() failed"))
		return callerInfo{}
	}

	funcName := runtime.FuncForPC(pc).Name()

	return callerInfo{file: file, funcName: funcName, lineNo: lineNo}
}
