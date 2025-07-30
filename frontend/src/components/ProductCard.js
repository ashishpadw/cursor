import React from 'react';
import { Link } from 'react-router-dom';

const ProductCard = ({ product, onQuickAdd }) => {
  return (
    <div className="product-card">
      <Link to={`/product/${product.id}`}>
        <img 
          src={product.image} 
          alt={product.name}
          className="product-image"
        />
        <div className="product-info">
          <h3 className="product-name">{product.name}</h3>
          <p className="product-description">
            {product.description.length > 100 
              ? `${product.description.substring(0, 100)}...` 
              : product.description
            }
          </p>
          <div className="product-price">${product.price}</div>
        </div>
      </Link>
      
      <div className="product-info">
        <div className="product-actions">
          <Link to={`/product/${product.id}`} className="btn btn-primary">
            View Details
          </Link>
          <button 
            onClick={(e) => {
              e.preventDefault();
              onQuickAdd(product);
            }}
            className="btn btn-secondary"
          >
            Quick Add
          </button>
        </div>
      </div>
    </div>
  );
};

export default ProductCard;