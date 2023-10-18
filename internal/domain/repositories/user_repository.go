package repositories

import (
	"InterviewPracticeBot-BE/internal/domain/entities"
	"errors"
)

var ErrUserNotFound = errors.New("user not found")

type UserRepository interface {
	FindByID(id string) (*entities.UserPrivate, error)

	FindByEmail(email string) (*entities.UserPrivate, error)

	Save(user *entities.UserPrivate) error

	Delete(id string) error
}
