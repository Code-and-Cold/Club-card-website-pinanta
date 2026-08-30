#!/bin/bash
set -euo pipefail

if [ -f .env ]; then
  set -a
  # shellcheck disable=SC1091
  source .env
  set +a
fi

DOMAIN="${DOMAIN:-cold-code.ru}"
EMAIL="${SSL_EMAIL:?Set SSL_EMAIL in .env}"

echo "Stopping nginx..."
docker compose -f docker-compose.prod.yml stop nginx

echo "Requesting SSL certificate for ${DOMAIN} and www.${DOMAIN}..."
docker run --rm -it \
  -v "$(pwd)/nginx/ssl:/etc/letsencrypt" \
  -p 80:80 \
  -p 443:443 \
  certbot/certbot certonly --standalone \
  -d "${DOMAIN}" \
  -d "www.${DOMAIN}" \
  --email "${EMAIL}" \
  --agree-tos \
  --no-eff-email

echo "Starting production stack with HTTPS..."
if grep -q '^NGINX_SSL=' .env 2>/dev/null; then
  sed -i 's/^NGINX_SSL=.*/NGINX_SSL=on/' .env
else
  echo 'NGINX_SSL=on' >> .env
fi
docker compose -f docker-compose.prod.yml up -d --build

echo "Done. Certificates are in nginx/ssl/live/${DOMAIN}/"
