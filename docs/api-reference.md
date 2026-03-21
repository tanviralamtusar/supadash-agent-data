# API Reference

Base URL: `http://localhost:8080`

All authenticated endpoints require a `Bearer` token in the `Authorization` header.

---

## Health & Monitoring

| Method | Endpoint | Auth | Description |
|--------|----------|------|-------------|
| GET | `/` | No | API index |
| GET | `/status` | No | Health status |
| GET | `/v1/health` | No | Health check |
| GET | `/v1/metrics` | No | Prometheus metrics |

---

## Authentication

| Method | Endpoint | Auth | Description |
|--------|----------|------|-------------|
| POST | `/v1/auth/token?grant_type=refresh_token` | No | Exchange refresh token for new access + refresh token |
| POST | `/v1/auth/logout` | Yes | Revoke all refresh tokens for the user |

### Token Exchange

```bash
curl -X POST http://localhost:8080/v1/auth/token?grant_type=refresh_token \
  -H "Content-Type: application/json" \
  -d '{"refresh_token": "your-refresh-token"}'
```

---

## Profile

| Method | Endpoint | Auth | Description |
|--------|----------|------|-------------|
| GET | `/profile` | Yes | Get current user profile |
| GET | `/profile/permissions` | Yes | Get user permissions |
| POST | `/profile/password-check` | Yes | Validate password strength |

---

## Organizations

| Method | Endpoint | Auth | Description |
|--------|----------|------|-------------|
| GET | `/organizations` | Yes | List user's organizations |
| GET | `/organizations/:slug/members/reached-free-project-limit` | Yes | Check free project limit |

### Team Management

| Method | Endpoint | Auth | Roles | Description |
|--------|----------|------|-------|-------------|
| GET | `/organizations/:slug/team` | Yes | Any | List team members |
| POST | `/organizations/:slug/team/invite` | Yes | Owner, Admin | Invite member via email |
| PUT | `/organizations/:slug/team/:id` | Yes | Owner, Admin | Update member role |
| DELETE | `/organizations/:slug/team/:id` | Yes | Owner, Admin | Remove team member |

---

## Projects

| Method | Endpoint | Auth | Description |
|--------|----------|------|-------------|
| GET | `/platform/projects` | Yes | List all projects |
| POST | `/platform/projects` | Yes | Create new project |
| GET | `/platform/projects/:ref` | Yes | Get project details |
| DELETE | `/platform/projects/:ref` | Yes | Delete project |
| POST | `/platform/projects/:ref/pause` | Yes | Pause project |
| POST | `/platform/projects/:ref/resume` | Yes | Resume project |

### Project Status & Config

| Method | Endpoint | Auth | Description |
|--------|----------|------|-------------|
| GET | `/platform/projects/:ref/status` | Yes | Project status |
| GET | `/platform/projects/:ref/settings` | Yes | Project settings |
| GET | `/platform/projects/:ref/api` | Yes | API configuration |

### Project Resources

| Method | Endpoint | Auth | Description |
|--------|----------|------|-------------|
| GET | `/platform/projects/:ref/resources` | Yes | Get resource limits |
| PUT | `/platform/projects/:ref/resources` | Yes | Update resource limits |

### Environment Variables

| Method | Endpoint | Auth | Description |
|--------|----------|------|-------------|
| GET | `/platform/projects/:ref/env` | Yes | List env vars |
| PUT | `/platform/projects/:ref/env` | Yes | Update env vars |

---

## Database Management

| Method | Endpoint | Auth | Description |
|--------|----------|------|-------------|
| GET | `/platform/pg-meta/:ref/tables` | Yes | List tables |
| POST | `/platform/pg-meta/:ref/tables` | Yes | Create table |
| PATCH | `/platform/pg-meta/:ref/tables` | Yes | Update table |
| DELETE | `/platform/pg-meta/:ref/tables` | Yes | Delete table |
| POST | `/platform/pg-meta/:ref/columns` | Yes | Create column |
| POST | `/platform/pg-meta/:ref/query` | Yes | Execute SQL query |
| GET | `/platform/pg-meta/:ref/types` | Yes | List types |

---

## Secrets & Security

| Method | Endpoint | Auth | Description |
|--------|----------|------|-------------|
| POST | `/v1/projects/:ref/secrets/rotate` | Yes | Rotate JWT secrets |
| GET | `/v1/projects/:ref/audit` | Yes | Get audit logs |

---

## Analytics & Monitoring

| Method | Endpoint | Auth | Description |
|--------|----------|------|-------------|
| GET | `/platform/projects/:ref/analytics` | Yes | Resource analytics |

---

## Integrations

| Method | Endpoint | Auth | Description |
|--------|----------|------|-------------|
| GET | `/platform/integrations` | Yes | List integrations |
| GET | `/platform/integrations/connections` | Yes | Integration connections |
| GET | `/platform/integrations/:id/authorization` | Yes | Authorization status |
| GET | `/platform/integrations/:id/repositories` | Yes | Connected repositories |

---

## Notifications & Billing

| Method | Endpoint | Auth | Description |
|--------|----------|------|-------------|
| GET | `/platform/notifications` | Yes | List notifications |
| GET | `/platform/notifications/summary` | Yes | Notification summary |
| GET | `/platform/organizations/:slug/billing/subscription` | Yes | Subscription status |
| GET | `/platform/organizations/:slug/billing/usage` | Yes | Usage details |
| GET | `/platform/organizations/:slug/billing/invoices/overdue` | Yes | Overdue invoices |
| GET | `/platform/projects/:ref/billing/addons` | Yes | Project add-ons |

---

## Error Responses

All errors follow this format:

```json
{
  "error": "Description of what went wrong"
}
```

| Status | Meaning |
|--------|---------|
| 400 | Bad Request — invalid input |
| 401 | Unauthorized — missing or invalid token |
| 403 | Forbidden — insufficient permissions |
| 404 | Not Found |
| 429 | Too Many Requests — rate limited |
| 500 | Internal Server Error |
