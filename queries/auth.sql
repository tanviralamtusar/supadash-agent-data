-- name: InsertRefreshToken :one
INSERT INTO public.refresh_tokens (account_id, token, expires_at)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetRefreshToken :one
SELECT * FROM public.refresh_tokens
WHERE token = $1 LIMIT 1;

-- name: RevokeRefreshToken :exec
UPDATE public.refresh_tokens
SET revoked = true, updated_at = now()
WHERE token = $1;

-- name: RevokeAllRefreshTokensForUser :exec
UPDATE public.refresh_tokens
SET revoked = true, updated_at = now()
WHERE account_id = $1 AND revoked = false;

-- name: InsertAuditLog :one
INSERT INTO public.audit_logs (target_project, actor_id, action, ip_address, user_agent, details)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetProjectAuditLogs :many
SELECT * FROM public.audit_logs
WHERE target_project = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: GetOrganizationAuditLogs :many
-- assuming audit logs without target_project or with org projects context.
-- For simplicity, let's keep it project-level or actor-level.
SELECT * FROM public.audit_logs
WHERE actor_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;
