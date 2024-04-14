# Build
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /cmd/main

EXPOSE 8080

# Build
FROM alpine:latest AS production

COPY --from=builder /app .

EXPOSE 9001

CMD ["./main"]