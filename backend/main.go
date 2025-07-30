package main

import (
	"fmt"
	"log"
	"net/http"

	"ecommerce-backend/config"
	"ecommerce-backend/routes"
	"github.com/rs/cors"
)

func main() {
	// Load application configuration
	cfg := config.LoadConfig()

	// Setup routes
	router := routes.SetupRoutes()

	// Configure CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: cfg.CORS.AllowedOrigins,
		AllowedMethods: cfg.CORS.AllowedMethods,
		AllowedHeaders: cfg.CORS.AllowedHeaders,
	})

	// Apply CORS middleware
	handler := corsHandler.Handler(router)

	// Start server
	serverAddr := cfg.GetServerAddress()
	fmt.Printf("🚀 E-commerce Backend Server starting on %s...\n", serverAddr)
	fmt.Printf("📱 Frontend URL: %s\n", cfg.CORS.AllowedOrigins[0])
	fmt.Println("📋 Available endpoints:")
	fmt.Println("   GET  /api/products")
	fmt.Println("   GET  /api/products/{id}")
	fmt.Println("   GET  /api/products/search?q={query}")
	fmt.Println("   GET  /api/products/price-range?min={min}&max={max}")
	fmt.Println("   GET  /api/categories")
	fmt.Println("   GET  /api/genders")
	
	log.Fatal(http.ListenAndServe(serverAddr, handler))
}