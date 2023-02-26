FROM golang:1.19.1-alpine3.16 as builder

WORKDIR /home
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && cd rpc \
    && go build -ldflags="-s -w" opsrpc.go

FROM alpine:latest

LABEL MAINTAINER="jiangbo202@gmail.com"

WORKDIR /home

COPY --from=builder /home/rpc/opsrpc ./
COPY --from=builder /home/rpc/etc/opsrpc.yaml ./

EXPOSE 9001
ENTRYPOINT ./opsrpc -f opsrpc.yaml
