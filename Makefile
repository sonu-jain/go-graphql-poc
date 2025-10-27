# GraphQL POC Makefile

.PHONY: help server client test clean

# Default target
help:
	@echo "Available targets:"
	@echo "  server    - Start the GraphQL server"
	@echo "  client    - Run the GraphQL client examples"
	@echo "  test      - Run tests"
	@echo "  clean     - Clean up generated files"
	@echo ""
	@echo "Client examples:"
	@echo "  make client              - Create a customer"
	@echo "  make client-create       - Create a customer"
	@echo "  make client-login        - Test login"
	@echo "  make client-demo-queries - Demo all queries"
	@echo "  make client-demo-mutations - Demo all mutations"
	@echo "  make client-demo-workflow - Demo complete workflow"
	@echo "  make client-demo-all     - Run all demos"

# Start the GraphQL server
server:
	@echo "🚀 Starting GraphQL server..."
	go run server.go

# Run client examples
client:
	@echo "🚀 Running GraphQL client examples..."
	cd cmd/client && go run main.go -action create

client-create:
	@echo "👥 Creating customer..."
	cd cmd/client && go run main.go -action create

client-login:
	@echo "🔐 Testing login..."
	cd cmd/client && go run main.go -action login

client-demo-queries:
	@echo "🔍 Running queries demo..."
	cd cmd/client && go run main.go -action demo-queries

client-demo-mutations:
	@echo "✏️ Running mutations demo..."
	cd cmd/client && go run main.go -action demo-mutations

client-demo-workflow:
	@echo "🔄 Running workflow demo..."
	cd cmd/client && go run main.go -action demo-workflow

client-demo-all:
	@echo "🎯 Running all demos..."
	cd cmd/client && go run main.go -action demo-all

# Run tests
test:
	@echo "🧪 Running tests..."
	go test ./...

# Clean up generated files
clean:
	@echo "🧹 Cleaning up..."
	go clean
	rm -f graph/generated.go
	rm -f graph/model/models_gen.go

# Install dependencies
deps:
	@echo "📦 Installing dependencies..."
	go mod tidy
	go mod download

# Generate GraphQL code
generate:
	@echo "🔧 Generating GraphQL code..."
	go run github.com/99designs/gqlgen generate

# Build client
build-client:
	@echo "🔨 Building client..."
	cd cmd/client && go build -o ../../bin/graphql-client main.go

# Build server
build-server:
	@echo "🔨 Building server..."
	go build -o bin/graphql-server server.go

# Build all
build: build-server build-client
	@echo "✅ Build completed!"

# Run with custom URL
client-url:
	@echo "🌐 Running client with custom URL..."
	cd cmd/client && go run main.go -url $(URL)
