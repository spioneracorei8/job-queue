package models

import _ "github.com/go-playground/validator/v10"

type Register struct {
	IdCardNumber string `json:"id_card_number" validate:"max=13"`
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required"`
	TitleNameTh  string `json:"title_name_th" validate:"max=255"`
	FirstNameTh  string `json:"first_name_th" validate:"max=255"`
	LastNameTh   string `json:"last_name_th" validate:"max=255"`
	TitleNameEn  string `json:"title_name_en" validate:"max=255"`
	FirstNameEn  string `json:"first_name_en" validate:"max=255"`
	LastNameEn   string `json:"last_name_en" validate:"max=255"`
	MobilePhone  string `json:"mobile_phone" validate:"max=10"`
	OfficePhone  string `json:"office_phone" validate:"max=10"`
	BOD          string `json:"bod"`
	Gender       string `json:"gender"`
}
