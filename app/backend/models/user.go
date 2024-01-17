package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"size:255;not null"`
	Email    string `gorm:"size:255;not null"`
	Password string `gorm:"size:255;not null"`
}

type SignUpInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (user *User) Create(db *gorm.DB) error {
	result := db.Create(&user)

	return result.Error
}

func (user *User) Validate() error {
	err := validation.ValidateStruct(user,
		validation.Field(&user.Name,
			validation.Required.Error("Name is required"),
			validation.Length(1, 255).Error("Name is too long"),
		),
		validation.Field(&user.Email,
			validation.Required.Error("Email is required"),
			validation.Length(1, 255).Error("Email is too long"),
		),
		validation.Field(&user.Password,
			validation.Required.Error("Password is required"),
			validation.Length(8, 255).Error("Password is less than 7 chars or more than 256 chars"),
		),
	)

	return err
}
