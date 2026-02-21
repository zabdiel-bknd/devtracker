package middlewares

import (
	"log/slog"
	"net/http"
	"time"
)

// A middleware for logging request and exetion time
func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
		start := time.Now()

		next.ServeHTTP(w, r)

		duration := time.Since(start)

		slog.Info("Incoming Request",
			"method", r.Method,
			"path", r.URL.Path,
			"duration", duration.String(),
			"ip", r.RemoteAddr,
		)
	})
}