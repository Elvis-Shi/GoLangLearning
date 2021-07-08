package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
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
		}

		err := auditStorage.Append(&auditLog)
		
		if err != nil {
			log.Println("Audit log append failed:", err)
		}
	}
}
