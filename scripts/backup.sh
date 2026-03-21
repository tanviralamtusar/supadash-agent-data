#!/usr/bin/env bash
# SupaDash Management Database Backup Script
# Usage: ./scripts/backup.sh [output_dir]
set -euo pipefail

BACKUP_DIR="${1:-./backups}"
TIMESTAMP=$(date +%Y%m%d_%H%M%S)
BACKUP_FILE="${BACKUP_DIR}/supadash_backup_${TIMESTAMP}.sql.gz"

# Database connection (from environment or defaults)
DB_HOST="${POSTGRES_HOST:-localhost}"
DB_PORT="${POSTGRES_PORT:-5432}"
DB_USER="${POSTGRES_USER:-postgres}"
DB_NAME="${POSTGRES_DB:-supadash}"

mkdir -p "$BACKUP_DIR"

echo "==> Backing up SupaDash database..."
echo "    Host: ${DB_HOST}:${DB_PORT}"
echo "    Database: ${DB_NAME}"
echo "    Output: ${BACKUP_FILE}"

pg_dump -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME" \
    --no-owner --no-privileges --clean --if-exists | gzip > "$BACKUP_FILE"

FILESIZE=$(du -h "$BACKUP_FILE" | cut -f1)
echo "==> Backup complete: ${BACKUP_FILE} (${FILESIZE})"

# Retention: delete backups older than 30 days
echo "==> Cleaning backups older than 30 days..."
find "$BACKUP_DIR" -name "supadash_backup_*.sql.gz" -mtime +30 -delete 2>/dev/null || true

echo "==> Done."
