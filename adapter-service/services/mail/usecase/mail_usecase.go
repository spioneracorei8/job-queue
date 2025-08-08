package usecase

import (
	"adapter-service/models"
	"adapter-service/services/mail"
	"context"
)

type mailUsecase struct {
	mailRepo mail.MailRepository
}

func NewMailUsecaseImpl(mailRepo mail.MailRepository) mail.MailUsecase {
	return &mailUsecase{
		mailRepo: mailRepo,
	}
}

func (m *mailUsecase) SendMail(ctx context.Context, form *models.MailForm) error {
	return m.mailRepo.SendMail(ctx, form)

}
