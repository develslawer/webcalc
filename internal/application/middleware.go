package application

import (
	"bytes"
	"go.uber.org/zap"
	"io"
	"net/http"
	"time"
)

type responseRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rec *responseRecorder) WriteHeader(code int) {
	rec.statusCode = code
	rec.ResponseWriter.WriteHeader(code)
}

var logger, _ = zap.NewProduction()

func LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rec := &responseRecorder{ResponseWriter: w, statusCode: http.StatusOK}
		requestBody := ""
		if r.Body != nil {
			bodyBytes, err := io.ReadAll(r.Body)
			if err == nil {
				requestBody = string(bodyBytes)
				r.Body = io.NopCloser(bytes.NewReader(bodyBytes))
			}
		}

		logger.Info("Request received",
			zap.String("timestamp", start.Format(time.RFC3339)),
			zap.String("method", r.Method),
			zap.String("uri", r.RequestURI),
			zap.String("remote_addr", r.RemoteAddr),
			zap.String("body", requestBody),
		)

		defer func() {
			duration := time.Since(start)
			if rec.statusCode >= 400 {
				logger.Error("Request failed",
					zap.String("method", r.Method),
					zap.String("uri", r.RequestURI),
					zap.Int("status", rec.statusCode),
					zap.String("body", requestBody),
					zap.String("remote_addr", r.RemoteAddr),
					zap.Duration("duration", duration),
				)
			} else {
				logger.Info("Request completed",
					zap.String("method", r.Method),
					zap.String("uri", r.RequestURI),
					zap.Int("status", rec.statusCode),
					zap.String("body", requestBody),
					zap.String("remote_addr", r.RemoteAddr),
					zap.Duration("duration", duration),
				)
			}
		}()

		next.ServeHTTP(rec, r)
	})
}
