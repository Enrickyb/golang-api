.PHONY: default run build test docs clean

# Variables

APP_NAME = "golang-api"

# tasks

default: run

run: 
	@echo "Running $(APP_NAME)..."
	@go run main.go

build: 
	@echo "Building $(APP_NAME)..."
	@go build -o $(APP_NAME) main.go

test: 
	@echo "Testing $(APP_NAME)..."
	@go test ./...

docs:
	@echo "Generating docs for $(APP_NAME)..."
	@swag init

clean:
	@echo "Cleaning $(APP_NAME)..."
	@rm -rf $(APP_NAME) docs