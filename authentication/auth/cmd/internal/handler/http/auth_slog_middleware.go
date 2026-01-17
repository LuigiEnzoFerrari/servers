package handlers

import (
	"context"
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SlogMiddleware(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		traceID := uuid.New().String()
		c.Header("X-Trace-ID", traceID)

		ctx := context.WithValue(c.Request.Context(), "trace_id", traceID)
		c.Request = c.Request.WithContext(ctx)

		reqLogger := logger.With(
			slog.String("trace_id", traceID),
			slog.String("method", c.Request.Method),
			slog.String("path", c.Request.URL.Path),
		)

		c.Set("logger", reqLogger)


		status := c.Writer.Status()
		latency := time.Since(start)

		logLevel := slog.LevelInfo
		if status >= 500 {
			logLevel = slog.LevelError
		} else if status >= 400 {
			logLevel = slog.LevelWarn
		}

		reqLogger.Log(c.Request.Context(), logLevel, "request completed",
			slog.Int("status", status),
			slog.Duration("latency", latency),
		)
		
		c.Next()
	}
}