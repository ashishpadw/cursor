package services

import (
	"strings"
	
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
	filteredProducts := ps.products

	// Filter by gender if specified
	if gender != "" {
		filteredProducts = ps.filterByGender(filteredProducts, gender)
	}

	// Filter by category if specified
	if category != "" {
		filteredProducts = ps.filterByCategory(filteredProducts, category)
	}

	return filteredProducts
}

// GetProductByID returns a product by its ID
func (ps *ProductService) GetProductByID(id int) (*models.Product, bool) {
	for _, product := range ps.products {
		if product.ID == id {
			return &product, true
		}
	}
	return nil, false
}

// GetCategories returns all unique categories
func (ps *ProductService) GetCategories() []string {
	categoryMap := make(map[string]bool)
	for _, product := range ps.products {
		categoryMap[product.Category] = true
	}

	var categories []string
	for category := range categoryMap {
		categories = append(categories, category)
	}

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
	var filtered []models.Product
	for _, product := range ps.products {
		if product.Price >= minPrice && product.Price <= maxPrice {
			filtered = append(filtered, product)
		}
	}
	return filtered
}

// SearchProducts searches products by name or description
func (ps *ProductService) SearchProducts(query string) []models.Product {
	var filtered []models.Product
	for _, product := range ps.products {
		if containsIgnoreCase(product.Name, query) || containsIgnoreCase(product.Description, query) {
			filtered = append(filtered, product)
		}
	}
	return filtered
}

// Helper function to check if a string contains another string (case insensitive)
func containsIgnoreCase(str, substr string) bool {
	strLower := strings.ToLower(str)
	substrLower := strings.ToLower(substr)
	return strings.Contains(strLower, substrLower)
}