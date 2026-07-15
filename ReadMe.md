# api-gateway

A lightweight API gateway written in Go using the [Gin](https://github.com/gin-gonic/gin) web framework. It receives incoming HTTP requests, logs them, and reverse-proxies them to a configured backend service.

## Features

- **Reverse proxy** — forwards any request under `/api/*` to a target service (`net/http/httputil`)
- **Request logging middleware** — logs method, path, status code, and response time for every request
- **Health check endpoint** — `GET /health` for uptime/liveness checks
- **Environment-based config** — configurable via `.env` file or system environment variables
- **Dockerized** — runs in a container via the included `Dockerfile`

## Project Structure

```
api-gateway/
├── cmd/
│   └── main.go                  # entry point — sets up router, middleware, routes
├── internal/
│   ├── config/
│   │   └── config.go            # loads .env / system env vars
│   ├── handlers/
│   │   └── health.go            # GET /health handler
│   ├── middleware/
│   │   └── logger.go            # request logging middleware
│   └── proxy/
│       └── proxy.go             # reverse proxy handler
├── docs/                        # dev notes
├── Dockerfile
├── go.mod
└── .env.example
```

## Prerequisites

- [Go](https://go.dev/dl/) 1.22+
- [Docker](https://docs.docker.com/get-docker/) (optional, for containerized runs)

## Setup

1. Clone the repo:
   ```bash
   git clone https://github.com/NeoOnyedire/api-gateway.git
   cd api-gateway
   ```

2. Copy the example env file and fill in your own values:
   ```bash
   cp .env.example .env
   ```

3. Install dependencies:
   ```bash
   go mod download
   ```

## Environment Variables

| Variable         | Description                              | Default                   |
|-------------------|-------------------------------------------|----------------------------|
| `PORT`            | Port the gateway listens on               | `8080`                     |
| `TARGET_SERVICE`  | Backend service requests are forwarded to | `http://localhost:9000`    |

If no `.env` file is found, the gateway falls back to system environment variables (or the defaults above).

## Running Locally

### Option A: Run directly with Go

```bash
go run ./cmd/main.go
```

### Option B: Run with Docker

```bash
docker build -t api-gateway .
docker run --env-file .env -p 8080:8080 api-gateway
```

## Usage

Once running, the gateway exposes:

- `GET /health` — returns `{"status": "ok"}`
- `ANY /api/*path` — proxies the request to `TARGET_SERVICE`, preserving the method and the rest of the path

Example:
```bash
curl http://localhost:8080/health
curl http://localhost:8080/api/users
# forwarded to TARGET_SERVICE/api/users
```

## Known Issues

- The proxy currently panics on an invalid `TARGET_SERVICE` URL (no error handling around `url.Parse`) — fix in progress.

## Roadmap

- [ ] Fix proxy panic on invalid target URL
- [ ] Add error-handling middleware
- [ ] JWT authentication
- [ ] Clean up duplicate log entries
- [ ] Configurable / dynamic routing rules (currently a single static target)
- [ ] Rate limiting

## License

Add a license of your choice (MIT is a common default for open-source Go projects).
