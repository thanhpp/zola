## ----------------------------------------------------------------------------
FROM golang:1.17-alpine AS builder

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

# ----------------------------------------------------------------------------

FROM alpine

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait
RUN chmod +x /wait

WORKDIR /run

COPY --from=builder ["/go/src/github.com/thanhpp/zola/cmd/laclongquan/laclongquan", "laclongquan"]
COPY --from=builder ["/go/src/github.com/thanhpp/zola/cmd/laclongquan/laclongquan.yml", "laclongquan.yml"]
COPY --from=builder ["/go/src/github.com/thanhpp/zola/cmd/laclongquan/pri.rsa", "pri.rsa"]
COPY --from=builder ["/go/src/github.com/thanhpp/zola/cmd/laclongquan/pri.rsa.pub", "pri.rsa.pub"]

EXPOSE 12100

CMD /wait && /run/laclongquan
