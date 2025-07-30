package main

import (
	"fmt"
	"log"
	"net/http"

	"ecommerce-backend/config"
	"ecommerce-backend/logger"
	"ecommerce-backend/routes"
	"github.com/rs/cors"
)

func main() {
	// Initialize structured logging
	logConfig := logger.GetDefaultConfig()
	logger.Init(logConfig)

	logger.LogStartup("main", map[string]interface{}{
		"log_level":  logConfig.Level,
		"log_format": logConfig.Format,
		"log_output": logConfig.Output,
	})

	// Load application configuration
	cfg := config.LoadConfig()

	logger.LogStartup("config", map[string]interface{}{
		"server_port":      cfg.Server.Port,
		"server_host":      cfg.Server.Host,
		"cors_origins":     cfg.CORS.AllowedOrigins,
		"cors_methods":     cfg.CORS.AllowedMethods,
	})

	// Setup routes
	router := routes.SetupRoutes()

	// Configure CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: cfg.CORS.AllowedOrigins,
		AllowedMethods: cfg.CORS.AllowedMethods,
		AllowedHeaders: cfg.CORS.AllowedHeaders,
	})

	logger.Info("CORS configured", map[string]interface{}{
		"allowed_origins": cfg.CORS.AllowedOrigins,
		"allowed_methods": cfg.CORS.AllowedMethods,
	})

	// Apply CORS middleware
	handler := corsHandler.Handler(router)

	// Start server
	serverAddr := cfg.GetServerAddress()
	
	// Log startup information
	logger.Info("🚀 E-commerce Backend Server starting", map[string]interface{}{
		"server_address": serverAddr,
		"frontend_url":   cfg.CORS.AllowedOrigins[0],
		"endpoints": map[string]string{
			"products":           "GET /api/products",
			"product_by_id":      "GET /api/products/{id}",
			"search_products":    "GET /api/products/search?q={query}",
			"price_range":        "GET /api/products/price-range?min={min}&max={max}",
			"categories":         "GET /api/categories",
			"genders":            "GET /api/genders",
		},
	})

	// Also print to console for development visibility
	fmt.Printf("🚀 E-commerce Backend Server starting on %s...\n", serverAddr)
	fmt.Printf("📱 Frontend URL: %s\n", cfg.CORS.AllowedOrigins[0])
	fmt.Printf("📋 Available endpoints:\n")
	fmt.Printf("   GET  /api/products\n")
	fmt.Printf("   GET  /api/products/{id}\n")
	fmt.Printf("   GET  /api/products/search?q={query}\n")
	fmt.Printf("   GET  /api/products/price-range?min={min}&max={max}\n")
	fmt.Printf("   GET  /api/categories\n")
	fmt.Printf("   GET  /api/genders\n")
	fmt.Printf("🔍 Log Level: %s | Format: %s\n", logConfig.Level, logConfig.Format)
	
	// Start the server with error logging
	logger.Info("Starting HTTP server", map[string]interface{}{
		"address": serverAddr,
	})

	if err := http.ListenAndServe(serverAddr, handler); err != nil {
		logger.LogError("main", "server_startup", err, map[string]interface{}{
			"server_address": serverAddr,
		})
		log.Fatal(err)
	}
}