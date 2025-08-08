package mail

import (
	"adapter-service/models"
	"context"
)

type MailRepository interface {
	SendMail(ctx context.Context, form *models.MailForm) error
}
