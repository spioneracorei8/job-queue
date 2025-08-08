package repository

import (
	"user-service/models"
	"user-service/services/user"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepoImpl(db *gorm.DB) user.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) FetchUserByIdCardNumber(idCardNumber string) (*models.User, error) {
	var user = new(models.User)
	if err := r.db.Where("id_card_number = ?", idCardNumber).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (r *userRepository) UpsertUser(user *models.User) error {
	tx := r.db.Begin()

	if err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		UpdateAll: true,
	}).Create(&user).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}
