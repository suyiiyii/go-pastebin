# syntax=docker/dockerfile:1

##############################################
# Frontend Build Stage
##############################################
FROM node:20-slim AS frontend-build

WORKDIR /app

# Install pnpm
RUN npm install -g pnpm

# Copy frontend files
COPY frontend/package.json frontend/pnpm-lock.yaml* ./frontend/
COPY frontend/ ./frontend/

# Build frontend
RUN cd frontend && pnpm install
RUN cd frontend && pnpm build

##############################################
# Backend Build Stage
##############################################
FROM golang:1.23.4 AS backend-build

WORKDIR /src

# Download dependencies as a separate step for better caching
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod/ \
    go mod download -x

# Copy the rest of the source code
COPY . .

# Copy frontend build artifacts from the frontend-build stage
COPY --from=frontend-build /app/frontend/dist/ ./biz/router/dist/

# Build the application
RUN --mount=type=cache,target=/go/pkg/mod/ \
    CGO_ENABLED=0 go build -o /bin/server .

##############################################
# Final Stage
##############################################
FROM debian:12-slim AS final

# Install CA certificates for HTTPS
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*


# Copy the executable from the build stage
COPY --from=backend-build /bin/server /bin/

# Use the non-privileged user
USER appuser

# Expose the port that the application listens on
EXPOSE 8080

# What the container should run when it is started
ENTRYPOINT ["/bin/server"]