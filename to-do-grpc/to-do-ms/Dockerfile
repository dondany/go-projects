#Build stage
FROM golang:1.24.0-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o to-do-ms .

#Run stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/to-do-ms .

EXPOSE 8081

CMD ["./to-do-ms"]