package user

import (
	"auth-service/models"

	"github.com/gofrs/uuid"
)

type UserRepository interface {
	RegisterUser(register *models.Register) (*uuid.UUID, error)
}
