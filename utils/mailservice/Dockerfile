# ------------------------------------------
# BUILD LAYER
FROM golang:1.15-alpine AS builder

# set ENV variables
ENV GO111MODULE=on \
  GOOS=linux \
  GOARCH=amd64

# Build directory
WORKDIR /build
RUN apk add git
RUN git config --global url."https://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@github.com/".insteadOf "https://github.com"

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./
RUN go build -v -o mailservice .

# ------------------------------------------
# RUN LAYER
FROM alpine

WORKDIR /
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=builder ["/build/mailservice", "mailservice"]

CMD ["./mailservice"]
