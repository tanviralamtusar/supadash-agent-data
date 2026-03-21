# Security Policy

## Reporting a Vulnerability

If you discover a security vulnerability in SupaDash, **please do not open a public GitHub issue.**

Instead, report it privately by emailing **security@supadash.dev** (or open a [private security advisory](https://github.com/tanviralamtusar/SupaDash/security/advisories/new)).

Include:
- Description of the vulnerability
- Steps to reproduce
- Potential impact
- Suggested fix (if any)

We will acknowledge your report within **48 hours** and aim to release a fix within **7 days** for critical issues.

## Supported Versions

| Version | Supported |
|---------|-----------|
| latest  | ✅        |
| < 1.0   | ⚠️ Best effort |

## Security Best Practices

When deploying SupaDash in production:

1. **Use strong secrets** — Set `JWT_SECRET` to a random 64+ character string
2. **Enable HTTPS** — Use Caddy or a reverse proxy with TLS
3. **Restrict CORS** — Set `ALLOWED_ORIGINS` to your domain(s)
4. **Protect Docker socket** — Only trusted processes should access `/var/run/docker.sock`
5. **Rotate secrets regularly** — Use the `/projects/:ref/secrets/rotate` endpoint
6. **Enable rate limiting** — Set `RATE_LIMIT_REQUESTS` appropriately
7. **Regular backups** — Use `scripts/backup.sh` on a cron schedule
