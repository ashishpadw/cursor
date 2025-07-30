package logger

import (
	"os"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	// Logger is the global logger instance
	Logger *logrus.Logger
)

// LogConfig holds logging configuration
type LogConfig struct {
	Level  string `json:"level"`
	Format string `json:"format"` // "json" or "text"
	Output string `json:"output"` // "stdout", "stderr", or file path
}

// Init initializes the global logger with configuration
func Init(config LogConfig) {
	Logger = logrus.New()

	// Set log level
	level, err := logrus.ParseLevel(config.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	Logger.SetLevel(level)

	// Set formatter
	if config.Format == "json" {
		Logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
	} else {
		// Determine if we should use colors (only in true development with TTY)
		useColors := strings.ToLower(os.Getenv("ENV")) == "development" && 
					strings.ToLower(os.Getenv("LOG_COLORS")) != "false"
					
		Logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:          true,
			TimestampFormat:        "2006-01-02 15:04:05",
			ForceColors:            useColors,
			DisableColors:          !useColors,     // Ensure consistent output
			DisableQuote:           true,           // Don't quote field values
			DisableSorting:         true,           // Keep field order as specified
			PadLevelText:           true,           // Pad level text for alignment
			DisableLevelTruncation: true,          // Don't truncate level names
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				// Clean caller information formatting
				return "", ""
			},
		})
	}

	// Set output
	switch config.Output {
	case "stderr":
		Logger.SetOutput(os.Stderr)
	case "stdout":
		Logger.SetOutput(os.Stdout)
	default:
		// For file output or default to stdout
		Logger.SetOutput(os.Stdout)
	}
}

// GetDefaultConfig returns default logging configuration
func GetDefaultConfig() LogConfig {
	env := strings.ToLower(os.Getenv("ENV"))
	
	config := LogConfig{
		Level:  getEnvOrDefault("LOG_LEVEL", "info"),
		Output: "stdout",
	}

	// Use JSON format in production, text format in development
	if env == "production" || env == "prod" {
		config.Format = "json"
	} else {
		config.Format = "text"
	}

	return config
}

// getEnvOrDefault gets environment variable or returns default value
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// Structured logging helper functions

// Info logs an info message with optional fields
func Info(message string, fields ...logrus.Fields) {
	entry := Logger.Info
	if len(fields) > 0 {
		entry = Logger.WithFields(fields[0]).Info
	}
	entry(message)
}

// Warn logs a warning message with optional fields
func Warn(message string, fields ...logrus.Fields) {
	entry := Logger.Warn
	if len(fields) > 0 {
		entry = Logger.WithFields(fields[0]).Warn
	}
	entry(message)
}

// Error logs an error message with optional fields
func Error(message string, err error, fields ...logrus.Fields) {
	logFields := logrus.Fields{}
	if len(fields) > 0 {
		logFields = fields[0]
	}
	if err != nil {
		logFields["error"] = err.Error()
	}
	Logger.WithFields(logFields).Error(message)
}

// Debug logs a debug message with optional fields
func Debug(message string, fields ...logrus.Fields) {
	entry := Logger.Debug
	if len(fields) > 0 {
		entry = Logger.WithFields(fields[0]).Debug
	}
	entry(message)
}

// HTTP Request Logging Helpers

// LogHTTPRequest logs HTTP request details
func LogHTTPRequest(method, path, userAgent, clientIP string, statusCode int, duration float64) {
	Logger.WithFields(logrus.Fields{
		"method":      method,
		"path":        path,
		"status_code": statusCode,
		"duration_ms": duration,
		"user_agent":  userAgent,
		"client_ip":   clientIP,
		"type":        "http_request",
	}).Info("HTTP Request")
}

// LogServiceCall logs service layer method calls
func LogServiceCall(service, method string, params map[string]interface{}) {
	Logger.WithFields(logrus.Fields{
		"service": service,
		"method":  method,
		"params":  params,
		"type":    "service_call",
	}).Debug("Service method called")
}

// LogServiceResult logs service layer method results
func LogServiceResult(service, method string, resultCount int, duration float64) {
	Logger.WithFields(logrus.Fields{
		"service":      service,
		"method":       method,
		"result_count": resultCount,
		"duration_ms":  duration,
		"type":         "service_result",
	}).Debug("Service method completed")
}

// LogDatabaseQuery logs database query information
func LogDatabaseQuery(query string, duration float64, rowsAffected int64) {
	Logger.WithFields(logrus.Fields{
		"query":         query,
		"duration_ms":   duration,
		"rows_affected": rowsAffected,
		"type":          "database_query",
	}).Debug("Database query executed")
}

// LogError logs application errors with context
func LogError(component, operation string, err error, context map[string]interface{}) {
	fields := logrus.Fields{
		"component": component,
		"operation": operation,
		"error":     err.Error(),
		"type":      "application_error",
	}
	
	// Add context fields
	for key, value := range context {
		fields[key] = value
	}
	
	Logger.WithFields(fields).Error("Application error occurred")
}

// LogStartup logs application startup information
func LogStartup(component string, config map[string]interface{}) {
	Logger.WithFields(logrus.Fields{
		"component": component,
		"config":    config,
		"type":      "startup",
	}).Info("Component started")
}