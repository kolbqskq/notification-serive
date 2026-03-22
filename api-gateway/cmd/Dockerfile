FROM golang:1.25.6-alpine AS builder

WORKDIR /app

COPY api-gateway/ ./api-gateway/
COPY proto/ ./proto/

WORKDIR /app/api-gateway

RUN go mod edit -replace github.com/kolbqskq/notification-service/proto=../proto
RUN go mod download

RUN CGO_ENABLED=0 go build -o /app/bin/api-gateway ./cmd

FROM alpine:3.23
WORKDIR /app
COPY --from=builder /app/bin/api-gateway ./api-gateway
ENV PROJECT_ROOT=/app
CMD ["./api-gateway"]