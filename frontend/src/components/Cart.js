import React from 'react';

const Cart = ({ items, onClose, onRemove, onUpdateQuantity, totalPrice }) => {
  const handleCheckout = () => {
    alert(`Thank you for your purchase! Total: $${totalPrice}`);
    // In a real app, this would integrate with a payment processor
  };

  return (
    <div className="cart-overlay" onClick={onClose}>
      <div className="cart-modal" onClick={(e) => e.stopPropagation()}>
        <div className="cart-header">
          <h2>Shopping Cart</h2>
          <button onClick={onClose} className="close-cart">
            <i className="fas fa-times"></i>
          </button>
        </div>

        {items.length === 0 ? (
          <div style={{ textAlign: 'center', padding: '2rem 0' }}>
            <p>Your cart is empty</p>
            <button onClick={onClose} className="btn btn-primary">
              Continue Shopping
            </button>
          </div>
        ) : (
          <>
            <div className="cart-items">
              {items.map((item, index) => (
                <div key={`${item.id}-${item.selectedSize}-${item.selectedColor}-${index}`} className="cart-item">
                  <img src={item.image} alt={item.name} />
                  
                  <div className="cart-item-info">
                    <div className="cart-item-name">{item.name}</div>
                    <div className="cart-item-price">${item.price}</div>
                    <div style={{ fontSize: '0.9rem', color: '#666' }}>
                      Size: {item.selectedSize} | Color: {item.selectedColor}
                    </div>
                    
                    <div className="quantity-controls">
                      <button 
                        className="quantity-btn"
                        onClick={() => onUpdateQuantity(item.id, item.selectedSize, item.selectedColor, item.quantity - 1)}
                      >
                        -
                      </button>
                      <span style={{ padding: '0 0.5rem' }}>{item.quantity}</span>
                      <button 
                        className="quantity-btn"
                        onClick={() => onUpdateQuantity(item.id, item.selectedSize, item.selectedColor, item.quantity + 1)}
                      >
                        +
                      </button>
                      <button 
                        onClick={() => onRemove(item.id, item.selectedSize, item.selectedColor)}
                        style={{ 
                          background: '#e74c3c', 
                          color: 'white', 
                          border: 'none', 
                          padding: '0.3rem 0.5rem', 
                          borderRadius: '3px',
                          marginLeft: '1rem',
                          cursor: 'pointer'
                        }}
                      >
                        Remove
                      </button>
                    </div>
                  </div>
                </div>
              ))}
            </div>

            <div className="cart-total">
              Total: ${totalPrice}
            </div>

            <button onClick={handleCheckout} className="checkout-btn">
              Proceed to Checkout
            </button>
          </>
        )}
      </div>
    </div>
  );
};

export default Cart;