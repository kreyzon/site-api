package user

import (
	"errors"
	"fmt"
	"site/api/pkg/controllers"
	"site/api/pkg/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type controller struct {
	db *gorm.DB
}

type Option func(*controller) error

func WithGorm(db *gorm.DB) Option {
	return func(c *controller) error {
		c.db = db
		return nil
	}
}

func New(options ...Option) (controllers.UserController, error) {
	c := new(controller)
	for _, opt := range options {
		if err := opt(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (c *controller) Login(username string, password string) (*controllers.UserDTO, error) {
	userDB := models.User{}
	result := c.db.First(&userDB, "username = ?", username)

	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("user not found: %w", result.Error)
	}
	err := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("incorrect password: %w", err)
	}
	return toDTO(userDB), nil
}

func toDTO(userDB models.User) *controllers.UserDTO {
	return &controllers.UserDTO{
		Id:       int(userDB.ID),
		Username: userDB.Username,
		Fullname: userDB.Fullname,
		Role:     userDB.Role,
	}
}
