FROM postgres:alpine AS db-auth

COPY ./postgres_init.sql /docker-entrypoint-initdb.d/postgres_init.sql
EXPOSE 5432