package models

import (
	"fmt"
	"os"
	"time"
)

type TraceLevel int

const (
    Error TraceLevel = iota
    Warning
    Info
    Debug
)

type TraceLog struct {
	ServiceId string
	Content string
	Level TraceLevel
}

type TraceLogger interface {
	Append(log *TraceLog)
}

type FileTraceStorage struct {
}

// TODO: Just a way to get understand Go's IO operation, need to figure out a standard way to write log later.
var traceLog, _ = os.OpenFile("Trace.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // TODO: what is 0600.

var traceLogger = FileTraceStorage {}

func (storage FileTraceStorage) Append(log *TraceLog)  {
	if traceLog == nil {
		panic("Failed to open trace log file.")
	}

	traceLog.WriteString(fmt.Sprintf("%v\t%v\t%v\n", time.Now().Format("2006/01/02 03:04:05.99999pm"), log.ServiceId, log.Level, log.Content))
}

func Log(log *TraceLog) {
	traceLogger.Append(log)
}

