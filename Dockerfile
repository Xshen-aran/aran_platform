FROM golang:1.19.9 as builder
USER root
WORKDIR /app
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
COPY . .
RUN go mod download
RUN go build -o main .


FROM debian:bullseye-slim
COPY --from=builder /app/main /usr/local/bin/platform