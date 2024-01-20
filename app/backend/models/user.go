package models

import (
	"backend/errors"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"golang.org/x/crypto/bcrypt"
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

type SignInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (user *User) CheckUser(db *gorm.DB) *errors.MyError {
	userInDB := &User{}
	if result := db.Where("email = ?", user.Email).First(userInDB); result.RowsAffected == 0 {
		return &errors.MyError{
			Message: "email or password is wrong!",
			Code:    http.StatusUnauthorized,
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userInDB.Password), []byte(user.Password)); err != nil {
		return &errors.MyError{
			Message: "email or password is wrong!",
			Code:    http.StatusUnauthorized,
		}
	}

	user.ID = userInDB.ID

	return nil
}

func (user *User) Create(db *gorm.DB) *errors.MyError {
	if result := db.Where("email = ?", user.Email).First(&User{}); result.RowsAffected != 0 {
		return &errors.MyError{
			Message: "already exist same email user!",
			Code:    http.StatusConflict,
		}
	}
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return &errors.MyError{
			Message: "internal server error!",
			Code:    http.StatusInternalServerError,
		}
	}
	user.Password = string(hashedPwd)
	result := db.Create(&user)
	if result.Error != nil {
		return &errors.MyError{
			Message: "internal server error!",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
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
