FROM alpine:latest AS api-gateway

RUN apk add --no-cache \
        libc6-compat

RUN mkdir -p /app
WORKDIR /app

ADD api-gateway /app
EXPOSE 8000
CMD ["./api-gateway"]