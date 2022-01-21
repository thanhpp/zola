FROM ghcr.io/thanhpp/zola/aucobuilder AS builder

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

WORKDIR /go/src/github.com/thanhpp/zola/cmd/auco
RUN go build -v -o auco

# WORKDIR /go/src/github.com/thanhpp/zola
# RUN rm -rf *

# ----------------------------------------------------------------------------

FROM alpine

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait
RUN chmod +x /wait

WORKDIR /run
COPY --from=builder ["/go/src/github.com/thanhpp/zola/cmd/auco/auco", "auco"]
COPY --from=builder ["/go/src/github.com/thanhpp/zola/cmd/auco/auco.yml", "auco.yml"]

WORKDIR /run/public
COPY --from=builder ["/go/src/github.com/thanhpp/zola/cmd/auco/public", "."]

EXPOSE 12000

WORKDIR /run
CMD /wait && /run/auco