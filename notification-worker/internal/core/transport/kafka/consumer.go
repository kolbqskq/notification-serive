package kafka

import (
	"context"
	"encoding/json"
	"runtime/debug"
	"time"

	"github.com/kolbqskq/notification-service/notification-worker/internal/core/events"
	"github.com/rs/zerolog"
	"github.com/segmentio/kafka-go"
)

type EventHandler func(ctx context.Context, event events.NotificationEvent) error

type KafkaConsumer struct {
	reader  *kafka.Reader
	logger  zerolog.Logger
	handler EventHandler
}

type KafkaConsumerDeps struct {
	Reader  *kafka.Reader
	Logger  zerolog.Logger
	Handler EventHandler
	Brokers []string
	Topic   string
	GroupID string
}

func NewKafkaConsumer(deps KafkaConsumerDeps) *KafkaConsumer {
	return &KafkaConsumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: deps.Brokers,
			Topic:   deps.Topic,
			GroupID: deps.GroupID,
		}),
		logger:  deps.Logger,
		handler: deps.Handler,
	}
}

func (k *KafkaConsumer) Run(ctx context.Context) error {
	k.logger.Info().Msg("consumer started")

	for {
		if ctx.Err() != nil {
			return ctx.Err()
		}

		msg, err := k.reader.FetchMessage(ctx)
		if err != nil {
			k.logger.Error().Err(err).Msg("fetch message error")
			select {
			case <-time.After(time.Second):
			case <-ctx.Done():
				return ctx.Err()
			}
			continue
		}
		k.processMessage(ctx, msg)
	}
}

func (k *KafkaConsumer) processMessage(ctx context.Context, msg kafka.Message) {
	defer func() {
		if r := recover(); r != nil {
			k.logger.Error().Interface("panic", r).Str("stack", string(debug.Stack())).Msg("panic recovered")
			k.reader.CommitMessages(ctx, msg)
		}
	}()
	var event events.NotificationEvent

	if err := json.Unmarshal(msg.Value, &event); err != nil {
		k.logger.Error().Err(err).Msg("unmarshal error")
		if err := k.reader.CommitMessages(ctx, msg); err != nil {
			k.logger.Error().Err(err).Msg("commit error failed unmarshal")
		}
		return
	}
	handlerCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := k.handler(handlerCtx, event); err != nil {
		k.logger.Error().Err(err).Str("event_id", event.ID.String()).Str("type", event.Type).Msg("handler error")
		return
	}
	if err := k.reader.CommitMessages(ctx, msg); err != nil {
		k.logger.Error().Err(err).Msg("commit error")
	}

}

func (k *KafkaConsumer) Close() error {
	return k.reader.Close()
}
