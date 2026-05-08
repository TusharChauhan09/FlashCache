# ! ---------- Build Stage ----------
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download 

COPY . .

RUN go build -o flashcache ./cmd/server 

#! ---------- Run Stage ----------
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/flashcache .

EXPOSE 6379

CMD ["./flashcache"]