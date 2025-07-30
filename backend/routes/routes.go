package routes

import (
	"net/http"

	"ecommerce-backend/handlers"
	"ecommerce-backend/logger"
	"ecommerce-backend/middleware"
	"ecommerce-backend/services"
	"github.com/gorilla/mux"
)

// SetupRoutes configures all the routes for the application
func SetupRoutes() *mux.Router {
	logger.Info("Setting up routes", map[string]interface{}{
		"component": "routes",
	})

	// Initialize services
	productService := services.NewProductService()

	// Initialize handlers
	productHandler := handlers.NewProductHandler(productService)

	// Create router
	router := mux.NewRouter()

	// Apply global middleware
	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.RecoveryMiddleware)

	// Setup product routes
	setupProductRoutes(router, productHandler)

	logger.Info("Routes setup completed", map[string]interface{}{
		"component": "routes",
		"endpoints": []string{
			"GET /api/products",
			"GET /api/products/{id}",
			"GET /api/products/search",
			"GET /api/products/price-range",
			"GET /api/categories",
			"GET /api/genders",
		},
	})

	return router
}

// setupProductRoutes configures all product-related routes
func setupProductRoutes(router *mux.Router, productHandler *handlers.ProductHandler) {
	// Product API routes
	api := router.PathPrefix("/api").Subrouter()

	// Extended endpoints for better functionality (must come BEFORE parameterized routes)
	api.HandleFunc("/products/search", productHandler.SearchProducts).Methods("GET")
	api.HandleFunc("/products/price-range", productHandler.GetProductsByPriceRange).Methods("GET")
	
	// Core product endpoints
	api.HandleFunc("/products", productHandler.GetProducts).Methods("GET")
	api.HandleFunc("/products/{id:[0-9]+}", productHandler.GetProduct).Methods("GET")
	
	// Category and gender endpoints
	api.HandleFunc("/categories", productHandler.GetCategories).Methods("GET")
	api.HandleFunc("/genders", productHandler.GetGenders).Methods("GET")

	// Add OPTIONS method for CORS preflight requests
	api.HandleFunc("/products/search", optionsHandler).Methods("OPTIONS")
	api.HandleFunc("/products/price-range", optionsHandler).Methods("OPTIONS")
	api.HandleFunc("/products", optionsHandler).Methods("OPTIONS")
	api.HandleFunc("/products/{id:[0-9]+}", optionsHandler).Methods("OPTIONS")
	api.HandleFunc("/categories", optionsHandler).Methods("OPTIONS")
	api.HandleFunc("/genders", optionsHandler).Methods("OPTIONS")
}

// optionsHandler handles CORS preflight requests
func optionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.WriteHeader(http.StatusOK)
}