#!/bin/bash

echo "Waiting for postgres..."
echo $DATABASE_HOST
echo $DATABASE_PORT
while ! nc -z $DATABASE_HOST $DATABASE_PORT; do
  sleep 0.1
done

echo "PostgreSQL started for fintech_trader"

exec "$@"