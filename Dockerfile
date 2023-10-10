# FROM apache/skywalking-go:0.2.0-go1.20 AS builder
FROM golang:latest AS builder
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app
COPY . .
RUN go mod tidy
# RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -toolexec="skywalking-go-agent" -a .
RUN CGO_ENABLED=0 GOOS=linux go build -toolexec="/app/agent/skywalking-go-agent-0.2.0-linux-amd64" -a -ldflags="-w -s" .
RUN mkdir publish  \
    && cp TuringBackend publish  \
    && cp -r config publish 
FROM busybox:1.28.4
WORKDIR /app
COPY --from=builder /app/publish .
ENV GIN_MODE=release
EXPOSE 5001
ENTRYPOINT ["./TuringBackend"]