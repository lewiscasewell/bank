#!/bin/sh

set -e

# Wait for the PostgreSQL database to be ready
/app/wait-for.sh postgres:5432 -t 60
/app/migrate -path /app/db/migration -database "$DB_SOURCE" -verbose up


echo "PostgreSQL is ready. Starting the app..."
exec "$@"