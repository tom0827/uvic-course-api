FROM golang:1.22-alpine AS builder
WORKDIR /app

# Copy go.mod and go.sum first (for caching downloads)
COPY go.mod .
RUN go mod download

# Copy the rest of the source code (including handler/)
COPY . .

RUN go build -o server main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/server .
CMD ["./server"]