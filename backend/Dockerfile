# Build Stage
FROM golang:1.25.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Run Stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

# this help me for permission denied from openshift 
RUN chmod 755 /app

EXPOSE 8000

CMD ["./main"]
