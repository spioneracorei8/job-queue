package usecase

import (
	"user-service/models"
	"user-service/services/user"
)

type userUsecase struct {
	userRepository user.UserRepository
}

func NewUserUsecaseImpl(userRepository user.UserRepository) user.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (u *userUsecase) FetchUserByIdCardNumber(idCardNumber string) (*models.User, error) {
	return u.userRepository.FetchUserByIdCardNumber(idCardNumber)
}

func (u *userUsecase) UpsertUser(user *models.User) error {
	return u.userRepository.UpsertUser(user)
}
