package controllers

import (
	"github.com/gin-gonic/gin"
)

type RequestLogger interface {
	LogRequest(c *gin.Context) string
	LogResponse(requestId string, c *gin.Context)
}

type FileLogger struct{}

// TODO: figure out how to persist request and response.
func (logger FileLogger) LogRequest(c *gin.Context) string {
	// return requestId, so LogResponse can use for correlation.

	// TODO: figure out how to generate GUID.
	return ""
}

func (logger FileLogger) LogResponse(requestId string, c *gin.Context) {

}
