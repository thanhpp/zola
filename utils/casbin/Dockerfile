# ----------------------------------------------------------
# Build layer
FROM golang:alpine AS builder

# set ENV variables
ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64

# Build directory
WORKDIR /build

# go mod download for better layer cache
COPY ["./", "./"]

RUN go mod tidy
RUN go mod download

# build app

RUN go build -v -o gt-casbin 

# ----------------------------------------------------------
# Run layer
FROM alpine

WORKDIR /

COPY --from=builder ["/build/gt-casbin", "gt-casbin"]
COPY --from=builder ["/build/dev.yml", "dev.yml"]
COPY --from=builder ["/build/model.conf", "model.conf"]
COPY --from=builder ["/build/rbac_with_domain_policy.csv", "rbac_with_domain_policy.csv"]

EXPOSE 80

CMD ["./gt-casbin"]
