package usecase

import (
	"auth-service/models"
	"auth-service/services/adapter"
)

type adapterUsecase struct {
	grpcAdapterRepo adapter.GrpcAdapterRepository
}

func NewAdapterUsecaseImpl(grpcAdapterRepo adapter.GrpcAdapterRepository) adapter.AdapterUsecase {
	return &adapterUsecase{
		grpcAdapterRepo: grpcAdapterRepo,
	}
}

func (u *adapterUsecase) SendLog(logForm *models.LogForm) {
	u.grpcAdapterRepo.SendLog(logForm)
}
