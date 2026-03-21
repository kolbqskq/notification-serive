FROM golang:1.25.6-alpine AS builder

WORKDIR /app

COPY go.work go.work.sum ./
COPY api-gateway/go.mod api-gateway/go.sum ./api-gateway/
COPY notification-worker/go.mod notification-worker/go.sum ./notification-worker/
COPY history-service/go.mod history-service/go.sum ./history-service/
COPY proto/go.mod proto/go.sum ./proto/

RUN go work sync

COPY . .

ARG SERVICE
RUN CGO_ENABLED=0 go build -o /app/bin/service ./${SERVICE}/cmd

FROM alpine:3.23
WORKDIR /app
COPY --from=builder /app/bin/service ./service
CMD ["./service"]