FROM postgres:16.1

LABEL author="Oleh Shalapskyi"
LABEL description="Postgres Image for Books database"
LABEL version="1.0"

ENV POSTGRES_USER="postgres"
ENV POSTGRES_PASSWORD=159357
ENV POSTGRES_DB="books"

COPY *.sql /docker-entrypoint-initdb.d/