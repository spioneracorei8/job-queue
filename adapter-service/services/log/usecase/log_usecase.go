package usecase

import (
	"adapter-service/proto/proto_models"
	"adapter-service/services/log"
	"context"
)

type logUsecase struct {
	grpcLogRepo log.GrpcLogRepository
}

func NewAdapterUsecaseImpl(grpcLogRepo log.GrpcLogRepository) log.LogUsecase {
	return &logUsecase{
		grpcLogRepo: grpcLogRepo,
	}
}

func (u *logUsecase) SaveLog(ctx context.Context, params *proto_models.SendLogRequest) error {
	return u.grpcLogRepo.SaveLog(ctx, params)
}
