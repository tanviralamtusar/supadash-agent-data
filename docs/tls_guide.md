# Setting up HTTPS/TLS for SupaDash

For production deployments, SupaDash should be served over HTTPS. You can easily achieve this using a reverse proxy with automatic TLS, such as Caddy or Traefik.

## Option 1: Using Caddy (Recommended for simplicity)

1. **Install Caddy**: Follow the official guide at https://caddyserver.com/docs/install
2. **Create a Caddyfile**:
   In the same directory where you run SupaDash, create a `Caddyfile` with the following content:

```caddyfile
supadash.yourdomain.com {
    reverse_proxy localhost:8080
}
```
   *(Replace `8080` with the port you configured for the SupaDash API/Dashboard)*

3. **Run Caddy**:
   Run `caddy start` or configure Caddy as a system service. Caddy will automatically provision a Let's Encrypt certificate and renew it.

## Option 2: Using Traefik (Best for Docker Compose)

If you are running SupaDash within a Docker Compose setup, Traefik is highly recommended since it integrates natively.

1. **Update `docker-compose.yml`**:
   Add a Traefik service and connect SupaDash to the Traefik network, mapping the Traefik ports to `80` and `443`.

```yaml
services:
  traefik:
    image: traefik:v2.10
    command:
      - "--providers.docker=true"
      - "--entrypoints.web.address=:80"
      - "--entrypoints.websecure.address=:443"
      - "--certificatesresolvers.myresolver.acme.tlschallenge=true"
      - "--certificatesresolvers.myresolver.acme.email=your-email@example.com"
      - "--certificatesresolvers.myresolver.acme.storage=/letsencrypt/acme.json"
    ports:
      - "80:80"
      - "443:443"
    volumes:
      - "/var/run/docker.sock:/var/run/docker.sock:ro"
      - "./letsencrypt:/letsencrypt"

  supadash:
    # ... your supadash configuration ...
    labels:
      - "traefik.enable=true"
      - "traefik.http.routers.supadash.rule=Host(`supadash.yourdomain.com`)"
      - "traefik.http.routers.supadash.entrypoints=websecure"
      - "traefik.http.routers.supadash.tls.certresolver=myresolver"
```

2. **Run `docker-compose up -d`**. Traefik will negotiate the certificate and route requests over HTTPS to SupaDash automatically.

## CORS Considerations

When moving to production, ensure that the `ALLOWED_ORIGINS` environment variable in your `.env` file accurately specifies the production frontend domains to protect the SupaDash APIs. Example:
```env
ALLOWED_ORIGINS=https://supadash.yourdomain.com,https://dashboard.yourdomain.com
```
