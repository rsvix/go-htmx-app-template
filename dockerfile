# Build
FROM golang:1.22 AS build-stage
ENV CGO_ENABLED=0
ENV GOOS=linux
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build ./cmd/main.go

# Deploy
# FROM alpine:latest AS production
FROM gcr.io/distroless/static-debian11 AS release-stage
COPY --from=build-stage /app ./
EXPOSE 9001
CMD ["./main"]