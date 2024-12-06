# Go Hello API

Simple Go API with Docker support and CI/CD pipeline.

## Local Development

### Prerequisites
- Go 1.20 or later
- Docker

### Running locally
```bash
# Run with Go
go run main.go

# Run with Docker
docker build -t go-hello-api .
docker run -p 8080:8080 go-hello-api
```

## CI/CD Pipeline

The project includes a GitHub Actions workflow that:
1. Builds and tests the application
2. Creates and pushes a Docker image to Docker Hub

### Setup

To enable CI/CD pipeline:

1. Fork this repository
2. Add the following secrets to your GitHub repository:
   - `DOCKERHUB_USERNAME`: Your Docker Hub username
   - `DOCKERHUB_TOKEN`: Your Docker Hub access token

### Pipeline Steps

1. **Build and Test**
   - Checks out code
   - Sets up Go environment
   - Builds the application
   - Runs tests

2. **Docker**
   - Builds Docker image
   - Pushes to Docker Hub

## API Endpoints

- `GET /`: Returns a JSON response with "Hello, Docker CI/CD!" message

## Testing

Run tests with:
```bash
go test -v ./...
```
