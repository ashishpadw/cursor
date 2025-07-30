# StyleHub - Modern Ecommerce Platform

A full-stack ecommerce application built with Go backend and React frontend, featuring men's and women's clothing collections.

## Features

- **Modern Tech Stack**: Go backend with React frontend
- **Responsive Design**: Mobile-first responsive design
- **Product Catalog**: Browse men's and women's clothing
- **Advanced Filtering**: Filter by gender and category
- **Product Details**: Detailed product pages with multiple images
- **Shopping Cart**: Full cart functionality with size/color selection
- **Real-time Updates**: Cart persists across sessions

## Tech Stack

### Backend (Go)
- Go 1.21+
- Gorilla Mux for routing
- CORS middleware
- RESTful API design

### Frontend (React)
- React 18
- React Router for navigation
- Modern CSS with Flexbox/Grid
- Font Awesome icons
- Responsive design

## Getting Started

### Prerequisites
- Go 1.21 or higher
- Node.js 16 or higher
- npm or yarn

### Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd node-typescript-mongodb
   ```

2. **Start the Go Backend**
   ```bash
   cd backend
   go mod tidy
   go run main.go
   ```
   The backend will start on `http://localhost:8080`

3. **Start the React Frontend**
   ```bash
   cd frontend
   npm install
   npm start
   ```
   The frontend will start on `http://localhost:3000`

4. **Access the Application**
   Open your browser and navigate to `http://localhost:3000`

## API Endpoints

### Products
- `GET /api/products` - Get all products (supports `?gender=` and `?category=` filters)
- `GET /api/products/{id}` - Get single product by ID
- `GET /api/categories` - Get all available categories

## Project Structure

```
├── backend/                 # Go backend
│   ├── main.go             # Main server file with API endpoints
│   └── go.mod              # Go module dependencies
├── frontend/               # React frontend
│   ├── public/             # Public assets
│   ├── src/
│   │   ├── components/     # React components
│   │   │   ├── Header.js   # Navigation header
│   │   │   ├── Home.js     # Main product listing page
│   │   │   ├── ProductCard.js  # Product card component
│   │   │   ├── ProductDetail.js # Product detail page
│   │   │   └── Cart.js     # Shopping cart modal
│   │   ├── App.js          # Main app component
│   │   ├── index.js        # App entry point
│   │   └── index.css       # Global styles
│   └── package.json        # Node.js dependencies
└── README.md               # This file
```

## Features in Detail

### Product Catalog
- 12 dummy products (6 men's, 6 women's clothing items)
- High-quality product images from Unsplash
- Detailed product information including sizes, colors, and descriptions

### Shopping Cart
- Add products with specific size and color selections
- Quantity management
- Persistent cart (survives page refresh)
- Real-time total calculation

### Responsive Design
- Mobile-first approach
- Works on desktop, tablet, and mobile devices
- Modern UI with hover effects and transitions

## Dummy Data

The application includes realistic dummy data for:
- **Men's Clothing**: T-shirts, jeans, button shirts, chinos, sweaters
- **Women's Clothing**: Dresses, jeans, blouses, cardigans, activewear, skirts, tops

Each product includes:
- Multiple high-quality images
- Detailed descriptions
- Size options
- Color variations
- Realistic pricing

## Development

### Backend Development
```bash
cd backend
go run main.go
```

### Frontend Development
```bash
cd frontend
npm start
```

The React app will automatically reload when you make changes.

## Future Enhancements

- User authentication and accounts
- Payment integration
- Product reviews and ratings
- Inventory management
- Order history
- Wishlist functionality
- Search functionality
- Admin panel for product management

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
