# FROM golang:1.21-alpine AS builder

FROM golang:1.24.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/main .
# COPY --from=builder /app/.env .

EXPOSE 8080

CMD ["./main"]