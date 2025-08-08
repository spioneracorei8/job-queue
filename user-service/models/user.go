package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	Id           *uuid.UUID `gorm:"index;type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	IdCardNumber string     `gorm:"size:13" json:"id_card_number"`
	TitleNameTH  string     `gorm:"size:255" json:"title_name_th"`
	FirstNameTH  string     `gorm:"size:255" json:"first_name_th"`
	LastNameTH   string     `gorm:"size:255" json:"last_name_th"`
	TitleNameEN  string     `gorm:"size:255" json:"title_name_en"`
	FirstNameEN  string     `gorm:"size:255" json:"first_name_en"`
	LastNameEN   string     `gorm:"size:255" json:"last_name_en"`
	MobilePhone  string     `gorm:"size:10" json:"mobile_phone"`
	OfficePhone  string     `gorm:"size:10" json:"office_phone"`
	Email        string     `gorm:"size:255" json:"email"`
	BOD          time.Time  `json:"bod"`                            // Birth of Date
	Gender       string     `gorm:"type:user_gender" json:"gender"` // MALE, FEMALE
	CreatedBy    string     `gorm:"type:text" json:"created_by"`
	UpdatedBy    string     `gorm:"type:text" json:"updated_by"`
	DeletedBy    *string    `gorm:"type:text" json:"deleted_by"`
	CreatedAt    time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}

func (User) TableName() string {
	return "user"
}

func (u *User) GenUUID() {
	id, _ := uuid.NewV4()
	u.Id = &id
}
