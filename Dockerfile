FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o main main.go

FROM alpine:3.20
WORKDIR /app

# Install openssl for cert fetching
RUN apk add --no-cache openssl

COPY --from=builder /app/main .
COPY entrypoint.sh .

RUN chmod +x entrypoint.sh

CMD ["./entrypoint.sh"]