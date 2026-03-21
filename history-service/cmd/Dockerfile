FROM golang:1.25.6-alpine AS builder

WORKDIR /app

COPY history-service/ ./history-service/
COPY proto/ ./proto/

WORKDIR /app/history-service

RUN go mod edit -replace github.com/kolbqskq/notification-service/proto=../proto
RUN go mod download

RUN CGO_ENABLED=0 go build -o /app/bin/history-service ./cmd

FROM alpine:3.23
WORKDIR /app
COPY --from=builder /app/bin/history-service ./history-service
CMD ["./history-service"]