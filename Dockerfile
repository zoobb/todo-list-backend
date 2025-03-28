FROM golang:1.23 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o main ./cmd

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/main /app/main

RUN apk add --no-cache libc6-compat

EXPOSE 8080

CMD ["/app/main"]