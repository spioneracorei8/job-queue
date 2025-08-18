package adapter

import "auth-service/models"

type AdapterUsecase interface {
	SendLog(logForm *models.LogForm)
}
