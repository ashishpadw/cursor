package middleware

import (
	"net/http"
	"strings"
	"time"

	"ecommerce-backend/logger"
)

// responseWriter wraps http.ResponseWriter to capture status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
	written    bool
}

func (rw *responseWriter) WriteHeader(code int) {
	if !rw.written {
		rw.statusCode = code
		rw.written = true
		rw.ResponseWriter.WriteHeader(code)
	}
}

func (rw *responseWriter) Write(data []byte) (int, error) {
	if !rw.written {
		rw.statusCode = http.StatusOK
		rw.written = true
	}
	return rw.ResponseWriter.Write(data)
}

// LoggingMiddleware logs HTTP requests and responses
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Wrap the response writer to capture status code
		wrapped := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
			written:        false,
		}

		// Get client IP
		clientIP := getClientIP(r)

		// Call the next handler
		next.ServeHTTP(wrapped, r)

		// Calculate duration
		duration := float64(time.Since(start).Nanoseconds()) / 1e6 // Convert to milliseconds

		// Log the request
		logger.LogHTTPRequest(
			r.Method,
			r.URL.Path,
			r.UserAgent(),
			clientIP,
			wrapped.statusCode,
			duration,
		)

		// Log query parameters if present (for debugging)
		if len(r.URL.RawQuery) > 0 {
			logger.Debug("HTTP Request Query Parameters", map[string]interface{}{
				"method":      r.Method,
				"path":        r.URL.Path,
				"query":       r.URL.RawQuery,
				"client_ip":   clientIP,
				"status_code": wrapped.statusCode,
			})
		}
	})
}

// getClientIP extracts the client IP address from the request
func getClientIP(r *http.Request) string {
	// Check for X-Forwarded-For header (proxy/load balancer)
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		// X-Forwarded-For can contain multiple IPs, get the first one
		ips := strings.Split(forwarded, ",")
		if len(ips) > 0 {
			return strings.TrimSpace(ips[0])
		}
	}

	// Check for X-Real-IP header (nginx proxy)
	realIP := r.Header.Get("X-Real-IP")
	if realIP != "" {
		return realIP
	}

	// Fall back to RemoteAddr
	ip := r.RemoteAddr
	// Remove port if present
	if idx := strings.LastIndex(ip, ":"); idx != -1 {
		ip = ip[:idx]
	}
	
	// Remove brackets for IPv6
	ip = strings.Trim(ip, "[]")
	
	return ip
}

// Recovery middleware with logging
func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger.LogError("middleware", "panic_recovery", 
					nil, // We don't have an actual error object
					map[string]interface{}{
						"panic":      err,
						"method":     r.Method,
						"path":       r.URL.Path,
						"client_ip":  getClientIP(r),
						"user_agent": r.UserAgent(),
					})

				// Return 500 Internal Server Error
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}