FROM postgres:alpine AS db-ads

COPY ./postgres_init.sql /docker-entrypoint-initdb.d/postgres_init.sql
EXPOSE 5432