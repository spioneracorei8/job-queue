package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type consumer struct{}

func NewQueueConsumerImpl() Consumer {
	return &consumer{}
}

func (c *consumer) StartWorker(brokerURL, topic, groupId string, worker func([]byte)) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerURL},
		GroupID: groupId,
		Topic:   topic,
	})
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			logrus.Fatalf("ReadMessage error: %s", err.Error())
			continue
		}
		worker(m.Value)
	}
}
