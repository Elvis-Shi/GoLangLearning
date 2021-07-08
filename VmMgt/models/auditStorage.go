package models

import (
	"fmt"
	"os"
	"time"
)

type AuditLog struct {
	StartTime      *time.Time
	URI            string
	Verb           string
	ResponseStatus int
	Duration       time.Duration
}

type AuditStorage interface {
	Append(log *AuditLog) error
}

type FileAuditStorage struct {
}

// TODO: Just a way to get understand Go's IO operation, need to figure out a standard way to write log later.
// global handle of audit log variable.
var auditLog, _ = os.OpenFile("Audit.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // TODO: what is 0600.

func (storage FileAuditStorage) Append(log *AuditLog) error {
	if auditLog == nil {
		return fmt.Errorf("Failed to open audit log file.")
	}

	auditLog.WriteString(fmt.Sprintf("%v\t%v\t%v\t%v\t%v\n", log.StartTime.Format("2006/01/02 03:04:05.99999pm"), log.URI, log.Verb, log.ResponseStatus, log.Duration.Nanoseconds()))
	return nil
}
