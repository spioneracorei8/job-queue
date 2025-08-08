package mail

import (
	"adapter-service/models"
	"context"
)

type MailUsecase interface {
	SendMail(ctx context.Context, form *models.MailForm) error
}
