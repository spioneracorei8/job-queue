package user

import "user-service/models"

type UserUsecase interface {
	FetchUserByIdCardNumber(idCardNumber string) (*models.User, error)
	UpsertUser(user *models.User) error
}
