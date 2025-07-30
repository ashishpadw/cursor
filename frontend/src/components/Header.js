import React from 'react';
import { Link, useLocation } from 'react-router-dom';

const Header = ({ cartItemsCount, onCartClick }) => {
  const location = useLocation();

  return (
    <header className="header">
      <div className="container">
        <div className="header-content">
          <Link to="/" className="logo">
            StyleHub
          </Link>
          
          <nav>
            <ul className="nav">
              <li>
                <Link 
                  to="/" 
                  className={location.pathname === '/' ? 'active' : ''}
                >
                  Home
                </Link>
              </li>
              <li>
                <Link to="/?gender=men">Men's Clothing</Link>
              </li>
              <li>
                <Link to="/?gender=women">Women's Clothing</Link>
              </li>
            </ul>
          </nav>

          <div className="header-actions">
            <button onClick={onCartClick} className="cart-button">
              <i className="fas fa-shopping-cart"></i>
              Cart
              {cartItemsCount > 0 && (
                <span className="cart-count">{cartItemsCount}</span>
              )}
            </button>
          </div>
        </div>
      </div>
    </header>
  );
};

export default Header;