FROM alpine:latest AS svc-auth

RUN apk add --no-cache \
        libc6-compat

RUN mkdir -p /app
WORKDIR /app

ADD svc-auth /app
EXPOSE 8001
CMD ["./svc-auth"]