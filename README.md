# golang-server

A minimal Go HTTP server that demonstrates a clean project layout and common server concerns (configuration, logging, middleware, and user handling).

## ✅ Overview

This repository contains a small, opinionated Go server implementation intended for learning and experimentation. It follows a conventional layout (e.g., `cmd/`, `internal/`, `pkg/`) and includes examples for configuration, logging, middleware, and a simple user service.

## ⚙️ Features

- Clean project structure (entrypoint, internal packages, and reusable code)
- Configuration via environment variables (`APP_PORT`)
- Structured logging (see `internal/logger`)
- Authentication middleware (see `internal/middleware`)
- Simple user handler, service, and repository implementations

## Requirements

- Go 1.25.5 or later
- `go` toolchain and modules enabled

## Installation & Usage

Clone the repo and build or run the server:

```bash
git clone https://github.com/nimpeboss/golang-server.git
cd golang-server
# Build
go build -o bin/server ./cmd/server
# Run (default port 8080)
APP_PORT=8080 ./bin/server

# Or run with 'go run' for development
go run ./cmd/server
```

## Configuration

The server reads the `APP_PORT` environment variable. If unset, the server defaults to port `8080`.

## Development

- Run tests: `go test ./...`
- Static analysis: `go vet ./...` (add linters as desired)

## Maintenance / Status

This project is actively developed and will continue to receive occasional updates and improvements. Expect periodic, often small, updates — frequently when the maintainer has spare time (or is feeling bored).

## Contributing

Contributions are welcome. Please open an issue to discuss larger changes or submit a pull request with clear tests and documentation.

## License

No license is specified in this repository. Add a `LICENSE` file to indicate how the project may be used and distributed.
