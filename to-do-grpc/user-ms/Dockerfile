#Build stage
FROM golang:1.24.0-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o user-ms .

#Run stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/user-ms .

EXPOSE 8082

CMD ["./user-ms"]