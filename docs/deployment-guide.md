# Deployment Guide

## Prerequisites

- Linux server (Ubuntu 22.04+ recommended)
- Docker Engine 24+ & Docker Compose v2
- 4+ CPU cores, 8+ GB RAM (for 1-5 projects)
- Domain name (optional, for HTTPS)

## Step 1: Clone & Configure

```bash
git clone https://github.com/tanviralamtusar/SupaDash.git
cd SupaDash
cp .env.example .env
```

Edit `.env` with secure values:

```bash
# Generate a strong JWT secret
openssl rand -hex 32  # → paste into JWT_SECRET

# Set database credentials
POSTGRES_USER=supadash
POSTGRES_PASSWORD=$(openssl rand -hex 16)
POSTGRES_DB=supadash

# Set your domain for CORS
ALLOWED_ORIGINS=https://your-domain.com
```

## Step 2: Configure HTTPS (Optional)

Edit `Caddyfile` for automatic TLS:

```
your-domain.com {
    reverse_proxy api:8080
}
```

## Step 3: Deploy

```bash
docker compose -f docker-compose.prod.yml up -d
```

Verify all services are running:

```bash
docker compose -f docker-compose.prod.yml ps
curl http://localhost:8080/v1/health
```

## Step 4: Set Up Backups

```bash
# Make scripts executable
chmod +x scripts/backup.sh scripts/restore.sh

# Add to crontab (daily at 2 AM)
echo "0 2 * * * cd /path/to/SupaDash && ./scripts/backup.sh" | crontab -
```

## Step 5: Monitoring (Optional)

Start Prometheus alongside the stack:

```bash
docker run -d --name prometheus \
  --network supadash_default \
  -p 9090:9090 \
  -v $(pwd)/monitoring/prometheus.yml:/etc/prometheus/prometheus.yml \
  prom/prometheus
```

Access Prometheus at `http://your-server:9090`.

## Upgrading

```bash
git pull origin main
docker compose -f docker-compose.prod.yml build
docker compose -f docker-compose.prod.yml up -d
```

## Troubleshooting

| Issue | Solution |
|-------|----------|
| API won't start | Check `docker logs supadash-api` |
| DB connection fails | Verify `DATABASE_URL` in `.env` |
| Port conflicts | Change ports in `docker-compose.prod.yml` |
| Provisioning fails | Ensure Docker socket is mounted |
