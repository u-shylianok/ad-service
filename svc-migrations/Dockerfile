FROM alpine:latest AS svc-migrations

RUN apk add --no-cache \
        libc6-compat

RUN mkdir -p /app
WORKDIR /app

ADD build/svc-migrations /app
ADD migrations /app/migrations
EXPOSE 8001
CMD ["./svc-migrations"]
