## ----------------------------------------------------------------------------
FROM golang:1.17-alpine

# set ENV variables
ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src/github.com/thanhpp/zola

COPY pkg pkg/
COPY cmd cmd/
COPY config config/
COPY internal internal/
COPY go.mod go.mod

RUN go mod tidy

WORKDIR /go/src/github.com/thanhpp/zola/cmd/laclongquan

RUN go build -v -o laclongquan

WORKDIR /go/src/github.com/thanhpp/zola

RUN rm -rf *