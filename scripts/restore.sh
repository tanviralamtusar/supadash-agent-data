#!/usr/bin/env bash
# SupaDash Management Database Restore Script
# Usage: ./scripts/restore.sh <backup_file.sql.gz>
set -euo pipefail

if [ $# -lt 1 ]; then
    echo "Usage: $0 <backup_file.sql.gz>"
    echo ""
    echo "Available backups:"
    ls -lh ./backups/supadash_backup_*.sql.gz 2>/dev/null || echo "  (none found in ./backups/)"
    exit 1
fi

BACKUP_FILE="$1"

if [ ! -f "$BACKUP_FILE" ]; then
    echo "Error: Backup file not found: ${BACKUP_FILE}"
    exit 1
fi

# Database connection
DB_HOST="${POSTGRES_HOST:-localhost}"
DB_PORT="${POSTGRES_PORT:-5432}"
DB_USER="${POSTGRES_USER:-postgres}"
DB_NAME="${POSTGRES_DB:-supadash}"

echo "==> Restoring SupaDash database from: ${BACKUP_FILE}"
echo "    Host: ${DB_HOST}:${DB_PORT}"
echo "    Database: ${DB_NAME}"
echo ""
echo "WARNING: This will overwrite the current database contents."
read -p "Continue? (y/N) " -n 1 -r
echo ""

if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo "Aborted."
    exit 0
fi

gunzip -c "$BACKUP_FILE" | psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME"

echo "==> Restore complete."
