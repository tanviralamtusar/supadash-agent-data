# SupaDash

> Self-host the Supabase platform. Create, manage, and monitor multiple Supabase projects on your own infrastructure.

[![Build](https://github.com/tanviralamtusar/SupaDash/actions/workflows/test.yml/badge.svg)](https://github.com/tanviralamtusar/SupaDash/actions)
[![Go](https://img.shields.io/badge/Go-1.25-00ADD8?logo=go)](https://go.dev)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

---

## Features

- **Project Provisioning** — Create, pause, resume, and delete Supabase projects with a single API call
- **Resource Management** — CPU/memory limits, burst pool, and dynamic scaling per project
- **Security** — JWT auth with refresh tokens, RBAC, rate limiting, input validation, CORS, audit logging
- **Team Management** — Organization-based access with role-based permissions and email invitations
- **Monitoring** — Real-time resource metrics, anomaly detection, and optimization recommendations
- **Production Ready** — Docker Compose deployment, Prometheus metrics, TLS via Caddy

## Architecture

```
┌────────────────────────────────────────────────┐
│              Caddy (HTTPS :443)                 │
└────────────────────┬───────────────────────────┘
                     │
          ┌──────────▼──────────┐
          │   SupaDash API       │
          │   (Go + Gin :8080)   │
          └──┬─────────────┬────┘
             │             │
   ┌─────────▼────┐  ┌────▼────────────────┐
   │ Management   │  │ Docker Engine        │
   │ PostgreSQL   │  │ ┌─── Project A ───┐  │
   │ (port 5432)  │  │ │ postgres, kong, │  │
   └──────────────┘  │ │ gotrue, studio  │  │
                     │ └─────────────────┘  │
                     └──────────────────────┘
```

## Quick Start

### Prerequisites

- Docker & Docker Compose
- Git

### 1. Clone & configure

```bash
git clone https://github.com/tanviralamtusar/SupaDash.git
cd SupaDash
cp .env.example .env
# Edit .env — set JWT_SECRET, POSTGRES_PASSWORD, etc.
```

### 2. Run

```bash
docker compose -f docker-compose.prod.yml up -d
```

### 3. Verify

```bash
curl http://localhost:8080/v1/health
# → {"is_healthy": true}
```

## Configuration

All configuration is via environment variables. See [`.env.example`](.env.example) for the full list.

| Variable | Description | Default |
|----------|-------------|---------|
| `DATABASE_URL` | PostgreSQL connection string | `postgres://postgres:postgres@localhost:5432/supadash` |
| `JWT_SECRET` | JWT signing secret (64+ chars) | — |
| `ALLOWED_ORIGINS` | CORS allowed origins (comma-separated) | `*` |
| `RATE_LIMIT_REQUESTS` | Max requests per second per IP | `100` |
| `PROVISIONING_ENABLED` | Enable Docker provisioning | `true` |

## Documentation

- [Deployment Guide](docs/deployment-guide.md) — Full production setup
- [API Reference](docs/api-reference.md) — All endpoints
- [TLS Guide](docs/tls_guide.md) — HTTPS setup with Caddy or Traefik

## Development

```bash
# Run locally
go run main.go

# Run tests
go test ./... -v

# Build
go build -o supadash .
```

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## Security

See [SECURITY.md](SECURITY.md) for reporting vulnerabilities.

## License

[MIT](LICENSE)
