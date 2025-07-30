import React, { useState, useEffect } from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import './index.css';
import Header from './components/Header';
import Home from './components/Home';
import ProductDetail from './components/ProductDetail';
import Cart from './components/Cart';

function App() {
  const [cartItems, setCartItems] = useState([]);
  const [showCart, setShowCart] = useState(false);

  // Load cart from localStorage on mount
  useEffect(() => {
    const savedCart = localStorage.getItem('cart');
    if (savedCart) {
      setCartItems(JSON.parse(savedCart));
    }
  }, []);

  // Save cart to localStorage whenever it changes
  useEffect(() => {
    localStorage.setItem('cart', JSON.stringify(cartItems));
  }, [cartItems]);

  const addToCart = (product, selectedSize, selectedColor, quantity = 1) => {
    const existingItem = cartItems.find(
      item => 
        item.id === product.id && 
        item.selectedSize === selectedSize && 
        item.selectedColor === selectedColor
    );

    if (existingItem) {
      setCartItems(cartItems.map(item =>
        item.id === product.id && 
        item.selectedSize === selectedSize && 
        item.selectedColor === selectedColor
          ? { ...item, quantity: item.quantity + quantity }
          : item
      ));
    } else {
      setCartItems([...cartItems, {
        ...product,
        selectedSize,
        selectedColor,
        quantity
      }]);
    }
  };

  const removeFromCart = (id, selectedSize, selectedColor) => {
    setCartItems(cartItems.filter(item => 
      !(item.id === id && item.selectedSize === selectedSize && item.selectedColor === selectedColor)
    ));
  };

  const updateQuantity = (id, selectedSize, selectedColor, newQuantity) => {
    if (newQuantity <= 0) {
      removeFromCart(id, selectedSize, selectedColor);
      return;
    }

    setCartItems(cartItems.map(item =>
      item.id === id && item.selectedSize === selectedSize && item.selectedColor === selectedColor
        ? { ...item, quantity: newQuantity }
        : item
    ));
  };

  const getTotalItems = () => {
    return cartItems.reduce((total, item) => total + item.quantity, 0);
  };

  const getTotalPrice = () => {
    return cartItems.reduce((total, item) => total + (item.price * item.quantity), 0).toFixed(2);
  };

  return (
    <Router>
      <div className="App">
        <Header 
          cartItemsCount={getTotalItems()} 
          onCartClick={() => setShowCart(true)} 
        />
        
        <Routes>
          <Route 
            path="/" 
            element={<Home addToCart={addToCart} />} 
          />
          <Route 
            path="/product/:id" 
            element={<ProductDetail addToCart={addToCart} />} 
          />
        </Routes>

        {showCart && (
          <Cart
            items={cartItems}
            onClose={() => setShowCart(false)}
            onRemove={removeFromCart}
            onUpdateQuantity={updateQuantity}
            totalPrice={getTotalPrice()}
          />
        )}
      </div>
    </Router>
  );
}

export default App;