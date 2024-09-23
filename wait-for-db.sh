#!/bin/bash

while ! pg_isready -U $POSTGRES_USER -h db; do
  echo "Waiting for database to become available..."
  sleep 1
done

echo "Database is available, starting application..."