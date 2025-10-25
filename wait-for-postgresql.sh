#!/bin/sh
set -e

host="$1"
shift
cmd="$@"

until pg_isready -h "$host" -U "$DB_USER" >/dev/null 2>&1; do
  echo "Waiting for PostgreSQL ($host)..."
  sleep 2
done

echo "PostgreSQL is ready, starting app..."
exec $cmd