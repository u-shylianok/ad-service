FROM postgres:alpine AS ad-service-db-pg

COPY ./postgres_init.sql /docker-entrypoint-initdb.d/postgres_init.sql
EXPOSE 5432