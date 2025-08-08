package kafka

import (
	"adapter-service/helper"
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

type producer struct {
	writers map[string]*kafka.Writer
}

func NewQueueProducerImpl() Producer {
	return &producer{
		writers: make(map[string]*kafka.Writer),
	}
}

func (p *producer) InitKafkaWriter(brokerURL, topic string) {
	p.writers[topic] = &kafka.Writer{
		Addr:     kafka.TCP(brokerURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
		// BatchSize:    10,
		// BatchTimeout: 1 * time.Second,
		Async: true,
	}
}

func (p *producer) PublishMessage(topic, payload string) error {
	writer, ok := p.writers[topic]
	if !ok {
		return fmt.Errorf("writer for topic %s not initialized", topic)
	}
	msgs := kafka.Message{
		Time:  helper.NewTimestampFromTime(time.Now()),
		Key:   []byte(time.Now().Format(time.RFC3339)),
		Value: []byte(payload),
	}
	return writer.WriteMessages(context.Background(), msgs)
}
