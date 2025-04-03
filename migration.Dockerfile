FROM alpine:3.21

RUN apk update && \
    apk upgrade && \
    apk add bash && \
    rm -rf /var/cache/apk/*

ADD https://github.com/pressly/goose/releases/download/v3.24.2/goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose

WORKDIR /root

ADD ./internal/infrastructure/db/migrations/*.sql migrations/
# ADD migration.sh .
# ADD .env .

# RUN chmod +x migration_prod.sh

ENTRYPOINT ["bash", "migration.sh"]