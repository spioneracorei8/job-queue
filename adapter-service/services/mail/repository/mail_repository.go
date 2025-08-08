package repository

import (
	"adapter-service/constants"
	my_kafka "adapter-service/kafka"
	"adapter-service/models"
	my_mail "adapter-service/services/mail"
	"context"
	"encoding/json"
)

type mailRepository struct {
	producer my_kafka.Producer
}

func NewMailRepositoryImpl(producer my_kafka.Producer) my_mail.MailRepository {
	return &mailRepository{
		producer: producer,
	}
}

func (m *mailRepository) SendMail(ctx context.Context, form *models.MailForm) error {
	byteVal, _ := json.Marshal(form)
	if err := m.producer.PublishMessage(constants.TOPIC_EMAIL_TOPIC, string(byteVal)); err != nil {
		return err
	}
	return nil
}
