package repository

import (
	"adapter-service/constants"
	"adapter-service/kafka"
	"encoding/json"

	"adapter-service/proto/proto_models"
	"adapter-service/services/log"
	"context"
)

type logRepository struct {
	producer kafka.Producer
}

func NewLogRepositoryImpl(producer kafka.Producer) log.GrpcLogRepository {
	return &logRepository{
		producer: producer,
	}
}

func (r *logRepository) SaveLog(ctx context.Context, params *proto_models.SendLogRequest) error {
	byteVal, _ := json.Marshal(params)
	if err := r.producer.PublishMessage(constants.TOPIC_LOG_TOPIC, string(byteVal)); err != nil {
		return err
	}
	return nil
}
