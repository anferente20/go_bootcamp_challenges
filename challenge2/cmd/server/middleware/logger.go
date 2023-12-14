package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// NewLogger returns a new logger middleware.
func NewLogger() *Logger {
	return &Logger{}
}

// Logger is a middleware that logs the request and response.
type Logger struct{}

// Log logs the request and response.
func (l *Logger) Log() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// before
		start := time.Now()

		// now
		ctx.Next()

		// after
		lapse := time.Since(start)
		// - request
		url := ctx.Request.URL.String()
		method := ctx.Request.Method
		reqContentSize := ctx.Request.ContentLength
		// - response
		statusCode := ctx.Writer.Status()
		resContentSize := ctx.Writer.Size()

		// log
		log.Printf("[%s] %s request-content-size: %d || code: %d response-content-size: %d time: %dns", method, url, reqContentSize, statusCode, resContentSize, lapse.Nanoseconds())
	}
}
