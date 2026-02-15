# This is for the docker-compose, which means, root path is /server
# Build Stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLE=0 GOOS=linux go build -p 1 -o main ./client/main
RUN CGO_ENABLE=0 GOOS=linux go build -o health_checker ./pkg/common/healthcheck/main.go

# Run Stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/health_checker /usr/local/bin/health_checker

EXPOSE 8082 9090

HEALTHCHECK --interval=10s --timeout=3s --retries=3 \
    CMD [ "executable" ]

CMD ["./main"]