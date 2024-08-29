package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// customResponseWriter wraps the standard http.ResponseWriter to capture the status code
type customResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newCustomResponseWriter(w http.ResponseWriter) *customResponseWriter {
	// Default status code to 200 OK
	return &customResponseWriter{w, http.StatusOK}
}

func (rw *customResponseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// Logger middleware logs request details and uses different log levels based on the status code
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Start time
		start := time.Now()

		// Wrap the ResponseWriter to capture the status code
		crw := newCustomResponseWriter(w)

		// Serve the request
		next.ServeHTTP(crw, r)

		// Determine the log level based on the status code
		duration := time.Since(start)
		statusCode := crw.statusCode

		// Log entry with timestamp, log level, and request details
		logEntry := fmt.Sprintf("%s method=%s url=%s status=%d duration=%s",
			getLogLevel(statusCode),
			r.Method,
			r.URL.Path,
			statusCode,
			duration,
		)

		// Log the entry
		log.Println(logEntry)
	})
}

// getLogLevel returns the log level based on the status code
func getLogLevel(statusCode int) string {
	switch {
	case statusCode >= 500:
		return colorize("ERROR", "red")
	case statusCode >= 400:
		return colorize("WARN", "yellow")
	default:
		return colorize("INFO", "blue")
	}
}

// colorize returns the text wrapped in ANSI escape codes for coloring
func colorize(text, color string) string {
	var colorCode string
	switch color {
	case "red":
		colorCode = "\033[31m"
	case "green":
		colorCode = "\033[32m"
	case "yellow":
		colorCode = "\033[33m"
	case "blue":
		colorCode = "\033[1;34m"
	default:
		colorCode = "\033[0m" // Reset color
	}
	return fmt.Sprintf("%s%s\033[0m", colorCode, text)
}
