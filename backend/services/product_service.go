package services

import (
	"strings"
	"time"
	
	"ecommerce-backend/logger"
	"ecommerce-backend/models"
)

// ProductService handles all product-related business logic
type ProductService struct {
	products []models.Product
}

// NewProductService creates a new instance of ProductService
func NewProductService() *ProductService {
	return &ProductService{
		products: models.GetSampleProducts(),
	}
}

// GetAllProducts returns all products with optional filtering
func (ps *ProductService) GetAllProducts(gender, category string) []models.Product {
	start := time.Now()
	
	// Log service call
	params := map[string]interface{}{
		"gender":   gender,
		"category": category,
	}
	logger.LogServiceCall("ProductService", "GetAllProducts", params)

	filteredProducts := ps.products

	// Filter by gender if specified
	if gender != "" {
		filteredProducts = ps.filterByGender(filteredProducts, gender)
		logger.Debug("Applied gender filter", map[string]interface{}{
			"gender":           gender,
			"filtered_count":   len(filteredProducts),
			"original_count":   len(ps.products),
		})
	}

	// Filter by category if specified
	if category != "" {
		originalCount := len(filteredProducts)
		filteredProducts = ps.filterByCategory(filteredProducts, category)
		logger.Debug("Applied category filter", map[string]interface{}{
			"category":         category,
			"filtered_count":   len(filteredProducts),
			"before_filter":    originalCount,
		})
	}

	// Log result
	duration := float64(time.Since(start).Nanoseconds()) / 1e6
	logger.LogServiceResult("ProductService", "GetAllProducts", len(filteredProducts), duration)

	logger.Info("Products retrieved successfully", map[string]interface{}{
		"count":    len(filteredProducts),
		"gender":   gender,
		"category": category,
	})

	return filteredProducts
}

// GetProductByID returns a product by its ID
func (ps *ProductService) GetProductByID(id int) (*models.Product, bool) {
	start := time.Now()
	
	// Log service call
	params := map[string]interface{}{
		"product_id": id,
	}
	logger.LogServiceCall("ProductService", "GetProductByID", params)

	for _, product := range ps.products {
		if product.ID == id {
			duration := float64(time.Since(start).Nanoseconds()) / 1e6
			logger.LogServiceResult("ProductService", "GetProductByID", 1, duration)
			
			logger.Info("Product found", map[string]interface{}{
				"product_id":   id,
				"product_name": product.Name,
				"category":     product.Category,
			})
			
			return &product, true
		}
	}
	
	duration := float64(time.Since(start).Nanoseconds()) / 1e6
	logger.LogServiceResult("ProductService", "GetProductByID", 0, duration)
	
	logger.Warn("Product not found", map[string]interface{}{
		"product_id": id,
	})
	
	return nil, false
}

// GetCategories returns all unique categories
func (ps *ProductService) GetCategories() []string {
	start := time.Now()
	
	logger.LogServiceCall("ProductService", "GetCategories", map[string]interface{}{})

	categoryMap := make(map[string]bool)
	for _, product := range ps.products {
		categoryMap[product.Category] = true
	}

	var categories []string
	for category := range categoryMap {
		categories = append(categories, category)
	}

	duration := float64(time.Since(start).Nanoseconds()) / 1e6
	logger.LogServiceResult("ProductService", "GetCategories", len(categories), duration)

	logger.Info("Categories retrieved", map[string]interface{}{
		"categories_count": len(categories),
		"categories":       categories,
	})

	return categories
}

// GetGenders returns all unique genders
func (ps *ProductService) GetGenders() []string {
	genderMap := make(map[string]bool)
	for _, product := range ps.products {
		genderMap[product.Gender] = true
	}

	var genders []string
	for gender := range genderMap {
		genders = append(genders, gender)
	}

	return genders
}

// filterByGender filters products by gender
func (ps *ProductService) filterByGender(products []models.Product, gender string) []models.Product {
	var filtered []models.Product
	for _, product := range products {
		if product.Gender == gender {
			filtered = append(filtered, product)
		}
	}
	return filtered
}

// filterByCategory filters products by category
func (ps *ProductService) filterByCategory(products []models.Product, category string) []models.Product {
	var filtered []models.Product
	for _, product := range products {
		if product.Category == category {
			filtered = append(filtered, product)
		}
	}
	return filtered
}

// GetProductsByPriceRange returns products within a price range
func (ps *ProductService) GetProductsByPriceRange(minPrice, maxPrice float64) []models.Product {
	start := time.Now()
	
	params := map[string]interface{}{
		"min_price": minPrice,
		"max_price": maxPrice,
	}
	logger.LogServiceCall("ProductService", "GetProductsByPriceRange", params)

	var filtered []models.Product
	for _, product := range ps.products {
		if product.Price >= minPrice && product.Price <= maxPrice {
			filtered = append(filtered, product)
		}
	}

	duration := float64(time.Since(start).Nanoseconds()) / 1e6
	logger.LogServiceResult("ProductService", "GetProductsByPriceRange", len(filtered), duration)

	logger.Info("Price range filter applied", map[string]interface{}{
		"min_price":     minPrice,
		"max_price":     maxPrice,
		"results_count": len(filtered),
		"total_products": len(ps.products),
	})

	return filtered
}

// SearchProducts searches products by name or description
func (ps *ProductService) SearchProducts(query string) []models.Product {
	start := time.Now()
	
	params := map[string]interface{}{
		"search_query": query,
	}
	logger.LogServiceCall("ProductService", "SearchProducts", params)

	var filtered []models.Product
	for _, product := range ps.products {
		if containsIgnoreCase(product.Name, query) || containsIgnoreCase(product.Description, query) {
			filtered = append(filtered, product)
		}
	}

	duration := float64(time.Since(start).Nanoseconds()) / 1e6
	logger.LogServiceResult("ProductService", "SearchProducts", len(filtered), duration)

	logger.Info("Product search completed", map[string]interface{}{
		"search_query":  query,
		"results_count": len(filtered),
		"total_products": len(ps.products),
	})

	return filtered
}

// Helper function to check if a string contains another string (case insensitive)
func containsIgnoreCase(str, substr string) bool {
	strLower := strings.ToLower(str)
	substrLower := strings.ToLower(substr)
	return strings.Contains(strLower, substrLower)
}