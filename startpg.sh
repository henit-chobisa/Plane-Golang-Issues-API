#!/bin/bash

# Start the PostgreSQL container
docker run -d \
  --name postgres \
  -e POSTGRES_USER=root \
  -e POSTGRES_PASSWORD=plane \
  -p 5432:5432 \
  postgres

# Wait for the PostgreSQL container to be ready
until docker exec -it postgres pg_isready; do
  echo "Waiting for PostgreSQL container to be ready..."
  sleep 1
done

# Execute the createdb command
docker exec -it postgres /bin/sh -c "createdb --username=root --owner=root plane"