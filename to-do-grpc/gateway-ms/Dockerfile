#Build stage
FROM golang:1.24.0-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o gateway-ms .

#Run stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/gateway-ms .

EXPOSE 8080

CMD ["./gateway-ms"]