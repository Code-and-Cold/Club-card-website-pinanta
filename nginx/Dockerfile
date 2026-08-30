FROM nginx:alpine

RUN apk add --no-cache gettext

COPY nginx.conf /etc/nginx/nginx.conf
COPY sites/dev.conf.template /etc/nginx/sites/dev.conf.template
COPY sites/prod.conf.template /etc/nginx/sites/prod.conf.template
COPY sites/prod.http.conf.template /etc/nginx/sites/prod.http.conf.template
COPY entrypoint.sh /entrypoint.sh

RUN chmod +x /entrypoint.sh \
    && mkdir -p /etc/nginx/conf.d

ENTRYPOINT ["/entrypoint.sh"]
