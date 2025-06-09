# EnvGuard Makefile

.PHONY: build test clean install fmt vet run-example demo-env clean-demo build-npm package-npm help

# Build the binary
build:
	go build -o envguard -ldflags="-s -w" .

# Build for multiple platforms
build-all:
	GOOS=linux GOARCH=amd64 go build -o dist/envguard-linux-amd64 -ldflags="-s -w" .
	GOOS=darwin GOARCH=amd64 go build -o dist/envguard-darwin-amd64 -ldflags="-s -w" .
	GOOS=darwin GOARCH=arm64 go build -o dist/envguard-darwin-arm64 -ldflags="-s -w" .
	GOOS=windows GOARCH=amd64 go build -o dist/envguard-windows-amd64.exe -ldflags="-s -w" .

# Build for npm distribution
build-npm:
	mkdir -p envguard-npm/bin
	GOOS=linux GOARCH=amd64 go build -o envguard-npm/bin/envguard-linux -ldflags="-s -w" .
	GOOS=darwin GOARCH=amd64 go build -o envguard-npm/bin/envguard-darwin -ldflags="-s -w" .
	GOOS=windows GOARCH=amd64 go build -o envguard-npm/bin/envguard-win.exe -ldflags="-s -w" .
	chmod +x envguard-npm/bin/envguard-linux envguard-npm/bin/envguard-darwin

# Run tests
test:
	go test -v ./...

# Format code
fmt:
	go fmt ./...

# Vet code
vet:
	go vet ./...

# Clean build artifacts
clean:
	rm -f envguard
	rm -rf dist/
	rm -f *.test *.prof *.out
	rm -rf coverage/
	rm -f envguard-npm/*.tgz

# Install locally
install: build
	sudo cp envguard /usr/local/bin/

# Download dependencies
deps:
	go mod download
	go mod tidy

# Run example scenarios
run-example:
	@echo "🔍 Testing with missing and extra variables:"
	./envguard --env example.env --example example.env.example || true
	@echo "\n🔍 Testing with perfect configuration:"
	./envguard --env perfect.env --example example.env.example || true

# Demo environment management
demo-env:
	@echo "🌍 Environment Management Demo:"
	@echo "Creating demo environments..."
	./envguard create -e development --from-current || true
	./envguard create -e staging --from-current || true
	./envguard create -e production --from-current || true
	@echo "\n📋 Listing environments:"
	./envguard list
	@echo "\n🔄 Using staging environment:"
	./envguard use staging || true
	@echo "\n📊 Checking status:"
	./envguard status || true
	@echo "\n🔄 Using development environment:"
	./envguard use development || true
	@echo "\n📊 Final status:"
	./envguard status || true

# Clean up demo environments
clean-demo:
	./envguard delete -e development --no-confirm || true
	./envguard delete -e staging --no-confirm || true
	./envguard delete -e production --no-confirm || true
	rm -rf .envguard/ || true

# Package for npm distribution
package-npm: build-npm
	@echo "📦 Packaging EnvGuard for npm..."
	cd envguard-npm && npm pack
	@echo "✅ Package created! Ready for publishing."
	@echo ""
	@echo "To publish:"
	@echo "  cd envguard-npm"
	@echo "  npm publish"
	@echo ""
	@echo "To test locally:"
	@echo "  cd envguard-npm"
	@echo "  npm install -g ."

# Development setup
dev: deps build

# Show help
help:
	@echo "EnvGuard Development Commands:"
	@echo "  build        - Build the binary"
	@echo "  build-all    - Build for multiple platforms"
	@echo "  test         - Run tests"
	@echo "  fmt          - Format code"
	@echo "  vet          - Vet code"
	@echo "  clean        - Clean build artifacts"
	@echo "  install      - Install binary to /usr/local/bin"
	@echo "  deps         - Download and tidy dependencies"
	@echo "  run-example  - Run example scenarios"
	@echo "  demo-env     - Demo environment management features"
	@echo "  clean-demo   - Clean up demo environments"
	@echo "  build-npm    - Build binaries for npm distribution"
	@echo "  package-npm  - Package for npm distribution"
	@echo "  dev          - Development setup (deps + build)"
	@echo "  help         - Show this help"
	@echo ""
	@echo "📁 Repository includes .gitignore files for:"
	@echo "  • Go build artifacts (binaries, test files)"
	@echo "  • Environment files (.env, .envguard/)"
	@echo "  • npm package artifacts (*.tgz, node_modules/)"
	@echo "  • IDE and OS files (.vscode/, .DS_Store)"

# Default target
all: fmt vet build test 