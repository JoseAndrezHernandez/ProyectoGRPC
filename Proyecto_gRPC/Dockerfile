FROM golang:1.23.3 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download


COPY  . .
RUN go build -o server .

FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /app/server .
COPY  .env /app/.env
RUN apt-get update && apt-get install -y ca-certificates && update-ca-certificates

EXPOSE 50051 8080
CMD ["./server"]