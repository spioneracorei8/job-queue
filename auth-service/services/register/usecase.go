package register

import (
	"auth-service/models"
	"context"
)

type RegisterUsecase interface {
	RegisterUser(ctx context.Context, register *models.Register, source string) error
}
