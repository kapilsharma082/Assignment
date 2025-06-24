# Build stage
FROM golang:1.20-alpine as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o service1

# Runtime stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/service1 .
EXPOSE 5001
ENTRYPOINT ["/app/service1"]
