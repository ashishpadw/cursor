import React, { useState, useEffect } from 'react';
import { Link, useSearchParams } from 'react-router-dom';
import ProductCard from './ProductCard';

const Home = ({ addToCart }) => {
  const [products, setProducts] = useState([]);
  const [categories, setCategories] = useState([]);
  const [filteredProducts, setFilteredProducts] = useState([]);
  const [selectedGender, setSelectedGender] = useState('');
  const [selectedCategory, setSelectedCategory] = useState('');
  const [loading, setLoading] = useState(true);
  const [searchParams] = useSearchParams();

  useEffect(() => {
    fetchProducts();
    fetchCategories();
    
    // Check for gender filter in URL params
    const genderFromUrl = searchParams.get('gender');
    if (genderFromUrl) {
      setSelectedGender(genderFromUrl);
    }
  }, [searchParams]);

  useEffect(() => {
    filterProducts();
  }, [products, selectedGender, selectedCategory]);

  const fetchProducts = async () => {
    try {
      const response = await fetch('http://localhost:8080/api/products');
      const data = await response.json();
      setProducts(data);
      setLoading(false);
    } catch (error) {
      console.error('Error fetching products:', error);
      setLoading(false);
    }
  };

  const fetchCategories = async () => {
    try {
      const response = await fetch('http://localhost:8080/api/categories');
      const data = await response.json();
      setCategories(data);
    } catch (error) {
      console.error('Error fetching categories:', error);
    }
  };

  const filterProducts = () => {
    let filtered = products;

    if (selectedGender) {
      filtered = filtered.filter(product => product.gender === selectedGender);
    }

    if (selectedCategory) {
      filtered = filtered.filter(product => product.category === selectedCategory);
    }

    setFilteredProducts(filtered);
  };

  const handleQuickAdd = (product) => {
    // Quick add with default size and color
    const defaultSize = product.sizes && product.sizes.length > 0 ? product.sizes[0] : '';
    const defaultColor = product.colors && product.colors.length > 0 ? product.colors[0] : '';
    addToCart(product, defaultSize, defaultColor, 1);
  };

  if (loading) {
    return (
      <div style={{ textAlign: 'center', padding: '4rem 0' }}>
        <h2>Loading products...</h2>
      </div>
    );
  }

  return (
    <div>
      {/* Hero Section */}
      <section className="hero">
        <div className="container">
          <h1>Discover Your Style</h1>
          <p>
            Explore our curated collection of premium clothing for men and women. 
            From casual wear to formal attire, find the perfect pieces to express your unique style.
          </p>
        </div>
      </section>

      {/* Filters */}
      <section className="filters">
        <div className="container">
          <div className="filter-controls">
            <div className="filter-group">
              <label>Gender</label>
              <select 
                value={selectedGender} 
                onChange={(e) => setSelectedGender(e.target.value)}
              >
                <option value="">All</option>
                <option value="men">Men</option>
                <option value="women">Women</option>
              </select>
            </div>

            <div className="filter-group">
              <label>Category</label>
              <select 
                value={selectedCategory} 
                onChange={(e) => setSelectedCategory(e.target.value)}
              >
                <option value="">All Categories</option>
                {categories.map(category => (
                  <option key={category} value={category}>
                    {category.charAt(0).toUpperCase() + category.slice(1)}
                  </option>
                ))}
              </select>
            </div>

            {(selectedGender || selectedCategory) && (
              <button 
                onClick={() => {
                  setSelectedGender('');
                  setSelectedCategory('');
                }}
                className="btn btn-secondary"
              >
                Clear Filters
              </button>
            )}
          </div>
        </div>
      </section>

      {/* Products */}
      <section className="products-section">
        <div className="container">
          <h2 className="section-title">
            {selectedGender ? `${selectedGender.charAt(0).toUpperCase() + selectedGender.slice(1)}'s` : 'Our'} Collection
            {selectedCategory && ` - ${selectedCategory.charAt(0).toUpperCase() + selectedCategory.slice(1)}`}
          </h2>

          {filteredProducts.length === 0 ? (
            <div style={{ textAlign: 'center', padding: '3rem 0' }}>
              <h3>No products found</h3>
              <p>Try adjusting your filters to see more products.</p>
            </div>
          ) : (
            <div className="products-grid">
              {filteredProducts.map(product => (
                <ProductCard
                  key={product.id}
                  product={product}
                  onQuickAdd={handleQuickAdd}
                />
              ))}
            </div>
          )}
        </div>
      </section>
    </div>
  );
};

export default Home;