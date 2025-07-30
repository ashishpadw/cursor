package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"ecommerce-backend/logger"
	"ecommerce-backend/services"
	"github.com/gorilla/mux"
)

// ProductHandler handles HTTP requests for products
type ProductHandler struct {
	productService *services.ProductService
}

// NewProductHandler creates a new instance of ProductHandler
func NewProductHandler(productService *services.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

// GetProducts handles GET /api/products requests
func (ph *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Content-Type", "application/json")

	// Parse query parameters
	gender := r.URL.Query().Get("gender")
	category := r.URL.Query().Get("category")

	logger.Info("Handling get products request", map[string]interface{}{
		"handler":  "GetProducts",
		"gender":   gender,
		"category": category,
		"method":   r.Method,
		"path":     r.URL.Path,
	})

	// Get filtered products from service
	products := ph.productService.GetAllProducts(gender, category)

	// Encode and send response
	if err := json.NewEncoder(w).Encode(products); err != nil {
		duration := float64(time.Since(start).Nanoseconds()) / 1e6
		logger.LogError("handlers", "GetProducts", err, map[string]interface{}{
			"gender":      gender,
			"category":    category,
			"duration_ms": duration,
		})
		http.Error(w, "Failed to encode products", http.StatusInternalServerError)
		return
	}

	duration := float64(time.Since(start).Nanoseconds()) / 1e6
	logger.Info("Products request completed successfully", map[string]interface{}{
		"handler":      "GetProducts",
		"products_count": len(products),
		"gender":       gender,
		"category":     category,
		"duration_ms":  duration,
	})
}

// GetProduct handles GET /api/products/{id} requests
func (ph *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Content-Type", "application/json")
	
	// Extract product ID from URL parameters
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		duration := float64(time.Since(start).Nanoseconds()) / 1e6
		logger.LogError("handlers", "GetProduct", err, map[string]interface{}{
			"invalid_id":  idStr,
			"duration_ms": duration,
		})
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	logger.Info("Handling get product request", map[string]interface{}{
		"handler":    "GetProduct",
		"product_id": id,
		"method":     r.Method,
		"path":       r.URL.Path,
	})

	// Get product from service
	product, found := ph.productService.GetProductByID(id)
	if !found {
		duration := float64(time.Since(start).Nanoseconds()) / 1e6
		logger.Warn("Product not found in handler", map[string]interface{}{
			"handler":     "GetProduct",
			"product_id":  id,
			"duration_ms": duration,
		})
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	// Encode and send response
	if err := json.NewEncoder(w).Encode(product); err != nil {
		duration := float64(time.Since(start).Nanoseconds()) / 1e6
		logger.LogError("handlers", "GetProduct", err, map[string]interface{}{
			"product_id":  id,
			"duration_ms": duration,
		})
		http.Error(w, "Failed to encode product", http.StatusInternalServerError)
		return
	}

	duration := float64(time.Since(start).Nanoseconds()) / 1e6
	logger.Info("Product request completed successfully", map[string]interface{}{
		"handler":      "GetProduct",
		"product_id":   id,
		"product_name": product.Name,
		"duration_ms":  duration,
	})
}

// GetCategories handles GET /api/categories requests
func (ph *ProductHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Content-Type", "application/json")

	logger.Info("Handling get categories request", map[string]interface{}{
		"handler": "GetCategories",
		"method":  r.Method,
		"path":    r.URL.Path,
	})

	// Get categories from service
	categories := ph.productService.GetCategories()

	// Encode and send response
	if err := json.NewEncoder(w).Encode(categories); err != nil {
		duration := float64(time.Since(start).Nanoseconds()) / 1e6
		logger.LogError("handlers", "GetCategories", err, map[string]interface{}{
			"duration_ms": duration,
		})
		http.Error(w, "Failed to encode categories", http.StatusInternalServerError)
		return
	}

	duration := float64(time.Since(start).Nanoseconds()) / 1e6
	logger.Info("Categories request completed successfully", map[string]interface{}{
		"handler":          "GetCategories",
		"categories_count": len(categories),
		"duration_ms":      duration,
	})
}

// GetGenders handles GET /api/genders requests
func (ph *ProductHandler) GetGenders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get genders from service
	genders := ph.productService.GetGenders()

	// Encode and send response
	if err := json.NewEncoder(w).Encode(genders); err != nil {
		http.Error(w, "Failed to encode genders", http.StatusInternalServerError)
		return
	}
}

// SearchProducts handles GET /api/products/search requests
func (ph *ProductHandler) SearchProducts(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Content-Type", "application/json")

	// Get search query from URL parameters
	query := r.URL.Query().Get("q")
	if query == "" {
		duration := float64(time.Since(start).Nanoseconds()) / 1e6
		logger.Warn("Search request missing query parameter", map[string]interface{}{
			"handler":     "SearchProducts",
			"duration_ms": duration,
		})
		http.Error(w, "Search query is required", http.StatusBadRequest)
		return
	}

	logger.Info("Handling search products request", map[string]interface{}{
		"handler":      "SearchProducts",
		"search_query": query,
		"method":       r.Method,
		"path":         r.URL.Path,
	})

	// Search products using service
	products := ph.productService.SearchProducts(query)

	// Encode and send response
	if err := json.NewEncoder(w).Encode(products); err != nil {
		duration := float64(time.Since(start).Nanoseconds()) / 1e6
		logger.LogError("handlers", "SearchProducts", err, map[string]interface{}{
			"search_query": query,
			"duration_ms":  duration,
		})
		http.Error(w, "Failed to encode search results", http.StatusInternalServerError)
		return
	}

	duration := float64(time.Since(start).Nanoseconds()) / 1e6
	logger.Info("Search request completed successfully", map[string]interface{}{
		"handler":       "SearchProducts",
		"search_query":  query,
		"results_count": len(products),
		"duration_ms":   duration,
	})
}

// GetProductsByPriceRange handles GET /api/products/price-range requests
func (ph *ProductHandler) GetProductsByPriceRange(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Content-Type", "application/json")

	// Parse price range parameters
	minPriceStr := r.URL.Query().Get("min")
	maxPriceStr := r.URL.Query().Get("max")

	if minPriceStr == "" || maxPriceStr == "" {
		duration := float64(time.Since(start).Nanoseconds()) / 1e6
		logger.Warn("Price range request missing parameters", map[string]interface{}{
			"handler":     "GetProductsByPriceRange",
			"min_price":   minPriceStr,
			"max_price":   maxPriceStr,
			"duration_ms": duration,
		})
		http.Error(w, "Both min and max price parameters are required", http.StatusBadRequest)
		return
	}

	minPrice, err := strconv.ParseFloat(minPriceStr, 64)
	if err != nil {
		duration := float64(time.Since(start).Nanoseconds()) / 1e6
		logger.LogError("handlers", "GetProductsByPriceRange", err, map[string]interface{}{
			"invalid_min_price": minPriceStr,
			"duration_ms":       duration,
		})
		http.Error(w, "Invalid min price format", http.StatusBadRequest)
		return
	}

	maxPrice, err := strconv.ParseFloat(maxPriceStr, 64)
	if err != nil {
		duration := float64(time.Since(start).Nanoseconds()) / 1e6
		logger.LogError("handlers", "GetProductsByPriceRange", err, map[string]interface{}{
			"invalid_max_price": maxPriceStr,
			"duration_ms":       duration,
		})
		http.Error(w, "Invalid max price format", http.StatusBadRequest)
		return
	}

	if minPrice > maxPrice {
		duration := float64(time.Since(start).Nanoseconds()) / 1e6
		logger.Warn("Invalid price range", map[string]interface{}{
			"handler":     "GetProductsByPriceRange",
			"min_price":   minPrice,
			"max_price":   maxPrice,
			"duration_ms": duration,
		})
		http.Error(w, "Min price cannot be greater than max price", http.StatusBadRequest)
		return
	}

	logger.Info("Handling price range request", map[string]interface{}{
		"handler":   "GetProductsByPriceRange",
		"min_price": minPrice,
		"max_price": maxPrice,
		"method":    r.Method,
		"path":      r.URL.Path,
	})

	// Get products by price range from service
	products := ph.productService.GetProductsByPriceRange(minPrice, maxPrice)

	// Encode and send response
	if err := json.NewEncoder(w).Encode(products); err != nil {
		duration := float64(time.Since(start).Nanoseconds()) / 1e6
		logger.LogError("handlers", "GetProductsByPriceRange", err, map[string]interface{}{
			"min_price":   minPrice,
			"max_price":   maxPrice,
			"duration_ms": duration,
		})
		http.Error(w, "Failed to encode products", http.StatusInternalServerError)
		return
	}

	duration := float64(time.Since(start).Nanoseconds()) / 1e6
	logger.Info("Price range request completed successfully", map[string]interface{}{
		"handler":       "GetProductsByPriceRange",
		"min_price":     minPrice,
		"max_price":     maxPrice,
		"results_count": len(products),
		"duration_ms":   duration,
	})
}