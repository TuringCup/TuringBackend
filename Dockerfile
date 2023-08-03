FROM golang:latest AS builder
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct
WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build  -ldflags="-w -s" -o ./main
RUN mkdir publish  \
    && cp main publish  \
    && cp -r config publish
FROM busybox:1.28.4
WORKDIR /app
COPY --from=builder /app/publish .
ENV GIN_MODE=release
EXPOSE 8080
ENTRYPOINT ["./main"]