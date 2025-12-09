FROM golang:1.24-alpine AS builder
WORKDIR /app

COPY go.mod .
RUN go mod download
COPY . .

RUN go build -o src/main ./src

FROM alpine:3.20
WORKDIR /app

COPY --from=builder /app/src/main .
COPY entrypoint.sh .
RUN chmod +x entrypoint.sh

CMD ["./entrypoint.sh"]