-- name: CreateOrganization :one
INSERT INTO public.organizations (name, created_at, updated_at)
VALUES ($1, now(), now())
RETURNING *;

-- name: GetOrganizationById :one
SELECT * FROM public.organizations WHERE slug = sqlc.arg('id');

-- name: GetOrganizationBySlug :one
SELECT id, public_id, name, slug, stripe_customer_id, billing_email, created_at, updated_at
FROM public.organizations
WHERE slug = $1;

-- name: GetOrganizationsForAccountId :many
SELECT o.*, om.role as member_role
FROM organization_membership om
         JOIN organizations o on o.id = om.organization_id
WHERE account_id = $1;

-- name: GetOrganizationIdsForAccountId :many
SELECT o.id
FROM organization_membership om
         JOIN organizations o on o.id = om.organization_id
WHERE account_id = $1;