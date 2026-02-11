# This is for the docker-compose, which means, root path is /server
# Build Stage
FROM golang:1.25-alphine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLE=0 GOOS=linux go build -p 1 -o main ./admin/main

# Run Stage
FROM alphine:latest

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]