FROM postgres

WORKDIR /docker-entrypoint-initdb.d

COPY ["./postgresmultipledb.sh", "."]

RUN chmod +x ./postgresmultipledb.sh