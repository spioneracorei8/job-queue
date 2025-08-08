package kafka

import "adapter-service/constants"

type Producer interface {
	InitKafkaWriter(brokerURL, topic string)
	PublishMessage(topic, payload string) error
}

type Consumer interface {
	StartWorker(brokerURL, topic, groupId string, worker func([]byte))
}

type kafkaQueue struct {
	queueProducer Producer
	queueConsumer Consumer
	SMTP_ADDRESS  string
	SMTP_HOST     string
	SMTP_PORT     int
	SMTP_USERNAME string
	SMTP_PASSWORD string
	SENDER_NAME   string
}

func NewKafkaQueue(queueProducer Producer, queueConsumer Consumer, SMTP_ADDRESS, SMTP_HOST string, SMTP_PORT int, SMTP_USERNAME, SMTP_PASSWORD, SENDER_NAME string) *kafkaQueue {
	return &kafkaQueue{
		queueProducer: queueProducer,
		queueConsumer: queueConsumer,
		SMTP_ADDRESS:  SMTP_ADDRESS,
		SMTP_HOST:     SMTP_HOST,
		SMTP_PORT:     SMTP_PORT,
		SMTP_USERNAME: SMTP_USERNAME,
		SMTP_PASSWORD: SMTP_PASSWORD,
		SENDER_NAME:   SENDER_NAME,
	}
}

func (k *kafkaQueue) StartKafkaQueue(brokerURL string) {
	k.queueProducer.InitKafkaWriter(brokerURL, constants.TOPIC_EMAIL_TOPIC)
	// k.queueProducer.InitKafkaWriter(brokerURL, "TEST-TOPIC")

	go k.queueConsumer.StartWorker(brokerURL, constants.TOPIC_EMAIL_TOPIC, constants.GROUP_EMAIL_GROUP, k.sendMailWorker)
	// go k.queueConsumer.StartWorker(brokerURL, "TEST-TOPIC", "TEST-GROUP", k.testWorker)
}
