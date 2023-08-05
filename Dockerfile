# FROM apache/skywalking-go:0.2.0-go1.20 AS builder
FROM golang:latest AS builder
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app
COPY . .
# RUN skywalking-go-agent -inject .
RUN go mod tidy
# RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -toolexec="skywalking-go-agent" -a .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" .
RUN ls
RUN mkdir publish  \
    && cp TuringBackend publish  \
    && cp -r config publish
FROM busybox:1.28.4
WORKDIR /app
COPY --from=builder /app/publish .
ENV GIN_MODE=release
EXPOSE 5001
ENTRYPOINT ["./TuringBackend"]