FROM alpine:latest AS ad-service-app

RUN apk add --no-cache \
        libc6-compat

RUN mkdir -p /app
WORKDIR /app

ADD ad-service /app
EXPOSE 8000
CMD ["./ad-service"]