FROM golang:1.25.6-alpine AS builder

WORKDIR /app

COPY notification-worker/ ./notification-worker/
COPY api-gateway/ ./api-gateway/

WORKDIR /app/notification-worker

RUN go mod edit -replace github.com/kolbqskq/notification-service/api-gateway=../api-gateway
RUN go mod download

RUN CGO_ENABLED=0 go build -o /app/bin/notification-worker ./cmd

FROM alpine:3.23
WORKDIR /app
COPY --from=builder /app/bin/notification-worker ./notification-worker
ENV PROJECT_ROOT=/app
CMD ["./notification-worker"]