package aggregates

import (
	"InterviewPracticeBot-BE/internal/domain/entities"
	"InterviewPracticeBot-BE/internal/domain/repositories"
)

type IUserAggregate interface {
	Register() (*entities.UserPrivate, error)
	UpdatePassword(oldPassword, NewPassword string, userRepo repositories.IUserPrivateRepository) error
}
