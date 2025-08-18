package kafka

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

type consumer struct {
	maxRetries int
}

func NewQueueConsumerImpl(maxRetries int) Consumer {
	return &consumer{maxRetries: maxRetries}
}

func (c *consumer) StartWorker(brokerURL, topic, groupId string, worker func([]byte) error) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerURL},
		GroupID: groupId,
		Topic:   topic,
	})
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			continue
		}
		for i := 1; i <= c.maxRetries; i++ {
			if err := worker(m.Value); err == nil {
				break
			}
			time.Sleep(time.Second * time.Duration(i))
		}
	}
}
