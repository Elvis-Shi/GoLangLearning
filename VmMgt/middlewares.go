package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"ssdd.com/vms/models"
)
var auditStorage = models.FileAuditStorage{}

func Audit() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		auditLog := models.AuditLog {
			StartTime: &startTime,
			URI: c.Request.URL.String(),
			Verb: c.Request.Method,
			ResponseStatus: c.Writer.Status(),
			Duration: time.Since(startTime),
			ServiceId: c.GetString("ServiceId"),
		}

		auditStorage.Append(&auditLog)
	}
}

func AssignServiceId() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("ServiceId", uuid.NewString());
		c.Header("ServiceId", c.GetString("ServiceId"));
	}
}
