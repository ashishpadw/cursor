package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Product struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Price       float64  `json:"price"`
	Description string   `json:"description"`
	Category    string   `json:"category"`
	Gender      string   `json:"gender"`
	Image       string   `json:"image"`
	Images      []string `json:"images"`
	Sizes       []string `json:"sizes"`
	Colors      []string `json:"colors"`
	InStock     bool     `json:"inStock"`
}

var products = []Product{
	// Men's Clothing
	{
		ID:          1,
		Name:        "Classic Fit Cotton T-Shirt",
		Price:       29.99,
		Description: "Comfortable everyday t-shirt made from 100% premium cotton. Perfect for casual wear.",
		Category:    "t-shirts",
		Gender:      "men",
		Image:       "https://images.unsplash.com/photo-1521572163474-6864f9cf17ab?w=400",
		Images:      []string{"https://images.unsplash.com/photo-1521572163474-6864f9cf17ab?w=400", "https://images.unsplash.com/photo-1583743814966-8936f37f82e6?w=400"},
		Sizes:       []string{"S", "M", "L", "XL", "XXL"},
		Colors:      []string{"White", "Black", "Navy", "Gray"},
		InStock:     true,
	},
	{
		ID:          2,
		Name:        "Slim Fit Denim Jeans",
		Price:       79.99,
		Description: "Modern slim-fit jeans with premium denim fabric. Versatile and durable.",
		Category:    "jeans",
		Gender:      "men",
		Image:       "https://images.unsplash.com/photo-1542272604-787c3835535d?w=400",
		Images:      []string{"https://images.unsplash.com/photo-1542272604-787c3835535d?w=400", "https://images.unsplash.com/photo-1605518216938-7c31b7b14ad0?w=400"},
		Sizes:       []string{"28", "30", "32", "34", "36", "38"},
		Colors:      []string{"Dark Blue", "Light Blue", "Black"},
		InStock:     true,
	},
	{
		ID:          3,
		Name:        "Business Casual Button Shirt",
		Price:       59.99,
		Description: "Professional button-down shirt perfect for office wear or formal occasions.",
		Category:    "shirts",
		Gender:      "men",
		Image:       "https://images.unsplash.com/photo-1596755094514-f87e34085b2c?w=400",
		Images:      []string{"https://images.unsplash.com/photo-1596755094514-f87e34085b2c?w=400", "https://images.unsplash.com/photo-1602810318383-e386cc2a3ccf?w=400"},
		Sizes:       []string{"S", "M", "L", "XL"},
		Colors:      []string{"White", "Light Blue", "Pink", "Gray"},
		InStock:     true,
	},
	{
		ID:          4,
		Name:        "Comfort Fit Chinos",
		Price:       69.99,
		Description: "Versatile chino pants that work for both casual and semi-formal occasions.",
		Category:    "pants",
		Gender:      "men",
		Image:       "https://images.unsplash.com/photo-1473966968600-fa801b869a1a?w=400",
		Images:      []string{"https://images.unsplash.com/photo-1473966968600-fa801b869a1a?w=400", "https://images.unsplash.com/photo-1624378439575-d8705ad7ae80?w=400"},
		Sizes:       []string{"28", "30", "32", "34", "36"},
		Colors:      []string{"Khaki", "Navy", "Black", "Olive"},
		InStock:     true,
	},
	{
		ID:          5,
		Name:        "Wool Blend Sweater",
		Price:       89.99,
		Description: "Warm and comfortable sweater made from premium wool blend. Perfect for cooler weather.",
		Category:    "sweaters",
		Gender:      "men",
		Image:       "https://images.unsplash.com/photo-1571945153237-4929e783af4a?w=400",
		Images:      []string{"https://images.unsplash.com/photo-1571945153237-4929e783af4a?w=400", "https://images.unsplash.com/photo-1578662996442-48f60103fc96?w=400"},
		Sizes:       []string{"S", "M", "L", "XL"},
		Colors:      []string{"Charcoal", "Navy", "Burgundy", "Cream"},
		InStock:     true,
	},

	// Women's Clothing
	{
		ID:          6,
		Name:        "Floral Summer Dress",
		Price:       89.99,
		Description: "Beautiful floral dress perfect for summer occasions. Lightweight and flowing design.",
		Category:    "dresses",
		Gender:      "women",
		Image:       "https://images.unsplash.com/photo-1515372039744-b8f02a3ae446?w=400",
		Images:      []string{"https://images.unsplash.com/photo-1515372039744-b8f02a3ae446?w=400", "https://images.unsplash.com/photo-1494578819711-4bbcce5dbf76?w=400"},
		Sizes:       []string{"XS", "S", "M", "L", "XL"},
		Colors:      []string{"Pink Floral", "Blue Floral", "White Floral"},
		InStock:     true,
	},
	{
		ID:          7,
		Name:        "High-Waisted Skinny Jeans",
		Price:       75.99,
		Description: "Flattering high-waisted jeans with skinny fit. Made from stretch denim for comfort.",
		Category:    "jeans",
		Gender:      "women",
		Image:       "https://images.unsplash.com/photo-1541099649105-f69ad21f3246?w=400",
		Images:      []string{"https://images.unsplash.com/photo-1541099649105-f69ad21f3246?w=400", "https://images.unsplash.com/photo-1582418702059-97ebafb35d09?w=400"},
		Sizes:       []string{"24", "25", "26", "27", "28", "29", "30"},
		Colors:      []string{"Dark Blue", "Light Blue", "Black", "White"},
		InStock:     true,
	},
	{
		ID:          8,
		Name:        "Silk Blouse",
		Price:       119.99,
		Description: "Elegant silk blouse perfect for professional or formal settings. Luxurious feel and drape.",
		Category:    "blouses",
		Gender:      "women",
		Image:       "https://images.unsplash.com/photo-1485462537746-965f33f7f6a7?w=400",
		Images:      []string{"https://images.unsplash.com/photo-1485462537746-965f33f7f6a7?w=400", "https://images.unsplash.com/photo-1551698618-1dfe5d97d256?w=400"},
		Sizes:       []string{"XS", "S", "M", "L"},
		Colors:      []string{"Ivory", "Black", "Blush", "Navy"},
		InStock:     true,
	},
	{
		ID:          9,
		Name:        "Cozy Knit Cardigan",
		Price:       95.99,
		Description: "Soft and comfortable cardigan perfect for layering. Made from premium knit fabric.",
		Category:    "cardigans",
		Gender:      "women",
		Image:       "https://images.unsplash.com/photo-1544966503-7cc5ac882d5f?w=400",
		Images:      []string{"https://images.unsplash.com/photo-1544966503-7cc5ac882d5f?w=400", "https://images.unsplash.com/photo-1583496661160-fb5886a13c8c?w=400"},
		Sizes:       []string{"XS", "S", "M", "L", "XL"},
		Colors:      []string{"Beige", "Gray", "Black", "Cream"},
		InStock:     true,
	},
	{
		ID:          10,
		Name:        "Active Leggings",
		Price:       49.99,
		Description: "High-performance leggings for workouts or casual wear. Moisture-wicking and stretchy.",
		Category:    "activewear",
		Gender:      "women",
		Image:       "https://images.unsplash.com/photo-1506629905607-bb5bdd92ff5e?w=400",
		Images:      []string{"https://images.unsplash.com/photo-1506629905607-bb5bdd92ff5e?w=400", "https://images.unsplash.com/photo-1571019613454-1cb2f99b2d8b?w=400"},
		Sizes:       []string{"XS", "S", "M", "L", "XL"},
		Colors:      []string{"Black", "Navy", "Gray", "Purple"},
		InStock:     true,
	},
	{
		ID:          11,
		Name:        "Midi A-Line Skirt",
		Price:       65.99,
		Description: "Versatile midi skirt with flattering A-line silhouette. Perfect for work or weekend.",
		Category:    "skirts",
		Gender:      "women",
		Image:       "https://images.unsplash.com/photo-1594633312681-425c7b97ccd1?w=400",
		Images:      []string{"https://images.unsplash.com/photo-1594633312681-425c7b97ccd1?w=400", "https://images.unsplash.com/photo-1583496661160-fb5886a13c8c?w=400"},
		Sizes:       []string{"XS", "S", "M", "L"},
		Colors:      []string{"Black", "Navy", "Burgundy", "Camel"},
		InStock:     true,
	},
	{
		ID:          12,
		Name:        "Casual Cotton Top",
		Price:       39.99,
		Description: "Comfortable cotton top perfect for everyday wear. Soft fabric with relaxed fit.",
		Category:    "tops",
		Gender:      "women",
		Image:       "https://images.unsplash.com/photo-1434389677669-e08b4cac3105?w=400",
		Images:      []string{"https://images.unsplash.com/photo-1434389677669-e08b4cac3105?w=400", "https://images.unsplash.com/photo-1551698618-1dfe5d97d256?w=400"},
		Sizes:       []string{"XS", "S", "M", "L", "XL"},
		Colors:      []string{"White", "Black", "Pink", "Blue", "Green"},
		InStock:     true,
	},
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Parse query parameters
	gender := r.URL.Query().Get("gender")
	category := r.URL.Query().Get("category")
	
	filteredProducts := products
	
	// Filter by gender if specified
	if gender != "" {
		var filtered []Product
		for _, product := range filteredProducts {
			if product.Gender == gender {
				filtered = append(filtered, product)
			}
		}
		filteredProducts = filtered
	}
	
	// Filter by category if specified
	if category != "" {
		var filtered []Product
		for _, product := range filteredProducts {
			if product.Category == category {
				filtered = append(filtered, product)
			}
		}
		filteredProducts = filtered
	}
	
	json.NewEncoder(w).Encode(filteredProducts)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	
	for _, product := range products {
		if product.ID == id {
			json.NewEncoder(w).Encode(product)
			return
		}
	}
	
	http.Error(w, "Product not found", http.StatusNotFound)
}

func getCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	categoryMap := make(map[string]bool)
	for _, product := range products {
		categoryMap[product.Category] = true
	}
	
	var categories []string
	for category := range categoryMap {
		categories = append(categories, category)
	}
	
	json.NewEncoder(w).Encode(categories)
}

func main() {
	r := mux.NewRouter()
	
	// API routes
	r.HandleFunc("/api/products", getProducts).Methods("GET")
	r.HandleFunc("/api/products/{id}", getProduct).Methods("GET")
	r.HandleFunc("/api/categories", getCategories).Methods("GET")
	
	// Enable CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})
	
	handler := c.Handler(r)
	
	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}