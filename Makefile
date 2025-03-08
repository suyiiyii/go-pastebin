.PHONY: all clean frontend backend

# Default target
.PHONY: all
all: frontend backend

# Build frontend
.PHONY: frontend
frontend:
	@echo "Building frontend..."
	cd frontend && pnpm i
	cd frontend && pnpm build
	@echo "Copying frontend build artifacts to biz/router/dist/..."
	mkdir -p biz/router/dist
	cp -r frontend/dist/* biz/router/dist/

# Build backend
.PHONY: backend
backend:
	@echo "Building backend..."
	go build .

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning..."
	rm -rf biz/router/dist
	rm -rf frontend/dist
