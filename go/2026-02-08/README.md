# OIDC Tutorial

A simple HTTP server written in Go using the Gin framework.

## Project Structure

```
<your-project-name>/
├── cmd/
│   └── server/
│       └── main.go       # Application entry point
├── internal/
│   └── http/
│       └── server.go     # HTTP server and route handlers
├── go.mod
└── go.sum
```

## Requirements

- Go 1.25.5 or later

## Setup from Scratch

### 1. Create project directory

```bash
mkdir <your-project-name>
cd <your-project-name>
```

### 2. Initialize Go module

```bash
go mod init <your-project-name>
```

### 3. Install Gin framework

```bash
go get -u github.com/gin-gonic/gin
```

### 4. Create directory structure

```bash
mkdir -p cmd/server
mkdir -p internal/http
```

### 5. Create source files

Create `cmd/server/main.go` and `internal/http/server.go` as shown in the project structure.

## Build and Run

### Build

```bash
go build -o server ./cmd/server
```

### Run

```bash
./server
```

The server will start on port 8080.

## API Endpoints

| Method | Path         | Description               |
|--------|--------------|---------------------------|
| GET    | /hello       | Returns a greeting        |
| GET    | /user/:name  | Returns personalized greeting |

## Examples

```bash
# Get hello message
curl http://localhost:8080/hello

# Get personalized greeting
curl http://localhost:8080/user/John
```
