package config

import (
	"os"
	"strconv"
)

// Config holds the application configuration
type Config struct {
	Server ServerConfig
	CORS   CORSConfig
}

// ServerConfig holds server-related configuration
type ServerConfig struct {
	Port int
	Host string
}

// CORSConfig holds CORS-related configuration
type CORSConfig struct {
	AllowedOrigins []string
	AllowedMethods []string
	AllowedHeaders []string
}

// LoadConfig loads configuration from environment variables with default values
func LoadConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Port: getEnvAsInt("PORT", 8080),
			Host: getEnv("HOST", ""),
		},
		CORS: CORSConfig{
			AllowedOrigins: []string{
				getEnv("FRONTEND_URL", "http://localhost:3000"),
				"http://localhost:3001", // Additional development URLs
			},
			AllowedMethods: []string{
				"GET", 
				"POST", 
				"PUT", 
				"DELETE", 
				"OPTIONS",
			},
			AllowedHeaders: []string{
				"*",
			},
		},
	}
}

// GetServerAddress returns the full server address
func (c *Config) GetServerAddress() string {
	if c.Server.Host == "" {
		return ":" + strconv.Itoa(c.Server.Port)
	}
	return c.Server.Host + ":" + strconv.Itoa(c.Server.Port)
}

// getEnv gets an environment variable with a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getEnvAsInt gets an environment variable as an integer with a default value
func getEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}