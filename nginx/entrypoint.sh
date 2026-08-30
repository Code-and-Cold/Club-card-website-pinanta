#!/bin/sh
set -e

export BACKEND_INTERNAL_PORT="${BACKEND_INTERNAL_PORT:-8080}"
export FRONTEND_INTERNAL_PORT="${FRONTEND_INTERNAL_PORT:-5173}"
export DOMAIN="${DOMAIN:-localhost}"
export NGINX_ENV="${NGINX_ENV:-dev}"
export NGINX_SSL="${NGINX_SSL:-off}"

if [ "$NGINX_ENV" = "prod" ] && [ "$NGINX_SSL" = "on" ]; then
  TEMPLATE="/etc/nginx/sites/prod.conf.template"
elif [ "$NGINX_ENV" = "prod" ]; then
  TEMPLATE="/etc/nginx/sites/prod.http.conf.template"
else
  TEMPLATE="/etc/nginx/sites/dev.conf.template"
fi
OUTPUT="/etc/nginx/conf.d/default.conf"

if [ ! -f "$TEMPLATE" ]; then
  echo "Nginx template not found: $TEMPLATE" >&2
  exit 1
fi

envsubst '${BACKEND_INTERNAL_PORT} ${FRONTEND_INTERNAL_PORT} ${DOMAIN}' \
  < "$TEMPLATE" > "$OUTPUT"

exec nginx -g 'daemon off;'
