FROM alpine:latest AS svc-ads

RUN apk add --no-cache \
        libc6-compat

RUN mkdir -p /app
WORKDIR /app

ADD svc-ads /app
EXPOSE 8001
CMD ["./svc-ads"]