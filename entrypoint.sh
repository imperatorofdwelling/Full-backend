#!/bin/bash

# Путь к файлу конфигурации PostgreSQL
POSTGRES_CONF="/var/lib/postgresql/data/postgresql.conf"

# Новый порт
NEW_PORT="5433"

# Изменение конфигурационного файла
if grep -q "^port" "$POSTGRES_CONF"; then
  sed -i "s/^port.*/port = $NEW_PORT/" "$POSTGRES_CONF"
else
  echo "port = $NEW_PORT" >> "$POSTGRES_CONF"
fi

# Перезапуск PostgreSQL
pg_ctl reload