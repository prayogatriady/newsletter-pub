package middleware

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"newsletter-pub/utils/config"
	l "newsletter-pub/utils/logger"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	Cfg = config.AppCfg
)

const (
	YYYYMMDD       = "2006-01-02"
	YYYYMMDDhhmmss = "2006-01-02 15:04:05.000"
)

type ResponseWriterInterceptor struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}

func (w *ResponseWriterInterceptor) Write(b []byte) (int, error) {
	// Capture the response body
	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now()

		requestBody, err := io.ReadAll(c.Request.Body)
		if err != nil {
			// Handle error
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		defer c.Request.Body.Close()

		requestSlog := slog.Group("request",
			slog.String("method", c.Request.Method),
			slog.String("path", c.Request.URL.Path),
			slog.String("body", string(requestBody)),
			slog.Any("query", c.Request.URL.Query()),
		)

		// Create an interceptor that captures the response
		interceptor := &ResponseWriterInterceptor{
			ResponseWriter: c.Writer,
			Body:           &bytes.Buffer{},
		}
		c.Writer = interceptor // Replace the default writer

		// Run the handlers
		c.Next()

		responseBody := interceptor.Body.String()
		duration := time.Since(start)

		responseSlog := slog.Group("response",
			slog.Int("status", c.Writer.Status()),
			slog.String("duration", duration.String()),
			slog.String("body", responseBody),
		)

		l.Slogger.Info("LoggerMiddleware", slog.Group("log", requestSlog, responseSlog))
	}
}
