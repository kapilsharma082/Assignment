FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o service1 main.go

FROM alpine:latest
WORKDIR /app
RUN apk add --no-cache curl
COPY --from=builder /app/service1 .
EXPOSE 5001
ENTRYPOINT ["/app/service1"]
