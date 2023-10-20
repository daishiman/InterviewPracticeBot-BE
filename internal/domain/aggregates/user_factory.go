package aggregates

import (
	"InterviewPracticeBot-BE/internal/domain/entities"
	"InterviewPracticeBot-BE/internal/domain/utilities"
	"InterviewPracticeBot-BE/internal/domain/value_objects"
	"time"
)

type UserFactory struct {}

func NewUserFactory() *UserFactory {
  return & UserFactory{}
}

func (uf *UserFactory) CreateUser(email, rawPassword string) (*entities.UserPrivate, error) {
  emailVO, err := value_objects.NewEmail(email)
  if err != nil {
    return nil, err
  }

  passwordVO, err := value_objects.NewPassword(rawPassword)
  if err != nil {
    return nil, err
  }

  user := &entities.UserPrivate{
    ID: utilities.GenerateUUID(),
    Email: *emailVO,
    Password: *passwordVO,
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
  }
  return user, nil
}
