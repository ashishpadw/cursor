package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

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
	w.Header().Set("Content-Type", "application/json")

	// Parse query parameters
	gender := r.URL.Query().Get("gender")
	category := r.URL.Query().Get("category")

	// Get filtered products from service
	products := ph.productService.GetAllProducts(gender, category)

	// Encode and send response
	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, "Failed to encode products", http.StatusInternalServerError)
		return
	}
}

// GetProduct handles GET /api/products/{id} requests
func (ph *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Extract product ID from URL parameters
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// Get product from service
	product, found := ph.productService.GetProductByID(id)
	if !found {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	// Encode and send response
	if err := json.NewEncoder(w).Encode(product); err != nil {
		http.Error(w, "Failed to encode product", http.StatusInternalServerError)
		return
	}
}

// GetCategories handles GET /api/categories requests
func (ph *ProductHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get categories from service
	categories := ph.productService.GetCategories()

	// Encode and send response
	if err := json.NewEncoder(w).Encode(categories); err != nil {
		http.Error(w, "Failed to encode categories", http.StatusInternalServerError)
		return
	}
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
	w.Header().Set("Content-Type", "application/json")

	// Get search query from URL parameters
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Search query is required", http.StatusBadRequest)
		return
	}

	// Search products using service
	products := ph.productService.SearchProducts(query)

	// Encode and send response
	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, "Failed to encode search results", http.StatusInternalServerError)
		return
	}
}

// GetProductsByPriceRange handles GET /api/products/price-range requests
func (ph *ProductHandler) GetProductsByPriceRange(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Parse price range parameters
	minPriceStr := r.URL.Query().Get("min")
	maxPriceStr := r.URL.Query().Get("max")

	if minPriceStr == "" || maxPriceStr == "" {
		http.Error(w, "Both min and max price parameters are required", http.StatusBadRequest)
		return
	}

	minPrice, err := strconv.ParseFloat(minPriceStr, 64)
	if err != nil {
		http.Error(w, "Invalid min price format", http.StatusBadRequest)
		return
	}

	maxPrice, err := strconv.ParseFloat(maxPriceStr, 64)
	if err != nil {
		http.Error(w, "Invalid max price format", http.StatusBadRequest)
		return
	}

	if minPrice > maxPrice {
		http.Error(w, "Min price cannot be greater than max price", http.StatusBadRequest)
		return
	}

	// Get products by price range from service
	products := ph.productService.GetProductsByPriceRange(minPrice, maxPrice)

	// Encode and send response
	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, "Failed to encode products", http.StatusInternalServerError)
		return
	}
}