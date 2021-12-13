FROM postgres

ENV POSTGRES_USER=thanhpp \
    POSTGRES_PASSWORD=testthanhpp \
    POSTGRES_MULTIPLE_DATABASES=prom,thanhpp:promuser,thanhpp

WORKDIR /docker-entrypoint-initdb.d

COPY ["./postgresmultipledb.sh", "."]

RUN chmod +x ./postgresmultipledb.sh