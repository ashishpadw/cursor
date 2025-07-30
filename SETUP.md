# Quick Setup Guide for StyleHub Ecommerce

## Prerequisites Installation

### 1. Install Go
```bash
# On macOS with Homebrew
brew install go

# On Ubuntu/Debian
sudo apt update
sudo apt install golang-go

# On Windows
# Download from https://golang.org/dl/
```

### 2. Install Node.js
```bash
# On macOS with Homebrew
brew install node

# On Ubuntu/Debian
curl -fsSL https://deb.nodesource.com/setup_18.x | sudo -E bash -
sudo apt-get install -y nodejs

# On Windows
# Download from https://nodejs.org/
```

## Quick Start

### Option 1: Manual Setup

1. **Terminal 1 - Backend (Go)**
   ```bash
   cd backend
   go mod tidy
   go run main.go
   ```

2. **Terminal 2 - Frontend (React)**
   ```bash
   cd frontend
   npm install
   npm start
   ```

3. **Access the Application**
   - Frontend: http://localhost:3000
   - Backend API: http://localhost:8080

### Option 2: Using Scripts (Coming Soon)

```bash
# Start both backend and frontend
./start.sh
```

## Verification

1. **Backend Test**
   ```bash
   curl http://localhost:8080/api/products
   ```

2. **Frontend Test**
   - Navigate to http://localhost:3000
   - You should see the StyleHub homepage

## Troubleshooting

### Go Issues
- Make sure Go 1.21+ is installed: `go version`
- If modules fail: `go clean -modcache && go mod tidy`

### React Issues
- Make sure Node.js 16+ is installed: `node --version`
- Clear npm cache: `npm cache clean --force`
- Delete node_modules: `rm -rf node_modules && npm install`

### CORS Issues
- Make sure the backend is running on port 8080
- Frontend should automatically proxy API calls

## Development Tips

- Backend hot reload: Use `air` for Go hot reload
- Frontend hot reload: Automatic with `npm start`
- API testing: Use curl, Postman, or browser dev tools