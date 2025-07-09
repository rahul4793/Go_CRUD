FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git ca-certificates

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o crud-app .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/crud-app .

COPY .env .

EXPOSE 8080

CMD ["./crud-app"]
