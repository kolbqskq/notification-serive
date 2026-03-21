package transport_kafka

import (
	"context"
	"encoding/json"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	writer *kafka.Writer
}

type KafkaProducerDeps struct {
	Topic   string
	Brokers []string
}

func NewKafkaProducer(deps KafkaProducerDeps) *KafkaProducer {
	writer := &kafka.Writer{
		Addr:         kafka.TCP(deps.Brokers...),
		Topic:        deps.Topic,
		Balancer:     &kafka.Hash{},
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,

		BatchSize:    100,
		BatchTimeout: 10 * time.Millisecond,
		RequiredAcks: kafka.RequireAll,
		MaxAttempts:  3,
	}
	return &KafkaProducer{
		writer: writer,
	}
}

func (k *KafkaProducer) Publish(ctx context.Context, key string, v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	return k.writer.WriteMessages(ctx, kafka.Message{
		Value: data,
		Key:   []byte(key),
	})
}

func (k *KafkaProducer) Close() error {
	return k.writer.Close()
}
