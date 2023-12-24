package entities

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
	UserRoleAdmin = "ADMIN"
	UserRoleStaff = "STAFF"
	UserRoleUser  = "USER"
)

type User struct {
	ID          uint      `gorm:"primarykey"`
	DisplayName string    `gorm:"unique;not null"`
	Password    string    `gorm:"not null"`
	Email       string    `gorm:"unique;not null"`
	Role        string    `gorm:"not null;default:USER"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}

type UserRegisterDTO struct {
	DisplayName string `json:"displayname" validate:"required,min=3,max=32"`
	Password    string `json:"password" validate:"required,min=6,max=32"`
	Email       string `json:"email" validate:"required,email,min=6,max=50"`
}

func ValidateUserRegisterDTO(c *fiber.Ctx) UserRegisterDTO {
	var dto UserRegisterDTO

	if err := c.BodyParser(&dto); err != nil {
		panic(err)
	}

	if err := validate.Struct(&dto); err != nil {
		panic(err)
	}

	return dto
}

type UserRegisterResponse struct {
	UserID      uint   `json:"userid"`
	DisplayName string `json:"displayname"`
	Email       string `json:"email"`
	Role        string `json:"role"`
}

type UserLoginDTO struct {
	Password string `json:"password" validate:"required,min=6,max=32"`
	Email    string `json:"email" validate:"required,email,min=6,max=50"`
}

func ValidateUserLoginDTO(c *fiber.Ctx) UserLoginDTO {
	var dto UserLoginDTO

	if err := c.BodyParser(&dto); err != nil {
		panic(err)
	}

	if err := validate.Struct(&dto); err != nil {
		panic(err)
	}

	return dto
}

type UserLoginResponse struct {
	Token string `json:"token"`
	UserRegisterResponse
}
