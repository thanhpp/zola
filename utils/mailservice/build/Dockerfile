# ----------------------------------------------------------
# Build layer
FROM golang:alpine AS builder

# Arguments
ARG GITTOKEN
ENV GITTOKEN=${GITTOKEN:-EMPTY}

# set ENV variables
ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64

# Build directory
WORKDIR /build

# git install
RUN apk --update add git less openssh && \
    rm -rf /var/lib/apt/lists/* && \
    rm /var/cache/apk/*

# Config for private repo
RUN echo ${GITTOKEN}
RUN git config --global url."https://"${GITTOKEN}"@github.com".insteadOf "https://github.com"

# go mod download for better layer cache
COPY ["./go.mod", "./go.sum", "./"]

RUN go mod download

# build app
COPY ["./", "./"]

RUN go build -v -o mailservice

# ----------------------------------------------------------
# Run layer
FROM alpine

WORKDIR /

COPY --from=builder ["/build/mailservice", "mailservice"]

EXPOSE 10000

CMD ["./mailservice"]