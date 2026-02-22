# This is for the docker-compose, which means, root path is /server
# Build Stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -p 1 -o main ./client/main

# Run Stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8082

HEALTHCHECK --interval=10s --timeout=3s --retries=3 \
    CMD ["/usr/local/bin/health_checker" ]

CMD ["./main"]